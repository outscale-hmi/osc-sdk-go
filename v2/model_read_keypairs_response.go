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

// ReadKeypairsResponse struct for ReadKeypairsResponse
type ReadKeypairsResponse struct {
	// Information about one or more keypairs.
	Keypairs        *[]Keypair       `json:"Keypairs,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadKeypairsResponse instantiates a new ReadKeypairsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadKeypairsResponse() *ReadKeypairsResponse {
	this := ReadKeypairsResponse{}
	return &this
}

// NewReadKeypairsResponseWithDefaults instantiates a new ReadKeypairsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadKeypairsResponseWithDefaults() *ReadKeypairsResponse {
	this := ReadKeypairsResponse{}
	return &this
}

// GetKeypairs returns the Keypairs field value if set, zero value otherwise.
func (o *ReadKeypairsResponse) GetKeypairs() []Keypair {
	if o == nil || o.Keypairs == nil {
		var ret []Keypair
		return ret
	}
	return *o.Keypairs
}

// GetKeypairsOk returns a tuple with the Keypairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadKeypairsResponse) GetKeypairsOk() (*[]Keypair, bool) {
	if o == nil || o.Keypairs == nil {
		return nil, false
	}
	return o.Keypairs, true
}

// HasKeypairs returns a boolean if a field has been set.
func (o *ReadKeypairsResponse) HasKeypairs() bool {
	if o != nil && o.Keypairs != nil {
		return true
	}

	return false
}

// SetKeypairs gets a reference to the given []Keypair and assigns it to the Keypairs field.
func (o *ReadKeypairsResponse) SetKeypairs(v []Keypair) {
	o.Keypairs = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadKeypairsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadKeypairsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadKeypairsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadKeypairsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadKeypairsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keypairs != nil {
		toSerialize["Keypairs"] = o.Keypairs
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadKeypairsResponse struct {
	value *ReadKeypairsResponse
	isSet bool
}

func (v NullableReadKeypairsResponse) Get() *ReadKeypairsResponse {
	return v.value
}

func (v *NullableReadKeypairsResponse) Set(val *ReadKeypairsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadKeypairsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadKeypairsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadKeypairsResponse(val *ReadKeypairsResponse) *NullableReadKeypairsResponse {
	return &NullableReadKeypairsResponse{value: val, isSet: true}
}

func (v NullableReadKeypairsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadKeypairsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
