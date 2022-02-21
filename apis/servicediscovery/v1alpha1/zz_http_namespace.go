/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// HTTPNamespaceParameters defines the desired state of HTTPNamespace
type HTTPNamespaceParameters struct {
	// Region is which region the HTTPNamespace will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A description for the namespace.
	Description *string `json:"description,omitempty"`
	// The name that you want to assign to this namespace.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value that you define. Tags keys can be up to 128 characters in length, and
	// tag values can be up to 256 characters in length.
	Tags                          []*Tag `json:"tags,omitempty"`
	CustomHTTPNamespaceParameters `json:",inline"`
}

// HTTPNamespaceSpec defines the desired state of HTTPNamespace
type HTTPNamespaceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       HTTPNamespaceParameters `json:"forProvider"`
}

// HTTPNamespaceObservation defines the observed state of HTTPNamespace
type HTTPNamespaceObservation struct {
	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationID *string `json:"operationID,omitempty"`
}

// HTTPNamespaceStatus defines the observed state of HTTPNamespace.
type HTTPNamespaceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          HTTPNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPNamespace is the Schema for the HTTPNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HTTPNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPNamespaceSpec   `json:"spec"`
	Status            HTTPNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPNamespaceList contains a list of HTTPNamespaces
type HTTPNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPNamespace `json:"items"`
}

// Repository type metadata.
var (
	HTTPNamespaceKind             = "HTTPNamespace"
	HTTPNamespaceGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPNamespaceKind}.String()
	HTTPNamespaceKindAPIVersion   = HTTPNamespaceKind + "." + GroupVersion.String()
	HTTPNamespaceGroupVersionKind = GroupVersion.WithKind(HTTPNamespaceKind)
)

func init() {
	SchemeBuilder.Register(&HTTPNamespace{}, &HTTPNamespaceList{})
}
