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

package targetgroup

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/elbv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/elbv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeTargetGroupsInput returns input for read
// operation.
func GenerateDescribeTargetGroupsInput(cr *svcapitypes.TargetGroup) *svcsdk.DescribeTargetGroupsInput {
	res := &svcsdk.DescribeTargetGroupsInput{}

	if cr.Spec.ForProvider.Name != nil {
		f2 := []*string{}
		f2 = append(f2, cr.Spec.ForProvider.Name)
		res.SetNames(f2)
	}

	return res
}

// GenerateTargetGroup returns the current state in the form of *svcapitypes.TargetGroup.
func GenerateTargetGroup(resp *svcsdk.DescribeTargetGroupsOutput) *svcapitypes.TargetGroup {
	cr := &svcapitypes.TargetGroup{}

	found := false
	for _, elem := range resp.TargetGroups {
		if elem.HealthCheckEnabled != nil {
			cr.Spec.ForProvider.HealthCheckEnabled = elem.HealthCheckEnabled
		} else {
			cr.Spec.ForProvider.HealthCheckEnabled = nil
		}
		if elem.HealthCheckIntervalSeconds != nil {
			cr.Spec.ForProvider.HealthCheckIntervalSeconds = elem.HealthCheckIntervalSeconds
		} else {
			cr.Spec.ForProvider.HealthCheckIntervalSeconds = nil
		}
		if elem.HealthCheckPath != nil {
			cr.Spec.ForProvider.HealthCheckPath = elem.HealthCheckPath
		} else {
			cr.Spec.ForProvider.HealthCheckPath = nil
		}
		if elem.HealthCheckPort != nil {
			cr.Spec.ForProvider.HealthCheckPort = elem.HealthCheckPort
		} else {
			cr.Spec.ForProvider.HealthCheckPort = nil
		}
		if elem.HealthCheckProtocol != nil {
			cr.Spec.ForProvider.HealthCheckProtocol = elem.HealthCheckProtocol
		} else {
			cr.Spec.ForProvider.HealthCheckProtocol = nil
		}
		if elem.HealthCheckTimeoutSeconds != nil {
			cr.Spec.ForProvider.HealthCheckTimeoutSeconds = elem.HealthCheckTimeoutSeconds
		} else {
			cr.Spec.ForProvider.HealthCheckTimeoutSeconds = nil
		}
		if elem.HealthyThresholdCount != nil {
			cr.Spec.ForProvider.HealthyThresholdCount = elem.HealthyThresholdCount
		} else {
			cr.Spec.ForProvider.HealthyThresholdCount = nil
		}
		if elem.IpAddressType != nil {
			cr.Spec.ForProvider.IPAddressType = elem.IpAddressType
		} else {
			cr.Spec.ForProvider.IPAddressType = nil
		}
		if elem.Matcher != nil {
			f9 := &svcapitypes.Matcher{}
			if elem.Matcher.GrpcCode != nil {
				f9.GrpcCode = elem.Matcher.GrpcCode
			}
			if elem.Matcher.HttpCode != nil {
				f9.HTTPCode = elem.Matcher.HttpCode
			}
			cr.Spec.ForProvider.Matcher = f9
		} else {
			cr.Spec.ForProvider.Matcher = nil
		}
		if elem.Port != nil {
			cr.Spec.ForProvider.Port = elem.Port
		} else {
			cr.Spec.ForProvider.Port = nil
		}
		if elem.Protocol != nil {
			cr.Spec.ForProvider.Protocol = elem.Protocol
		} else {
			cr.Spec.ForProvider.Protocol = nil
		}
		if elem.ProtocolVersion != nil {
			cr.Spec.ForProvider.ProtocolVersion = elem.ProtocolVersion
		} else {
			cr.Spec.ForProvider.ProtocolVersion = nil
		}
		if elem.TargetType != nil {
			cr.Spec.ForProvider.TargetType = elem.TargetType
		} else {
			cr.Spec.ForProvider.TargetType = nil
		}
		if elem.UnhealthyThresholdCount != nil {
			cr.Spec.ForProvider.UnhealthyThresholdCount = elem.UnhealthyThresholdCount
		} else {
			cr.Spec.ForProvider.UnhealthyThresholdCount = nil
		}
		if elem.VpcId != nil {
			cr.Spec.ForProvider.VPCID = elem.VpcId
		} else {
			cr.Spec.ForProvider.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateTargetGroupInput returns a create input.
func GenerateCreateTargetGroupInput(cr *svcapitypes.TargetGroup) *svcsdk.CreateTargetGroupInput {
	res := &svcsdk.CreateTargetGroupInput{}

	if cr.Spec.ForProvider.HealthCheckEnabled != nil {
		res.SetHealthCheckEnabled(*cr.Spec.ForProvider.HealthCheckEnabled)
	}
	if cr.Spec.ForProvider.HealthCheckIntervalSeconds != nil {
		res.SetHealthCheckIntervalSeconds(*cr.Spec.ForProvider.HealthCheckIntervalSeconds)
	}
	if cr.Spec.ForProvider.HealthCheckPath != nil {
		res.SetHealthCheckPath(*cr.Spec.ForProvider.HealthCheckPath)
	}
	if cr.Spec.ForProvider.HealthCheckPort != nil {
		res.SetHealthCheckPort(*cr.Spec.ForProvider.HealthCheckPort)
	}
	if cr.Spec.ForProvider.HealthCheckProtocol != nil {
		res.SetHealthCheckProtocol(*cr.Spec.ForProvider.HealthCheckProtocol)
	}
	if cr.Spec.ForProvider.HealthCheckTimeoutSeconds != nil {
		res.SetHealthCheckTimeoutSeconds(*cr.Spec.ForProvider.HealthCheckTimeoutSeconds)
	}
	if cr.Spec.ForProvider.HealthyThresholdCount != nil {
		res.SetHealthyThresholdCount(*cr.Spec.ForProvider.HealthyThresholdCount)
	}
	if cr.Spec.ForProvider.IPAddressType != nil {
		res.SetIpAddressType(*cr.Spec.ForProvider.IPAddressType)
	}
	if cr.Spec.ForProvider.Matcher != nil {
		f8 := &svcsdk.Matcher{}
		if cr.Spec.ForProvider.Matcher.GrpcCode != nil {
			f8.SetGrpcCode(*cr.Spec.ForProvider.Matcher.GrpcCode)
		}
		if cr.Spec.ForProvider.Matcher.HTTPCode != nil {
			f8.SetHttpCode(*cr.Spec.ForProvider.Matcher.HTTPCode)
		}
		res.SetMatcher(f8)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.Protocol != nil {
		res.SetProtocol(*cr.Spec.ForProvider.Protocol)
	}
	if cr.Spec.ForProvider.ProtocolVersion != nil {
		res.SetProtocolVersion(*cr.Spec.ForProvider.ProtocolVersion)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f13 := []*svcsdk.Tag{}
		for _, f13iter := range cr.Spec.ForProvider.Tags {
			f13elem := &svcsdk.Tag{}
			if f13iter.Key != nil {
				f13elem.SetKey(*f13iter.Key)
			}
			if f13iter.Value != nil {
				f13elem.SetValue(*f13iter.Value)
			}
			f13 = append(f13, f13elem)
		}
		res.SetTags(f13)
	}
	if cr.Spec.ForProvider.TargetType != nil {
		res.SetTargetType(*cr.Spec.ForProvider.TargetType)
	}
	if cr.Spec.ForProvider.UnhealthyThresholdCount != nil {
		res.SetUnhealthyThresholdCount(*cr.Spec.ForProvider.UnhealthyThresholdCount)
	}
	if cr.Spec.ForProvider.VPCID != nil {
		res.SetVpcId(*cr.Spec.ForProvider.VPCID)
	}

	return res
}

// GenerateModifyTargetGroupInput returns an update input.
func GenerateModifyTargetGroupInput(cr *svcapitypes.TargetGroup) *svcsdk.ModifyTargetGroupInput {
	res := &svcsdk.ModifyTargetGroupInput{}

	if cr.Spec.ForProvider.HealthCheckEnabled != nil {
		res.SetHealthCheckEnabled(*cr.Spec.ForProvider.HealthCheckEnabled)
	}
	if cr.Spec.ForProvider.HealthCheckIntervalSeconds != nil {
		res.SetHealthCheckIntervalSeconds(*cr.Spec.ForProvider.HealthCheckIntervalSeconds)
	}
	if cr.Spec.ForProvider.HealthCheckPath != nil {
		res.SetHealthCheckPath(*cr.Spec.ForProvider.HealthCheckPath)
	}
	if cr.Spec.ForProvider.HealthCheckPort != nil {
		res.SetHealthCheckPort(*cr.Spec.ForProvider.HealthCheckPort)
	}
	if cr.Spec.ForProvider.HealthCheckProtocol != nil {
		res.SetHealthCheckProtocol(*cr.Spec.ForProvider.HealthCheckProtocol)
	}
	if cr.Spec.ForProvider.HealthCheckTimeoutSeconds != nil {
		res.SetHealthCheckTimeoutSeconds(*cr.Spec.ForProvider.HealthCheckTimeoutSeconds)
	}
	if cr.Spec.ForProvider.HealthyThresholdCount != nil {
		res.SetHealthyThresholdCount(*cr.Spec.ForProvider.HealthyThresholdCount)
	}
	if cr.Spec.ForProvider.Matcher != nil {
		f7 := &svcsdk.Matcher{}
		if cr.Spec.ForProvider.Matcher.GrpcCode != nil {
			f7.SetGrpcCode(*cr.Spec.ForProvider.Matcher.GrpcCode)
		}
		if cr.Spec.ForProvider.Matcher.HTTPCode != nil {
			f7.SetHttpCode(*cr.Spec.ForProvider.Matcher.HTTPCode)
		}
		res.SetMatcher(f7)
	}
	if cr.Spec.ForProvider.UnhealthyThresholdCount != nil {
		res.SetUnhealthyThresholdCount(*cr.Spec.ForProvider.UnhealthyThresholdCount)
	}

	return res
}

// GenerateDeleteTargetGroupInput returns a deletion input.
func GenerateDeleteTargetGroupInput(cr *svcapitypes.TargetGroup) *svcsdk.DeleteTargetGroupInput {
	res := &svcsdk.DeleteTargetGroupInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "TargetGroupNotFound"
}
