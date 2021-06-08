/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.8
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadLoadBalancerTagsResponse struct for ReadLoadBalancerTagsResponse
type ReadLoadBalancerTagsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more load balancer tags.
	Tags *[]LoadBalancerTag `json:"Tags,omitempty"`
}

// NewReadLoadBalancerTagsResponse instantiates a new ReadLoadBalancerTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadLoadBalancerTagsResponse() *ReadLoadBalancerTagsResponse {
	this := ReadLoadBalancerTagsResponse{}
	return &this
}

// NewReadLoadBalancerTagsResponseWithDefaults instantiates a new ReadLoadBalancerTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadLoadBalancerTagsResponseWithDefaults() *ReadLoadBalancerTagsResponse {
	this := ReadLoadBalancerTagsResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadLoadBalancerTagsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancerTagsResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadLoadBalancerTagsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadLoadBalancerTagsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ReadLoadBalancerTagsResponse) GetTags() []LoadBalancerTag {
	if o == nil || o.Tags == nil {
		var ret []LoadBalancerTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancerTagsResponse) GetTagsOk() (*[]LoadBalancerTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ReadLoadBalancerTagsResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []LoadBalancerTag and assigns it to the Tags field.
func (o *ReadLoadBalancerTagsResponse) SetTags(v []LoadBalancerTag) {
	o.Tags = &v
}

func (o ReadLoadBalancerTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableReadLoadBalancerTagsResponse struct {
	value *ReadLoadBalancerTagsResponse
	isSet bool
}

func (v NullableReadLoadBalancerTagsResponse) Get() *ReadLoadBalancerTagsResponse {
	return v.value
}

func (v *NullableReadLoadBalancerTagsResponse) Set(val *ReadLoadBalancerTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReadLoadBalancerTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReadLoadBalancerTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadLoadBalancerTagsResponse(val *ReadLoadBalancerTagsResponse) *NullableReadLoadBalancerTagsResponse {
	return &NullableReadLoadBalancerTagsResponse{value: val, isSet: true}
}

func (v NullableReadLoadBalancerTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadLoadBalancerTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}