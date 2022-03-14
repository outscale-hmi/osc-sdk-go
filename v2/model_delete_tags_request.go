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

// DeleteTagsRequest struct for DeleteTagsRequest
type DeleteTagsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more resource IDs.
	ResourceIds []string `json:"ResourceIds"`
	// One or more tags to delete (if you set a tag value, only the tags matching exactly this value are deleted).
	Tags []ResourceTag `json:"Tags"`
}

// NewDeleteTagsRequest instantiates a new DeleteTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTagsRequest(resourceIds []string, tags []ResourceTag) *DeleteTagsRequest {
	this := DeleteTagsRequest{}
	this.ResourceIds = resourceIds
	this.Tags = tags
	return &this
}

// NewDeleteTagsRequestWithDefaults instantiates a new DeleteTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTagsRequestWithDefaults() *DeleteTagsRequest {
	this := DeleteTagsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteTagsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTagsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteTagsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteTagsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetResourceIds returns the ResourceIds field value
func (o *DeleteTagsRequest) GetResourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *DeleteTagsRequest) GetResourceIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// SetResourceIds sets field value
func (o *DeleteTagsRequest) SetResourceIds(v []string) {
	o.ResourceIds = v
}

// GetTags returns the Tags field value
func (o *DeleteTagsRequest) GetTags() []ResourceTag {
	if o == nil {
		var ret []ResourceTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *DeleteTagsRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *DeleteTagsRequest) SetTags(v []ResourceTag) {
	o.Tags = v
}

func (o DeleteTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["ResourceIds"] = o.ResourceIds
	}
	if true {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteTagsRequest struct {
	value *DeleteTagsRequest
	isSet bool
}

func (v NullableDeleteTagsRequest) Get() *DeleteTagsRequest {
	return v.value
}

func (v *NullableDeleteTagsRequest) Set(val *DeleteTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTagsRequest(val *DeleteTagsRequest) *NullableDeleteTagsRequest {
	return &NullableDeleteTagsRequest{value: val, isSet: true}
}

func (v NullableDeleteTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
