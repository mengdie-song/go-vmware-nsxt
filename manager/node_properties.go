/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type NodeProperties struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Node BIOS Unique Indentifier
	BiosUuid string `json:"bios_uuid,omitempty"`

	// NSX CLI inactivity timeout, set to 0 to configure no timeout
	CliTimeout int64 `json:"cli_timeout,omitempty"`

	// Host name or fully qualified domain name of node
	Hostname string `json:"hostname,omitempty"`

	// Kernel version
	KernelVersion string `json:"kernel_version,omitempty"`

	// Node Unique Identifier
	NodeUuid string `json:"node_uuid,omitempty"`

	// Node version
	NodeVersion string `json:"node_version,omitempty"`

	// Current time expressed in milliseconds since epoch
	SystemTime int64 `json:"system_time,omitempty"`

	// Timezone
	Timezone string `json:"timezone,omitempty"`
}
