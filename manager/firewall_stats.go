/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type FirewallStats struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Aggregated number of bytes processed by the rule.
	ByteCount int64 `json:"byte_count,omitempty"`

	// Aggregated number of packets processed by the rule.
	PacketCount int64 `json:"packet_count,omitempty"`

	// Rule Identifier of the Firewall rule. This is a globally unique number.
	RuleId string `json:"rule_id,omitempty"`

	// Aggregated number of sessions processed by the rule
	SessionCount int64 `json:"session_count,omitempty"`
}
