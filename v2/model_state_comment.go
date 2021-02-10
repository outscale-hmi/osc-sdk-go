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

// StateComment Information about the change of state.
type StateComment struct {
	// The code of the change of state.
	StateCode *string `json:"StateCode,omitempty"`
	// A message explaining the change of state.
	StateMessage *string `json:"StateMessage,omitempty"`
}

// NewStateComment instantiates a new StateComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateComment() *StateComment {
	this := StateComment{}
	return &this
}

// NewStateCommentWithDefaults instantiates a new StateComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateCommentWithDefaults() *StateComment {
	this := StateComment{}
	return &this
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *StateComment) GetStateCode() string {
	if o == nil || o.StateCode == nil {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateComment) GetStateCodeOk() (*string, bool) {
	if o == nil || o.StateCode == nil {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *StateComment) HasStateCode() bool {
	if o != nil && o.StateCode != nil {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *StateComment) SetStateCode(v string) {
	o.StateCode = &v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise.
func (o *StateComment) GetStateMessage() string {
	if o == nil || o.StateMessage == nil {
		var ret string
		return ret
	}
	return *o.StateMessage
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateComment) GetStateMessageOk() (*string, bool) {
	if o == nil || o.StateMessage == nil {
		return nil, false
	}
	return o.StateMessage, true
}

// HasStateMessage returns a boolean if a field has been set.
func (o *StateComment) HasStateMessage() bool {
	if o != nil && o.StateMessage != nil {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given string and assigns it to the StateMessage field.
func (o *StateComment) SetStateMessage(v string) {
	o.StateMessage = &v
}

func (o StateComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StateCode != nil {
		toSerialize["StateCode"] = o.StateCode
	}
	if o.StateMessage != nil {
		toSerialize["StateMessage"] = o.StateMessage
	}
	return json.Marshal(toSerialize)
}

type NullableStateComment struct {
	value *StateComment
	isSet bool
}

func (v NullableStateComment) Get() *StateComment {
	return v.value
}

func (v *NullableStateComment) Set(val *StateComment) {
	v.value = val
	v.isSet = true
}

func (v NullableStateComment) IsSet() bool {
	return v.isSet
}

func (v *NullableStateComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateComment(val *StateComment) *NullableStateComment {
	return &NullableStateComment{value: val, isSet: true}
}

func (v NullableStateComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
