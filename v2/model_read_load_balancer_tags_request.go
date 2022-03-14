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

// ReadLoadBalancerTagsRequest struct for ReadLoadBalancerTagsRequest
type ReadLoadBalancerTagsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more load balancer names.
	LoadBalancerNames []string `json:"LoadBalancerNames"`
}

// NewReadLoadBalancerTagsRequest instantiates a new ReadLoadBalancerTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLoadBalancerTagsRequest(loadBalancerNames []string) *ReadLoadBalancerTagsRequest {
	this := ReadLoadBalancerTagsRequest{}
	this.LoadBalancerNames = loadBalancerNames
	return &this
}

// NewReadLoadBalancerTagsRequestWithDefaults instantiates a new ReadLoadBalancerTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLoadBalancerTagsRequestWithDefaults() *ReadLoadBalancerTagsRequest {
	this := ReadLoadBalancerTagsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadLoadBalancerTagsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadLoadBalancerTagsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadLoadBalancerTagsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerNames returns the LoadBalancerNames field value
func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LoadBalancerNames
}

// GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field value
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerNames, true
}

// SetLoadBalancerNames sets field value
func (o *ReadLoadBalancerTagsRequest) SetLoadBalancerNames(v []string) {
	o.LoadBalancerNames = v
}

func (o ReadLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerNames"] = o.LoadBalancerNames
	}
	return json.Marshal(toSerialize)
}

type NullableReadLoadBalancerTagsRequest struct {
	value *ReadLoadBalancerTagsRequest
	isSet bool
}

func (v NullableReadLoadBalancerTagsRequest) Get() *ReadLoadBalancerTagsRequest {
	return v.value
}

func (v *NullableReadLoadBalancerTagsRequest) Set(val *ReadLoadBalancerTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLoadBalancerTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLoadBalancerTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLoadBalancerTagsRequest(val *ReadLoadBalancerTagsRequest) *NullableReadLoadBalancerTagsRequest {
	return &NullableReadLoadBalancerTagsRequest{value: val, isSet: true}
}

func (v NullableReadLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLoadBalancerTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
