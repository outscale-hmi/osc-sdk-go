/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateNetPeeringRequest struct for CreateNetPeeringRequest
type CreateNetPeeringRequest struct {
	// The ID of the Net you want to connect with.
	AccepterNetId string `json:"AccepterNetId"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net you send the peering request from.
	SourceNetId string `json:"SourceNetId"`
}

// NewCreateNetPeeringRequest instantiates a new CreateNetPeeringRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetPeeringRequest(accepterNetId string, sourceNetId string) *CreateNetPeeringRequest {
	this := CreateNetPeeringRequest{}
	this.AccepterNetId = accepterNetId
	this.SourceNetId = sourceNetId
	return &this
}

// NewCreateNetPeeringRequestWithDefaults instantiates a new CreateNetPeeringRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetPeeringRequestWithDefaults() *CreateNetPeeringRequest {
	this := CreateNetPeeringRequest{}
	return &this
}

// GetAccepterNetId returns the AccepterNetId field value
func (o *CreateNetPeeringRequest) GetAccepterNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccepterNetId
}

// GetAccepterNetIdOk returns a tuple with the AccepterNetId field value
// and a boolean to check if the value has been set.
func (o *CreateNetPeeringRequest) GetAccepterNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccepterNetId, true
}

// SetAccepterNetId sets field value
func (o *CreateNetPeeringRequest) SetAccepterNetId(v string) {
	o.AccepterNetId = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNetPeeringRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetPeeringRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNetPeeringRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNetPeeringRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetSourceNetId returns the SourceNetId field value
func (o *CreateNetPeeringRequest) GetSourceNetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceNetId
}

// GetSourceNetIdOk returns a tuple with the SourceNetId field value
// and a boolean to check if the value has been set.
func (o *CreateNetPeeringRequest) GetSourceNetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceNetId, true
}

// SetSourceNetId sets field value
func (o *CreateNetPeeringRequest) SetSourceNetId(v string) {
	o.SourceNetId = v
}

func (o CreateNetPeeringRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AccepterNetId"] = o.AccepterNetId
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["SourceNetId"] = o.SourceNetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetPeeringRequest struct {
	value *CreateNetPeeringRequest
	isSet bool
}

func (v NullableCreateNetPeeringRequest) Get() *CreateNetPeeringRequest {
	return v.value
}

func (v *NullableCreateNetPeeringRequest) Set(val *CreateNetPeeringRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetPeeringRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetPeeringRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetPeeringRequest(val *CreateNetPeeringRequest) *NullableCreateNetPeeringRequest {
	return &NullableCreateNetPeeringRequest{value: val, isSet: true}
}

func (v NullableCreateNetPeeringRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetPeeringRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
