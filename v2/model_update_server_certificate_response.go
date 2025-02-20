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

// UpdateServerCertificateResponse struct for UpdateServerCertificateResponse
type UpdateServerCertificateResponse struct {
	ResponseContext   *ResponseContext   `json:"ResponseContext,omitempty"`
	ServerCertificate *ServerCertificate `json:"ServerCertificate,omitempty"`
}

// NewUpdateServerCertificateResponse instantiates a new UpdateServerCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerCertificateResponse() *UpdateServerCertificateResponse {
	this := UpdateServerCertificateResponse{}
	return &this
}

// NewUpdateServerCertificateResponseWithDefaults instantiates a new UpdateServerCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerCertificateResponseWithDefaults() *UpdateServerCertificateResponse {
	this := UpdateServerCertificateResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateServerCertificateResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateServerCertificateResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateServerCertificateResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetServerCertificate returns the ServerCertificate field value if set, zero value otherwise.
func (o *UpdateServerCertificateResponse) GetServerCertificate() ServerCertificate {
	if o == nil || o.ServerCertificate == nil {
		var ret ServerCertificate
		return ret
	}
	return *o.ServerCertificate
}

// GetServerCertificateOk returns a tuple with the ServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateResponse) GetServerCertificateOk() (*ServerCertificate, bool) {
	if o == nil || o.ServerCertificate == nil {
		return nil, false
	}
	return o.ServerCertificate, true
}

// HasServerCertificate returns a boolean if a field has been set.
func (o *UpdateServerCertificateResponse) HasServerCertificate() bool {
	if o != nil && o.ServerCertificate != nil {
		return true
	}

	return false
}

// SetServerCertificate gets a reference to the given ServerCertificate and assigns it to the ServerCertificate field.
func (o *UpdateServerCertificateResponse) SetServerCertificate(v ServerCertificate) {
	o.ServerCertificate = &v
}

func (o UpdateServerCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.ServerCertificate != nil {
		toSerialize["ServerCertificate"] = o.ServerCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateServerCertificateResponse struct {
	value *UpdateServerCertificateResponse
	isSet bool
}

func (v NullableUpdateServerCertificateResponse) Get() *UpdateServerCertificateResponse {
	return v.value
}

func (v *NullableUpdateServerCertificateResponse) Set(val *UpdateServerCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerCertificateResponse(val *UpdateServerCertificateResponse) *NullableUpdateServerCertificateResponse {
	return &NullableUpdateServerCertificateResponse{value: val, isSet: true}
}

func (v NullableUpdateServerCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
