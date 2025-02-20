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

// UpdateRoutePropagationRequest struct for UpdateRoutePropagationRequest
type UpdateRoutePropagationRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// If true, a virtual gateway can propagate routes to a specified route table of a Net. If false, the propagation is disabled.
	Enable bool `json:"Enable"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId"`
	// The ID of the virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId"`
}

// NewUpdateRoutePropagationRequest instantiates a new UpdateRoutePropagationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoutePropagationRequest(enable bool, routeTableId string, virtualGatewayId string) *UpdateRoutePropagationRequest {
	this := UpdateRoutePropagationRequest{}
	this.Enable = enable
	this.RouteTableId = routeTableId
	this.VirtualGatewayId = virtualGatewayId
	return &this
}

// NewUpdateRoutePropagationRequestWithDefaults instantiates a new UpdateRoutePropagationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoutePropagationRequestWithDefaults() *UpdateRoutePropagationRequest {
	this := UpdateRoutePropagationRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateRoutePropagationRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoutePropagationRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateRoutePropagationRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateRoutePropagationRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEnable returns the Enable field value
func (o *UpdateRoutePropagationRequest) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *UpdateRoutePropagationRequest) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *UpdateRoutePropagationRequest) SetEnable(v bool) {
	o.Enable = v
}

// GetRouteTableId returns the RouteTableId field value
func (o *UpdateRoutePropagationRequest) GetRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value
// and a boolean to check if the value has been set.
func (o *UpdateRoutePropagationRequest) GetRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableId, true
}

// SetRouteTableId sets field value
func (o *UpdateRoutePropagationRequest) SetRouteTableId(v string) {
	o.RouteTableId = v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value
func (o *UpdateRoutePropagationRequest) GetVirtualGatewayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value
// and a boolean to check if the value has been set.
func (o *UpdateRoutePropagationRequest) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualGatewayId, true
}

// SetVirtualGatewayId sets field value
func (o *UpdateRoutePropagationRequest) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = v
}

func (o UpdateRoutePropagationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Enable"] = o.Enable
	}
	if true {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if true {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRoutePropagationRequest struct {
	value *UpdateRoutePropagationRequest
	isSet bool
}

func (v NullableUpdateRoutePropagationRequest) Get() *UpdateRoutePropagationRequest {
	return v.value
}

func (v *NullableUpdateRoutePropagationRequest) Set(val *UpdateRoutePropagationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoutePropagationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoutePropagationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoutePropagationRequest(val *UpdateRoutePropagationRequest) *NullableUpdateRoutePropagationRequest {
	return &NullableUpdateRoutePropagationRequest{value: val, isSet: true}
}

func (v NullableUpdateRoutePropagationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoutePropagationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
