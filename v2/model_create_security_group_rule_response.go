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

// CreateSecurityGroupRuleResponse struct for CreateSecurityGroupRuleResponse
type CreateSecurityGroupRuleResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	SecurityGroup   *SecurityGroup   `json:"SecurityGroup,omitempty"`
}

// NewCreateSecurityGroupRuleResponse instantiates a new CreateSecurityGroupRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityGroupRuleResponse() *CreateSecurityGroupRuleResponse {
	this := CreateSecurityGroupRuleResponse{}
	return &this
}

// NewCreateSecurityGroupRuleResponseWithDefaults instantiates a new CreateSecurityGroupRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityGroupRuleResponseWithDefaults() *CreateSecurityGroupRuleResponse {
	this := CreateSecurityGroupRuleResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateSecurityGroupRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSecurityGroup returns the SecurityGroup field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleResponse) GetSecurityGroup() SecurityGroup {
	if o == nil || o.SecurityGroup == nil {
		var ret SecurityGroup
		return ret
	}
	return *o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleResponse) GetSecurityGroupOk() (*SecurityGroup, bool) {
	if o == nil || o.SecurityGroup == nil {
		return nil, false
	}
	return o.SecurityGroup, true
}

// HasSecurityGroup returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleResponse) HasSecurityGroup() bool {
	if o != nil && o.SecurityGroup != nil {
		return true
	}

	return false
}

// SetSecurityGroup gets a reference to the given SecurityGroup and assigns it to the SecurityGroup field.
func (o *CreateSecurityGroupRuleResponse) SetSecurityGroup(v SecurityGroup) {
	o.SecurityGroup = &v
}

func (o CreateSecurityGroupRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.SecurityGroup != nil {
		toSerialize["SecurityGroup"] = o.SecurityGroup
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSecurityGroupRuleResponse struct {
	value *CreateSecurityGroupRuleResponse
	isSet bool
}

func (v NullableCreateSecurityGroupRuleResponse) Get() *CreateSecurityGroupRuleResponse {
	return v.value
}

func (v *NullableCreateSecurityGroupRuleResponse) Set(val *CreateSecurityGroupRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityGroupRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityGroupRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityGroupRuleResponse(val *CreateSecurityGroupRuleResponse) *NullableCreateSecurityGroupRuleResponse {
	return &NullableCreateSecurityGroupRuleResponse{value: val, isSet: true}
}

func (v NullableCreateSecurityGroupRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityGroupRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
