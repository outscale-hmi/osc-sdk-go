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

// SecurityGroupsMember Information about the member of a security group.
type SecurityGroupsMember struct {
	// The account ID of a user.
	AccountId *string `json:"AccountId,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
}

// NewSecurityGroupsMember instantiates a new SecurityGroupsMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupsMember() *SecurityGroupsMember {
	this := SecurityGroupsMember{}
	return &this
}

// NewSecurityGroupsMemberWithDefaults instantiates a new SecurityGroupsMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupsMemberWithDefaults() *SecurityGroupsMember {
	this := SecurityGroupsMember{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SecurityGroupsMember) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroupsMember) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetSecurityGroupNameOk() (*string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		return nil, false
	}
	return o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *SecurityGroupsMember) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

func (o SecurityGroupsMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.SecurityGroupId != nil {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupName != nil {
		toSerialize["SecurityGroupName"] = o.SecurityGroupName
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroupsMember struct {
	value *SecurityGroupsMember
	isSet bool
}

func (v NullableSecurityGroupsMember) Get() *SecurityGroupsMember {
	return v.value
}

func (v *NullableSecurityGroupsMember) Set(val *SecurityGroupsMember) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupsMember) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupsMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupsMember(val *SecurityGroupsMember) *NullableSecurityGroupsMember {
	return &NullableSecurityGroupsMember{value: val, isSet: true}
}

func (v NullableSecurityGroupsMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupsMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
