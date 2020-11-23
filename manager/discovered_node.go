/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type DiscoveredNode struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Display name of discovered node
	DisplayName string `json:"display_name,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// Certificate of the discovered node
	Certificate string `json:"certificate,omitempty"`

	// Local Id of the discovered node in the Compute Manager
	CmLocalId string `json:"cm_local_id,omitempty"`

	// External id of the discovered node, ex. a mo-ref from VC
	ExternalId string `json:"external_id,omitempty"`

	// IP Addresses of the the discovered node.
	IpAddresses []string `json:"ip_addresses,omitempty"`

	// Discovered Node type like Host
	NodeType string `json:"node_type,omitempty"`

	// Id of the compute manager from where this node was discovered
	OriginId string `json:"origin_id,omitempty"`

	// Key-Value map of additional specific properties of discovered node in the Compute Manager
	OriginProperties []common.KeyValuePair `json:"origin_properties,omitempty"`

	// OS type of the discovered node
	OsType string `json:"os_type,omitempty"`

	// OS version of the discovered node
	OsVersion string `json:"os_version,omitempty"`

	// External id of the compute collection to which this node belongs
	ParentComputeCollection string `json:"parent_compute_collection,omitempty"`
}
