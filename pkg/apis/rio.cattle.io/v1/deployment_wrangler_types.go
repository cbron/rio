package v1

import (
	"github.com/rancher/wrangler/pkg/genericcondition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DeploymentWrangler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentWranglerSpec   `json:"spec,omitempty"`
	Status DeploymentWranglerStatus `json:"status,omitempty"`
}

type DeploymentWranglerSpec struct {
	// The exposed app name, if no value is set, then metadata.name of the Service is used
	App string `json:"app,omitempty"`

	// Version of this service
	Version string `json:"version,omitempty"`

	// The weight among services with matching app field to determine how much traffic is load balanced
	// to this service.  If rollout is set, the weight becomes the target weight of the rollout.
	Weight *int `json:"weight,omitempty"`

	// Autoscale the replicas based on the amount of traffic received by this service
	Autoscale *AutoscaleConfig `json:"autoscale,omitempty"`

	// RolloutDuration specifies time for template service to reach 100% weight, used to set rollout config
	RolloutDuration *metav1.Duration `json:"rolloutDuration,omitempty" mapper:"duration"`
}

type DeploymentWranglerStatus struct {
	// Represents the latest available observations of current state.
	Conditions []genericcondition.GenericCondition `json:"conditions,omitempty"`
}
