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

// ReadImagesResponse struct for ReadImagesResponse
type ReadImagesResponse struct {
	// Information about one or more OMIs.
	Images          *[]Image         `json:"Images,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadImagesResponse instantiates a new ReadImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadImagesResponse() *ReadImagesResponse {
	this := ReadImagesResponse{}
	return &this
}

// NewReadImagesResponseWithDefaults instantiates a new ReadImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadImagesResponseWithDefaults() *ReadImagesResponse {
	this := ReadImagesResponse{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ReadImagesResponse) GetImages() []Image {
	if o == nil || o.Images == nil {
		var ret []Image
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesResponse) GetImagesOk() (*[]Image, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ReadImagesResponse) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *ReadImagesResponse) SetImages(v []Image) {
	o.Images = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadImagesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadImagesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadImagesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["Images"] = o.Images
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadImagesResponse struct {
	value *ReadImagesResponse
	isSet bool
}

func (v NullableReadImagesResponse) Get() *ReadImagesResponse {
	return v.value
}

func (v *NullableReadImagesResponse) Set(val *ReadImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadImagesResponse(val *ReadImagesResponse) *NullableReadImagesResponse {
	return &NullableReadImagesResponse{value: val, isSet: true}
}

func (v NullableReadImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
