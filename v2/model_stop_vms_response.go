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

// StopVmsResponse struct for StopVmsResponse
type StopVmsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more stopped VMs.
	Vms *[]VmState `json:"Vms,omitempty"`
}

// NewStopVmsResponse instantiates a new StopVmsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopVmsResponse() *StopVmsResponse {
	this := StopVmsResponse{}
	return &this
}

// NewStopVmsResponseWithDefaults instantiates a new StopVmsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopVmsResponseWithDefaults() *StopVmsResponse {
	this := StopVmsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *StopVmsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *StopVmsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *StopVmsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *StopVmsResponse) GetVms() []VmState {
	if o == nil || o.Vms == nil {
		var ret []VmState
		return ret
	}
	return *o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVmsResponse) GetVmsOk() (*[]VmState, bool) {
	if o == nil || o.Vms == nil {
		return nil, false
	}
	return o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *StopVmsResponse) HasVms() bool {
	if o != nil && o.Vms != nil {
		return true
	}

	return false
}

// SetVms gets a reference to the given []VmState and assigns it to the Vms field.
func (o *StopVmsResponse) SetVms(v []VmState) {
	o.Vms = &v
}

func (o StopVmsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Vms != nil {
		toSerialize["Vms"] = o.Vms
	}
	return json.Marshal(toSerialize)
}

type NullableStopVmsResponse struct {
	value *StopVmsResponse
	isSet bool
}

func (v NullableStopVmsResponse) Get() *StopVmsResponse {
	return v.value
}

func (v *NullableStopVmsResponse) Set(val *StopVmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStopVmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStopVmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopVmsResponse(val *StopVmsResponse) *NullableStopVmsResponse {
	return &NullableStopVmsResponse{value: val, isSet: true}
}

func (v NullableStopVmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopVmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
