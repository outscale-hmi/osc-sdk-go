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

// DeleteRouteRequest struct for DeleteRouteRequest
type DeleteRouteRequest struct {
	// The exact IP range for the route.
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the route table from which you want to delete a route.
	RouteTableId string `json:"RouteTableId"`
}

// NewDeleteRouteRequest instantiates a new DeleteRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRouteRequest(destinationIpRange string, routeTableId string) *DeleteRouteRequest {
	this := DeleteRouteRequest{}
	this.DestinationIpRange = destinationIpRange
	this.RouteTableId = routeTableId
	return &this
}

// NewDeleteRouteRequestWithDefaults instantiates a new DeleteRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRouteRequestWithDefaults() *DeleteRouteRequest {
	this := DeleteRouteRequest{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value
func (o *DeleteRouteRequest) GetDestinationIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value
// and a boolean to check if the value has been set.
func (o *DeleteRouteRequest) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIpRange, true
}

// SetDestinationIpRange sets field value
func (o *DeleteRouteRequest) SetDestinationIpRange(v string) {
	o.DestinationIpRange = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteRouteRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteRouteRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteRouteRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteRouteRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetRouteTableId returns the RouteTableId field value
func (o *DeleteRouteRequest) GetRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value
// and a boolean to check if the value has been set.
func (o *DeleteRouteRequest) GetRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableId, true
}

// SetRouteTableId sets field value
func (o *DeleteRouteRequest) SetRouteTableId(v string) {
	o.RouteTableId = v
}

func (o DeleteRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteRouteRequest struct {
	value *DeleteRouteRequest
	isSet bool
}

func (v NullableDeleteRouteRequest) Get() *DeleteRouteRequest {
	return v.value
}

func (v *NullableDeleteRouteRequest) Set(val *DeleteRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRouteRequest(val *DeleteRouteRequest) *NullableDeleteRouteRequest {
	return &NullableDeleteRouteRequest{value: val, isSet: true}
}

func (v NullableDeleteRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
