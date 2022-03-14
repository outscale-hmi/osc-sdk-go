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

// RejectNetPeeringRequest struct for RejectNetPeeringRequest
type RejectNetPeeringRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net peering connection you want to reject.
	NetPeeringId string `json:"NetPeeringId"`
}

// NewRejectNetPeeringRequest instantiates a new RejectNetPeeringRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectNetPeeringRequest(netPeeringId string) *RejectNetPeeringRequest {
	this := RejectNetPeeringRequest{}
	this.NetPeeringId = netPeeringId
	return &this
}

// NewRejectNetPeeringRequestWithDefaults instantiates a new RejectNetPeeringRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectNetPeeringRequestWithDefaults() *RejectNetPeeringRequest {
	this := RejectNetPeeringRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *RejectNetPeeringRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RejectNetPeeringRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *RejectNetPeeringRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *RejectNetPeeringRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetPeeringId returns the NetPeeringId field value
func (o *RejectNetPeeringRequest) GetNetPeeringId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value
// and a boolean to check if the value has been set.
func (o *RejectNetPeeringRequest) GetNetPeeringIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetPeeringId, true
}

// SetNetPeeringId sets field value
func (o *RejectNetPeeringRequest) SetNetPeeringId(v string) {
	o.NetPeeringId = v
}

func (o RejectNetPeeringRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	return json.Marshal(toSerialize)
}

type NullableRejectNetPeeringRequest struct {
	value *RejectNetPeeringRequest
	isSet bool
}

func (v NullableRejectNetPeeringRequest) Get() *RejectNetPeeringRequest {
	return v.value
}

func (v *NullableRejectNetPeeringRequest) Set(val *RejectNetPeeringRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectNetPeeringRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectNetPeeringRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectNetPeeringRequest(val *RejectNetPeeringRequest) *NullableRejectNetPeeringRequest {
	return &NullableRejectNetPeeringRequest{value: val, isSet: true}
}

func (v NullableRejectNetPeeringRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRejectNetPeeringRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
