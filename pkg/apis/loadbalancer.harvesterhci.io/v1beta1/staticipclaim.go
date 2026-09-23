package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StaticIPClaimState represents the lifecycle phase of a StaticIPClaim.
//
// Lifecycle State Machine:
//
//	                  ┌──────────────────────┐
//	                  │      Unassigned      │ (Initial state on creation)
//	                  └──────────┬───────────┘
//	                             │
//	            ┌────────────────┴────────────────┐
//	            │ IPPool Controller allocates IP  │ Pool error (e.g., OOM)
//	            ▼                                 ▼
//	   ┌─────────────────┐               ┌────────────────┐
//	   │    Reserved     │               │     Failed     │
//	   │ (Claimable by   │               │ (No IP issued) │
//	   │    Client)      │               └────────────────┘
//	   └───────┬▲────────┘
//	           ││
//	Client     ││ Client
//	claims IP  ││ releases IP
//	           ▼│
//	     ┌─────────────┐
//	     │    InUse    │
//	     │ (Active in  │
//	     │ Workload)   │
//	     └─────────────┘
//
// Transition Rules & Guard Conditions:
//   - Unassigned -> Reserved  : Managed by IPPool controller upon successfully leasing IP(s) from the pool.
//   - Unassigned -> Failed    : Managed by IPPool controller when pool is exhausted or target IP is invalid.
//   - Reserved   -> InUse     : Managed by Client/Guest when presenting matching spec.identifier.
//     Clients MUST ONLY attempt to claim IPs when state is 'Reserved'.
//   - InUse      -> Reserved  : Managed by Client/Guest upon workload detach/release.
//   - Failed     -> Unassigned: Re-evaluated by IPPool controller if pool capacity expands or config updates.
//
// +kubebuilder:validation:Enum=Unassigned;Reserved;InUse;Failed
type StaticIPClaimState string

const (
	// ClaimStateUnassigned indicates the claim has been created, but the IPPool controller
	// has not yet allocated IP address(es) from the referenced pool.
	ClaimStateUnassigned StaticIPClaimState = "Unassigned"

	// ClaimStateReserved indicates IP address(es) have been successfully allocated from
	// the IPPool and staged in status. The claim is now eligible to be claimed by a client.
	//
	// Guard: Clients MUST ONLY attempt to claim and attach IPs when state == Reserved.
	ClaimStateReserved StaticIPClaimState = "Reserved"

	// ClaimStateInUse indicates a client has successfully validated spec.identifier,
	// claimed the reserved IP(s), and attached them to an active workload interface.
	ClaimStateInUse StaticIPClaimState = "InUse"

	// ClaimStateFailed indicates the IPPool controller could not fulfill the allocation
	// request (e.g., Pool OOM or target IP out of range). No IP address is granted.
	ClaimStateFailed StaticIPClaimState = "Failed"
)

// StaticIPClaim represents an admin- or user-defined static IP reservation claim.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Pool",type="string",JSONPath=".spec.poolRef.name"
// +kubebuilder:printcolumn:name="IPv4",type="string",JSONPath=".status.allocatedIPv4"
// +kubebuilder:printcolumn:name="IPv6",type="string",JSONPath=".status.allocatedIPv6"
// +kubebuilder:printcolumn:name="Identifier",type="string",JSONPath=".spec.identifier"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
type StaticIPClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StaticIPClaimSpec   `json:"spec"`
	Status StaticIPClaimStatus `json:"status,omitempty"`
}

// StaticIPClaimList contains a list of StaticIPClaim.
// +kubebuilder:object:root=true
type StaticIPClaimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticIPClaim `json:"items"`
}

// StaticIPClaimSpec defines the desired state of StaticIPClaim.
type StaticIPClaimSpec struct {
	// Reference to the target IPPool. Immutable once created.
	// The IP family strategy (IPv4, IPv6, or DualStack) is governed by the referenced pool.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.poolRef is immutable"
	PoolRef LocalObjectReference `json:"poolRef"`

	// Static IPv4 address requested by user. Optional on creation.
	// If empty and referenced pool supports IPv4, controller auto-allocates an IPv4 address.
	// Immutable once set by the user.
	// +optional
	// +kubebuilder:validation:Format=ipv4
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.targetIPv4 is immutable once set. Delete and recreate to change IPs."
	TargetIPv4 string `json:"targetIPv4,omitempty"`

	// Static IPv6 address requested by user. Optional on creation.
	// If empty and referenced pool supports IPv6, controller auto-allocates an IPv6 address.
	// Immutable once set by the user.
	// +optional
	// +kubebuilder:validation:Format=ipv6
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.targetIPv6 is immutable once set. Delete and recreate to change IPs."
	TargetIPv6 string `json:"targetIPv6,omitempty"`

	// Unique secret token/identifier presented by the guest. Immutable once created.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=128
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9\-_.]+$`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="spec.identifier is immutable"
	Identifier string `json:"identifier"`
}

// StaticIPClaimStatus defines the observed state of StaticIPClaim.
type StaticIPClaimStatus struct {
	// State represents the current lifecycle phase of the claim: Unassigned | Reserved | InUse | Failed.
	// +optional
	State StaticIPClaimState `json:"state,omitempty"`

	// Allocated IPv4 address from the pool (populated in Reserved/InUse states).
	// +optional
	AllocatedIPv4 string `json:"allocatedIPv4,omitempty"`

	// Allocated IPv6 address from the pool (populated in Reserved/InUse states).
	// +optional
	AllocatedIPv6 string `json:"allocatedIPv6,omitempty"`

	// Standard Kubernetes conditions detailing programmatic transition reasons.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
