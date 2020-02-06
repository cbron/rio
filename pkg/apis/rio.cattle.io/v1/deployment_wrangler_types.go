package v1

import (
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
	WorkloadSpec
}

type DeploymentWranglerStatus struct {
	WorkloadStatus
}

func (dw DeploymentWrangler) GetMeta() metav1.ObjectMeta {
	return dw.ObjectMeta
}

func (dw DeploymentWrangler) GetSpec() WorkloadSpec {
	return dw.Spec.WorkloadSpec
}

func (dw DeploymentWrangler) GetStatus() WorkloadStatus {
	return dw.Status.WorkloadStatus
}

func (dw DeploymentWrangler) GetType() string {
	return "DeploymentWrangler"
}

func DeploymentWranglerWorkloadSlice(items []*DeploymentWrangler) []Workload {
	wItems := make([]Workload, len(items))
	for i, v := range items {
		wItems[i] = Workload(v)
	}
	return wItems
}
