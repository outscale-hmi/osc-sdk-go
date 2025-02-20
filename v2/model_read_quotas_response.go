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

// ReadQuotasResponse struct for ReadQuotasResponse
type ReadQuotasResponse struct {
	// Information about one or more quotas.
	QuotaTypes      *[]QuotaTypes    `json:"QuotaTypes,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewReadQuotasResponse instantiates a new ReadQuotasResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadQuotasResponse() *ReadQuotasResponse {
	this := ReadQuotasResponse{}
	return &this
}

// NewReadQuotasResponseWithDefaults instantiates a new ReadQuotasResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadQuotasResponseWithDefaults() *ReadQuotasResponse {
	this := ReadQuotasResponse{}
	return &this
}

// GetQuotaTypes returns the QuotaTypes field value if set, zero value otherwise.
func (o *ReadQuotasResponse) GetQuotaTypes() []QuotaTypes {
	if o == nil || o.QuotaTypes == nil {
		var ret []QuotaTypes
		return ret
	}
	return *o.QuotaTypes
}

// GetQuotaTypesOk returns a tuple with the QuotaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadQuotasResponse) GetQuotaTypesOk() (*[]QuotaTypes, bool) {
	if o == nil || o.QuotaTypes == nil {
		return nil, false
	}
	return o.QuotaTypes, true
}

// HasQuotaTypes returns a boolean if a field has been set.
func (o *ReadQuotasResponse) HasQuotaTypes() bool {
	if o != nil && o.QuotaTypes != nil {
		return true
	}

	return false
}

// SetQuotaTypes gets a reference to the given []QuotaTypes and assigns it to the QuotaTypes field.
func (o *ReadQuotasResponse) SetQuotaTypes(v []QuotaTypes) {
	o.QuotaTypes = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadQuotasResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadQuotasResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadQuotasResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadQuotasResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o ReadQuotasResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuotaTypes != nil {
		toSerialize["QuotaTypes"] = o.QuotaTypes
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableReadQuotasResponse struct {
	value *ReadQuotasResponse
	isSet bool
}

func (v NullableReadQuotasResponse) Get() *ReadQuotasResponse {
	return v.value
}

func (v *NullableReadQuotasResponse) Set(val *ReadQuotasResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadQuotasResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadQuotasResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadQuotasResponse(val *ReadQuotasResponse) *NullableReadQuotasResponse {
	return &NullableReadQuotasResponse{value: val, isSet: true}
}

func (v NullableReadQuotasResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadQuotasResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
