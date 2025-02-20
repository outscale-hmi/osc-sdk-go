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

// DeleteSnapshotRequest struct for DeleteSnapshotRequest
type DeleteSnapshotRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the snapshot you want to delete.
	SnapshotId string `json:"SnapshotId"`
}

// NewDeleteSnapshotRequest instantiates a new DeleteSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSnapshotRequest(snapshotId string) *DeleteSnapshotRequest {
	this := DeleteSnapshotRequest{}
	this.SnapshotId = snapshotId
	return &this
}

// NewDeleteSnapshotRequestWithDefaults instantiates a new DeleteSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSnapshotRequestWithDefaults() *DeleteSnapshotRequest {
	this := DeleteSnapshotRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteSnapshotRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSnapshotRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteSnapshotRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteSnapshotRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetSnapshotId returns the SnapshotId field value
func (o *DeleteSnapshotRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *DeleteSnapshotRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *DeleteSnapshotRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

func (o DeleteSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSnapshotRequest struct {
	value *DeleteSnapshotRequest
	isSet bool
}

func (v NullableDeleteSnapshotRequest) Get() *DeleteSnapshotRequest {
	return v.value
}

func (v *NullableDeleteSnapshotRequest) Set(val *DeleteSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSnapshotRequest(val *DeleteSnapshotRequest) *NullableDeleteSnapshotRequest {
	return &NullableDeleteSnapshotRequest{value: val, isSet: true}
}

func (v NullableDeleteSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
