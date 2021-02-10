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

// ReadSubregionsResponse struct for ReadSubregionsResponse
type ReadSubregionsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more Subregions.
	Subregions *[]Subregion `json:"Subregions,omitempty"`
}

// NewReadSubregionsResponse instantiates a new ReadSubregionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadSubregionsResponse() *ReadSubregionsResponse {
	this := ReadSubregionsResponse{}
	return &this
}

// NewReadSubregionsResponseWithDefaults instantiates a new ReadSubregionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadSubregionsResponseWithDefaults() *ReadSubregionsResponse {
	this := ReadSubregionsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSubregionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubregionsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSubregionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSubregionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubregions returns the Subregions field value if set, zero value otherwise.
func (o *ReadSubregionsResponse) GetSubregions() []Subregion {
	if o == nil || o.Subregions == nil {
		var ret []Subregion
		return ret
	}
	return *o.Subregions
}

// GetSubregionsOk returns a tuple with the Subregions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubregionsResponse) GetSubregionsOk() (*[]Subregion, bool) {
	if o == nil || o.Subregions == nil {
		return nil, false
	}
	return o.Subregions, true
}

// HasSubregions returns a boolean if a field has been set.
func (o *ReadSubregionsResponse) HasSubregions() bool {
	if o != nil && o.Subregions != nil {
		return true
	}

	return false
}

// SetSubregions gets a reference to the given []Subregion and assigns it to the Subregions field.
func (o *ReadSubregionsResponse) SetSubregions(v []Subregion) {
	o.Subregions = &v
}

func (o ReadSubregionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Subregions != nil {
		toSerialize["Subregions"] = o.Subregions
	}
	return json.Marshal(toSerialize)
}

type NullableReadSubregionsResponse struct {
	value *ReadSubregionsResponse
	isSet bool
}

func (v NullableReadSubregionsResponse) Get() *ReadSubregionsResponse {
	return v.value
}

func (v *NullableReadSubregionsResponse) Set(val *ReadSubregionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadSubregionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadSubregionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadSubregionsResponse(val *ReadSubregionsResponse) *NullableReadSubregionsResponse {
	return &NullableReadSubregionsResponse{value: val, isSet: true}
}

func (v NullableReadSubregionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadSubregionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
