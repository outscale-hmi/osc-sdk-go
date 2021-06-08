/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateImageResponse struct for UpdateImageResponse
type UpdateImageResponse struct {
	Image           *Image           `json:"Image,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateImageResponse instantiates a new UpdateImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageResponse() *UpdateImageResponse {
	this := UpdateImageResponse{}
	return &this
}

// NewUpdateImageResponseWithDefaults instantiates a new UpdateImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageResponseWithDefaults() *UpdateImageResponse {
	this := UpdateImageResponse{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpdateImageResponse) GetImage() Image {
	if o == nil || o.Image == nil {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageResponse) GetImageOk() (*Image, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpdateImageResponse) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *UpdateImageResponse) SetImage(v Image) {
	o.Image = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateImageResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateImageResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateImageResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateImageResponse struct {
	value *UpdateImageResponse
	isSet bool
}

func (v NullableUpdateImageResponse) Get() *UpdateImageResponse {
	return v.value
}

func (v *NullableUpdateImageResponse) Set(val *UpdateImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageResponse(val *UpdateImageResponse) *NullableUpdateImageResponse {
	return &NullableUpdateImageResponse{value: val, isSet: true}
}

func (v NullableUpdateImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
