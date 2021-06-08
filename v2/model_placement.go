/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// Placement Information about the placement of the VM.
type Placement struct {
	// The name of the Subregion.
	SubregionName *string `json:"SubregionName,omitempty"`
	// The tenancy of the VM (`default` \\| `dedicated`).
	Tenancy *string `json:"Tenancy,omitempty"`
}

// NewPlacement instantiates a new Placement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlacement() *Placement {
	this := Placement{}
	return &this
}

// NewPlacementWithDefaults instantiates a new Placement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlacementWithDefaults() *Placement {
	this := Placement{}
	return &this
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *Placement) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placement) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *Placement) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *Placement) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *Placement) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Placement) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *Placement) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *Placement) SetTenancy(v string) {
	o.Tenancy = &v
}

func (o Placement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	return json.Marshal(toSerialize)
}

type NullablePlacement struct {
	value *Placement
	isSet bool
}

func (v NullablePlacement) Get() *Placement {
	return v.value
}

func (v *NullablePlacement) Set(val *Placement) {
	v.value = val
	v.isSet = true
}

func (v NullablePlacement) IsSet() bool {
	return v.isSet
}

func (v *NullablePlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlacement(val *Placement) *NullablePlacement {
	return &NullablePlacement{value: val, isSet: true}
}

func (v NullablePlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
