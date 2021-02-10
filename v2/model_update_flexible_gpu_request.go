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

// UpdateFlexibleGpuRequest struct for UpdateFlexibleGpuRequest
type UpdateFlexibleGpuRequest struct {
	// If `true`, the fGPU is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the fGPU you want to modify.
	FlexibleGpuId string `json:"FlexibleGpuId"`
}

// NewUpdateFlexibleGpuRequest instantiates a new UpdateFlexibleGpuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFlexibleGpuRequest(flexibleGpuId string) *UpdateFlexibleGpuRequest {
	this := UpdateFlexibleGpuRequest{}
	this.FlexibleGpuId = flexibleGpuId
	return &this
}

// NewUpdateFlexibleGpuRequestWithDefaults instantiates a new UpdateFlexibleGpuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFlexibleGpuRequestWithDefaults() *UpdateFlexibleGpuRequest {
	this := UpdateFlexibleGpuRequest{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *UpdateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlexibleGpuId returns the FlexibleGpuId field value
func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlexibleGpuId
}

// GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field value
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlexibleGpuId, true
}

// SetFlexibleGpuId sets field value
func (o *UpdateFlexibleGpuRequest) SetFlexibleGpuId(v string) {
	o.FlexibleGpuId = v
}

func (o UpdateFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FlexibleGpuId"] = o.FlexibleGpuId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFlexibleGpuRequest struct {
	value *UpdateFlexibleGpuRequest
	isSet bool
}

func (v NullableUpdateFlexibleGpuRequest) Get() *UpdateFlexibleGpuRequest {
	return v.value
}

func (v *NullableUpdateFlexibleGpuRequest) Set(val *UpdateFlexibleGpuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFlexibleGpuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFlexibleGpuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFlexibleGpuRequest(val *UpdateFlexibleGpuRequest) *NullableUpdateFlexibleGpuRequest {
	return &NullableUpdateFlexibleGpuRequest{value: val, isSet: true}
}

func (v NullableUpdateFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
