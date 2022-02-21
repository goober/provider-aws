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

// DBClusterParameterGroupParameters defines the desired state of DBClusterParameterGroup
type DBClusterParameterGroupParameters struct {
	// Region is which region the DBClusterParameterGroup will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The DB cluster parameter group family name. A DB cluster parameter group
	// can be associated with one and only one DB cluster parameter group family,
	// and can be applied only to a DB cluster running a database engine and engine
	// version compatible with that DB cluster parameter group family.
	//
	// Aurora MySQL
	//
	// Example: aurora5.6, aurora-mysql5.7
	//
	// Aurora PostgreSQL
	//
	// Example: aurora-postgresql9.6
	//
	// To list all of the available parameter group families for a DB engine, use
	// the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine <engine>
	//
	// For example, to list all of the available parameter group families for the
	// Aurora PostgreSQL DB engine, use the following command:
	//
	// aws rds describe-db-engine-versions --query "DBEngineVersions[].DBParameterGroupFamily"
	// --engine aurora-postgresql
	//
	// The output contains duplicates.
	//
	// The following are the valid DB engine values:
	//
	//    * aurora (for MySQL 5.6-compatible Aurora)
	//
	//    * aurora-mysql (for MySQL 5.7-compatible Aurora)
	//
	//    * aurora-postgresql
	// +kubebuilder:validation:Required
	DBParameterGroupFamily *string `json:"dbParameterGroupFamily"`
	// The description for the DB cluster parameter group.
	// +kubebuilder:validation:Required
	Description *string `json:"description"`
	// Tags to assign to the DB cluster parameter group.
	Tags                                    []*Tag `json:"tags,omitempty"`
	CustomDBClusterParameterGroupParameters `json:",inline"`
}

// DBClusterParameterGroupSpec defines the desired state of DBClusterParameterGroup
type DBClusterParameterGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DBClusterParameterGroupParameters `json:"forProvider"`
}

// DBClusterParameterGroupObservation defines the observed state of DBClusterParameterGroup
type DBClusterParameterGroupObservation struct {
	// The Amazon Resource Name (ARN) for the DB cluster parameter group.
	DBClusterParameterGroupARN *string `json:"dbClusterParameterGroupARN,omitempty"`
	// The name of the DB cluster parameter group.
	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`
}

// DBClusterParameterGroupStatus defines the observed state of DBClusterParameterGroup.
type DBClusterParameterGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DBClusterParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DBClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBClusterParameterGroupSpec   `json:"spec"`
	Status            DBClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBClusterParameterGroupList contains a list of DBClusterParameterGroups
type DBClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBClusterParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	DBClusterParameterGroupKind             = "DBClusterParameterGroup"
	DBClusterParameterGroupGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBClusterParameterGroupKind}.String()
	DBClusterParameterGroupKindAPIVersion   = DBClusterParameterGroupKind + "." + GroupVersion.String()
	DBClusterParameterGroupGroupVersionKind = GroupVersion.WithKind(DBClusterParameterGroupKind)
)

func init() {
	SchemeBuilder.Register(&DBClusterParameterGroup{}, &DBClusterParameterGroupList{})
}
