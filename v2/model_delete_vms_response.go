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

// DeleteVmsResponse struct for DeleteVmsResponse
type DeleteVmsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more terminated VMs.
	Vms *[]VmState `json:"Vms,omitempty"`
}

// NewDeleteVmsResponse instantiates a new DeleteVmsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVmsResponse() *DeleteVmsResponse {
	this := DeleteVmsResponse{}
	return &this
}

// NewDeleteVmsResponseWithDefaults instantiates a new DeleteVmsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVmsResponseWithDefaults() *DeleteVmsResponse {
	this := DeleteVmsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteVmsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVmsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteVmsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteVmsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *DeleteVmsResponse) GetVms() []VmState {
	if o == nil || o.Vms == nil {
		var ret []VmState
		return ret
	}
	return *o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVmsResponse) GetVmsOk() (*[]VmState, bool) {
	if o == nil || o.Vms == nil {
		return nil, false
	}
	return o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *DeleteVmsResponse) HasVms() bool {
	if o != nil && o.Vms != nil {
		return true
	}

	return false
}

// SetVms gets a reference to the given []VmState and assigns it to the Vms field.
func (o *DeleteVmsResponse) SetVms(v []VmState) {
	o.Vms = &v
}

func (o DeleteVmsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Vms != nil {
		toSerialize["Vms"] = o.Vms
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteVmsResponse struct {
	value *DeleteVmsResponse
	isSet bool
}

func (v NullableDeleteVmsResponse) Get() *DeleteVmsResponse {
	return v.value
}

func (v *NullableDeleteVmsResponse) Set(val *DeleteVmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVmsResponse(val *DeleteVmsResponse) *NullableDeleteVmsResponse {
	return &NullableDeleteVmsResponse{value: val, isSet: true}
}

func (v NullableDeleteVmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
