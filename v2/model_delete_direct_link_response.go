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

// DeleteDirectLinkResponse struct for DeleteDirectLinkResponse
type DeleteDirectLinkResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewDeleteDirectLinkResponse instantiates a new DeleteDirectLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDirectLinkResponse() *DeleteDirectLinkResponse {
	this := DeleteDirectLinkResponse{}
	return &this
}

// NewDeleteDirectLinkResponseWithDefaults instantiates a new DeleteDirectLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDirectLinkResponseWithDefaults() *DeleteDirectLinkResponse {
	this := DeleteDirectLinkResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *DeleteDirectLinkResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDirectLinkResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *DeleteDirectLinkResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *DeleteDirectLinkResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o DeleteDirectLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDirectLinkResponse struct {
	value *DeleteDirectLinkResponse
	isSet bool
}

func (v NullableDeleteDirectLinkResponse) Get() *DeleteDirectLinkResponse {
	return v.value
}

func (v *NullableDeleteDirectLinkResponse) Set(val *DeleteDirectLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDirectLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDirectLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDirectLinkResponse(val *DeleteDirectLinkResponse) *NullableDeleteDirectLinkResponse {
	return &NullableDeleteDirectLinkResponse{value: val, isSet: true}
}

func (v NullableDeleteDirectLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDirectLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
