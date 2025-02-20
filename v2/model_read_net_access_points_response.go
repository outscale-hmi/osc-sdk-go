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

// ReadNetAccessPointsResponse struct for ReadNetAccessPointsResponse
type ReadNetAccessPointsResponse struct {
	// One or more Net access points.
	NetAccessPoints *[]NetAccessPoint `json:"NetAccessPoints,omitempty"`
	ResponseContext *ResponseContext  `json:"ResponseContext,omitempty"`
}

// NewReadNetAccessPointsResponse instantiates a new ReadNetAccessPointsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadNetAccessPointsResponse() *ReadNetAccessPointsResponse {
	this := ReadNetAccessPointsResponse{}
	return &this
}

// NewReadNetAccessPointsResponseWithDefaults instantiates a new ReadNetAccessPointsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadNetAccessPointsResponseWithDefaults() *ReadNetAccessPointsResponse {
	this := ReadNetAccessPointsResponse{}
	return &this
}

// GetNetAccessPoints returns the NetAccessPoints field value if set, zero value otherwise.
func (o *ReadNetAccessPointsResponse) GetNetAccessPoints() []NetAccessPoint {
	if o == nil || o.NetAccessPoints == nil {
		var ret []NetAccessPoint
		return ret
	}
	return *o.NetAccessPoints
}

// GetNetAccessPointsOk returns a tuple with the NetAccessPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointsResponse) GetNetAccessPointsOk() (*[]NetAccessPoint, bool) {
	if o == nil || o.NetAccessPoints == nil {
		return nil, false
	}
	return o.NetAccessPoints, true
}

// HasNetAccessPoints returns a boolean if a field has been set.
func (o *ReadNetAccessPointsResponse) HasNetAccessPoints() bool {
	if o != nil && o.NetAccessPoints != nil {
		return true
	}

	return false
}

// SetNetAccessPoints gets a reference to the given []NetAccessPoint and assigns it to the NetAccessPoints field.
func (o *ReadNetAccessPointsResponse) SetNetAccessPoints(v []NetAccessPoint) {
	o.NetAccessPoints = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetAccessPointsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetAccessPointsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetAccessPointsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetAccessPointsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadNetAccessPointsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetAccessPoints != nil {
		toSerialize["NetAccessPoints"] = o.NetAccessPoints
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadNetAccessPointsResponse struct {
	value *ReadNetAccessPointsResponse
	isSet bool
}

func (v NullableReadNetAccessPointsResponse) Get() *ReadNetAccessPointsResponse {
	return v.value
}

func (v *NullableReadNetAccessPointsResponse) Set(val *ReadNetAccessPointsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadNetAccessPointsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadNetAccessPointsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadNetAccessPointsResponse(val *ReadNetAccessPointsResponse) *NullableReadNetAccessPointsResponse {
	return &NullableReadNetAccessPointsResponse{value: val, isSet: true}
}

func (v NullableReadNetAccessPointsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadNetAccessPointsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
