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

// ReadSecretAccessKeyResponse struct for ReadSecretAccessKeyResponse
type ReadSecretAccessKeyResponse struct {
	AccessKey       *AccessKeySecretKey `json:"AccessKey,omitempty"`
	ResponseContext *ResponseContext    `json:"ResponseContext,omitempty"`
}

// NewReadSecretAccessKeyResponse instantiates a new ReadSecretAccessKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSecretAccessKeyResponse() *ReadSecretAccessKeyResponse {
	this := ReadSecretAccessKeyResponse{}
	return &this
}

// NewReadSecretAccessKeyResponseWithDefaults instantiates a new ReadSecretAccessKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSecretAccessKeyResponseWithDefaults() *ReadSecretAccessKeyResponse {
	this := ReadSecretAccessKeyResponse{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ReadSecretAccessKeyResponse) GetAccessKey() AccessKeySecretKey {
	if o == nil || o.AccessKey == nil {
		var ret AccessKeySecretKey
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSecretAccessKeyResponse) GetAccessKeyOk() (*AccessKeySecretKey, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ReadSecretAccessKeyResponse) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given AccessKeySecretKey and assigns it to the AccessKey field.
func (o *ReadSecretAccessKeyResponse) SetAccessKey(v AccessKeySecretKey) {
	o.AccessKey = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSecretAccessKeyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSecretAccessKeyResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSecretAccessKeyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSecretAccessKeyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadSecretAccessKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["AccessKey"] = o.AccessKey
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadSecretAccessKeyResponse struct {
	value *ReadSecretAccessKeyResponse
	isSet bool
}

func (v NullableReadSecretAccessKeyResponse) Get() *ReadSecretAccessKeyResponse {
	return v.value
}

func (v *NullableReadSecretAccessKeyResponse) Set(val *ReadSecretAccessKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSecretAccessKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSecretAccessKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSecretAccessKeyResponse(val *ReadSecretAccessKeyResponse) *NullableReadSecretAccessKeyResponse {
	return &NullableReadSecretAccessKeyResponse{value: val, isSet: true}
}

func (v NullableReadSecretAccessKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSecretAccessKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
