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

// CreateLoadBalancerListenersRequest struct for CreateLoadBalancerListenersRequest
type CreateLoadBalancerListenersRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more listeners for the load balancer.
	Listeners []ListenerForCreation `json:"Listeners"`
	// The name of the load balancer for which you want to create listeners.
	LoadBalancerName string `json:"LoadBalancerName"`
}

// NewCreateLoadBalancerListenersRequest instantiates a new CreateLoadBalancerListenersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerListenersRequest(listeners []ListenerForCreation, loadBalancerName string) *CreateLoadBalancerListenersRequest {
	this := CreateLoadBalancerListenersRequest{}
	this.Listeners = listeners
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewCreateLoadBalancerListenersRequestWithDefaults instantiates a new CreateLoadBalancerListenersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerListenersRequestWithDefaults() *CreateLoadBalancerListenersRequest {
	this := CreateLoadBalancerListenersRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateLoadBalancerListenersRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerListenersRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateLoadBalancerListenersRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateLoadBalancerListenersRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetListeners returns the Listeners field value
func (o *CreateLoadBalancerListenersRequest) GetListeners() []ListenerForCreation {
	if o == nil {
		var ret []ListenerForCreation
		return ret
	}

	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerListenersRequest) GetListenersOk() (*[]ListenerForCreation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Listeners, true
}

// SetListeners sets field value
func (o *CreateLoadBalancerListenersRequest) SetListeners(v []ListenerForCreation) {
	o.Listeners = v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *CreateLoadBalancerListenersRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

func (o CreateLoadBalancerListenersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Listeners"] = o.Listeners
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerListenersRequest struct {
	value *CreateLoadBalancerListenersRequest
	isSet bool
}

func (v NullableCreateLoadBalancerListenersRequest) Get() *CreateLoadBalancerListenersRequest {
	return v.value
}

func (v *NullableCreateLoadBalancerListenersRequest) Set(val *CreateLoadBalancerListenersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerListenersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerListenersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerListenersRequest(val *CreateLoadBalancerListenersRequest) *NullableCreateLoadBalancerListenersRequest {
	return &NullableCreateLoadBalancerListenersRequest{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerListenersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerListenersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
