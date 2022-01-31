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

// ReadApiAccessRulesResponse struct for ReadApiAccessRulesResponse
type ReadApiAccessRulesResponse struct {
	// A list of API access rules.
	ApiAccessRules  *[]ApiAccessRule `json:"ApiAccessRules,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadApiAccessRulesResponse instantiates a new ReadApiAccessRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApiAccessRulesResponse() *ReadApiAccessRulesResponse {
	this := ReadApiAccessRulesResponse{}
	return &this
}

// NewReadApiAccessRulesResponseWithDefaults instantiates a new ReadApiAccessRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApiAccessRulesResponseWithDefaults() *ReadApiAccessRulesResponse {
	this := ReadApiAccessRulesResponse{}
	return &this
}

// GetApiAccessRules returns the ApiAccessRules field value if set, zero value otherwise.
func (o *ReadApiAccessRulesResponse) GetApiAccessRules() []ApiAccessRule {
	if o == nil || o.ApiAccessRules == nil {
		var ret []ApiAccessRule
		return ret
	}
	return *o.ApiAccessRules
}

// GetApiAccessRulesOk returns a tuple with the ApiAccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiAccessRulesResponse) GetApiAccessRulesOk() (*[]ApiAccessRule, bool) {
	if o == nil || o.ApiAccessRules == nil {
		return nil, false
	}
	return o.ApiAccessRules, true
}

// HasApiAccessRules returns a boolean if a field has been set.
func (o *ReadApiAccessRulesResponse) HasApiAccessRules() bool {
	if o != nil && o.ApiAccessRules != nil {
		return true
	}

	return false
}

// SetApiAccessRules gets a reference to the given []ApiAccessRule and assigns it to the ApiAccessRules field.
func (o *ReadApiAccessRulesResponse) SetApiAccessRules(v []ApiAccessRule) {
	o.ApiAccessRules = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadApiAccessRulesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiAccessRulesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadApiAccessRulesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadApiAccessRulesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadApiAccessRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessRules != nil {
		toSerialize["ApiAccessRules"] = o.ApiAccessRules
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadApiAccessRulesResponse struct {
	value *ReadApiAccessRulesResponse
	isSet bool
}

func (v NullableReadApiAccessRulesResponse) Get() *ReadApiAccessRulesResponse {
	return v.value
}

func (v *NullableReadApiAccessRulesResponse) Set(val *ReadApiAccessRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApiAccessRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApiAccessRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApiAccessRulesResponse(val *ReadApiAccessRulesResponse) *NullableReadApiAccessRulesResponse {
	return &NullableReadApiAccessRulesResponse{value: val, isSet: true}
}

func (v NullableReadApiAccessRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApiAccessRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
