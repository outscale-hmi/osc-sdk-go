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

// FiltersNetAccessPoint One or more filters.
type FiltersNetAccessPoint struct {
	// The IDs of the Net access points.
	NetAccessPointIds *[]string `json:"NetAccessPointIds,omitempty"`
	// The IDs of the Nets.
	NetIds *[]string `json:"NetIds,omitempty"`
	// The names of the services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceNames *[]string `json:"ServiceNames,omitempty"`
	// The states of the Net access points (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States *[]string `json:"States,omitempty"`
	// The keys of the tags associated with the Net access points.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Net access points.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Net access points, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags *[]string `json:"Tags,omitempty"`
}

// NewFiltersNetAccessPoint instantiates a new FiltersNetAccessPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersNetAccessPoint() *FiltersNetAccessPoint {
	this := FiltersNetAccessPoint{}
	return &this
}

// NewFiltersNetAccessPointWithDefaults instantiates a new FiltersNetAccessPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersNetAccessPointWithDefaults() *FiltersNetAccessPoint {
	this := FiltersNetAccessPoint{}
	return &this
}

// GetNetAccessPointIds returns the NetAccessPointIds field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetNetAccessPointIds() []string {
	if o == nil || o.NetAccessPointIds == nil {
		var ret []string
		return ret
	}
	return *o.NetAccessPointIds
}

// GetNetAccessPointIdsOk returns a tuple with the NetAccessPointIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetNetAccessPointIdsOk() (*[]string, bool) {
	if o == nil || o.NetAccessPointIds == nil {
		return nil, false
	}
	return o.NetAccessPointIds, true
}

// HasNetAccessPointIds returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasNetAccessPointIds() bool {
	if o != nil && o.NetAccessPointIds != nil {
		return true
	}

	return false
}

// SetNetAccessPointIds gets a reference to the given []string and assigns it to the NetAccessPointIds field.
func (o *FiltersNetAccessPoint) SetNetAccessPointIds(v []string) {
	o.NetAccessPointIds = &v
}

// GetNetIds returns the NetIds field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetNetIds() []string {
	if o == nil || o.NetIds == nil {
		var ret []string
		return ret
	}
	return *o.NetIds
}

// GetNetIdsOk returns a tuple with the NetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetNetIdsOk() (*[]string, bool) {
	if o == nil || o.NetIds == nil {
		return nil, false
	}
	return o.NetIds, true
}

// HasNetIds returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasNetIds() bool {
	if o != nil && o.NetIds != nil {
		return true
	}

	return false
}

// SetNetIds gets a reference to the given []string and assigns it to the NetIds field.
func (o *FiltersNetAccessPoint) SetNetIds(v []string) {
	o.NetIds = &v
}

// GetServiceNames returns the ServiceNames field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetServiceNames() []string {
	if o == nil || o.ServiceNames == nil {
		var ret []string
		return ret
	}
	return *o.ServiceNames
}

// GetServiceNamesOk returns a tuple with the ServiceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetServiceNamesOk() (*[]string, bool) {
	if o == nil || o.ServiceNames == nil {
		return nil, false
	}
	return o.ServiceNames, true
}

// HasServiceNames returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasServiceNames() bool {
	if o != nil && o.ServiceNames != nil {
		return true
	}

	return false
}

// SetServiceNames gets a reference to the given []string and assigns it to the ServiceNames field.
func (o *FiltersNetAccessPoint) SetServiceNames(v []string) {
	o.ServiceNames = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetStates() []string {
	if o == nil || o.States == nil {
		var ret []string
		return ret
	}
	return *o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetStatesOk() (*[]string, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasStates() bool {
	if o != nil && o.States != nil {
		return true
	}

	return false
}

// SetStates gets a reference to the given []string and assigns it to the States field.
func (o *FiltersNetAccessPoint) SetStates(v []string) {
	o.States = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersNetAccessPoint) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetTagValuesOk() (*[]string, bool) {
	if o == nil || o.TagValues == nil {
		return nil, false
	}
	return o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersNetAccessPoint) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersNetAccessPoint) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersNetAccessPoint) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersNetAccessPoint) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersNetAccessPoint) SetTags(v []string) {
	o.Tags = &v
}

func (o FiltersNetAccessPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetAccessPointIds != nil {
		toSerialize["NetAccessPointIds"] = o.NetAccessPointIds
	}
	if o.NetIds != nil {
		toSerialize["NetIds"] = o.NetIds
	}
	if o.ServiceNames != nil {
		toSerialize["ServiceNames"] = o.ServiceNames
	}
	if o.States != nil {
		toSerialize["States"] = o.States
	}
	if o.TagKeys != nil {
		toSerialize["TagKeys"] = o.TagKeys
	}
	if o.TagValues != nil {
		toSerialize["TagValues"] = o.TagValues
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersNetAccessPoint struct {
	value *FiltersNetAccessPoint
	isSet bool
}

func (v NullableFiltersNetAccessPoint) Get() *FiltersNetAccessPoint {
	return v.value
}

func (v *NullableFiltersNetAccessPoint) Set(val *FiltersNetAccessPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersNetAccessPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersNetAccessPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersNetAccessPoint(val *FiltersNetAccessPoint) *NullableFiltersNetAccessPoint {
	return &NullableFiltersNetAccessPoint{value: val, isSet: true}
}

func (v NullableFiltersNetAccessPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersNetAccessPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
