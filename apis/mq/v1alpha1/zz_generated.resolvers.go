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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta1 "github.com/crossplane/provider-aws/apis/ec2/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Broker.
func (mg *Broker) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDs),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDs")
	}
	mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomBrokerParameters.SubnetIDRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroups")
	}
	mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomBrokerParameters.SecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomUserParameters.BrokerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomUserParameters.BrokerIDRef,
		Selector:     mg.Spec.ForProvider.CustomUserParameters.BrokerIDSelector,
		To: reference.To{
			List:    &BrokerList{},
			Managed: &Broker{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomUserParameters.BrokerID")
	}
	mg.Spec.ForProvider.CustomUserParameters.BrokerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomUserParameters.BrokerIDRef = rsp.ResolvedReference

	return nil
}
