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

// FiltersSubregion One or more filters.
type FiltersSubregion struct {
	// The names of the Subregions.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
}

// NewFiltersSubregion instantiates a new FiltersSubregion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersSubregion() *FiltersSubregion {
	this := FiltersSubregion{}
	return &this
}

// NewFiltersSubregionWithDefaults instantiates a new FiltersSubregion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersSubregionWithDefaults() *FiltersSubregion {
	this := FiltersSubregion{}
	return &this
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *FiltersSubregion) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersSubregion) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *FiltersSubregion) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *FiltersSubregion) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

func (o FiltersSubregion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubregionNames != nil {
		toSerialize["SubregionNames"] = o.SubregionNames
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersSubregion struct {
	value *FiltersSubregion
	isSet bool
}

func (v NullableFiltersSubregion) Get() *FiltersSubregion {
	return v.value
}

func (v *NullableFiltersSubregion) Set(val *FiltersSubregion) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersSubregion) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersSubregion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersSubregion(val *FiltersSubregion) *NullableFiltersSubregion {
	return &NullableFiltersSubregion{value: val, isSet: true}
}

func (v NullableFiltersSubregion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersSubregion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
