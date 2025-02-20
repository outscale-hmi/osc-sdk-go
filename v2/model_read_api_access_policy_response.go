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

// ReadApiAccessPolicyResponse struct for ReadApiAccessPolicyResponse
type ReadApiAccessPolicyResponse struct {
	ApiAccessPolicy *ApiAccessPolicy `json:"ApiAccessPolicy,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadApiAccessPolicyResponse instantiates a new ReadApiAccessPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApiAccessPolicyResponse() *ReadApiAccessPolicyResponse {
	this := ReadApiAccessPolicyResponse{}
	return &this
}

// NewReadApiAccessPolicyResponseWithDefaults instantiates a new ReadApiAccessPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApiAccessPolicyResponseWithDefaults() *ReadApiAccessPolicyResponse {
	this := ReadApiAccessPolicyResponse{}
	return &this
}

// GetApiAccessPolicy returns the ApiAccessPolicy field value if set, zero value otherwise.
func (o *ReadApiAccessPolicyResponse) GetApiAccessPolicy() ApiAccessPolicy {
	if o == nil || o.ApiAccessPolicy == nil {
		var ret ApiAccessPolicy
		return ret
	}
	return *o.ApiAccessPolicy
}

// GetApiAccessPolicyOk returns a tuple with the ApiAccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiAccessPolicyResponse) GetApiAccessPolicyOk() (*ApiAccessPolicy, bool) {
	if o == nil || o.ApiAccessPolicy == nil {
		return nil, false
	}
	return o.ApiAccessPolicy, true
}

// HasApiAccessPolicy returns a boolean if a field has been set.
func (o *ReadApiAccessPolicyResponse) HasApiAccessPolicy() bool {
	if o != nil && o.ApiAccessPolicy != nil {
		return true
	}

	return false
}

// SetApiAccessPolicy gets a reference to the given ApiAccessPolicy and assigns it to the ApiAccessPolicy field.
func (o *ReadApiAccessPolicyResponse) SetApiAccessPolicy(v ApiAccessPolicy) {
	o.ApiAccessPolicy = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadApiAccessPolicyResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiAccessPolicyResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadApiAccessPolicyResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadApiAccessPolicyResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadApiAccessPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessPolicy != nil {
		toSerialize["ApiAccessPolicy"] = o.ApiAccessPolicy
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadApiAccessPolicyResponse struct {
	value *ReadApiAccessPolicyResponse
	isSet bool
}

func (v NullableReadApiAccessPolicyResponse) Get() *ReadApiAccessPolicyResponse {
	return v.value
}

func (v *NullableReadApiAccessPolicyResponse) Set(val *ReadApiAccessPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApiAccessPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApiAccessPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApiAccessPolicyResponse(val *ReadApiAccessPolicyResponse) *NullableReadApiAccessPolicyResponse {
	return &NullableReadApiAccessPolicyResponse{value: val, isSet: true}
}

func (v NullableReadApiAccessPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApiAccessPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
