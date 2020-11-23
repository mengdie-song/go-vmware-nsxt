/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package upgrade

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type CsmConfig struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Csm appliance status
	CsmApplianceStatus string `json:"csm_appliance_status,omitempty"`

	// Specifies whether this is a single/multi region deployment
	SingleRegion bool `json:"single_region,omitempty"`
}
