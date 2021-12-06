// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime/schema")

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for AuthConfig
var AuthConfigGVK = schema.GroupVersionKind{
    Group: "enterprise.gloo.apis.solo.io",
    Version: "v1",
    Kind: "AuthConfig",
}

// AuthConfig is the Schema for the authConfig API
type AuthConfig struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec AuthConfigSpec `json:"spec,omitempty"`
    Status AuthConfigStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (AuthConfig)  GVK() schema.GroupVersionKind {
	return AuthConfigGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AuthConfigList contains a list of AuthConfig
type AuthConfigList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []AuthConfig `json:"items"`
}

func init() {
    SchemeBuilder.Register(&AuthConfig{}, &AuthConfigList{})
}
