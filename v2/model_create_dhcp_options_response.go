/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.18
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateDhcpOptionsResponse struct for CreateDhcpOptionsResponse
type CreateDhcpOptionsResponse struct {
	DhcpOptionsSet  *DhcpOptionsSet  `json:"DhcpOptionsSet,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateDhcpOptionsResponse instantiates a new CreateDhcpOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDhcpOptionsResponse() *CreateDhcpOptionsResponse {
	this := CreateDhcpOptionsResponse{}
	return &this
}

// NewCreateDhcpOptionsResponseWithDefaults instantiates a new CreateDhcpOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDhcpOptionsResponseWithDefaults() *CreateDhcpOptionsResponse {
	this := CreateDhcpOptionsResponse{}
	return &this
}

// GetDhcpOptionsSet returns the DhcpOptionsSet field value if set, zero value otherwise.
func (o *CreateDhcpOptionsResponse) GetDhcpOptionsSet() DhcpOptionsSet {
	if o == nil || o.DhcpOptionsSet == nil {
		var ret DhcpOptionsSet
		return ret
	}
	return *o.DhcpOptionsSet
}

// GetDhcpOptionsSetOk returns a tuple with the DhcpOptionsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsResponse) GetDhcpOptionsSetOk() (*DhcpOptionsSet, bool) {
	if o == nil || o.DhcpOptionsSet == nil {
		return nil, false
	}
	return o.DhcpOptionsSet, true
}

// HasDhcpOptionsSet returns a boolean if a field has been set.
func (o *CreateDhcpOptionsResponse) HasDhcpOptionsSet() bool {
	if o != nil && o.DhcpOptionsSet != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSet gets a reference to the given DhcpOptionsSet and assigns it to the DhcpOptionsSet field.
func (o *CreateDhcpOptionsResponse) SetDhcpOptionsSet(v DhcpOptionsSet) {
	o.DhcpOptionsSet = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateDhcpOptionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDhcpOptionsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateDhcpOptionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateDhcpOptionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateDhcpOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DhcpOptionsSet != nil {
		toSerialize["DhcpOptionsSet"] = o.DhcpOptionsSet
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDhcpOptionsResponse struct {
	value *CreateDhcpOptionsResponse
	isSet bool
}

func (v NullableCreateDhcpOptionsResponse) Get() *CreateDhcpOptionsResponse {
	return v.value
}

func (v *NullableCreateDhcpOptionsResponse) Set(val *CreateDhcpOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDhcpOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDhcpOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDhcpOptionsResponse(val *CreateDhcpOptionsResponse) *NullableCreateDhcpOptionsResponse {
	return &NullableCreateDhcpOptionsResponse{value: val, isSet: true}
}

func (v NullableCreateDhcpOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDhcpOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
