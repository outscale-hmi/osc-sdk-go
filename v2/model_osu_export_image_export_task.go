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

// OsuExportImageExportTask Information about the OMI export task.
type OsuExportImageExportTask struct {
	// The format of the export disk (`qcow2` \\| `raw`).
	DiskImageFormat string `json:"DiskImageFormat"`
	// The name of the OOS bucket the OMI is exported to.
	OsuBucket string `json:"OsuBucket"`
	// The URL of the manifest file.
	OsuManifestUrl *string `json:"OsuManifestUrl,omitempty"`
	// The prefix for the key of the OOS object corresponding to the image.
	OsuPrefix *string `json:"OsuPrefix,omitempty"`
}

// NewOsuExportImageExportTask instantiates a new OsuExportImageExportTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsuExportImageExportTask(diskImageFormat string, osuBucket string) *OsuExportImageExportTask {
	this := OsuExportImageExportTask{}
	this.DiskImageFormat = diskImageFormat
	this.OsuBucket = osuBucket
	return &this
}

// NewOsuExportImageExportTaskWithDefaults instantiates a new OsuExportImageExportTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsuExportImageExportTaskWithDefaults() *OsuExportImageExportTask {
	this := OsuExportImageExportTask{}
	return &this
}

// GetDiskImageFormat returns the DiskImageFormat field value
func (o *OsuExportImageExportTask) GetDiskImageFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskImageFormat
}

// GetDiskImageFormatOk returns a tuple with the DiskImageFormat field value
// and a boolean to check if the value has been set.
func (o *OsuExportImageExportTask) GetDiskImageFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskImageFormat, true
}

// SetDiskImageFormat sets field value
func (o *OsuExportImageExportTask) SetDiskImageFormat(v string) {
	o.DiskImageFormat = v
}

// GetOsuBucket returns the OsuBucket field value
func (o *OsuExportImageExportTask) GetOsuBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsuBucket
}

// GetOsuBucketOk returns a tuple with the OsuBucket field value
// and a boolean to check if the value has been set.
func (o *OsuExportImageExportTask) GetOsuBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsuBucket, true
}

// SetOsuBucket sets field value
func (o *OsuExportImageExportTask) SetOsuBucket(v string) {
	o.OsuBucket = v
}

// GetOsuManifestUrl returns the OsuManifestUrl field value if set, zero value otherwise.
func (o *OsuExportImageExportTask) GetOsuManifestUrl() string {
	if o == nil || o.OsuManifestUrl == nil {
		var ret string
		return ret
	}
	return *o.OsuManifestUrl
}

// GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportImageExportTask) GetOsuManifestUrlOk() (*string, bool) {
	if o == nil || o.OsuManifestUrl == nil {
		return nil, false
	}
	return o.OsuManifestUrl, true
}

// HasOsuManifestUrl returns a boolean if a field has been set.
func (o *OsuExportImageExportTask) HasOsuManifestUrl() bool {
	if o != nil && o.OsuManifestUrl != nil {
		return true
	}

	return false
}

// SetOsuManifestUrl gets a reference to the given string and assigns it to the OsuManifestUrl field.
func (o *OsuExportImageExportTask) SetOsuManifestUrl(v string) {
	o.OsuManifestUrl = &v
}

// GetOsuPrefix returns the OsuPrefix field value if set, zero value otherwise.
func (o *OsuExportImageExportTask) GetOsuPrefix() string {
	if o == nil || o.OsuPrefix == nil {
		var ret string
		return ret
	}
	return *o.OsuPrefix
}

// GetOsuPrefixOk returns a tuple with the OsuPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsuExportImageExportTask) GetOsuPrefixOk() (*string, bool) {
	if o == nil || o.OsuPrefix == nil {
		return nil, false
	}
	return o.OsuPrefix, true
}

// HasOsuPrefix returns a boolean if a field has been set.
func (o *OsuExportImageExportTask) HasOsuPrefix() bool {
	if o != nil && o.OsuPrefix != nil {
		return true
	}

	return false
}

// SetOsuPrefix gets a reference to the given string and assigns it to the OsuPrefix field.
func (o *OsuExportImageExportTask) SetOsuPrefix(v string) {
	o.OsuPrefix = &v
}

func (o OsuExportImageExportTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DiskImageFormat"] = o.DiskImageFormat
	}
	if true {
		toSerialize["OsuBucket"] = o.OsuBucket
	}
	if o.OsuManifestUrl != nil {
		toSerialize["OsuManifestUrl"] = o.OsuManifestUrl
	}
	if o.OsuPrefix != nil {
		toSerialize["OsuPrefix"] = o.OsuPrefix
	}
	return json.Marshal(toSerialize)
}

type NullableOsuExportImageExportTask struct {
	value *OsuExportImageExportTask
	isSet bool
}

func (v NullableOsuExportImageExportTask) Get() *OsuExportImageExportTask {
	return v.value
}

func (v *NullableOsuExportImageExportTask) Set(val *OsuExportImageExportTask) {
	v.value = val
	v.isSet = true
}

func (v NullableOsuExportImageExportTask) IsSet() bool {
	return v.isSet
}

func (v *NullableOsuExportImageExportTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsuExportImageExportTask(val *OsuExportImageExportTask) *NullableOsuExportImageExportTask {
	return &NullableOsuExportImageExportTask{value: val, isSet: true}
}

func (v NullableOsuExportImageExportTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsuExportImageExportTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
