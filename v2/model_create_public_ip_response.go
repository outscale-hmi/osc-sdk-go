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

// CreatePublicIpResponse struct for CreatePublicIpResponse
type CreatePublicIpResponse struct {
	PublicIp        *PublicIp        `json:"PublicIp,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreatePublicIpResponse instantiates a new CreatePublicIpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePublicIpResponse() *CreatePublicIpResponse {
	this := CreatePublicIpResponse{}
	return &this
}

// NewCreatePublicIpResponseWithDefaults instantiates a new CreatePublicIpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePublicIpResponseWithDefaults() *CreatePublicIpResponse {
	this := CreatePublicIpResponse{}
	return &this
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *CreatePublicIpResponse) GetPublicIp() PublicIp {
	if o == nil || o.PublicIp == nil {
		var ret PublicIp
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePublicIpResponse) GetPublicIpOk() (*PublicIp, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *CreatePublicIpResponse) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given PublicIp and assigns it to the PublicIp field.
func (o *CreatePublicIpResponse) SetPublicIp(v PublicIp) {
	o.PublicIp = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreatePublicIpResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePublicIpResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreatePublicIpResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreatePublicIpResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreatePublicIpResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePublicIpResponse struct {
	value *CreatePublicIpResponse
	isSet bool
}

func (v NullableCreatePublicIpResponse) Get() *CreatePublicIpResponse {
	return v.value
}

func (v *NullableCreatePublicIpResponse) Set(val *CreatePublicIpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePublicIpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePublicIpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePublicIpResponse(val *CreatePublicIpResponse) *NullableCreatePublicIpResponse {
	return &NullableCreatePublicIpResponse{value: val, isSet: true}
}

func (v NullableCreatePublicIpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePublicIpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
