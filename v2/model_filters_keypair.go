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

// FiltersKeypair One or more filters.
type FiltersKeypair struct {
	// The fingerprints of the keypairs.
	KeypairFingerprints *[]string `json:"KeypairFingerprints,omitempty"`
	// The names of the keypairs.
	KeypairNames *[]string `json:"KeypairNames,omitempty"`
}

// NewFiltersKeypair instantiates a new FiltersKeypair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersKeypair() *FiltersKeypair {
	this := FiltersKeypair{}
	return &this
}

// NewFiltersKeypairWithDefaults instantiates a new FiltersKeypair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersKeypairWithDefaults() *FiltersKeypair {
	this := FiltersKeypair{}
	return &this
}

// GetKeypairFingerprints returns the KeypairFingerprints field value if set, zero value otherwise.
func (o *FiltersKeypair) GetKeypairFingerprints() []string {
	if o == nil || o.KeypairFingerprints == nil {
		var ret []string
		return ret
	}
	return *o.KeypairFingerprints
}

// GetKeypairFingerprintsOk returns a tuple with the KeypairFingerprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersKeypair) GetKeypairFingerprintsOk() (*[]string, bool) {
	if o == nil || o.KeypairFingerprints == nil {
		return nil, false
	}
	return o.KeypairFingerprints, true
}

// HasKeypairFingerprints returns a boolean if a field has been set.
func (o *FiltersKeypair) HasKeypairFingerprints() bool {
	if o != nil && o.KeypairFingerprints != nil {
		return true
	}

	return false
}

// SetKeypairFingerprints gets a reference to the given []string and assigns it to the KeypairFingerprints field.
func (o *FiltersKeypair) SetKeypairFingerprints(v []string) {
	o.KeypairFingerprints = &v
}

// GetKeypairNames returns the KeypairNames field value if set, zero value otherwise.
func (o *FiltersKeypair) GetKeypairNames() []string {
	if o == nil || o.KeypairNames == nil {
		var ret []string
		return ret
	}
	return *o.KeypairNames
}

// GetKeypairNamesOk returns a tuple with the KeypairNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersKeypair) GetKeypairNamesOk() (*[]string, bool) {
	if o == nil || o.KeypairNames == nil {
		return nil, false
	}
	return o.KeypairNames, true
}

// HasKeypairNames returns a boolean if a field has been set.
func (o *FiltersKeypair) HasKeypairNames() bool {
	if o != nil && o.KeypairNames != nil {
		return true
	}

	return false
}

// SetKeypairNames gets a reference to the given []string and assigns it to the KeypairNames field.
func (o *FiltersKeypair) SetKeypairNames(v []string) {
	o.KeypairNames = &v
}

func (o FiltersKeypair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeypairFingerprints != nil {
		toSerialize["KeypairFingerprints"] = o.KeypairFingerprints
	}
	if o.KeypairNames != nil {
		toSerialize["KeypairNames"] = o.KeypairNames
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersKeypair struct {
	value *FiltersKeypair
	isSet bool
}

func (v NullableFiltersKeypair) Get() *FiltersKeypair {
	return v.value
}

func (v *NullableFiltersKeypair) Set(val *FiltersKeypair) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersKeypair) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersKeypair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersKeypair(val *FiltersKeypair) *NullableFiltersKeypair {
	return &NullableFiltersKeypair{value: val, isSet: true}
}

func (v NullableFiltersKeypair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersKeypair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
