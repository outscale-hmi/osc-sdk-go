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

// LinkRouteTable One or more associations between the route table and the Subnets.
type LinkRouteTable struct {
	// The ID of the association between the route table and the Subnet.
	LinkRouteTableId *string `json:"LinkRouteTableId,omitempty"`
	// If true, the route table is the main one.
	Main *bool `json:"Main,omitempty"`
	// The ID of the route table.
	RouteTableId *string `json:"RouteTableId,omitempty"`
	// The ID of the Subnet.
	SubnetId *string `json:"SubnetId,omitempty"`
}

// NewLinkRouteTable instantiates a new LinkRouteTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkRouteTable() *LinkRouteTable {
	this := LinkRouteTable{}
	return &this
}

// NewLinkRouteTableWithDefaults instantiates a new LinkRouteTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkRouteTableWithDefaults() *LinkRouteTable {
	this := LinkRouteTable{}
	return &this
}

// GetLinkRouteTableId returns the LinkRouteTableId field value if set, zero value otherwise.
func (o *LinkRouteTable) GetLinkRouteTableId() string {
	if o == nil || o.LinkRouteTableId == nil {
		var ret string
		return ret
	}
	return *o.LinkRouteTableId
}

// GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTable) GetLinkRouteTableIdOk() (*string, bool) {
	if o == nil || o.LinkRouteTableId == nil {
		return nil, false
	}
	return o.LinkRouteTableId, true
}

// HasLinkRouteTableId returns a boolean if a field has been set.
func (o *LinkRouteTable) HasLinkRouteTableId() bool {
	if o != nil && o.LinkRouteTableId != nil {
		return true
	}

	return false
}

// SetLinkRouteTableId gets a reference to the given string and assigns it to the LinkRouteTableId field.
func (o *LinkRouteTable) SetLinkRouteTableId(v string) {
	o.LinkRouteTableId = &v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *LinkRouteTable) GetMain() bool {
	if o == nil || o.Main == nil {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTable) GetMainOk() (*bool, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *LinkRouteTable) HasMain() bool {
	if o != nil && o.Main != nil {
		return true
	}

	return false
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *LinkRouteTable) SetMain(v bool) {
	o.Main = &v
}

// GetRouteTableId returns the RouteTableId field value if set, zero value otherwise.
func (o *LinkRouteTable) GetRouteTableId() string {
	if o == nil || o.RouteTableId == nil {
		var ret string
		return ret
	}
	return *o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTable) GetRouteTableIdOk() (*string, bool) {
	if o == nil || o.RouteTableId == nil {
		return nil, false
	}
	return o.RouteTableId, true
}

// HasRouteTableId returns a boolean if a field has been set.
func (o *LinkRouteTable) HasRouteTableId() bool {
	if o != nil && o.RouteTableId != nil {
		return true
	}

	return false
}

// SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.
func (o *LinkRouteTable) SetRouteTableId(v string) {
	o.RouteTableId = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *LinkRouteTable) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTable) GetSubnetIdOk() (*string, bool) {
	if o == nil || o.SubnetId == nil {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *LinkRouteTable) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *LinkRouteTable) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o LinkRouteTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkRouteTableId != nil {
		toSerialize["LinkRouteTableId"] = o.LinkRouteTableId
	}
	if o.Main != nil {
		toSerialize["Main"] = o.Main
	}
	if o.RouteTableId != nil {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if o.SubnetId != nil {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkRouteTable struct {
	value *LinkRouteTable
	isSet bool
}

func (v NullableLinkRouteTable) Get() *LinkRouteTable {
	return v.value
}

func (v *NullableLinkRouteTable) Set(val *LinkRouteTable) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkRouteTable) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkRouteTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkRouteTable(val *LinkRouteTable) *NullableLinkRouteTable {
	return &NullableLinkRouteTable{value: val, isSet: true}
}

func (v NullableLinkRouteTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkRouteTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
