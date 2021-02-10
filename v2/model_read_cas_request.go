/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadCasRequest struct for ReadCasRequest
type ReadCasRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun  *bool      `json:"DryRun,omitempty"`
	Filters *FiltersCa `json:"Filters,omitempty"`
}

// NewReadCasRequest instantiates a new ReadCasRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadCasRequest() *ReadCasRequest {
	this := ReadCasRequest{}
	return &this
}

// NewReadCasRequestWithDefaults instantiates a new ReadCasRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadCasRequestWithDefaults() *ReadCasRequest {
	this := ReadCasRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadCasRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCasRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadCasRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadCasRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadCasRequest) GetFilters() FiltersCa {
	if o == nil || o.Filters == nil {
		var ret FiltersCa
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadCasRequest) GetFiltersOk() (*FiltersCa, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadCasRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersCa and assigns it to the Filters field.
func (o *ReadCasRequest) SetFilters(v FiltersCa) {
	o.Filters = &v
}

func (o ReadCasRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadCasRequest struct {
	value *ReadCasRequest
	isSet bool
}

func (v NullableReadCasRequest) Get() *ReadCasRequest {
	return v.value
}

func (v *NullableReadCasRequest) Set(val *ReadCasRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadCasRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadCasRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadCasRequest(val *ReadCasRequest) *NullableReadCasRequest {
	return &NullableReadCasRequest{value: val, isSet: true}
}

func (v NullableReadCasRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadCasRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
