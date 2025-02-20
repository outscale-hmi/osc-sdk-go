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

// ReadServerCertificatesResponse struct for ReadServerCertificatesResponse
type ReadServerCertificatesResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more server certificates.
	ServerCertificates *[]ServerCertificate `json:"ServerCertificates,omitempty"`
}

// NewReadServerCertificatesResponse instantiates a new ReadServerCertificatesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadServerCertificatesResponse() *ReadServerCertificatesResponse {
	this := ReadServerCertificatesResponse{}
	return &this
}

// NewReadServerCertificatesResponseWithDefaults instantiates a new ReadServerCertificatesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadServerCertificatesResponseWithDefaults() *ReadServerCertificatesResponse {
	this := ReadServerCertificatesResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadServerCertificatesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadServerCertificatesResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadServerCertificatesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadServerCertificatesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetServerCertificates returns the ServerCertificates field value if set, zero value otherwise.
func (o *ReadServerCertificatesResponse) GetServerCertificates() []ServerCertificate {
	if o == nil || o.ServerCertificates == nil {
		var ret []ServerCertificate
		return ret
	}
	return *o.ServerCertificates
}

// GetServerCertificatesOk returns a tuple with the ServerCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadServerCertificatesResponse) GetServerCertificatesOk() (*[]ServerCertificate, bool) {
	if o == nil || o.ServerCertificates == nil {
		return nil, false
	}
	return o.ServerCertificates, true
}

// HasServerCertificates returns a boolean if a field has been set.
func (o *ReadServerCertificatesResponse) HasServerCertificates() bool {
	if o != nil && o.ServerCertificates != nil {
		return true
	}

	return false
}

// SetServerCertificates gets a reference to the given []ServerCertificate and assigns it to the ServerCertificates field.
func (o *ReadServerCertificatesResponse) SetServerCertificates(v []ServerCertificate) {
	o.ServerCertificates = &v
}

func (o ReadServerCertificatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.ServerCertificates != nil {
		toSerialize["ServerCertificates"] = o.ServerCertificates
	}
	return json.Marshal(toSerialize)
}

type NullableReadServerCertificatesResponse struct {
	value *ReadServerCertificatesResponse
	isSet bool
}

func (v NullableReadServerCertificatesResponse) Get() *ReadServerCertificatesResponse {
	return v.value
}

func (v *NullableReadServerCertificatesResponse) Set(val *ReadServerCertificatesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadServerCertificatesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadServerCertificatesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadServerCertificatesResponse(val *ReadServerCertificatesResponse) *NullableReadServerCertificatesResponse {
	return &NullableReadServerCertificatesResponse{value: val, isSet: true}
}

func (v NullableReadServerCertificatesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadServerCertificatesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
