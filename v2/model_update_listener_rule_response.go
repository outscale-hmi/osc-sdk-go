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

// UpdateListenerRuleResponse struct for UpdateListenerRuleResponse
type UpdateListenerRuleResponse struct {
	ListenerRule    *ListenerRule    `json:"ListenerRule,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewUpdateListenerRuleResponse instantiates a new UpdateListenerRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateListenerRuleResponse() *UpdateListenerRuleResponse {
	this := UpdateListenerRuleResponse{}
	return &this
}

// NewUpdateListenerRuleResponseWithDefaults instantiates a new UpdateListenerRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateListenerRuleResponseWithDefaults() *UpdateListenerRuleResponse {
	this := UpdateListenerRuleResponse{}
	return &this
}

// GetListenerRule returns the ListenerRule field value if set, zero value otherwise.
func (o *UpdateListenerRuleResponse) GetListenerRule() ListenerRule {
	if o == nil || o.ListenerRule == nil {
		var ret ListenerRule
		return ret
	}
	return *o.ListenerRule
}

// GetListenerRuleOk returns a tuple with the ListenerRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleResponse) GetListenerRuleOk() (*ListenerRule, bool) {
	if o == nil || o.ListenerRule == nil {
		return nil, false
	}
	return o.ListenerRule, true
}

// HasListenerRule returns a boolean if a field has been set.
func (o *UpdateListenerRuleResponse) HasListenerRule() bool {
	if o != nil && o.ListenerRule != nil {
		return true
	}

	return false
}

// SetListenerRule gets a reference to the given ListenerRule and assigns it to the ListenerRule field.
func (o *UpdateListenerRuleResponse) SetListenerRule(v ListenerRule) {
	o.ListenerRule = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateListenerRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateListenerRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateListenerRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateListenerRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o UpdateListenerRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListenerRule != nil {
		toSerialize["ListenerRule"] = o.ListenerRule
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateListenerRuleResponse struct {
	value *UpdateListenerRuleResponse
	isSet bool
}

func (v NullableUpdateListenerRuleResponse) Get() *UpdateListenerRuleResponse {
	return v.value
}

func (v *NullableUpdateListenerRuleResponse) Set(val *UpdateListenerRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateListenerRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateListenerRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateListenerRuleResponse(val *UpdateListenerRuleResponse) *NullableUpdateListenerRuleResponse {
	return &NullableUpdateListenerRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateListenerRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateListenerRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
