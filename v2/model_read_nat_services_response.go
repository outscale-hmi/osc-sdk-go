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

// ReadNatServicesResponse struct for ReadNatServicesResponse
type ReadNatServicesResponse struct {
	// Information about one or more NAT services.
	NatServices     *[]NatService    `json:"NatServices,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadNatServicesResponse instantiates a new ReadNatServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadNatServicesResponse() *ReadNatServicesResponse {
	this := ReadNatServicesResponse{}
	return &this
}

// NewReadNatServicesResponseWithDefaults instantiates a new ReadNatServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadNatServicesResponseWithDefaults() *ReadNatServicesResponse {
	this := ReadNatServicesResponse{}
	return &this
}

// GetNatServices returns the NatServices field value if set, zero value otherwise.
func (o *ReadNatServicesResponse) GetNatServices() []NatService {
	if o == nil || o.NatServices == nil {
		var ret []NatService
		return ret
	}
	return *o.NatServices
}

// GetNatServicesOk returns a tuple with the NatServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNatServicesResponse) GetNatServicesOk() (*[]NatService, bool) {
	if o == nil || o.NatServices == nil {
		return nil, false
	}
	return o.NatServices, true
}

// HasNatServices returns a boolean if a field has been set.
func (o *ReadNatServicesResponse) HasNatServices() bool {
	if o != nil && o.NatServices != nil {
		return true
	}

	return false
}

// SetNatServices gets a reference to the given []NatService and assigns it to the NatServices field.
func (o *ReadNatServicesResponse) SetNatServices(v []NatService) {
	o.NatServices = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNatServicesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadNatServicesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNatServicesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNatServicesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadNatServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NatServices != nil {
		toSerialize["NatServices"] = o.NatServices
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadNatServicesResponse struct {
	value *ReadNatServicesResponse
	isSet bool
}

func (v NullableReadNatServicesResponse) Get() *ReadNatServicesResponse {
	return v.value
}

func (v *NullableReadNatServicesResponse) Set(val *ReadNatServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadNatServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadNatServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadNatServicesResponse(val *ReadNatServicesResponse) *NullableReadNatServicesResponse {
	return &NullableReadNatServicesResponse{value: val, isSet: true}
}

func (v NullableReadNatServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadNatServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
