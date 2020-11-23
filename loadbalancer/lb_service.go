/*
 * NSX API
 *
 * VMware NSX REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

import (
	"github.com/mengdie-song/go-vmware-nsxt/common"
)

type LbService struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	// Schema for this resource
	Schema string `json:"_schema,omitempty"`

	// Link to this resource
	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision"`

	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`

	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`

	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`

	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`

	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity.
	Protection string `json:"_protection,omitempty"`

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

	// whether access log is enabled
	AccessLogEnabled bool `json:"access_log_enabled"`

	// LBS could be instantiated (or created) on the Logical router, etc. Typically, it could be applied to Tier1 LogicalRouter. It can be attached to Tier0 LogicalRouter either in non-multi-tenant environments or to provide load balancing for infrastructure services offered by the provider.
	Attachment *common.ResourceReference `json:"attachment,omitempty"`

	// whether the load balancer service is enabled
	Enabled bool `json:"enabled"`

	// Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log.
	ErrorLogLevel string `json:"error_log_level,omitempty"`

	// the size of load balancer service
	Size string `json:"size,omitempty"`

	// virtual servers can be associated to LbService(which is similar to physical/virtual load balancer), Lb virtual servers, pools and other entities could be defined independently, the virtual server identifier list here would be used to maintain the relationship of LbService and other Lb entities.
	VirtualServerIds []string `json:"virtual_server_ids,omitempty"`
}
