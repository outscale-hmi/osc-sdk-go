/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersClientGateway One or more filters.
type FiltersClientGateway struct {
	// The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
	BgpAsns *[]int32 `json:"BgpAsns,omitempty"`
	// The IDs of the client gateways.
	ClientGatewayIds *[]string `json:"ClientGatewayIds,omitempty"`
	// The types of communication tunnels used by the client gateways (only `ipsec.1` is supported).
	ConnectionTypes *[]string `json:"ConnectionTypes,omitempty"`
	// The public IPv4 addresses of the client gateways.
	PublicIps *[]string `json:"PublicIps,omitempty"`
	// The states of the client gateways (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the client gateways.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the client gateways.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the client gateways, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersClientGateway instantiates a new FiltersClientGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersClientGateway() *FiltersClientGateway {
	this := FiltersClientGateway{}
	return &this
}

// NewFiltersClientGatewayWithDefaults instantiates a new FiltersClientGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersClientGatewayWithDefaults() *FiltersClientGateway {
	this := FiltersClientGateway{}
	return &this
}

// GetBgpAsns returns the BgpAsns field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetBgpAsns() []int32 {
	if o == nil || o.BgpAsns == nil {
		var ret []int32
		return ret
	}
	return *o.BgpAsns
}

// GetBgpAsnsOk returns a tuple with the BgpAsns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetBgpAsnsOk() (*[]int32, bool) {
	if o == nil || o.BgpAsns == nil {
		return nil, false
	}
	return o.BgpAsns, true
}

// HasBgpAsns returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasBgpAsns() bool {
	if o != nil && o.BgpAsns != nil {
		return true
	}

	return false
}

// SetBgpAsns gets a reference to the given []int32 and assigns it to the BgpAsns field.
func (o *FiltersClientGateway) SetBgpAsns(v []int32) {
	o.BgpAsns = &v
}

// GetClientGatewayIds returns the ClientGatewayIds field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetClientGatewayIds() []string {
	if o == nil || o.ClientGatewayIds == nil {
		var ret []string
		return ret
	}
	return *o.ClientGatewayIds
}

// GetClientGatewayIdsOk returns a tuple with the ClientGatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetClientGatewayIdsOk() (*[]string, bool) {
	if o == nil || o.ClientGatewayIds == nil {
		return nil, false
	}
	return o.ClientGatewayIds, true
}

// HasClientGatewayIds returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasClientGatewayIds() bool {
	if o != nil && o.ClientGatewayIds != nil {
		return true
	}

	return false
}

// SetClientGatewayIds gets a reference to the given []string and assigns it to the ClientGatewayIds field.
func (o *FiltersClientGateway) SetClientGatewayIds(v []string) {
	o.ClientGatewayIds = &v
}

// GetConnectionTypes returns the ConnectionTypes field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetConnectionTypes() []string {
	if o == nil || o.ConnectionTypes == nil {
		var ret []string
		return ret
	}
	return *o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetConnectionTypesOk() (*[]string, bool) {
	if o == nil || o.ConnectionTypes == nil {
		return nil, false
	}
	return o.ConnectionTypes, true
}

// HasConnectionTypes returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasConnectionTypes() bool {
	if o != nil && o.ConnectionTypes != nil {
		return true
	}

	return false
}

// SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.
func (o *FiltersClientGateway) SetConnectionTypes(v []string) {
	o.ConnectionTypes = &v
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetPublicIps() []string {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetPublicIpsOk() (*[]string, bool) {
	if o == nil || o.PublicIps == nil {
		return nil, false
	}
	return o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.
func (o *FiltersClientGateway) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersClientGateway) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersClientGateway) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersClientGateway) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersClientGateway) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersClientGateway) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersClientGateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersClientGateway) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersClientGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BgpAsns != nil {
		toSerialize["BgpAsns"] = o.BgpAsns
	}
	if o.ClientGatewayIds != nil {
		toSerialize["ClientGatewayIds"] = o.ClientGatewayIds
	}
	if o.ConnectionTypes != nil {
		toSerialize["ConnectionTypes"] = o.ConnectionTypes
	}
	if o.PublicIps != nil {
		toSerialize["PublicIps"] = o.PublicIps
	}
	if o.States != nil {
		toSerialize["States"] = o.States
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersClientGateway struct {
	value *FiltersClientGateway
	isSet bool
}

func (v NullableFiltersClientGateway) Get() *FiltersClientGateway {
	return v.value
}

func (v *NullableFiltersClientGateway) Set(val *FiltersClientGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersClientGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersClientGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersClientGateway(val *FiltersClientGateway) *NullableFiltersClientGateway {
	return &NullableFiltersClientGateway{value: val, isSet: true}
}

func (v NullableFiltersClientGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersClientGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
