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

// UpdateVolumeRequest struct for UpdateVolumeRequest
type UpdateVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The new number of I/O operations per second (IOPS). This parameter can be specified only if you update an `io1` volume. The maximum number of IOPS allowed for `io1` volumes is `13000`. This modification is instantaneous on a cold volume, not on a hot one.
	Iops *int32 `json:"Iops,omitempty"`
	// (cold volume only) The new size of the volume, in gibibytes (GiB). This value must be equal to or greater than the current size of the volume. This modification is not instantaneous.
	Size *int32 `json:"Size,omitempty"`
	// The ID of the volume you want to update.
	VolumeId string `json:"VolumeId"`
	// (cold volume only) The new type of the volume (`standard` \\| `io1` \\| `gpu2`). This modification is instantaneous. If you update to an `io1`volume, you must also specify the `Iops` parameter.
	VolumeType *string `json:"VolumeType,omitempty"`
}

// NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeRequest(volumeId string) *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	this.VolumeId = volumeId
	return &this
}

// NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetIops() int32 {
	if o == nil || o.Iops == nil {
		var ret int32
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetIopsOk() (*int32, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int32 and assigns it to the Iops field.
func (o *UpdateVolumeRequest) SetIops(v int32) {
	o.Iops = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *UpdateVolumeRequest) SetSize(v int32) {
	o.Size = &v
}

// GetVolumeId returns the VolumeId field value
func (o *UpdateVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *UpdateVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *UpdateVolumeRequest) SetVolumeType(v string) {
	o.VolumeType = &v
}

func (o UpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if true {
		toSerialize["VolumeId"] = o.VolumeId
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVolumeRequest struct {
	value *UpdateVolumeRequest
	isSet bool
}

func (v NullableUpdateVolumeRequest) Get() *UpdateVolumeRequest {
	return v.value
}

func (v *NullableUpdateVolumeRequest) Set(val *UpdateVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeRequest(val *UpdateVolumeRequest) *NullableUpdateVolumeRequest {
	return &NullableUpdateVolumeRequest{value: val, isSet: true}
}

func (v NullableUpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
