/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /> The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br /> You can learn more about errors returned by the API in the dedicated [errors page](api/errors).<br /><br /> Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but there are [differences in resource names](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html) between AWS and the OUTSCALE API.<br /> You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.21
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// VirtualGateway Information about the virtual gateway.
type VirtualGateway struct {
	// The type of VPN connection supported by the virtual gateway (only `ipsec.1` is supported).
	ConnectionType *string `json:"ConnectionType,omitempty"`
	// The Net to which the virtual gateway is attached.
	NetToVirtualGatewayLinks *[]NetToVirtualGatewayLink `json:"NetToVirtualGatewayLinks,omitempty"`
	// The state of the virtual gateway (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the virtual gateway.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
	// The ID of the virtual gateway.
	VirtualGatewayId *string `json:"VirtualGatewayId,omitempty"`
}

// NewVirtualGateway instantiates a new VirtualGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualGateway() *VirtualGateway {
	this := VirtualGateway{}
	return &this
}

// NewVirtualGatewayWithDefaults instantiates a new VirtualGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualGatewayWithDefaults() *VirtualGateway {
	this := VirtualGateway{}
	return &this
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *VirtualGateway) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualGateway) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *VirtualGateway) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *VirtualGateway) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetNetToVirtualGatewayLinks returns the NetToVirtualGatewayLinks field value if set, zero value otherwise.
func (o *VirtualGateway) GetNetToVirtualGatewayLinks() []NetToVirtualGatewayLink {
	if o == nil || o.NetToVirtualGatewayLinks == nil {
		var ret []NetToVirtualGatewayLink
		return ret
	}
	return *o.NetToVirtualGatewayLinks
}

// GetNetToVirtualGatewayLinksOk returns a tuple with the NetToVirtualGatewayLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualGateway) GetNetToVirtualGatewayLinksOk() (*[]NetToVirtualGatewayLink, bool) {
	if o == nil || o.NetToVirtualGatewayLinks == nil {
		return nil, false
	}
	return o.NetToVirtualGatewayLinks, true
}

// HasNetToVirtualGatewayLinks returns a boolean if a field has been set.
func (o *VirtualGateway) HasNetToVirtualGatewayLinks() bool {
	if o != nil && o.NetToVirtualGatewayLinks != nil {
		return true
	}

	return false
}

// SetNetToVirtualGatewayLinks gets a reference to the given []NetToVirtualGatewayLink and assigns it to the NetToVirtualGatewayLinks field.
func (o *VirtualGateway) SetNetToVirtualGatewayLinks(v []NetToVirtualGatewayLink) {
	o.NetToVirtualGatewayLinks = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *VirtualGateway) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualGateway) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *VirtualGateway) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *VirtualGateway) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualGateway) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualGateway) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualGateway) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *VirtualGateway) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value if set, zero value otherwise.
func (o *VirtualGateway) GetVirtualGatewayId() string {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret
	}
	return *o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualGateway) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil || o.VirtualGatewayId == nil {
		return nil, false
	}
	return o.VirtualGatewayId, true
}

// HasVirtualGatewayId returns a boolean if a field has been set.
func (o *VirtualGateway) HasVirtualGatewayId() bool {
	if o != nil && o.VirtualGatewayId != nil {
		return true
	}

	return false
}

// SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.
func (o *VirtualGateway) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = &v
}

func (o VirtualGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionType != nil {
		toSerialize["ConnectionType"] = o.ConnectionType
	}
	if o.NetToVirtualGatewayLinks != nil {
		toSerialize["NetToVirtualGatewayLinks"] = o.NetToVirtualGatewayLinks
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.VirtualGatewayId != nil {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualGateway struct {
	value *VirtualGateway
	isSet bool
}

func (v NullableVirtualGateway) Get() *VirtualGateway {
	return v.value
}

func (v *NullableVirtualGateway) Set(val *VirtualGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualGateway(val *VirtualGateway) *NullableVirtualGateway {
	return &NullableVirtualGateway{value: val, isSet: true}
}

func (v NullableVirtualGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
