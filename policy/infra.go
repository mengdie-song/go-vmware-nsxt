/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package policy

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type Infra struct {

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

	// Absolute path of this object
	Path string `json:"path,omitempty"`

	// Path relative from its parent
	RelativePath string `json:"relative_path,omitempty"`

	// This field is used while creating or updating the space.
	CommunicationProfiles []CommunicationProfile `json:"communication_profiles,omitempty"`

	// This field indicates the default connectivity policy for the infra or tenant space WHITELIST - Adds a default drop rule. Administrator can then use \"allow\" rules (aka whitelist) to allow traffic between groups BLACKLIST - Adds a default allow rule. Admin can then use \"drop\" rules (aka blacklist) to block traffic between groups NONE - No default rules are added. This is the default setting
	ConnectivityStrategy string `json:"connectivity_strategy,omitempty"`

	// This field is used while creating or updating the infra space.
	Domains []Domain `json:"domains,omitempty"`
}
