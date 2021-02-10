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

// NetPeeringState Information about the state of the Net peering connection.
type NetPeeringState struct {
	// Additional information about the state of the Net peering connection.
	Message *string `json:"Message,omitempty"`
	// The state of the Net peering connection (`pending-acceptance` \\| `active` \\| `rejected` \\| `failed` \\| `expired` \\| `deleted`).
	Name *string `json:"Name,omitempty"`
}

// NewNetPeeringState instantiates a new NetPeeringState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetPeeringState() *NetPeeringState {
	this := NetPeeringState{}
	return &this
}

// NewNetPeeringStateWithDefaults instantiates a new NetPeeringState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetPeeringStateWithDefaults() *NetPeeringState {
	this := NetPeeringState{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NetPeeringState) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeeringState) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NetPeeringState) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NetPeeringState) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetPeeringState) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetPeeringState) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetPeeringState) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetPeeringState) SetName(v string) {
	o.Name = &v
}

func (o NetPeeringState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetPeeringState struct {
	value *NetPeeringState
	isSet bool
}

func (v NullableNetPeeringState) Get() *NetPeeringState {
	return v.value
}

func (v *NullableNetPeeringState) Set(val *NetPeeringState) {
	v.value = val
	v.isSet = true
}

func (v NullableNetPeeringState) IsSet() bool {
	return v.isSet
}

func (v *NullableNetPeeringState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetPeeringState(val *NetPeeringState) *NullableNetPeeringState {
	return &NullableNetPeeringState{value: val, isSet: true}
}

func (v NullableNetPeeringState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetPeeringState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
