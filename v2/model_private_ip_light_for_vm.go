/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.16
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// PrivateIpLightForVm Information about the private IP of the NIC.
type PrivateIpLightForVm struct {
	// If true, the IP address is the primary private IP address of the NIC.
	IsPrimary    *bool                   `json:"IsPrimary,omitempty"`
	LinkPublicIp *LinkPublicIpLightForVm `json:"LinkPublicIp,omitempty"`
	// The name of the private DNS.
	PrivateDnsName *string `json:"PrivateDnsName,omitempty"`
	// The private IP address.
	PrivateIp *string `json:"PrivateIp,omitempty"`
}

// NewPrivateIpLightForVm instantiates a new PrivateIpLightForVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateIpLightForVm() *PrivateIpLightForVm {
	this := PrivateIpLightForVm{}
	return &this
}

// NewPrivateIpLightForVmWithDefaults instantiates a new PrivateIpLightForVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateIpLightForVmWithDefaults() *PrivateIpLightForVm {
	this := PrivateIpLightForVm{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *PrivateIpLightForVm) GetIsPrimary() bool {
	if o == nil || o.IsPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLightForVm) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || o.IsPrimary == nil {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *PrivateIpLightForVm) HasIsPrimary() bool {
	if o != nil && o.IsPrimary != nil {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *PrivateIpLightForVm) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetLinkPublicIp returns the LinkPublicIp field value if set, zero value otherwise.
func (o *PrivateIpLightForVm) GetLinkPublicIp() LinkPublicIpLightForVm {
	if o == nil || o.LinkPublicIp == nil {
		var ret LinkPublicIpLightForVm
		return ret
	}
	return *o.LinkPublicIp
}

// GetLinkPublicIpOk returns a tuple with the LinkPublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLightForVm) GetLinkPublicIpOk() (*LinkPublicIpLightForVm, bool) {
	if o == nil || o.LinkPublicIp == nil {
		return nil, false
	}
	return o.LinkPublicIp, true
}

// HasLinkPublicIp returns a boolean if a field has been set.
func (o *PrivateIpLightForVm) HasLinkPublicIp() bool {
	if o != nil && o.LinkPublicIp != nil {
		return true
	}

	return false
}

// SetLinkPublicIp gets a reference to the given LinkPublicIpLightForVm and assigns it to the LinkPublicIp field.
func (o *PrivateIpLightForVm) SetLinkPublicIp(v LinkPublicIpLightForVm) {
	o.LinkPublicIp = &v
}

// GetPrivateDnsName returns the PrivateDnsName field value if set, zero value otherwise.
func (o *PrivateIpLightForVm) GetPrivateDnsName() string {
	if o == nil || o.PrivateDnsName == nil {
		var ret string
		return ret
	}
	return *o.PrivateDnsName
}

// GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLightForVm) GetPrivateDnsNameOk() (*string, bool) {
	if o == nil || o.PrivateDnsName == nil {
		return nil, false
	}
	return o.PrivateDnsName, true
}

// HasPrivateDnsName returns a boolean if a field has been set.
func (o *PrivateIpLightForVm) HasPrivateDnsName() bool {
	if o != nil && o.PrivateDnsName != nil {
		return true
	}

	return false
}

// SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.
func (o *PrivateIpLightForVm) SetPrivateDnsName(v string) {
	o.PrivateDnsName = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PrivateIpLightForVm) GetPrivateIp() string {
	if o == nil || o.PrivateIp == nil {
		var ret string
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateIpLightForVm) GetPrivateIpOk() (*string, bool) {
	if o == nil || o.PrivateIp == nil {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PrivateIpLightForVm) HasPrivateIp() bool {
	if o != nil && o.PrivateIp != nil {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.
func (o *PrivateIpLightForVm) SetPrivateIp(v string) {
	o.PrivateIp = &v
}

func (o PrivateIpLightForVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsPrimary != nil {
		toSerialize["IsPrimary"] = o.IsPrimary
	}
	if o.LinkPublicIp != nil {
		toSerialize["LinkPublicIp"] = o.LinkPublicIp
	}
	if o.PrivateDnsName != nil {
		toSerialize["PrivateDnsName"] = o.PrivateDnsName
	}
	if o.PrivateIp != nil {
		toSerialize["PrivateIp"] = o.PrivateIp
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateIpLightForVm struct {
	value *PrivateIpLightForVm
	isSet bool
}

func (v NullablePrivateIpLightForVm) Get() *PrivateIpLightForVm {
	return v.value
}

func (v *NullablePrivateIpLightForVm) Set(val *PrivateIpLightForVm) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateIpLightForVm) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateIpLightForVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateIpLightForVm(val *PrivateIpLightForVm) *NullablePrivateIpLightForVm {
	return &NullablePrivateIpLightForVm{value: val, isSet: true}
}

func (v NullablePrivateIpLightForVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateIpLightForVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
