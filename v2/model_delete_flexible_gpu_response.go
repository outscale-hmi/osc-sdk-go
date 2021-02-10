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

// DeleteFlexibleGpuResponse struct for DeleteFlexibleGpuResponse
type DeleteFlexibleGpuResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteFlexibleGpuResponse instantiates a new DeleteFlexibleGpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFlexibleGpuResponse() *DeleteFlexibleGpuResponse {
	this := DeleteFlexibleGpuResponse{}
	return &this
}

// NewDeleteFlexibleGpuResponseWithDefaults instantiates a new DeleteFlexibleGpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFlexibleGpuResponseWithDefaults() *DeleteFlexibleGpuResponse {
	this := DeleteFlexibleGpuResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteFlexibleGpuResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFlexibleGpuResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteFlexibleGpuResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteFlexibleGpuResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteFlexibleGpuResponse struct {
	value *DeleteFlexibleGpuResponse
	isSet bool
}

func (v NullableDeleteFlexibleGpuResponse) Get() *DeleteFlexibleGpuResponse {
	return v.value
}

func (v *NullableDeleteFlexibleGpuResponse) Set(val *DeleteFlexibleGpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFlexibleGpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFlexibleGpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFlexibleGpuResponse(val *DeleteFlexibleGpuResponse) *NullableDeleteFlexibleGpuResponse {
	return &NullableDeleteFlexibleGpuResponse{value: val, isSet: true}
}

func (v NullableDeleteFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFlexibleGpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
