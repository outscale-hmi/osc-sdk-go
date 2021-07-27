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

// BlockDeviceMappingImage One or more parameters used to automatically set up volumes when the VM is created.
type BlockDeviceMappingImage struct {
	Bsu *BsuToCreate `json:"Bsu,omitempty"`
	// The name of the device.
	DeviceName *string `json:"DeviceName,omitempty"`
	// The name of the virtual device (ephemeralN).
	VirtualDeviceName *string `json:"VirtualDeviceName,omitempty"`
}

// NewBlockDeviceMappingImage instantiates a new BlockDeviceMappingImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDeviceMappingImage() *BlockDeviceMappingImage {
	this := BlockDeviceMappingImage{}
	return &this
}

// NewBlockDeviceMappingImageWithDefaults instantiates a new BlockDeviceMappingImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDeviceMappingImageWithDefaults() *BlockDeviceMappingImage {
	this := BlockDeviceMappingImage{}
	return &this
}

// GetBsu returns the Bsu field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetBsu() BsuToCreate {
	if o == nil || o.Bsu == nil {
		var ret BsuToCreate
		return ret
	}
	return *o.Bsu
}

// GetBsuOk returns a tuple with the Bsu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetBsuOk() (*BsuToCreate, bool) {
	if o == nil || o.Bsu == nil {
		return nil, false
	}
	return o.Bsu, true
}

// HasBsu returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasBsu() bool {
	if o != nil && o.Bsu != nil {
		return true
	}

	return false
}

// SetBsu gets a reference to the given BsuToCreate and assigns it to the Bsu field.
func (o *BlockDeviceMappingImage) SetBsu(v BsuToCreate) {
	o.Bsu = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *BlockDeviceMappingImage) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetVirtualDeviceName returns the VirtualDeviceName field value if set, zero value otherwise.
func (o *BlockDeviceMappingImage) GetVirtualDeviceName() string {
	if o == nil || o.VirtualDeviceName == nil {
		var ret string
		return ret
	}
	return *o.VirtualDeviceName
}

// GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockDeviceMappingImage) GetVirtualDeviceNameOk() (*string, bool) {
	if o == nil || o.VirtualDeviceName == nil {
		return nil, false
	}
	return o.VirtualDeviceName, true
}

// HasVirtualDeviceName returns a boolean if a field has been set.
func (o *BlockDeviceMappingImage) HasVirtualDeviceName() bool {
	if o != nil && o.VirtualDeviceName != nil {
		return true
	}

	return false
}

// SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.
func (o *BlockDeviceMappingImage) SetVirtualDeviceName(v string) {
	o.VirtualDeviceName = &v
}

func (o BlockDeviceMappingImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bsu != nil {
		toSerialize["Bsu"] = o.Bsu
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.VirtualDeviceName != nil {
		toSerialize["VirtualDeviceName"] = o.VirtualDeviceName
	}
	return json.Marshal(toSerialize)
}

type NullableBlockDeviceMappingImage struct {
	value *BlockDeviceMappingImage
	isSet bool
}

func (v NullableBlockDeviceMappingImage) Get() *BlockDeviceMappingImage {
	return v.value
}

func (v *NullableBlockDeviceMappingImage) Set(val *BlockDeviceMappingImage) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDeviceMappingImage) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDeviceMappingImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDeviceMappingImage(val *BlockDeviceMappingImage) *NullableBlockDeviceMappingImage {
	return &NullableBlockDeviceMappingImage{value: val, isSet: true}
}

func (v NullableBlockDeviceMappingImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDeviceMappingImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
