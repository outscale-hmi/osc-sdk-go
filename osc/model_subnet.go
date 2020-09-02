/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.3
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// Subnet Information about the Subnet.
type Subnet struct {
	// The number of available IP addresses in the Subnets.
	AvailableIpsCount int32 `json:"AvailableIpsCount,omitempty"`
	// The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange,omitempty"`
	// If `true`, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.
	MapPublicIpOnLaunch bool `json:"MapPublicIpOnLaunch,omitempty"`
	// The ID of the Net in which the Subnet is.
	NetId string `json:"NetId,omitempty"`
	// The state of the Subnet (`pending` \\| `available`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId,omitempty"`
	// The name of the Subregion in which the Subnet is located.
	SubregionName string `json:"SubregionName,omitempty"`
	// One or more tags associated with the Subnet.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
