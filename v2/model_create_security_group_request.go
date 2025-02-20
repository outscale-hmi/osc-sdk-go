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

// CreateSecurityGroupRequest struct for CreateSecurityGroupRequest
type CreateSecurityGroupRequest struct {
	// A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters).
	Description string `json:"Description"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net for the security group.
	NetId *string `json:"NetId,omitempty"`
	// The name of the security group.<br /> This name must not start with `sg-`.</br> This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed.
	SecurityGroupName string `json:"SecurityGroupName"`
}

// NewCreateSecurityGroupRequest instantiates a new CreateSecurityGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityGroupRequest(description string, securityGroupName string) *CreateSecurityGroupRequest {
	this := CreateSecurityGroupRequest{}
	this.Description = description
	this.SecurityGroupName = securityGroupName
	return &this
}

// NewCreateSecurityGroupRequestWithDefaults instantiates a new CreateSecurityGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityGroupRequestWithDefaults() *CreateSecurityGroupRequest {
	this := CreateSecurityGroupRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateSecurityGroupRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateSecurityGroupRequest) SetDescription(v string) {
	o.Description = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateSecurityGroupRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateSecurityGroupRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateSecurityGroupRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *CreateSecurityGroupRequest) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetNetIdOk() (*string, bool) {
	if o == nil || o.NetId == nil {
		return nil, false
	}
	return o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *CreateSecurityGroupRequest) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *CreateSecurityGroupRequest) SetNetId(v string) {
	o.NetId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value
func (o *CreateSecurityGroupRequest) GetSecurityGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRequest) GetSecurityGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecurityGroupName, true
}

// SetSecurityGroupName sets field value
func (o *CreateSecurityGroupRequest) SetSecurityGroupName(v string) {
	o.SecurityGroupName = v
}

func (o CreateSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.NetId != nil {
		toSerialize["NetId"] = o.NetId
	}
	if true {
		toSerialize["SecurityGroupName"] = o.SecurityGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSecurityGroupRequest struct {
	value *CreateSecurityGroupRequest
	isSet bool
}

func (v NullableCreateSecurityGroupRequest) Get() *CreateSecurityGroupRequest {
	return v.value
}

func (v *NullableCreateSecurityGroupRequest) Set(val *CreateSecurityGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityGroupRequest(val *CreateSecurityGroupRequest) *NullableCreateSecurityGroupRequest {
	return &NullableCreateSecurityGroupRequest{value: val, isSet: true}
}

func (v NullableCreateSecurityGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
