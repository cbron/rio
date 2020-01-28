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
	WranglerSpec
}

type DeploymentWranglerStatus struct {
	WranglerStatus
}
