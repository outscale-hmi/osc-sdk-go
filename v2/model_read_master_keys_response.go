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

// ReadMasterKeysResponse struct for ReadMasterKeysResponse
type ReadMasterKeysResponse struct {
	// Information about one or more master keys.
	MasterKeys      *[]MasterKey     `json:"MasterKeys,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadMasterKeysResponse instantiates a new ReadMasterKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadMasterKeysResponse() *ReadMasterKeysResponse {
	this := ReadMasterKeysResponse{}
	return &this
}

// NewReadMasterKeysResponseWithDefaults instantiates a new ReadMasterKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadMasterKeysResponseWithDefaults() *ReadMasterKeysResponse {
	this := ReadMasterKeysResponse{}
	return &this
}

// GetMasterKeys returns the MasterKeys field value if set, zero value otherwise.
func (o *ReadMasterKeysResponse) GetMasterKeys() []MasterKey {
	if o == nil || o.MasterKeys == nil {
		var ret []MasterKey
		return ret
	}
	return *o.MasterKeys
}

// GetMasterKeysOk returns a tuple with the MasterKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadMasterKeysResponse) GetMasterKeysOk() (*[]MasterKey, bool) {
	if o == nil || o.MasterKeys == nil {
		return nil, false
	}
	return o.MasterKeys, true
}

// HasMasterKeys returns a boolean if a field has been set.
func (o *ReadMasterKeysResponse) HasMasterKeys() bool {
	if o != nil && o.MasterKeys != nil {
		return true
	}

	return false
}

// SetMasterKeys gets a reference to the given []MasterKey and assigns it to the MasterKeys field.
func (o *ReadMasterKeysResponse) SetMasterKeys(v []MasterKey) {
	o.MasterKeys = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadMasterKeysResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadMasterKeysResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadMasterKeysResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadMasterKeysResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadMasterKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MasterKeys != nil {
		toSerialize["MasterKeys"] = o.MasterKeys
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadMasterKeysResponse struct {
	value *ReadMasterKeysResponse
	isSet bool
}

func (v NullableReadMasterKeysResponse) Get() *ReadMasterKeysResponse {
	return v.value
}

func (v *NullableReadMasterKeysResponse) Set(val *ReadMasterKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadMasterKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadMasterKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadMasterKeysResponse(val *ReadMasterKeysResponse) *NullableReadMasterKeysResponse {
	return &NullableReadMasterKeysResponse{value: val, isSet: true}
}

func (v NullableReadMasterKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadMasterKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
