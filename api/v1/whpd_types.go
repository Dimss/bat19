
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


type WhPlatformSpec struct {
	Foo string `json:"foo,omitempty"`
}


type WhPlatformStatus struct {
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Foo",type=string,JSONPath=`.spec.foo`

type WhPlatform struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WhPlatformSpec   `json:"spec,omitempty"`
	Status WhPlatformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type WhPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhPlatform `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WhPlatform{}, &WhPlatformList{})
}
