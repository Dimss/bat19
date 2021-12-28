package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

)

type WhPlatformSpec struct {
	Foo string `json:"foo,omitempty"`
}

type WhPlatformStatus struct {
	Message string `json:"message,omitempty"`
}


type WhPlatform struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Spec   WhPlatformSpec   `json:"spec,omitempty"`
	Status WhPlatformStatus `json:"status,omitempty"`
}

type WhPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhPlatform `json:"items"`
}

func init(){
	scheme.Builder.Register(&WhPlatform{}, &WhPlatformList{})
	//runtime.SchemeBuilder.Register(&WhPlatform{}, &WhPlatformList{})
}