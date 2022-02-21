//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Addon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonHealth) DeepCopyInto(out *AddonHealth) {
	*out = *in
	if in.Issues != nil {
		in, out := &in.Issues, &out.Issues
		*out = make([]*AddonIssue, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AddonIssue)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonHealth.
func (in *AddonHealth) DeepCopy() *AddonHealth {
	if in == nil {
		return nil
	}
	out := new(AddonHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInfo) DeepCopyInto(out *AddonInfo) {
	*out = *in
	if in.AddonName != nil {
		in, out := &in.AddonName, &out.AddonName
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInfo.
func (in *AddonInfo) DeepCopy() *AddonInfo {
	if in == nil {
		return nil
	}
	out := new(AddonInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonIssue) DeepCopyInto(out *AddonIssue) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDs != nil {
		in, out := &in.ResourceIDs, &out.ResourceIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonIssue.
func (in *AddonIssue) DeepCopy() *AddonIssue {
	if in == nil {
		return nil
	}
	out := new(AddonIssue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonList) DeepCopyInto(out *AddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonList.
func (in *AddonList) DeepCopy() *AddonList {
	if in == nil {
		return nil
	}
	out := new(AddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonObservation) DeepCopyInto(out *AddonObservation) {
	*out = *in
	if in.AddonARN != nil {
		in, out := &in.AddonARN, &out.AddonARN
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.Health != nil {
		in, out := &in.Health, &out.Health
		*out = new(AddonHealth)
		(*in).DeepCopyInto(*out)
	}
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		*out = (*in).DeepCopy()
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonObservation.
func (in *AddonObservation) DeepCopy() *AddonObservation {
	if in == nil {
		return nil
	}
	out := new(AddonObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonParameters) DeepCopyInto(out *AddonParameters) {
	*out = *in
	if in.AddonName != nil {
		in, out := &in.AddonName, &out.AddonName
		*out = new(string)
		**out = **in
	}
	if in.AddonVersion != nil {
		in, out := &in.AddonVersion, &out.AddonVersion
		*out = new(string)
		**out = **in
	}
	if in.ResolveConflicts != nil {
		in, out := &in.ResolveConflicts, &out.ResolveConflicts
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountRoleARN != nil {
		in, out := &in.ServiceAccountRoleARN, &out.ServiceAccountRoleARN
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	in.CustomAddonParameters.DeepCopyInto(&out.CustomAddonParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonParameters.
func (in *AddonParameters) DeepCopy() *AddonParameters {
	if in == nil {
		return nil
	}
	out := new(AddonParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonStatus) DeepCopyInto(out *AddonStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonStatus.
func (in *AddonStatus) DeepCopy() *AddonStatus {
	if in == nil {
		return nil
	}
	out := new(AddonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonVersionInfo) DeepCopyInto(out *AddonVersionInfo) {
	*out = *in
	if in.AddonVersion != nil {
		in, out := &in.AddonVersion, &out.AddonVersion
		*out = new(string)
		**out = **in
	}
	if in.Architecture != nil {
		in, out := &in.Architecture, &out.Architecture
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonVersionInfo.
func (in *AddonVersionInfo) DeepCopy() *AddonVersionInfo {
	if in == nil {
		return nil
	}
	out := new(AddonVersionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon_SDK) DeepCopyInto(out *Addon_SDK) {
	*out = *in
	if in.AddonARN != nil {
		in, out := &in.AddonARN, &out.AddonARN
		*out = new(string)
		**out = **in
	}
	if in.AddonName != nil {
		in, out := &in.AddonName, &out.AddonName
		*out = new(string)
		**out = **in
	}
	if in.AddonVersion != nil {
		in, out := &in.AddonVersion, &out.AddonVersion
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.Health != nil {
		in, out := &in.Health, &out.Health
		*out = new(AddonHealth)
		(*in).DeepCopyInto(*out)
	}
	if in.ModifiedAt != nil {
		in, out := &in.ModifiedAt, &out.ModifiedAt
		*out = (*in).DeepCopy()
	}
	if in.ServiceAccountRoleARN != nil {
		in, out := &in.ServiceAccountRoleARN, &out.ServiceAccountRoleARN
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon_SDK.
func (in *Addon_SDK) DeepCopy() *Addon_SDK {
	if in == nil {
		return nil
	}
	out := new(Addon_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingGroup) DeepCopyInto(out *AutoScalingGroup) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingGroup.
func (in *AutoScalingGroup) DeepCopy() *AutoScalingGroup {
	if in == nil {
		return nil
	}
	out := new(AutoScalingGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Compatibility) DeepCopyInto(out *Compatibility) {
	*out = *in
	if in.ClusterVersion != nil {
		in, out := &in.ClusterVersion, &out.ClusterVersion
		*out = new(string)
		**out = **in
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.PlatformVersions != nil {
		in, out := &in.PlatformVersions, &out.PlatformVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Compatibility.
func (in *Compatibility) DeepCopy() *Compatibility {
	if in == nil {
		return nil
	}
	out := new(Compatibility)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorConfigRequest) DeepCopyInto(out *ConnectorConfigRequest) {
	*out = *in
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorConfigRequest.
func (in *ConnectorConfigRequest) DeepCopy() *ConnectorConfigRequest {
	if in == nil {
		return nil
	}
	out := new(ConnectorConfigRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorConfigResponse) DeepCopyInto(out *ConnectorConfigResponse) {
	*out = *in
	if in.ActivationCode != nil {
		in, out := &in.ActivationCode, &out.ActivationCode
		*out = new(string)
		**out = **in
	}
	if in.ActivationExpiry != nil {
		in, out := &in.ActivationExpiry, &out.ActivationExpiry
		*out = (*in).DeepCopy()
	}
	if in.ActivationID != nil {
		in, out := &in.ActivationID, &out.ActivationID
		*out = new(string)
		**out = **in
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(string)
		**out = **in
	}
	if in.RoleARN != nil {
		in, out := &in.RoleARN, &out.RoleARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorConfigResponse.
func (in *ConnectorConfigResponse) DeepCopy() *ConnectorConfigResponse {
	if in == nil {
		return nil
	}
	out := new(ConnectorConfigResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAddonParameters) DeepCopyInto(out *CustomAddonParameters) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.ClusterNameRef != nil {
		in, out := &in.ClusterNameRef, &out.ClusterNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ClusterNameSelector != nil {
		in, out := &in.ClusterNameSelector, &out.ClusterNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAddonParameters.
func (in *CustomAddonParameters) DeepCopy() *CustomAddonParameters {
	if in == nil {
		return nil
	}
	out := new(CustomAddonParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfig) DeepCopyInto(out *EncryptionConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfig.
func (in *EncryptionConfig) DeepCopy() *EncryptionConfig {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorDetail) DeepCopyInto(out *ErrorDetail) {
	*out = *in
	if in.ErrorCode != nil {
		in, out := &in.ErrorCode, &out.ErrorCode
		*out = new(string)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDs != nil {
		in, out := &in.ResourceIDs, &out.ResourceIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorDetail.
func (in *ErrorDetail) DeepCopy() *ErrorDetail {
	if in == nil {
		return nil
	}
	out := new(ErrorDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FargateProfileSelector) DeepCopyInto(out *FargateProfileSelector) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FargateProfileSelector.
func (in *FargateProfileSelector) DeepCopy() *FargateProfileSelector {
	if in == nil {
		return nil
	}
	out := new(FargateProfileSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issue) DeepCopyInto(out *Issue) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.ResourceIDs != nil {
		in, out := &in.ResourceIDs, &out.ResourceIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issue.
func (in *Issue) DeepCopy() *Issue {
	if in == nil {
		return nil
	}
	out := new(Issue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNetworkConfigRequest) DeepCopyInto(out *KubernetesNetworkConfigRequest) {
	*out = *in
	if in.ServiceIPv4CIDR != nil {
		in, out := &in.ServiceIPv4CIDR, &out.ServiceIPv4CIDR
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNetworkConfigRequest.
func (in *KubernetesNetworkConfigRequest) DeepCopy() *KubernetesNetworkConfigRequest {
	if in == nil {
		return nil
	}
	out := new(KubernetesNetworkConfigRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesNetworkConfigResponse) DeepCopyInto(out *KubernetesNetworkConfigResponse) {
	*out = *in
	if in.ServiceIPv4CIDR != nil {
		in, out := &in.ServiceIPv4CIDR, &out.ServiceIPv4CIDR
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesNetworkConfigResponse.
func (in *KubernetesNetworkConfigResponse) DeepCopy() *KubernetesNetworkConfigResponse {
	if in == nil {
		return nil
	}
	out := new(KubernetesNetworkConfigResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateSpecification) DeepCopyInto(out *LaunchTemplateSpecification) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateSpecification.
func (in *LaunchTemplateSpecification) DeepCopy() *LaunchTemplateSpecification {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodegroupResources) DeepCopyInto(out *NodegroupResources) {
	*out = *in
	if in.RemoteAccessSecurityGroup != nil {
		in, out := &in.RemoteAccessSecurityGroup, &out.RemoteAccessSecurityGroup
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodegroupResources.
func (in *NodegroupResources) DeepCopy() *NodegroupResources {
	if in == nil {
		return nil
	}
	out := new(NodegroupResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDC) DeepCopyInto(out *OIDC) {
	*out = *in
	if in.Issuer != nil {
		in, out := &in.Issuer, &out.Issuer
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDC.
func (in *OIDC) DeepCopy() *OIDC {
	if in == nil {
		return nil
	}
	out := new(OIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCIdentityProviderConfig) DeepCopyInto(out *OIDCIdentityProviderConfig) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.GroupsClaim != nil {
		in, out := &in.GroupsClaim, &out.GroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.GroupsPrefix != nil {
		in, out := &in.GroupsPrefix, &out.GroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderConfigARN != nil {
		in, out := &in.IdentityProviderConfigARN, &out.IdentityProviderConfigARN
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderConfigName != nil {
		in, out := &in.IdentityProviderConfigName, &out.IdentityProviderConfigName
		*out = new(string)
		**out = **in
	}
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.UsernameClaim != nil {
		in, out := &in.UsernameClaim, &out.UsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.UsernamePrefix != nil {
		in, out := &in.UsernamePrefix, &out.UsernamePrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCIdentityProviderConfig.
func (in *OIDCIdentityProviderConfig) DeepCopy() *OIDCIdentityProviderConfig {
	if in == nil {
		return nil
	}
	out := new(OIDCIdentityProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCIdentityProviderConfigRequest) DeepCopyInto(out *OIDCIdentityProviderConfigRequest) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.GroupsClaim != nil {
		in, out := &in.GroupsClaim, &out.GroupsClaim
		*out = new(string)
		**out = **in
	}
	if in.GroupsPrefix != nil {
		in, out := &in.GroupsPrefix, &out.GroupsPrefix
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderConfigName != nil {
		in, out := &in.IdentityProviderConfigName, &out.IdentityProviderConfigName
		*out = new(string)
		**out = **in
	}
	if in.IssuerURL != nil {
		in, out := &in.IssuerURL, &out.IssuerURL
		*out = new(string)
		**out = **in
	}
	if in.UsernameClaim != nil {
		in, out := &in.UsernameClaim, &out.UsernameClaim
		*out = new(string)
		**out = **in
	}
	if in.UsernamePrefix != nil {
		in, out := &in.UsernamePrefix, &out.UsernamePrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCIdentityProviderConfigRequest.
func (in *OIDCIdentityProviderConfigRequest) DeepCopy() *OIDCIdentityProviderConfigRequest {
	if in == nil {
		return nil
	}
	out := new(OIDCIdentityProviderConfigRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	if in.KeyARN != nil {
		in, out := &in.KeyARN, &out.KeyARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteAccessConfig) DeepCopyInto(out *RemoteAccessConfig) {
	*out = *in
	if in.EC2SshKey != nil {
		in, out := &in.EC2SshKey, &out.EC2SshKey
		*out = new(string)
		**out = **in
	}
	if in.SourceSecurityGroups != nil {
		in, out := &in.SourceSecurityGroups, &out.SourceSecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteAccessConfig.
func (in *RemoteAccessConfig) DeepCopy() *RemoteAccessConfig {
	if in == nil {
		return nil
	}
	out := new(RemoteAccessConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Update) DeepCopyInto(out *Update) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]*ErrorDetail, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ErrorDetail)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make([]*UpdateParam, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(UpdateParam)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Update.
func (in *Update) DeepCopy() *Update {
	if in == nil {
		return nil
	}
	out := new(Update)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdateParam) DeepCopyInto(out *UpdateParam) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdateParam.
func (in *UpdateParam) DeepCopy() *UpdateParam {
	if in == nil {
		return nil
	}
	out := new(UpdateParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigRequest) DeepCopyInto(out *VPCConfigRequest) {
	*out = *in
	if in.PublicAccessCIDRs != nil {
		in, out := &in.PublicAccessCIDRs, &out.PublicAccessCIDRs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigRequest.
func (in *VPCConfigRequest) DeepCopy() *VPCConfigRequest {
	if in == nil {
		return nil
	}
	out := new(VPCConfigRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfigResponse) DeepCopyInto(out *VPCConfigResponse) {
	*out = *in
	if in.ClusterSecurityGroupID != nil {
		in, out := &in.ClusterSecurityGroupID, &out.ClusterSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.EndpointPrivateAccess != nil {
		in, out := &in.EndpointPrivateAccess, &out.EndpointPrivateAccess
		*out = new(bool)
		**out = **in
	}
	if in.EndpointPublicAccess != nil {
		in, out := &in.EndpointPublicAccess, &out.EndpointPublicAccess
		*out = new(bool)
		**out = **in
	}
	if in.PublicAccessCIDRs != nil {
		in, out := &in.PublicAccessCIDRs, &out.PublicAccessCIDRs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfigResponse.
func (in *VPCConfigResponse) DeepCopy() *VPCConfigResponse {
	if in == nil {
		return nil
	}
	out := new(VPCConfigResponse)
	in.DeepCopyInto(out)
	return out
}
