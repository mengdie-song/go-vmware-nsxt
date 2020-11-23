/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type AwsAccount struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int64 `json:"_revision"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// Unique identifier of this resource
	Id string `json:"id,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// List of authorized users
	AuthUsers []CloudUserInfo `json:"auth_users,omitempty"`

	// Cloud Type
	CloudType string `json:"cloud_type"`

	// Instance statistics
	InstanceStats *InstanceStats `json:"instance_stats,omitempty"`

	// Tenant ID of the cloud account
	TenantId string `json:"tenant_id,omitempty"`

	// Access key of cloud account
	AccessKey string `json:"access_key,omitempty"`

	// Is the AWS authorization mechanism based on Identity and Access Management(IAM) service?
	AuthMechanismIam bool `json:"auth_mechanism_iam,omitempty"`

	// External id for the IAM role csm needs to assume
	ExternalId string `json:"external_id,omitempty"`

	// Service Role Name for IAM role csm needs to assume
	GatewayRoleName string `json:"gateway_role_name,omitempty"`

	// Has a managed VPC?
	HasManagedVpc string `json:"has_managed_vpc,omitempty"`

	// Amazon Resource Names (ARNs) uniquely identify AWS resources. We will use it here to identify the IAM role csm needs to assume.
	IamRoleArn string `json:"iam_role_arn,omitempty"`

	// Count of the regions available
	RegionsCount int64 `json:"regions_count,omitempty"`

	// Secret key of cloud account
	SecretKey string `json:"secret_key,omitempty"`

	// Status of the account
	Status *AwsAccountStatus `json:"status,omitempty"`

	// VPC statistics
	VpcStats *VpcStats `json:"vpc_stats,omitempty"`
}
