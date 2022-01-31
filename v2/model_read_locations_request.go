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

// ReadLocationsRequest struct for ReadLocationsRequest
type ReadLocationsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewReadLocationsRequest instantiates a new ReadLocationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLocationsRequest() *ReadLocationsRequest {
	this := ReadLocationsRequest{}
	return &this
}

// NewReadLocationsRequestWithDefaults instantiates a new ReadLocationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLocationsRequestWithDefaults() *ReadLocationsRequest {
	this := ReadLocationsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadLocationsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLocationsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadLocationsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadLocationsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o ReadLocationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableReadLocationsRequest struct {
	value *ReadLocationsRequest
	isSet bool
}

func (v NullableReadLocationsRequest) Get() *ReadLocationsRequest {
	return v.value
}

func (v *NullableReadLocationsRequest) Set(val *ReadLocationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLocationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLocationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLocationsRequest(val *ReadLocationsRequest) *NullableReadLocationsRequest {
	return &NullableReadLocationsRequest{value: val, isSet: true}
}

func (v NullableReadLocationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLocationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
