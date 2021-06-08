/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateInternetServiceResponse struct for CreateInternetServiceResponse
type CreateInternetServiceResponse struct {
	InternetService *InternetService `json:"InternetService,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateInternetServiceResponse instantiates a new CreateInternetServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInternetServiceResponse() *CreateInternetServiceResponse {
	this := CreateInternetServiceResponse{}
	return &this
}

// NewCreateInternetServiceResponseWithDefaults instantiates a new CreateInternetServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInternetServiceResponseWithDefaults() *CreateInternetServiceResponse {
	this := CreateInternetServiceResponse{}
	return &this
}

// GetInternetService returns the InternetService field value if set, zero value otherwise.
func (o *CreateInternetServiceResponse) GetInternetService() InternetService {
	if o == nil || o.InternetService == nil {
		var ret InternetService
		return ret
	}
	return *o.InternetService
}

// GetInternetServiceOk returns a tuple with the InternetService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInternetServiceResponse) GetInternetServiceOk() (*InternetService, bool) {
	if o == nil || o.InternetService == nil {
		return nil, false
	}
	return o.InternetService, true
}

// HasInternetService returns a boolean if a field has been set.
func (o *CreateInternetServiceResponse) HasInternetService() bool {
	if o != nil && o.InternetService != nil {
		return true
	}

	return false
}

// SetInternetService gets a reference to the given InternetService and assigns it to the InternetService field.
func (o *CreateInternetServiceResponse) SetInternetService(v InternetService) {
	o.InternetService = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateInternetServiceResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInternetServiceResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateInternetServiceResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateInternetServiceResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateInternetServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InternetService != nil {
		toSerialize["InternetService"] = o.InternetService
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateInternetServiceResponse struct {
	value *CreateInternetServiceResponse
	isSet bool
}

func (v NullableCreateInternetServiceResponse) Get() *CreateInternetServiceResponse {
	return v.value
}

func (v *NullableCreateInternetServiceResponse) Set(val *CreateInternetServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInternetServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInternetServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInternetServiceResponse(val *CreateInternetServiceResponse) *NullableCreateInternetServiceResponse {
	return &NullableCreateInternetServiceResponse{value: val, isSet: true}
}

func (v NullableCreateInternetServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInternetServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
