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

// LinkNicRequest struct for LinkNicRequest
type LinkNicRequest struct {
	// The index of the VM device for the NIC attachment (between 1 and 7, both included).
	DeviceNumber int32 `json:"DeviceNumber"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the NIC you want to attach.
	NicId string `json:"NicId"`
	// The ID of the VM to which you want to attach the NIC.
	VmId string `json:"VmId"`
}

// NewLinkNicRequest instantiates a new LinkNicRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkNicRequest(deviceNumber int32, nicId string, vmId string) *LinkNicRequest {
	this := LinkNicRequest{}
	this.DeviceNumber = deviceNumber
	this.NicId = nicId
	this.VmId = vmId
	return &this
}

// NewLinkNicRequestWithDefaults instantiates a new LinkNicRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkNicRequestWithDefaults() *LinkNicRequest {
	this := LinkNicRequest{}
	return &this
}

// GetDeviceNumber returns the DeviceNumber field value
func (o *LinkNicRequest) GetDeviceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value
// and a boolean to check if the value has been set.
func (o *LinkNicRequest) GetDeviceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceNumber, true
}

// SetDeviceNumber sets field value
func (o *LinkNicRequest) SetDeviceNumber(v int32) {
	o.DeviceNumber = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkNicRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkNicRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkNicRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value
func (o *LinkNicRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value
// and a boolean to check if the value has been set.
func (o *LinkNicRequest) GetNicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NicId, true
}

// SetNicId sets field value
func (o *LinkNicRequest) SetNicId(v string) {
	o.NicId = v
}

// GetVmId returns the VmId field value
func (o *LinkNicRequest) GetVmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value
// and a boolean to check if the value has been set.
func (o *LinkNicRequest) GetVmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmId, true
}

// SetVmId sets field value
func (o *LinkNicRequest) SetVmId(v string) {
	o.VmId = v
}

func (o LinkNicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DeviceNumber"] = o.DeviceNumber
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["NicId"] = o.NicId
	}
	if true {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkNicRequest struct {
	value *LinkNicRequest
	isSet bool
}

func (v NullableLinkNicRequest) Get() *LinkNicRequest {
	return v.value
}

func (v *NullableLinkNicRequest) Set(val *LinkNicRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkNicRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkNicRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkNicRequest(val *LinkNicRequest) *NullableLinkNicRequest {
	return &NullableLinkNicRequest{value: val, isSet: true}
}

func (v NullableLinkNicRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkNicRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
