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

// CrawlerParameters defines the desired state of Crawler
type CrawlerParameters struct {
	// Region is which region the Crawler will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A list of custom classifiers that the user has registered. By default, all
	// built-in classifiers are included in a crawl, but these custom classifiers
	// always override the default classifiers for a given classification.
	Classifiers []*string `json:"classifiers,omitempty"`
	// Crawler configuration information. This versioned JSON string allows users
	// to specify aspects of a crawler's behavior. For more information, see Configuring
	// a Crawler (https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html).
	Configuration *string `json:"configuration,omitempty"`
	// The name of the SecurityConfiguration structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `json:"crawlerSecurityConfiguration,omitempty"`
	// The Glue database where results are written, such as: arn:aws:daylight:us-east-1::database/sometable/*.
	DatabaseName *string `json:"databaseName,omitempty"`
	// A description of the new crawler.
	Description *string `json:"description,omitempty"`
	// Specifies data lineage configuration settings for the crawler.
	LineageConfiguration *LineageConfiguration `json:"lineageConfiguration,omitempty"`
	// A policy that specifies whether to crawl the entire dataset again, or to
	// crawl only folders that were added since the last crawler run.
	RecrawlPolicy *RecrawlPolicy `json:"recrawlPolicy,omitempty"`
	// A cron expression used to specify the schedule (see Time-Based Schedules
	// for Jobs and Crawlers (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify:
	// cron(15 12 * * ? *).
	Schedule *string `json:"schedule,omitempty"`
	// The policy for the crawler's update and deletion behavior.
	SchemaChangePolicy *SchemaChangePolicy `json:"schemaChangePolicy,omitempty"`
	// The table prefix used for catalog tables that are created.
	TablePrefix *string `json:"tablePrefix,omitempty"`
	// The tags to use with this crawler request. You may use tags to limit access
	// to the crawler. For more information about tags in Glue, see Amazon Web Services
	// Tags in Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html)
	// in the developer guide.
	Tags map[string]*string `json:"tags,omitempty"`
	// A list of collection of targets to crawl.
	// +kubebuilder:validation:Required
	Targets                 *CrawlerTargets `json:"targets"`
	CustomCrawlerParameters `json:",inline"`
}

// CrawlerSpec defines the desired state of Crawler
type CrawlerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CrawlerParameters `json:"forProvider"`
}

// CrawlerObservation defines the observed state of Crawler
type CrawlerObservation struct {
}

// CrawlerStatus defines the observed state of Crawler.
type CrawlerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CrawlerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Crawler is the Schema for the Crawlers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Crawler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CrawlerSpec   `json:"spec"`
	Status            CrawlerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CrawlerList contains a list of Crawlers
type CrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Crawler `json:"items"`
}

// Repository type metadata.
var (
	CrawlerKind             = "Crawler"
	CrawlerGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CrawlerKind}.String()
	CrawlerKindAPIVersion   = CrawlerKind + "." + GroupVersion.String()
	CrawlerGroupVersionKind = GroupVersion.WithKind(CrawlerKind)
)

func init() {
	SchemeBuilder.Register(&Crawler{}, &CrawlerList{})
}
