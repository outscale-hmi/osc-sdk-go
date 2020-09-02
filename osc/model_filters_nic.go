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
// FiltersNic One or more filters.
type FiltersNic struct {
	// The device numbers the NICs are attached to.
	LinkNicSortNumbers []int32 `json:"LinkNicSortNumbers,omitempty"`
	// The IDs of the VMs the NICs are attached to.
	LinkNicVmIds []string `json:"LinkNicVmIds,omitempty"`
	// The IDs of the NICs.
	NicIds []string `json:"NicIds,omitempty"`
	// The private IP addresses of the NICs.
	PrivateIpsPrivateIps []string `json:"PrivateIpsPrivateIps,omitempty"`
	// The IDs of the Subnets for the NICs.
	SubnetIds []string `json:"SubnetIds,omitempty"`
}
