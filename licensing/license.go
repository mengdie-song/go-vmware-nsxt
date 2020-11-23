/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package licensing

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type License struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// license metric
	CapacityType string `json:"capacity_type,omitempty"`

	// license edition
	Description string `json:"description,omitempty"`

	// date that license expires
	Expiry int64 `json:"expiry,omitempty"`

	// semicolon delimited feature list
	Features string `json:"features,omitempty"`

	// true for evalution license
	IsEval bool `json:"is_eval,omitempty"`

	// whether the license has expired
	IsExpired bool `json:"is_expired,omitempty"`

	// multi-hypervisor support
	IsMh bool `json:"is_mh,omitempty"`

	// license key
	LicenseKey string `json:"license_key"`

	// product name
	ProductName string `json:"product_name,omitempty"`

	// product version
	ProductVersion string `json:"product_version,omitempty"`

	// license capacity; 0 for unlimited
	Quantity int64 `json:"quantity,omitempty"`
}
