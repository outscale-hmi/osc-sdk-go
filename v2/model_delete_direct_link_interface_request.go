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

// DeleteDirectLinkInterfaceRequest struct for DeleteDirectLinkInterfaceRequest
type DeleteDirectLinkInterfaceRequest struct {
	// The ID of the DirectLink interface you want to delete.
	DirectLinkInterfaceId string `json:"DirectLinkInterfaceId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewDeleteDirectLinkInterfaceRequest instantiates a new DeleteDirectLinkInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDirectLinkInterfaceRequest(directLinkInterfaceId string) *DeleteDirectLinkInterfaceRequest {
	this := DeleteDirectLinkInterfaceRequest{}
	this.DirectLinkInterfaceId = directLinkInterfaceId
	return &this
}

// NewDeleteDirectLinkInterfaceRequestWithDefaults instantiates a new DeleteDirectLinkInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDirectLinkInterfaceRequestWithDefaults() *DeleteDirectLinkInterfaceRequest {
	this := DeleteDirectLinkInterfaceRequest{}
	return &this
}

// GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field value
func (o *DeleteDirectLinkInterfaceRequest) GetDirectLinkInterfaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DirectLinkInterfaceId
}

// GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field value
// and a boolean to check if the value has been set.
func (o *DeleteDirectLinkInterfaceRequest) GetDirectLinkInterfaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectLinkInterfaceId, true
}

// SetDirectLinkInterfaceId sets field value
func (o *DeleteDirectLinkInterfaceRequest) SetDirectLinkInterfaceId(v string) {
	o.DirectLinkInterfaceId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteDirectLinkInterfaceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDirectLinkInterfaceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteDirectLinkInterfaceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteDirectLinkInterfaceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o DeleteDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DirectLinkInterfaceId"] = o.DirectLinkInterfaceId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDirectLinkInterfaceRequest struct {
	value *DeleteDirectLinkInterfaceRequest
	isSet bool
}

func (v NullableDeleteDirectLinkInterfaceRequest) Get() *DeleteDirectLinkInterfaceRequest {
	return v.value
}

func (v *NullableDeleteDirectLinkInterfaceRequest) Set(val *DeleteDirectLinkInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDirectLinkInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDirectLinkInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDirectLinkInterfaceRequest(val *DeleteDirectLinkInterfaceRequest) *NullableDeleteDirectLinkInterfaceRequest {
	return &NullableDeleteDirectLinkInterfaceRequest{value: val, isSet: true}
}

func (v NullableDeleteDirectLinkInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDirectLinkInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
