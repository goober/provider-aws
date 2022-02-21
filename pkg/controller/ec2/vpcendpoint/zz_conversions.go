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

package vpcendpoint

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ec2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/ec2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeVpcEndpointsInput returns input for read
// operation.
func GenerateDescribeVpcEndpointsInput(cr *svcapitypes.VPCEndpoint) *svcsdk.DescribeVpcEndpointsInput {
	res := &svcsdk.DescribeVpcEndpointsInput{}

	if cr.Status.AtProvider.VPCEndpointID != nil {
		f4 := []*string{}
		f4 = append(f4, cr.Status.AtProvider.VPCEndpointID)
		res.SetVpcEndpointIds(f4)
	}

	return res
}

// GenerateVPCEndpoint returns the current state in the form of *svcapitypes.VPCEndpoint.
func GenerateVPCEndpoint(resp *svcsdk.DescribeVpcEndpointsOutput) *svcapitypes.VPCEndpoint {
	cr := &svcapitypes.VPCEndpoint{}

	found := false
	for _, elem := range resp.VpcEndpoints {
		if elem.CreationTimestamp != nil {
			cr.Status.AtProvider.CreationTimestamp = &metav1.Time{*elem.CreationTimestamp}
		} else {
			cr.Status.AtProvider.CreationTimestamp = nil
		}
		if elem.DnsEntries != nil {
			f1 := []*svcapitypes.DNSEntry{}
			for _, f1iter := range elem.DnsEntries {
				f1elem := &svcapitypes.DNSEntry{}
				if f1iter.DnsName != nil {
					f1elem.DNSName = f1iter.DnsName
				}
				if f1iter.HostedZoneId != nil {
					f1elem.HostedZoneID = f1iter.HostedZoneId
				}
				f1 = append(f1, f1elem)
			}
			cr.Status.AtProvider.DNSEntries = f1
		} else {
			cr.Status.AtProvider.DNSEntries = nil
		}
		if elem.Groups != nil {
			f2 := []*svcapitypes.SecurityGroupIdentifier{}
			for _, f2iter := range elem.Groups {
				f2elem := &svcapitypes.SecurityGroupIdentifier{}
				if f2iter.GroupId != nil {
					f2elem.GroupID = f2iter.GroupId
				}
				if f2iter.GroupName != nil {
					f2elem.GroupName = f2iter.GroupName
				}
				f2 = append(f2, f2elem)
			}
			cr.Status.AtProvider.Groups = f2
		} else {
			cr.Status.AtProvider.Groups = nil
		}
		if elem.LastError != nil {
			f3 := &svcapitypes.LastError{}
			if elem.LastError.Code != nil {
				f3.Code = elem.LastError.Code
			}
			if elem.LastError.Message != nil {
				f3.Message = elem.LastError.Message
			}
			cr.Status.AtProvider.LastError = f3
		} else {
			cr.Status.AtProvider.LastError = nil
		}
		if elem.NetworkInterfaceIds != nil {
			f4 := []*string{}
			for _, f4iter := range elem.NetworkInterfaceIds {
				var f4elem string
				f4elem = *f4iter
				f4 = append(f4, &f4elem)
			}
			cr.Status.AtProvider.NetworkInterfaceIDs = f4
		} else {
			cr.Status.AtProvider.NetworkInterfaceIDs = nil
		}
		if elem.OwnerId != nil {
			cr.Status.AtProvider.OwnerID = elem.OwnerId
		} else {
			cr.Status.AtProvider.OwnerID = nil
		}
		if elem.PolicyDocument != nil {
			cr.Spec.ForProvider.PolicyDocument = elem.PolicyDocument
		} else {
			cr.Spec.ForProvider.PolicyDocument = nil
		}
		if elem.PrivateDnsEnabled != nil {
			cr.Spec.ForProvider.PrivateDNSEnabled = elem.PrivateDnsEnabled
		} else {
			cr.Spec.ForProvider.PrivateDNSEnabled = nil
		}
		if elem.RequesterManaged != nil {
			cr.Status.AtProvider.RequesterManaged = elem.RequesterManaged
		} else {
			cr.Status.AtProvider.RequesterManaged = nil
		}
		if elem.RouteTableIds != nil {
			f9 := []*string{}
			for _, f9iter := range elem.RouteTableIds {
				var f9elem string
				f9elem = *f9iter
				f9 = append(f9, &f9elem)
			}
			cr.Status.AtProvider.RouteTableIDs = f9
		} else {
			cr.Status.AtProvider.RouteTableIDs = nil
		}
		if elem.ServiceName != nil {
			cr.Spec.ForProvider.ServiceName = elem.ServiceName
		} else {
			cr.Spec.ForProvider.ServiceName = nil
		}
		if elem.State != nil {
			cr.Status.AtProvider.State = elem.State
		} else {
			cr.Status.AtProvider.State = nil
		}
		if elem.SubnetIds != nil {
			f12 := []*string{}
			for _, f12iter := range elem.SubnetIds {
				var f12elem string
				f12elem = *f12iter
				f12 = append(f12, &f12elem)
			}
			cr.Status.AtProvider.SubnetIDs = f12
		} else {
			cr.Status.AtProvider.SubnetIDs = nil
		}
		if elem.Tags != nil {
			f13 := []*svcapitypes.Tag{}
			for _, f13iter := range elem.Tags {
				f13elem := &svcapitypes.Tag{}
				if f13iter.Key != nil {
					f13elem.Key = f13iter.Key
				}
				if f13iter.Value != nil {
					f13elem.Value = f13iter.Value
				}
				f13 = append(f13, f13elem)
			}
			cr.Status.AtProvider.Tags = f13
		} else {
			cr.Status.AtProvider.Tags = nil
		}
		if elem.VpcEndpointId != nil {
			cr.Status.AtProvider.VPCEndpointID = elem.VpcEndpointId
		} else {
			cr.Status.AtProvider.VPCEndpointID = nil
		}
		if elem.VpcEndpointType != nil {
			cr.Spec.ForProvider.VPCEndpointType = elem.VpcEndpointType
		} else {
			cr.Spec.ForProvider.VPCEndpointType = nil
		}
		if elem.VpcId != nil {
			cr.Status.AtProvider.VPCID = elem.VpcId
		} else {
			cr.Status.AtProvider.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateVpcEndpointInput returns a create input.
func GenerateCreateVpcEndpointInput(cr *svcapitypes.VPCEndpoint) *svcsdk.CreateVpcEndpointInput {
	res := &svcsdk.CreateVpcEndpointInput{}

	if cr.Spec.ForProvider.PolicyDocument != nil {
		res.SetPolicyDocument(*cr.Spec.ForProvider.PolicyDocument)
	}
	if cr.Spec.ForProvider.PrivateDNSEnabled != nil {
		res.SetPrivateDnsEnabled(*cr.Spec.ForProvider.PrivateDNSEnabled)
	}
	if cr.Spec.ForProvider.ServiceName != nil {
		res.SetServiceName(*cr.Spec.ForProvider.ServiceName)
	}
	if cr.Spec.ForProvider.TagSpecifications != nil {
		f3 := []*svcsdk.TagSpecification{}
		for _, f3iter := range cr.Spec.ForProvider.TagSpecifications {
			f3elem := &svcsdk.TagSpecification{}
			if f3iter.ResourceType != nil {
				f3elem.SetResourceType(*f3iter.ResourceType)
			}
			if f3iter.Tags != nil {
				f3elemf1 := []*svcsdk.Tag{}
				for _, f3elemf1iter := range f3iter.Tags {
					f3elemf1elem := &svcsdk.Tag{}
					if f3elemf1iter.Key != nil {
						f3elemf1elem.SetKey(*f3elemf1iter.Key)
					}
					if f3elemf1iter.Value != nil {
						f3elemf1elem.SetValue(*f3elemf1iter.Value)
					}
					f3elemf1 = append(f3elemf1, f3elemf1elem)
				}
				f3elem.SetTags(f3elemf1)
			}
			f3 = append(f3, f3elem)
		}
		res.SetTagSpecifications(f3)
	}
	if cr.Spec.ForProvider.VPCEndpointType != nil {
		res.SetVpcEndpointType(*cr.Spec.ForProvider.VPCEndpointType)
	}

	return res
}

// GenerateModifyVpcEndpointInput returns an update input.
func GenerateModifyVpcEndpointInput(cr *svcapitypes.VPCEndpoint) *svcsdk.ModifyVpcEndpointInput {
	res := &svcsdk.ModifyVpcEndpointInput{}

	if cr.Spec.ForProvider.PolicyDocument != nil {
		res.SetPolicyDocument(*cr.Spec.ForProvider.PolicyDocument)
	}
	if cr.Spec.ForProvider.PrivateDNSEnabled != nil {
		res.SetPrivateDnsEnabled(*cr.Spec.ForProvider.PrivateDNSEnabled)
	}
	if cr.Status.AtProvider.VPCEndpointID != nil {
		res.SetVpcEndpointId(*cr.Status.AtProvider.VPCEndpointID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
