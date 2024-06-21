package v1beta1

import (
	"github.com/rancher/wrangler/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=ipa;ipas,scope=Namespaced
// +kubebuilder:printcolumn:name="DESCRIPTION",type=string,JSONPath=`.spec.description`
// +kubebuilder:printcolumn:name="IPAM",type=string,JSONPath=`.spec.ipam`
// +kubebuilder:printcolumn:name="ALLOCATEDIP",type=string,JSONPath=`.status.AllocatedAddress.IP`

type IPAllocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IPAllocationSpec   `json:"spec"`
	Status            IPAllocationStatus `json:"status,omitempty"`
}

type IPAllocationSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	IPAM IPAM `json:"ipam,omitempty"`
	// +optional
	IPPool    string     `json:"ipPool,omitempty"`
	// +optional
	ExpectedAddress string `json:"expectedAddress,omitempty"`
	// +optional
	BindingKey string `json:"bindingKey,omitempty"`
}

type IPAllocationStatus struct {
	// +optional
	AllocatedAddress AllocatedAddress `json:"allocatedAddress,omitempty"`
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}

var (
	IPAllocationRequest condition.Cond = "Request"
	IPAllocationReply condition.Cond = "Reply"
)
