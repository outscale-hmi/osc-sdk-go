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

// CreateApiAccessRuleRequest struct for CreateApiAccessRuleRequest
type CreateApiAccessRuleRequest struct {
	//  One or more IDs of Client Certificate Authorities (CAs).
	CaIds *[]string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs). If this parameter is specified, you must also specify the `CaIds` parameter.
	Cns *[]string `json:"Cns,omitempty"`
	// A description for the API access rule.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16).
	IpRanges *[]string `json:"IpRanges,omitempty"`
}

// NewCreateApiAccessRuleRequest instantiates a new CreateApiAccessRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiAccessRuleRequest() *CreateApiAccessRuleRequest {
	this := CreateApiAccessRuleRequest{}
	return &this
}

// NewCreateApiAccessRuleRequestWithDefaults instantiates a new CreateApiAccessRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiAccessRuleRequestWithDefaults() *CreateApiAccessRuleRequest {
	this := CreateApiAccessRuleRequest{}
	return &this
}

// GetCaIds returns the CaIds field value if set, zero value otherwise.
func (o *CreateApiAccessRuleRequest) GetCaIds() []string {
	if o == nil || o.CaIds == nil {
		var ret []string
		return ret
	}
	return *o.CaIds
}

// GetCaIdsOk returns a tuple with the CaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleRequest) GetCaIdsOk() (*[]string, bool) {
	if o == nil || o.CaIds == nil {
		return nil, false
	}
	return o.CaIds, true
}

// HasCaIds returns a boolean if a field has been set.
func (o *CreateApiAccessRuleRequest) HasCaIds() bool {
	if o != nil && o.CaIds != nil {
		return true
	}

	return false
}

// SetCaIds gets a reference to the given []string and assigns it to the CaIds field.
func (o *CreateApiAccessRuleRequest) SetCaIds(v []string) {
	o.CaIds = &v
}

// GetCns returns the Cns field value if set, zero value otherwise.
func (o *CreateApiAccessRuleRequest) GetCns() []string {
	if o == nil || o.Cns == nil {
		var ret []string
		return ret
	}
	return *o.Cns
}

// GetCnsOk returns a tuple with the Cns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleRequest) GetCnsOk() (*[]string, bool) {
	if o == nil || o.Cns == nil {
		return nil, false
	}
	return o.Cns, true
}

// HasCns returns a boolean if a field has been set.
func (o *CreateApiAccessRuleRequest) HasCns() bool {
	if o != nil && o.Cns != nil {
		return true
	}

	return false
}

// SetCns gets a reference to the given []string and assigns it to the Cns field.
func (o *CreateApiAccessRuleRequest) SetCns(v []string) {
	o.Cns = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateApiAccessRuleRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateApiAccessRuleRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateApiAccessRuleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateApiAccessRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateApiAccessRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateApiAccessRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *CreateApiAccessRuleRequest) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleRequest) GetIpRangesOk() (*[]string, bool) {
	if o == nil || o.IpRanges == nil {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *CreateApiAccessRuleRequest) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *CreateApiAccessRuleRequest) SetIpRanges(v []string) {
	o.IpRanges = &v
}

func (o CreateApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaIds != nil {
		toSerialize["CaIds"] = o.CaIds
	}
	if o.Cns != nil {
		toSerialize["Cns"] = o.Cns
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.IpRanges != nil {
		toSerialize["IpRanges"] = o.IpRanges
	}
	return json.Marshal(toSerialize)
}

type NullableCreateApiAccessRuleRequest struct {
	value *CreateApiAccessRuleRequest
	isSet bool
}

func (v NullableCreateApiAccessRuleRequest) Get() *CreateApiAccessRuleRequest {
	return v.value
}

func (v *NullableCreateApiAccessRuleRequest) Set(val *CreateApiAccessRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiAccessRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiAccessRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiAccessRuleRequest(val *CreateApiAccessRuleRequest) *NullableCreateApiAccessRuleRequest {
	return &NullableCreateApiAccessRuleRequest{value: val, isSet: true}
}

func (v NullableCreateApiAccessRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiAccessRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
