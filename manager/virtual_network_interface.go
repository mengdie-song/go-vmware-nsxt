/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type VirtualNetworkInterface struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`

	// Description of this resource
	Description string `json:"description,omitempty"`

	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`

	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`

	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`

	// Device key of the virtual network interface.
	DeviceKey string `json:"device_key"`

	// Device name of the virtual network interface.
	DeviceName string `json:"device_name,omitempty"`

	// External Id of the virtual network inferface.
	ExternalId string `json:"external_id"`

	// Id of the host on which the vm exists.
	HostId string `json:"host_id"`

	// IP Addresses of the the virtual network interface, from various sources.
	IpAddressInfo []IpAddressInfo `json:"ip_address_info,omitempty"`

	// LPort Attachment Id of the virtual network interface.
	LportAttachmentId string `json:"lport_attachment_id,omitempty"`

	// MAC address of the virtual network interface.
	MacAddress string `json:"mac_address"`

	// Id of the vm to which this virtual network interface belongs.
	OwnerVmId string `json:"owner_vm_id"`

	// Id of the vm unique within the host.
	VmLocalIdOnHost string `json:"vm_local_id_on_host"`
}
