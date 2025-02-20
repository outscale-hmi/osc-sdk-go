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

// ApiAccessRule Information about the API access rule.
type ApiAccessRule struct {
	//  The ID of the API access rule.
	ApiAccessRuleId *string `json:"ApiAccessRuleId,omitempty"`
	// One or more IDs of Client Certificate Authorities (CAs) used for the API access rule.
	CaIds *[]string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs).
	Cns *[]string `json:"Cns,omitempty"`
	// The description of the API access rule.
	Description *string `json:"Description,omitempty"`
	// One or more IP ranges used for the API access rule, in CIDR notation (for example, `192.0.2.0/16`).
	IpRanges *[]string `json:"IpRanges,omitempty"`
}

// NewApiAccessRule instantiates a new ApiAccessRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccessRule() *ApiAccessRule {
	this := ApiAccessRule{}
	return &this
}

// NewApiAccessRuleWithDefaults instantiates a new ApiAccessRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccessRuleWithDefaults() *ApiAccessRule {
	this := ApiAccessRule{}
	return &this
}

// GetApiAccessRuleId returns the ApiAccessRuleId field value if set, zero value otherwise.
func (o *ApiAccessRule) GetApiAccessRuleId() string {
	if o == nil || o.ApiAccessRuleId == nil {
		var ret string
		return ret
	}
	return *o.ApiAccessRuleId
}

// GetApiAccessRuleIdOk returns a tuple with the ApiAccessRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessRule) GetApiAccessRuleIdOk() (*string, bool) {
	if o == nil || o.ApiAccessRuleId == nil {
		return nil, false
	}
	return o.ApiAccessRuleId, true
}

// HasApiAccessRuleId returns a boolean if a field has been set.
func (o *ApiAccessRule) HasApiAccessRuleId() bool {
	if o != nil && o.ApiAccessRuleId != nil {
		return true
	}

	return false
}

// SetApiAccessRuleId gets a reference to the given string and assigns it to the ApiAccessRuleId field.
func (o *ApiAccessRule) SetApiAccessRuleId(v string) {
	o.ApiAccessRuleId = &v
}

// GetCaIds returns the CaIds field value if set, zero value otherwise.
func (o *ApiAccessRule) GetCaIds() []string {
	if o == nil || o.CaIds == nil {
		var ret []string
		return ret
	}
	return *o.CaIds
}

// GetCaIdsOk returns a tuple with the CaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessRule) GetCaIdsOk() (*[]string, bool) {
	if o == nil || o.CaIds == nil {
		return nil, false
	}
	return o.CaIds, true
}

// HasCaIds returns a boolean if a field has been set.
func (o *ApiAccessRule) HasCaIds() bool {
	if o != nil && o.CaIds != nil {
		return true
	}

	return false
}

// SetCaIds gets a reference to the given []string and assigns it to the CaIds field.
func (o *ApiAccessRule) SetCaIds(v []string) {
	o.CaIds = &v
}

// GetCns returns the Cns field value if set, zero value otherwise.
func (o *ApiAccessRule) GetCns() []string {
	if o == nil || o.Cns == nil {
		var ret []string
		return ret
	}
	return *o.Cns
}

// GetCnsOk returns a tuple with the Cns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessRule) GetCnsOk() (*[]string, bool) {
	if o == nil || o.Cns == nil {
		return nil, false
	}
	return o.Cns, true
}

// HasCns returns a boolean if a field has been set.
func (o *ApiAccessRule) HasCns() bool {
	if o != nil && o.Cns != nil {
		return true
	}

	return false
}

// SetCns gets a reference to the given []string and assigns it to the Cns field.
func (o *ApiAccessRule) SetCns(v []string) {
	o.Cns = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiAccessRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAccessRule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAccessRule) SetDescription(v string) {
	o.Description = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *ApiAccessRule) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessRule) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *ApiAccessRule) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *ApiAccessRule) SetIpRanges(v []string) {
	o.IpRanges = &v
}

func (o ApiAccessRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessRuleId != nil {
		toSerialize["ApiAccessRuleId"] = o.ApiAccessRuleId
	}
	if o.CaIds != nil {
		toSerialize["CaIds"] = o.CaIds
	}
	if o.Cns != nil {
		toSerialize["Cns"] = o.Cns
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableApiAccessRule struct {
	value *ApiAccessRule
	isSet bool
}

func (v NullableApiAccessRule) Get() *ApiAccessRule {
	return v.value
}

func (v *NullableApiAccessRule) Set(val *ApiAccessRule) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessRule) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessRule(val *ApiAccessRule) *NullableApiAccessRule {
	return &NullableApiAccessRule{value: val, isSet: true}
}

func (v NullableApiAccessRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
