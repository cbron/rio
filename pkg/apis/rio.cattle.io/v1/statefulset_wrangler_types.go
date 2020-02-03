package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type StatefulSetWrangler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StatefulSetWranglerSpec   `json:"spec,omitempty"`
	Status StatefulSetWranglerStatus `json:"status,omitempty"`
}

type StatefulSetWranglerSpec struct {
	WorkloadSpec

	// maybe not here exactly but you get idea
	VolumeTemplates []VolumeTemplate `json:"volumeTemplates,omitempty"`
}

type StatefulSetWranglerStatus struct {
	WorkloadStatus
}

func (ssw StatefulSetWrangler) GetMeta() metav1.ObjectMeta {
	return ssw.ObjectMeta
}

func (ssw StatefulSetWrangler) GetSpec() WorkloadSpec {
	return ssw.Spec.WorkloadSpec
}

func (ssw StatefulSetWrangler) GetStatus() WorkloadStatus {
	return ssw.Status.WorkloadStatus
}

type VolumeTemplate struct {
	// Labels to be applied to the created PVC
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations to be applied to the created PVC
	Annotations map[string]string `json:"annotations,omitempty"`

	// Name of the VolumeTemplate. A volume entry will use this name to refer to the created volume
	Name string

	// AccessModes contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	AccessModes []v1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`
	// Resources represents the minimum resources the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	// +optional
	StorageRequest int64 `json:"storage,omitempty" mapper:"quantity"`
	// Name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName string `json:"storageClassName,omitempty"`
	// volumeMode defines what type of volume is required by the claim.
	// Value of Filesystem is implied when not included in claim spec.
	// This is a beta feature.
	// +optional
	VolumeMode *v1.PersistentVolumeMode `json:"volumeMode,omitempty"`
}

func StatefulSetWranglerWorkloadSlice(items []*StatefulSetWrangler) []Workload {
	wItems := make([]Workload, len(items))
	for i, v := range items {
		wItems[i] = Workload(v)
	}
	return wItems
}
