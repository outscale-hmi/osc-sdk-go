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

// ReadNetPeeringsResponse struct for ReadNetPeeringsResponse
type ReadNetPeeringsResponse struct {
	// Information about one or more Net peering connections.
	NetPeerings     *[]NetPeering    `json:"NetPeerings,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadNetPeeringsResponse instantiates a new ReadNetPeeringsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadNetPeeringsResponse() *ReadNetPeeringsResponse {
	this := ReadNetPeeringsResponse{}
	return &this
}

// NewReadNetPeeringsResponseWithDefaults instantiates a new ReadNetPeeringsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadNetPeeringsResponseWithDefaults() *ReadNetPeeringsResponse {
	this := ReadNetPeeringsResponse{}
	return &this
}

// GetNetPeerings returns the NetPeerings field value if set, zero value otherwise.
func (o *ReadNetPeeringsResponse) GetNetPeerings() []NetPeering {
	if o == nil || o.NetPeerings == nil {
		var ret []NetPeering
		return ret
	}
	return *o.NetPeerings
}

// GetNetPeeringsOk returns a tuple with the NetPeerings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetPeeringsResponse) GetNetPeeringsOk() (*[]NetPeering, bool) {
	if o == nil || o.NetPeerings == nil {
		return nil, false
	}
	return o.NetPeerings, true
}

// HasNetPeerings returns a boolean if a field has been set.
func (o *ReadNetPeeringsResponse) HasNetPeerings() bool {
	if o != nil && o.NetPeerings != nil {
		return true
	}

	return false
}

// SetNetPeerings gets a reference to the given []NetPeering and assigns it to the NetPeerings field.
func (o *ReadNetPeeringsResponse) SetNetPeerings(v []NetPeering) {
	o.NetPeerings = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetPeeringsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetPeeringsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetPeeringsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetPeeringsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadNetPeeringsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetPeerings != nil {
		toSerialize["NetPeerings"] = o.NetPeerings
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadNetPeeringsResponse struct {
	value *ReadNetPeeringsResponse
	isSet bool
}

func (v NullableReadNetPeeringsResponse) Get() *ReadNetPeeringsResponse {
	return v.value
}

func (v *NullableReadNetPeeringsResponse) Set(val *ReadNetPeeringsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadNetPeeringsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadNetPeeringsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadNetPeeringsResponse(val *ReadNetPeeringsResponse) *NullableReadNetPeeringsResponse {
	return &NullableReadNetPeeringsResponse{value: val, isSet: true}
}

func (v NullableReadNetPeeringsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadNetPeeringsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
