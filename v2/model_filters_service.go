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

// FiltersService One or more filters.
type FiltersService struct {
	// The IDs of the services.
	ServiceIds *[]string `json:"ServiceIds,omitempty"`
	// The names of the services.
	ServiceNames *[]string `json:"ServiceNames,omitempty"`
}

// NewFiltersService instantiates a new FiltersService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersService() *FiltersService {
	this := FiltersService{}
	return &this
}

// NewFiltersServiceWithDefaults instantiates a new FiltersService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersServiceWithDefaults() *FiltersService {
	this := FiltersService{}
	return &this
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *FiltersService) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersService) GetServiceIdsOk() (*[]string, bool) {
	if o == nil || o.ServiceIds == nil {
		return nil, false
	}
	return o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *FiltersService) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *FiltersService) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

// GetServiceNames returns the ServiceNames field value if set, zero value otherwise.
func (o *FiltersService) GetServiceNames() []string {
	if o == nil || o.ServiceNames == nil {
		var ret []string
		return ret
	}
	return *o.ServiceNames
}

// GetServiceNamesOk returns a tuple with the ServiceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersService) GetServiceNamesOk() (*[]string, bool) {
	if o == nil || o.ServiceNames == nil {
		return nil, false
	}
	return o.ServiceNames, true
}

// HasServiceNames returns a boolean if a field has been set.
func (o *FiltersService) HasServiceNames() bool {
	if o != nil && o.ServiceNames != nil {
		return true
	}

	return false
}

// SetServiceNames gets a reference to the given []string and assigns it to the ServiceNames field.
func (o *FiltersService) SetServiceNames(v []string) {
	o.ServiceNames = &v
}

func (o FiltersService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceIds != nil {
		toSerialize["ServiceIds"] = o.ServiceIds
	}
	if o.ServiceNames != nil {
		toSerialize["ServiceNames"] = o.ServiceNames
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersService struct {
	value *FiltersService
	isSet bool
}

func (v NullableFiltersService) Get() *FiltersService {
	return v.value
}

func (v *NullableFiltersService) Set(val *FiltersService) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersService) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersService(val *FiltersService) *NullableFiltersService {
	return &NullableFiltersService{value: val, isSet: true}
}

func (v NullableFiltersService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
