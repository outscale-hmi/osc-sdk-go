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

// ReadRegionsResponse struct for ReadRegionsResponse
type ReadRegionsResponse struct {
	// Information about one or more Regions.
	Regions         *[]Region        `json:"Regions,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadRegionsResponse instantiates a new ReadRegionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRegionsResponse() *ReadRegionsResponse {
	this := ReadRegionsResponse{}
	return &this
}

// NewReadRegionsResponseWithDefaults instantiates a new ReadRegionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRegionsResponseWithDefaults() *ReadRegionsResponse {
	this := ReadRegionsResponse{}
	return &this
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *ReadRegionsResponse) GetRegions() []Region {
	if o == nil || o.Regions == nil {
		var ret []Region
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRegionsResponse) GetRegionsOk() (*[]Region, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *ReadRegionsResponse) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []Region and assigns it to the Regions field.
func (o *ReadRegionsResponse) SetRegions(v []Region) {
	o.Regions = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadRegionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRegionsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadRegionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadRegionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadRegionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Regions != nil {
		toSerialize["Regions"] = o.Regions
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadRegionsResponse struct {
	value *ReadRegionsResponse
	isSet bool
}

func (v NullableReadRegionsResponse) Get() *ReadRegionsResponse {
	return v.value
}

func (v *NullableReadRegionsResponse) Set(val *ReadRegionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRegionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRegionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRegionsResponse(val *ReadRegionsResponse) *NullableReadRegionsResponse {
	return &NullableReadRegionsResponse{value: val, isSet: true}
}

func (v NullableReadRegionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRegionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
