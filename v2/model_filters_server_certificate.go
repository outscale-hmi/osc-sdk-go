/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// FiltersServerCertificate One or more filters.
type FiltersServerCertificate struct {
	// The paths to the server certificates.
	Paths *[]string `json:"Paths,omitempty"`
}

// NewFiltersServerCertificate instantiates a new FiltersServerCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersServerCertificate() *FiltersServerCertificate {
	this := FiltersServerCertificate{}
	return &this
}

// NewFiltersServerCertificateWithDefaults instantiates a new FiltersServerCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersServerCertificateWithDefaults() *FiltersServerCertificate {
	this := FiltersServerCertificate{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *FiltersServerCertificate) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return *o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersServerCertificate) GetPathsOk() (*[]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *FiltersServerCertificate) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *FiltersServerCertificate) SetPaths(v []string) {
	o.Paths = &v
}

func (o FiltersServerCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paths != nil {
		toSerialize["Paths"] = o.Paths
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersServerCertificate struct {
	value *FiltersServerCertificate
	isSet bool
}

func (v NullableFiltersServerCertificate) Get() *FiltersServerCertificate {
	return v.value
}

func (v *NullableFiltersServerCertificate) Set(val *FiltersServerCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersServerCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersServerCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersServerCertificate(val *FiltersServerCertificate) *NullableFiltersServerCertificate {
	return &NullableFiltersServerCertificate{value: val, isSet: true}
}

func (v NullableFiltersServerCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersServerCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
