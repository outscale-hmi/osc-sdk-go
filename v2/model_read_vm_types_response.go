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

// ReadVmTypesResponse struct for ReadVmTypesResponse
type ReadVmTypesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more VM types.
	VmTypes *[]VmType `json:"VmTypes,omitempty"`
}

// NewReadVmTypesResponse instantiates a new ReadVmTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadVmTypesResponse() *ReadVmTypesResponse {
	this := ReadVmTypesResponse{}
	return &this
}

// NewReadVmTypesResponseWithDefaults instantiates a new ReadVmTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadVmTypesResponseWithDefaults() *ReadVmTypesResponse {
	this := ReadVmTypesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadVmTypesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadVmTypesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadVmTypesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVmTypes returns the VmTypes field value if set, zero value otherwise.
func (o *ReadVmTypesResponse) GetVmTypes() []VmType {
	if o == nil || o.VmTypes == nil {
		var ret []VmType
		return ret
	}
	return *o.VmTypes
}

// GetVmTypesOk returns a tuple with the VmTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesResponse) GetVmTypesOk() (*[]VmType, bool) {
	if o == nil || o.VmTypes == nil {
		return nil, false
	}
	return o.VmTypes, true
}

// HasVmTypes returns a boolean if a field has been set.
func (o *ReadVmTypesResponse) HasVmTypes() bool {
	if o != nil && o.VmTypes != nil {
		return true
	}

	return false
}

// SetVmTypes gets a reference to the given []VmType and assigns it to the VmTypes field.
func (o *ReadVmTypesResponse) SetVmTypes(v []VmType) {
	o.VmTypes = &v
}

func (o ReadVmTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.VmTypes != nil {
		toSerialize["VmTypes"] = o.VmTypes
	}
	return json.Marshal(toSerialize)
}

type NullableReadVmTypesResponse struct {
	value *ReadVmTypesResponse
	isSet bool
}

func (v NullableReadVmTypesResponse) Get() *ReadVmTypesResponse {
	return v.value
}

func (v *NullableReadVmTypesResponse) Set(val *ReadVmTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadVmTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadVmTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadVmTypesResponse(val *ReadVmTypesResponse) *NullableReadVmTypesResponse {
	return &NullableReadVmTypesResponse{value: val, isSet: true}
}

func (v NullableReadVmTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadVmTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
