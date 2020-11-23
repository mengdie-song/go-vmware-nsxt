/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package association

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type AssociationListResult struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`

	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`

	SortAscending bool `json:"sort_ascending,omitempty"`

	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`

	// Paged Collection of ResourceReference
	Results []common.ResourceReference `json:"results"`
}
