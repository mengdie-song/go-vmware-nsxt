/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

import (
	"github.com/mengdie-song/go-vmware-nsxt/manager"
)

type ManagementNodeAggregateInfo struct {

	// Array of Node interface statistic properties
	NodeInterfaceProperties []manager.NodeInterfaceProperties `json:"node_interface_properties,omitempty"`

	// Array of Node network interface statistic properties
	NodeInterfaceStatistics []manager.NodeInterfaceStatisticsProperties `json:"node_interface_statistics,omitempty"`

	NodeStatus *ClusterNodeStatus `json:"node_status,omitempty"`

	// Time series of the node's system properties
	NodeStatusProperties []manager.NodeStatusProperties `json:"node_status_properties,omitempty"`

	RoleConfig *ManagementClusterRoleConfig `json:"role_config,omitempty"`

	TransportNodesConnected int64 `json:"transport_nodes_connected,omitempty"`
}
