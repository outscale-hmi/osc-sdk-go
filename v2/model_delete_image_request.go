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

// DeleteImageRequest struct for DeleteImageRequest
type DeleteImageRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the OMI you want to delete.
	ImageId string `json:"ImageId"`
}

// NewDeleteImageRequest instantiates a new DeleteImageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteImageRequest(imageId string) *DeleteImageRequest {
	this := DeleteImageRequest{}
	this.ImageId = imageId
	return &this
}

// NewDeleteImageRequestWithDefaults instantiates a new DeleteImageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteImageRequestWithDefaults() *DeleteImageRequest {
	this := DeleteImageRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteImageRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteImageRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteImageRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteImageRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetImageId returns the ImageId field value
func (o *DeleteImageRequest) GetImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value
// and a boolean to check if the value has been set.
func (o *DeleteImageRequest) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageId, true
}

// SetImageId sets field value
func (o *DeleteImageRequest) SetImageId(v string) {
	o.ImageId = v
}

func (o DeleteImageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ImageId"] = o.ImageId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteImageRequest struct {
	value *DeleteImageRequest
	isSet bool
}

func (v NullableDeleteImageRequest) Get() *DeleteImageRequest {
	return v.value
}

func (v *NullableDeleteImageRequest) Set(val *DeleteImageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteImageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteImageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteImageRequest(val *DeleteImageRequest) *NullableDeleteImageRequest {
	return &NullableDeleteImageRequest{value: val, isSet: true}
}

func (v NullableDeleteImageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteImageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
