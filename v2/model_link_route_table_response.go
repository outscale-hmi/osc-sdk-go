/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.14
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkRouteTableResponse struct for LinkRouteTableResponse
type LinkRouteTableResponse struct {
	// The ID of the association between the route table and the Subnet.
	LinkRouteTableId *string          `json:"LinkRouteTableId,omitempty"`
	ResponseContext  *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkRouteTableResponse instantiates a new LinkRouteTableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkRouteTableResponse() *LinkRouteTableResponse {
	this := LinkRouteTableResponse{}
	return &this
}

// NewLinkRouteTableResponseWithDefaults instantiates a new LinkRouteTableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkRouteTableResponseWithDefaults() *LinkRouteTableResponse {
	this := LinkRouteTableResponse{}
	return &this
}

// GetLinkRouteTableId returns the LinkRouteTableId field value if set, zero value otherwise.
func (o *LinkRouteTableResponse) GetLinkRouteTableId() string {
	if o == nil || o.LinkRouteTableId == nil {
		var ret string
		return ret
	}
	return *o.LinkRouteTableId
}

// GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTableResponse) GetLinkRouteTableIdOk() (*string, bool) {
	if o == nil || o.LinkRouteTableId == nil {
		return nil, false
	}
	return o.LinkRouteTableId, true
}

// HasLinkRouteTableId returns a boolean if a field has been set.
func (o *LinkRouteTableResponse) HasLinkRouteTableId() bool {
	if o != nil && o.LinkRouteTableId != nil {
		return true
	}

	return false
}

// SetLinkRouteTableId gets a reference to the given string and assigns it to the LinkRouteTableId field.
func (o *LinkRouteTableResponse) SetLinkRouteTableId(v string) {
	o.LinkRouteTableId = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkRouteTableResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkRouteTableResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkRouteTableResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkRouteTableResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkRouteTableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LinkRouteTableId != nil {
		toSerialize["LinkRouteTableId"] = o.LinkRouteTableId
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkRouteTableResponse struct {
	value *LinkRouteTableResponse
	isSet bool
}

func (v NullableLinkRouteTableResponse) Get() *LinkRouteTableResponse {
	return v.value
}

func (v *NullableLinkRouteTableResponse) Set(val *LinkRouteTableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkRouteTableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkRouteTableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkRouteTableResponse(val *LinkRouteTableResponse) *NullableLinkRouteTableResponse {
	return &NullableLinkRouteTableResponse{value: val, isSet: true}
}

func (v NullableLinkRouteTableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkRouteTableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
