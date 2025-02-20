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

// BsuToCreate Information about the BSU volume to create.
type BsuToCreate struct {
	// By default or if set to true, the volume is deleted when terminating the VM. If false, the volume is not deleted when terminating the VM.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The number of I/O operations per second (IOPS). This parameter must be specified only if you create an `io1` volume. The maximum number of IOPS allowed for `io1` volumes is `13000`.
	Iops *int32 `json:"Iops,omitempty"`
	// The ID of the snapshot used to create the volume.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The size of the volume, in gibibytes (GiB).<br /> If you specify a snapshot ID, the volume size must be at least equal to the snapshot size.<br /> If you specify a snapshot ID but no volume size, the volume is created with a size similar to the snapshot one.
	VolumeSize *int32 `json:"VolumeSize,omitempty"`
	// The type of the volume (`standard` \\| `io1` \\| `gp2`). If not specified in the request, a `standard` volume is created.<br /> For more information about volume types, see [About Volumes > Volume Types and IOPS](https://docs.outscale.com/en/userguide/About-Volumes.html#_volume_types_and_iops).
	VolumeType *string `json:"VolumeType,omitempty"`
}

// NewBsuToCreate instantiates a new BsuToCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsuToCreate() *BsuToCreate {
	this := BsuToCreate{}
	var deleteOnVmDeletion bool = true
	this.DeleteOnVmDeletion = &deleteOnVmDeletion
	return &this
}

// NewBsuToCreateWithDefaults instantiates a new BsuToCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsuToCreateWithDefaults() *BsuToCreate {
	this := BsuToCreate{}
	var deleteOnVmDeletion bool = true
	this.DeleteOnVmDeletion = &deleteOnVmDeletion
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *BsuToCreate) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *BsuToCreate) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *BsuToCreate) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *BsuToCreate) GetIops() int32 {
	if o == nil || o.Iops == nil {
		var ret int32
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetIopsOk() (*int32, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *BsuToCreate) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int32 and assigns it to the Iops field.
func (o *BsuToCreate) SetIops(v int32) {
	o.Iops = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *BsuToCreate) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *BsuToCreate) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *BsuToCreate) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *BsuToCreate) GetVolumeSize() int32 {
	if o == nil || o.VolumeSize == nil {
		var ret int32
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetVolumeSizeOk() (*int32, bool) {
	if o == nil || o.VolumeSize == nil {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *BsuToCreate) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int32 and assigns it to the VolumeSize field.
func (o *BsuToCreate) SetVolumeSize(v int32) {
	o.VolumeSize = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *BsuToCreate) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsuToCreate) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *BsuToCreate) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *BsuToCreate) SetVolumeType(v string) {
	o.VolumeType = &v
}

func (o BsuToCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.SnapshotId != nil {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	if o.VolumeSize != nil {
		toSerialize["VolumeSize"] = o.VolumeSize
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	return json.Marshal(toSerialize)
}

type NullableBsuToCreate struct {
	value *BsuToCreate
	isSet bool
}

func (v NullableBsuToCreate) Get() *BsuToCreate {
	return v.value
}

func (v *NullableBsuToCreate) Set(val *BsuToCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBsuToCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBsuToCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsuToCreate(val *BsuToCreate) *NullableBsuToCreate {
	return &NullableBsuToCreate{value: val, isSet: true}
}

func (v NullableBsuToCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsuToCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
