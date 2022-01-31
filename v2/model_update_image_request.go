/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://docs.outscale.com/en/userguide/OUTSCALE-APIs-Reference.html).<br /><br />  You can also manage your resources using the [Cockpit](https://docs.outscale.com/en/userguide/About-Cockpit.html) web interface.
 *
 * API version: 1.17
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateImageRequest struct for UpdateImageRequest
type UpdateImageRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the OMI you want to modify.
	ImageId             string                        `json:"ImageId"`
	PermissionsToLaunch PermissionsOnResourceCreation `json:"PermissionsToLaunch"`
}

// NewUpdateImageRequest instantiates a new UpdateImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateImageRequest(imageId string, permissionsToLaunch PermissionsOnResourceCreation) *UpdateImageRequest {
	this := UpdateImageRequest{}
	this.ImageId = imageId
	this.PermissionsToLaunch = permissionsToLaunch
	return &this
}

// NewUpdateImageRequestWithDefaults instantiates a new UpdateImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateImageRequestWithDefaults() *UpdateImageRequest {
	this := UpdateImageRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateImageRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateImageRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateImageRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateImageRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetImageId returns the ImageId field value
func (o *UpdateImageRequest) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *UpdateImageRequest) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *UpdateImageRequest) SetImageId(v string) {
	o.ImageId = v
}

// GetPermissionsToLaunch returns the PermissionsToLaunch field value
func (o *UpdateImageRequest) GetPermissionsToLaunch() PermissionsOnResourceCreation {
	if o == nil {
		var ret PermissionsOnResourceCreation
		return ret
	}

	return o.PermissionsToLaunch
}

// GetPermissionsToLaunchOk returns a tuple with the PermissionsToLaunch field value
// and a boolean to check if the value has been set.
func (o *UpdateImageRequest) GetPermissionsToLaunchOk() (*PermissionsOnResourceCreation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionsToLaunch, true
}

// SetPermissionsToLaunch sets field value
func (o *UpdateImageRequest) SetPermissionsToLaunch(v PermissionsOnResourceCreation) {
	o.PermissionsToLaunch = v
}

func (o UpdateImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ImageId"] = o.ImageId
	}
	if true {
		toSerialize["PermissionsToLaunch"] = o.PermissionsToLaunch
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateImageRequest struct {
	value *UpdateImageRequest
	isSet bool
}

func (v NullableUpdateImageRequest) Get() *UpdateImageRequest {
	return v.value
}

func (v *NullableUpdateImageRequest) Set(val *UpdateImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateImageRequest(val *UpdateImageRequest) *NullableUpdateImageRequest {
	return &NullableUpdateImageRequest{value: val, isSet: true}
}

func (v NullableUpdateImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
