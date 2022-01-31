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

// ReadClientGatewaysRequest struct for ReadClientGatewaysRequest
type ReadClientGatewaysRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool                 `json:"DryRun,omitempty"`
	Filters *FiltersClientGateway `json:"Filters,omitempty"`
}

// NewReadClientGatewaysRequest instantiates a new ReadClientGatewaysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadClientGatewaysRequest() *ReadClientGatewaysRequest {
	this := ReadClientGatewaysRequest{}
	return &this
}

// NewReadClientGatewaysRequestWithDefaults instantiates a new ReadClientGatewaysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadClientGatewaysRequestWithDefaults() *ReadClientGatewaysRequest {
	this := ReadClientGatewaysRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadClientGatewaysRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadClientGatewaysRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadClientGatewaysRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadClientGatewaysRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadClientGatewaysRequest) GetFilters() FiltersClientGateway {
	if o == nil || o.Filters == nil {
		var ret FiltersClientGateway
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadClientGatewaysRequest) GetFiltersOk() (*FiltersClientGateway, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadClientGatewaysRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersClientGateway and assigns it to the Filters field.
func (o *ReadClientGatewaysRequest) SetFilters(v FiltersClientGateway) {
	o.Filters = &v
}

func (o ReadClientGatewaysRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadClientGatewaysRequest struct {
	value *ReadClientGatewaysRequest
	isSet bool
}

func (v NullableReadClientGatewaysRequest) Get() *ReadClientGatewaysRequest {
	return v.value
}

func (v *NullableReadClientGatewaysRequest) Set(val *ReadClientGatewaysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadClientGatewaysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadClientGatewaysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadClientGatewaysRequest(val *ReadClientGatewaysRequest) *NullableReadClientGatewaysRequest {
	return &NullableReadClientGatewaysRequest{value: val, isSet: true}
}

func (v NullableReadClientGatewaysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadClientGatewaysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
