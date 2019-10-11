package testutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	riov1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
)

type TestService struct {
	Name    string // namespace/name:version
	AppName string // namespace/name
	App     riov1.App
	Service riov1.Service
	Version string
	Kubeconfig string
	T       *testing.T
}

// Create generates a new rio service, named randomly in the testing namespace, and
// returns a new TestService with it attached. Guarantees ready state but not live endpoint
func (ts *TestService) Create(t *testing.T, source ...string) {
	args := ts.createArgs(t, source...)
	var envs []string
	if ts.Kubeconfig != "" {
		envs = []string{fmt.Sprintf("KUBECONFIG=%s", ts.Kubeconfig)}
	}
	_, err := RioCmd(args, envs...)
	if err != nil {
		ts.T.Fatalf("Failed to create service:  %v", err.Error())
	}
	err = ts.waitForReadyService()
	if err != nil {
		ts.T.Fatalf(err.Error())
	}
}

func (ts *TestService) CreateWithError(t *testing.T, source ...string) error {
	args := ts.createArgs(t, source...)
	var envs []string
	if ts.Kubeconfig != "" {
		envs = []string{fmt.Sprintf("KUBECONFIG=%s", ts.Kubeconfig)}
	}
	_, err := RioCmd(args, envs...)
	if err != nil {
		return err
	}
	return nil
}

// Takes name and version of existing service and returns loaded TestService
func GetService(t *testing.T, name string, version string) TestService {
	ts := TestService{
		App:     riov1.App{},
		Service: riov1.Service{},
		Version: version,
		T:       t,
	}
	ts.AppName = fmt.Sprintf(
		"%s/%s",
		testingNamespace,
		name,
	)
	ts.Name = fmt.Sprintf("%s:%s", ts.AppName, ts.Version)
	err := ts.waitForReadyService()
	if err != nil {
		ts.T.Fatalf(err.Error())
	}
	return ts
}

// Remove calls "rio rm" on this service. Logs error but does not fail test.
func (ts *TestService) Remove() {
	if ts.Kubeconfig != "" {
		err := os.RemoveAll(ts.Kubeconfig)
		if err != nil {
			ts.T.Log(err.Error())
		}
	}
	if ts.Service.Status.DeploymentStatus != nil {
		_, err := RioCmd([]string{"rm", "--type", "service", ts.Name})
		if err != nil {
			ts.T.Log(err.Error())
		}
	}
}

// Call "rio scale ns/service={scaleTo}"
func (ts *TestService) Scale(scaleTo int) {
	_, err := RioCmd([]string{
		"scale",
		fmt.Sprintf("%s=%d", ts.Name, scaleTo),
	})
	if err != nil {
		ts.T.Fatalf("scale command failed:  %v", err.Error())
	}
	// First wait for scale on service to update
	err = ts.waitForScale(scaleTo)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
	// Then wait for actual replicas to come up
	err = ts.waitForAvailableReplicas(scaleTo)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
}

// Call "rio weight ns/service:version={percentage}" on this service
func (ts *TestService) Weight(percentage int) {
	_, err := RioCmd([]string{
		"weight",
		fmt.Sprintf("%s=%d", ts.Name, percentage),
	})
	if err != nil {
		ts.T.Fatalf("weight command failed:  %v", err.Error())
	}
	err = ts.waitForWeight(percentage)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
}

// Call "rio stage --image={source} ns/name:{version}", this will return a new TestService
func (ts *TestService) Stage(source, version string) TestService {
	name := fmt.Sprintf("%s:%s", ts.AppName, version)
	_, err := RioCmd([]string{"stage", "--image", source, name})
	if err != nil {
		ts.T.Fatalf("stage command failed:  %v", err.Error())
	}
	stagedService := TestService{
		T:       ts.T,
		AppName: ts.AppName,
		Name:    name,
		Version: version,
	}
	err = stagedService.waitForReadyService()
	if err != nil {
		ts.T.Fatalf(err.Error())
	}
	return stagedService
}

// GetEndpoint performs an http.get against the service endpoint and returns response if
// status code is 200, otherwise it errors out
func (ts *TestService) GetEndpoint() string {
	endpoint, err := ts.waitForEndpointDNS()
	if err != nil {
		ts.T.Fatal(err.Error())
	}
	response, err := WaitForURLResponse(endpoint)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
	return response
}

// Export calls "rio export {serviceName}" and returns that in a new TestService object
func (ts *TestService) Export() TestService {
	args := []string{"export", "--type", "service", "--format", "json", ts.Name}
	service, err := ts.loadExport(args)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
	return service
}

// ExportRaw works the same as export, but with --raw flag
func (ts *TestService) ExportRaw() TestService {
	args := []string{"export", "--raw", "--type", "service", "--format", "json", ts.Name}
	service, err := ts.loadExport(args)
	if err != nil {
		ts.T.Fatal(err.Error())
	}
	return service
}

//////////
// Getters
//////////

// Returns count of ready and available pods
func (ts *TestService) GetAvailableReplicas() int {
	if ts.Service.Status.DeploymentStatus != nil {
		return int(ts.Service.Status.DeploymentStatus.AvailableReplicas)
	}
	return 0
}

// Returns desired scale, different from current available replicas
func (ts *TestService) GetScale() int {
	if ts.Service.Spec.Scale != nil {
		return *ts.Service.Spec.Scale
	}
	return 0
}

// Return service's goal weight, this is different from weight service is currently at
func (ts *TestService) GetSpecWeight() int {
	return ts.Service.Spec.Weight
}

// Return service's actual current weight, not the spec (end-goal) weight
func (ts *TestService) GetCurrentWeight() int {
	ts.reloadApp()
	if val, ok := ts.App.Status.RevisionWeight[ts.Version]; ok {
		return val.Weight
	}
	return 0
}

func (ts *TestService) GetImage() string {
	return ts.Service.Spec.Image
}

// GetRunningPods returns all running pods, separated by new lines, for this service's app
// Each value, separated by spaces, will have the Pod's NAME  READY  STATUS  RESTARTS  AGE in that order.
func (ts *TestService) GetRunningPods() string {
	ts.reloadApp()
	args := append([]string{"get", "pods",
		"-n", testingNamespace,
		"-l", fmt.Sprintf("app=%s", ts.App.Name),
		"--field-selector", "status.phase=Running",
		"--no-headers"})
	out, err := KubectlCmd(args)
	if err != nil {
		ts.T.Fatalf("Failed to get running pods:  %v", err.Error())
	}
	return out
}

//////////////////
// Private methods
//////////////////

func (ts *TestService) createArgs(t *testing.T, source ...string) []string {
	ts.T = t
	ts.Version = "v0"
	ts.AppName = fmt.Sprintf(
		"%s/%s",
		testingNamespace,
		RandomString(5),
	)
	ts.Name = fmt.Sprintf("%s:%s", ts.AppName, ts.Version)
	if len(source) == 0 {
		source = []string{"nginx"}
	}
	args := append([]string{"run", "-p", "80/http", "-n", ts.AppName}, source...)
	return args
}

// reload calls inspect on the service and uses that to reload our object
func (ts *TestService) reload() error {
	out, err := RioCmd([]string{"inspect", "--type", "service", "--format", "json", ts.Name})
	if err != nil {
		return err
	}
	ts.Service = riov1.Service{}
	if err := json.Unmarshal([]byte(out), &ts.Service); err != nil {
		return err
	}
	return nil
}

// reload calls inspect on the service's app and uses that to reload the app obj
func (ts *TestService) reloadApp() error {
	out, err := RioCmd([]string{"inspect", "--type", "app", "--format", "json", ts.AppName})
	if err != nil {
		return err
	}
	ts.App = riov1.App{}
	if err := json.Unmarshal([]byte(out), &ts.App); err != nil {
		return err
	}
	return nil
}

// load a "rio export..." response into a new TestService obj
func (ts *TestService) loadExport(args []string) (TestService, error) {
	out, err := RioCmd(args)
	if err != nil {
		return TestService{}, fmt.Errorf("export command failed: %v", err.Error())
	}
	exportedService := riov1.Service{}
	if err := json.Unmarshal([]byte(out), &exportedService); err != nil {
		return TestService{}, fmt.Errorf("failed to parse export data: %v", err.Error())
	}
	stagedService := TestService{
		Service: exportedService,
		AppName: ts.AppName,
		Name:    ts.Name,
		Version: ts.Version,
	}
	return stagedService, nil
}

// Wait longer for higher scale, always wait at least 120 seconds
func (ts *TestService) getScalingTimeout() time.Duration {
	scalingTimeout := time.Duration(math.Max(float64(ts.GetScale())*20, 120))
	return time.Second * scalingTimeout
}

//////////////////
// Wait helpers
//////////////////

// Wait until a service hits ready state, or error out
func (ts *TestService) waitForReadyService() error {
	f := wait.ConditionFunc(func() (bool, error) {
		err := ts.reload()
		if err == nil {
			if ts.Service.Status.DeploymentStatus != nil && ts.Service.Status.DeploymentStatus.AvailableReplicas > 0 {
				return true, nil
			}
		}
		return false, nil
	})
	err := wait.Poll(2*time.Second, ts.getScalingTimeout(), f)
	if err != nil {
		return fmt.Errorf("service %v never reached ready status", ts.Name)
	}
	ts.reload()
	return nil
}

// Wait until a service hits wanted number of available replicas, or error out
func (ts *TestService) waitForAvailableReplicas(want int) error {
	f := wait.ConditionFunc(func() (bool, error) {
		err := ts.reload()
		if err == nil {
			if ts.GetAvailableReplicas() == want {
				return true, nil
			}
		}
		return false, nil
	})
	err := wait.Poll(2*time.Second, ts.getScalingTimeout(), f)
	if err != nil {
		return fmt.Errorf("service %v failed to scale up available replicas", ts.Name)
	}
	return nil
}

// Wait until a service's scale field hits wanted int, or error out
func (ts *TestService) waitForScale(want int) error {
	f := wait.ConditionFunc(func() (bool, error) {
		err := ts.reload()
		if err == nil {
			if ts.GetScale() == want {
				return true, nil
			}
		}
		return false, nil
	})
	err := wait.Poll(2*time.Second, 60*time.Second, f)
	if err != nil {
		return errors.New("service failed to scale")
	}
	return nil
}

// Wait until a service has actual weight we want, note this is not spec weight which changes immediately
func (ts *TestService) waitForWeight(percentage int) error {
	f := wait.ConditionFunc(func() (bool, error) {
		ts.reloadApp()
		if val, ok := ts.App.Status.RevisionWeight[ts.Version]; ok {
			if val.Weight == percentage {
				return true, nil
			}
		}
		return false, nil
	})
	err := wait.Poll(2*time.Second, 60*time.Second, f)
	if err != nil {
		return errors.New("service revision never reached goal weight")
	}
	return nil
}

func (ts *TestService) waitForEndpointDNS() (string, error) {
	if len(ts.Service.Status.Endpoints) > 0 {
		return ts.Service.Status.Endpoints[0], nil
	}
	f := wait.ConditionFunc(func() (bool, error) {
		err := ts.reload()
		if err == nil {
			if len(ts.Service.Status.Endpoints) > 0 {
				return true, nil
			}
		}
		return false, nil
	})
	err := wait.Poll(2*time.Second, 60*time.Second, f)
	if err != nil {
		return "", errors.New("service endpoint never created")
	}
	return ts.Service.Status.Endpoints[0], nil
}
