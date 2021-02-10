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

// CreateDirectLinkInterfaceRequest struct for CreateDirectLinkInterfaceRequest
type CreateDirectLinkInterfaceRequest struct {
	// The ID of the existing DirectLink for which you want to create the DirectLink interface.
	DirectLinkId        string              `json:"DirectLinkId"`
	DirectLinkInterface DirectLinkInterface `json:"DirectLinkInterface"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewCreateDirectLinkInterfaceRequest instantiates a new CreateDirectLinkInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDirectLinkInterfaceRequest(directLinkId string, directLinkInterface DirectLinkInterface) *CreateDirectLinkInterfaceRequest {
	this := CreateDirectLinkInterfaceRequest{}
	this.DirectLinkId = directLinkId
	this.DirectLinkInterface = directLinkInterface
	return &this
}

// NewCreateDirectLinkInterfaceRequestWithDefaults instantiates a new CreateDirectLinkInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDirectLinkInterfaceRequestWithDefaults() *CreateDirectLinkInterfaceRequest {
	this := CreateDirectLinkInterfaceRequest{}
	return &this
}

// GetDirectLinkId returns the DirectLinkId field value
func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectLinkId
}

// GetDirectLinkIdOk returns a tuple with the DirectLinkId field value
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkId, true
}

// SetDirectLinkId sets field value
func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkId(v string) {
	o.DirectLinkId = v
}

// GetDirectLinkInterface returns the DirectLinkInterface field value
func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterface() DirectLinkInterface {
	if o == nil {
		var ret DirectLinkInterface
		return ret
	}

	return o.DirectLinkInterface
}

// GetDirectLinkInterfaceOk returns a tuple with the DirectLinkInterface field value
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterfaceOk() (*DirectLinkInterface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkInterface, true
}

// SetDirectLinkInterface sets field value
func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkInterface(v DirectLinkInterface) {
	o.DirectLinkInterface = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateDirectLinkInterfaceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateDirectLinkInterfaceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateDirectLinkInterfaceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o CreateDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DirectLinkId"] = o.DirectLinkId
	}
	if true {
		toSerialize["DirectLinkInterface"] = o.DirectLinkInterface
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDirectLinkInterfaceRequest struct {
	value *CreateDirectLinkInterfaceRequest
	isSet bool
}

func (v NullableCreateDirectLinkInterfaceRequest) Get() *CreateDirectLinkInterfaceRequest {
	return v.value
}

func (v *NullableCreateDirectLinkInterfaceRequest) Set(val *CreateDirectLinkInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDirectLinkInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDirectLinkInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDirectLinkInterfaceRequest(val *CreateDirectLinkInterfaceRequest) *NullableCreateDirectLinkInterfaceRequest {
	return &NullableCreateDirectLinkInterfaceRequest{value: val, isSet: true}
}

func (v NullableCreateDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDirectLinkInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
