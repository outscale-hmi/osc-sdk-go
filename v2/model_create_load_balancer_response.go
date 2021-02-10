/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateLoadBalancerResponse struct for CreateLoadBalancerResponse
type CreateLoadBalancerResponse struct {
	LoadBalancer    *LoadBalancer    `json:"LoadBalancer,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateLoadBalancerResponse instantiates a new CreateLoadBalancerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerResponse() *CreateLoadBalancerResponse {
	this := CreateLoadBalancerResponse{}
	return &this
}

// NewCreateLoadBalancerResponseWithDefaults instantiates a new CreateLoadBalancerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerResponseWithDefaults() *CreateLoadBalancerResponse {
	this := CreateLoadBalancerResponse{}
	return &this
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *CreateLoadBalancerResponse) GetLoadBalancer() LoadBalancer {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerResponse) GetLoadBalancerOk() (*LoadBalancer, bool) {
	if o == nil || o.LoadBalancer == nil {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *CreateLoadBalancerResponse) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer != nil {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given LoadBalancer and assigns it to the LoadBalancer field.
func (o *CreateLoadBalancerResponse) SetLoadBalancer(v LoadBalancer) {
	o.LoadBalancer = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateLoadBalancerResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateLoadBalancerResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateLoadBalancerResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateLoadBalancerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoadBalancer != nil {
		toSerialize["LoadBalancer"] = o.LoadBalancer
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerResponse struct {
	value *CreateLoadBalancerResponse
	isSet bool
}

func (v NullableCreateLoadBalancerResponse) Get() *CreateLoadBalancerResponse {
	return v.value
}

func (v *NullableCreateLoadBalancerResponse) Set(val *CreateLoadBalancerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerResponse(val *CreateLoadBalancerResponse) *NullableCreateLoadBalancerResponse {
	return &NullableCreateLoadBalancerResponse{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
