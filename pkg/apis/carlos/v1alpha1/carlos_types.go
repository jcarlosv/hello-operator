package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CarlosSpec defines the desired state of Carlos
// +k8s:openapi-gen=true
type CarlosSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// CarlosStatus defines the observed state of Carlos
// +k8s:openapi-gen=true
type CarlosStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Carlos is the Schema for the carlos API
// +k8s:openapi-gen=true
type Carlos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CarlosSpec   `json:"spec,omitempty"`
	Status CarlosStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CarlosList contains a list of Carlos
type CarlosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Carlos `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Carlos{}, &CarlosList{})
}
