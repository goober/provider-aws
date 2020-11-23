/*
Copyright 2020 The Crossplane Authors.

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

package backup

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"

	svcapitypes "github.com/crossplane/provider-aws/apis/dynamodb/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.

// GenerateDescribeBackupInput returns input for read
// operation.
func GenerateDescribeBackupInput(cr *svcapitypes.Backup) *svcsdk.DescribeBackupInput {
	res := preGenerateDescribeBackupInput(cr, &svcsdk.DescribeBackupInput{})

	if cr.Status.AtProvider.BackupARN != nil {
		res.SetBackupArn(*cr.Status.AtProvider.BackupARN)
	}

	return postGenerateDescribeBackupInput(cr, res)
}

// GenerateBackup returns the current state in the form of *svcapitypes.Backup.
func GenerateBackup(resp *svcsdk.DescribeBackupOutput) *svcapitypes.Backup {
	cr := &svcapitypes.Backup{}

	return cr
}

// GenerateCreateBackupInput returns a create input.
func GenerateCreateBackupInput(cr *svcapitypes.Backup) *svcsdk.CreateBackupInput {
	res := preGenerateCreateBackupInput(cr, &svcsdk.CreateBackupInput{})

	if cr.Spec.ForProvider.BackupName != nil {
		res.SetBackupName(*cr.Spec.ForProvider.BackupName)
	}

	return postGenerateCreateBackupInput(cr, res)
}

// GenerateDeleteBackupInput returns a deletion input.
func GenerateDeleteBackupInput(cr *svcapitypes.Backup) *svcsdk.DeleteBackupInput {
	res := preGenerateDeleteBackupInput(cr, &svcsdk.DeleteBackupInput{})

	if cr.Status.AtProvider.BackupARN != nil {
		res.SetBackupArn(*cr.Status.AtProvider.BackupARN)
	}

	return postGenerateDeleteBackupInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "BackupNotFoundException"
}
