/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadNicsResponse struct for ReadNicsResponse
type ReadNicsResponse struct {
	// Information about one or more NICs.
	Nics            *[]Nic           `json:"Nics,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadNicsResponse instantiates a new ReadNicsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadNicsResponse() *ReadNicsResponse {
	this := ReadNicsResponse{}
	return &this
}

// NewReadNicsResponseWithDefaults instantiates a new ReadNicsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadNicsResponseWithDefaults() *ReadNicsResponse {
	this := ReadNicsResponse{}
	return &this
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *ReadNicsResponse) GetNics() []Nic {
	if o == nil || o.Nics == nil {
		var ret []Nic
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNicsResponse) GetNicsOk() (*[]Nic, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *ReadNicsResponse) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []Nic and assigns it to the Nics field.
func (o *ReadNicsResponse) SetNics(v []Nic) {
	o.Nics = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNicsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNicsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNicsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNicsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadNicsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nics != nil {
		toSerialize["Nics"] = o.Nics
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadNicsResponse struct {
	value *ReadNicsResponse
	isSet bool
}

func (v NullableReadNicsResponse) Get() *ReadNicsResponse {
	return v.value
}

func (v *NullableReadNicsResponse) Set(val *ReadNicsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadNicsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadNicsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadNicsResponse(val *ReadNicsResponse) *NullableReadNicsResponse {
	return &NullableReadNicsResponse{value: val, isSet: true}
}

func (v NullableReadNicsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadNicsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
