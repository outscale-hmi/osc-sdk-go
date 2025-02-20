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

// CreateFlexibleGpuResponse struct for CreateFlexibleGpuResponse
type CreateFlexibleGpuResponse struct {
	FlexibleGpu     *FlexibleGpu     `json:"FlexibleGpu,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateFlexibleGpuResponse instantiates a new CreateFlexibleGpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFlexibleGpuResponse() *CreateFlexibleGpuResponse {
	this := CreateFlexibleGpuResponse{}
	return &this
}

// NewCreateFlexibleGpuResponseWithDefaults instantiates a new CreateFlexibleGpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFlexibleGpuResponseWithDefaults() *CreateFlexibleGpuResponse {
	this := CreateFlexibleGpuResponse{}
	return &this
}

// GetFlexibleGpu returns the FlexibleGpu field value if set, zero value otherwise.
func (o *CreateFlexibleGpuResponse) GetFlexibleGpu() FlexibleGpu {
	if o == nil || o.FlexibleGpu == nil {
		var ret FlexibleGpu
		return ret
	}
	return *o.FlexibleGpu
}

// GetFlexibleGpuOk returns a tuple with the FlexibleGpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuResponse) GetFlexibleGpuOk() (*FlexibleGpu, bool) {
	if o == nil || o.FlexibleGpu == nil {
		return nil, false
	}
	return o.FlexibleGpu, true
}

// HasFlexibleGpu returns a boolean if a field has been set.
func (o *CreateFlexibleGpuResponse) HasFlexibleGpu() bool {
	if o != nil && o.FlexibleGpu != nil {
		return true
	}

	return false
}

// SetFlexibleGpu gets a reference to the given FlexibleGpu and assigns it to the FlexibleGpu field.
func (o *CreateFlexibleGpuResponse) SetFlexibleGpu(v FlexibleGpu) {
	o.FlexibleGpu = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateFlexibleGpuResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFlexibleGpuResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateFlexibleGpuResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateFlexibleGpuResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlexibleGpu != nil {
		toSerialize["FlexibleGpu"] = o.FlexibleGpu
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFlexibleGpuResponse struct {
	value *CreateFlexibleGpuResponse
	isSet bool
}

func (v NullableCreateFlexibleGpuResponse) Get() *CreateFlexibleGpuResponse {
	return v.value
}

func (v *NullableCreateFlexibleGpuResponse) Set(val *CreateFlexibleGpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFlexibleGpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFlexibleGpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFlexibleGpuResponse(val *CreateFlexibleGpuResponse) *NullableCreateFlexibleGpuResponse {
	return &NullableCreateFlexibleGpuResponse{value: val, isSet: true}
}

func (v NullableCreateFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFlexibleGpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
