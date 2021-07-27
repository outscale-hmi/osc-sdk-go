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

// DeleteSecurityGroupRequest struct for DeleteSecurityGroupRequest
type DeleteSecurityGroupRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the security group you want to delete.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
}

// NewDeleteSecurityGroupRequest instantiates a new DeleteSecurityGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSecurityGroupRequest() *DeleteSecurityGroupRequest {
	this := DeleteSecurityGroupRequest{}
	return &this
}

// NewDeleteSecurityGroupRequestWithDefaults instantiates a new DeleteSecurityGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSecurityGroupRequestWithDefaults() *DeleteSecurityGroupRequest {
	this := DeleteSecurityGroupRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteSecurityGroupRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRequest) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRequest) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRequest) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *DeleteSecurityGroupRequest) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSecurityGroupRequest) GetSecurityGroupNameOk() (*string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		return nil, false
	}
	return o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *DeleteSecurityGroupRequest) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *DeleteSecurityGroupRequest) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

func (o DeleteSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.SecurityGroupId != nil {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupName != nil {
		toSerialize["SecurityGroupName"] = o.SecurityGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSecurityGroupRequest struct {
	value *DeleteSecurityGroupRequest
	isSet bool
}

func (v NullableDeleteSecurityGroupRequest) Get() *DeleteSecurityGroupRequest {
	return v.value
}

func (v *NullableDeleteSecurityGroupRequest) Set(val *DeleteSecurityGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSecurityGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSecurityGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSecurityGroupRequest(val *DeleteSecurityGroupRequest) *NullableDeleteSecurityGroupRequest {
	return &NullableDeleteSecurityGroupRequest{value: val, isSet: true}
}

func (v NullableDeleteSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSecurityGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
