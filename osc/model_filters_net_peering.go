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
// FiltersNetPeering One or more filters.
type FiltersNetPeering struct {
	// The account IDs of the owners of the peer Nets.
	AccepterNetAccountIds []string `json:"AccepterNetAccountIds,omitempty"`
	// The IP ranges of the peer Nets, in CIDR notation (for example, 10.0.0.0/24).
	AccepterNetIpRanges []string `json:"AccepterNetIpRanges,omitempty"`
	// The IDs of the peer Nets.
	AccepterNetNetIds []string `json:"AccepterNetNetIds,omitempty"`
	// The IDs of the Net peering connections.
	NetPeeringIds []string `json:"NetPeeringIds,omitempty"`
	// The account IDs of the owners of the peer Nets.
	SourceNetAccountIds []string `json:"SourceNetAccountIds,omitempty"`
	// The IP ranges of the peer Nets.
	SourceNetIpRanges []string `json:"SourceNetIpRanges,omitempty"`
	// The IDs of the peer Nets.
	SourceNetNetIds []string `json:"SourceNetNetIds,omitempty"`
	// Additional information about the states of the Net peering connections.
	StateMessages []string `json:"StateMessages,omitempty"`
	// The states of the Net peering connections (`pending-acceptance` \\| `active` \\| `rejected` \\| `failed` \\| `expired` \\| `deleted`).
	StateNames []string `json:"StateNames,omitempty"`
	// The keys of the tags associated with the Net peering connections.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Net peering connections.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Net peering connections, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
}
