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

// RebootVmsResponse struct for RebootVmsResponse
type RebootVmsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewRebootVmsResponse instantiates a new RebootVmsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRebootVmsResponse() *RebootVmsResponse {
	this := RebootVmsResponse{}
	return &this
}

// NewRebootVmsResponseWithDefaults instantiates a new RebootVmsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRebootVmsResponseWithDefaults() *RebootVmsResponse {
	this := RebootVmsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *RebootVmsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RebootVmsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *RebootVmsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *RebootVmsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o RebootVmsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableRebootVmsResponse struct {
	value *RebootVmsResponse
	isSet bool
}

func (v NullableRebootVmsResponse) Get() *RebootVmsResponse {
	return v.value
}

func (v *NullableRebootVmsResponse) Set(val *RebootVmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRebootVmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRebootVmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRebootVmsResponse(val *RebootVmsResponse) *NullableRebootVmsResponse {
	return &NullableRebootVmsResponse{value: val, isSet: true}
}

func (v NullableRebootVmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRebootVmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
