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

// FiltersApiLog One or more filters.
type FiltersApiLog struct {
	// The access keys used for the logged calls.
	QueryAccessKeys *[]string `json:"QueryAccessKeys,omitempty"`
	// The names of the APIs of the logged calls (always `oapi` for the OUTSCALE API).
	QueryApiNames *[]string `json:"QueryApiNames,omitempty"`
	// The names of the logged calls.
	QueryCallNames *[]string `json:"QueryCallNames,omitempty"`
	// The date after which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-14`). By default, this date is set to 48 hours before the `QueryDateBefore` parameter value.
	QueryDateAfter *string `json:"QueryDateAfter,omitempty"`
	// The date before which you want to retrieve logged calls, in ISO 8601 format (for example, `2020-06-30`). By default, this date is set to now, or 48 hours after the `QueryDateAfter` parameter value.
	QueryDateBefore *string `json:"QueryDateBefore,omitempty"`
	// The IPs used for the logged calls.
	QueryIpAddresses *[]string `json:"QueryIpAddresses,omitempty"`
	// The user agents of the HTTP requests of the logged calls.
	QueryUserAgents *[]string `json:"QueryUserAgents,omitempty"`
	// The request IDs provided in the responses of the logged calls.
	RequestIds *[]string `json:"RequestIds,omitempty"`
	// The HTTP status codes of the logged calls.
	ResponseStatusCodes *[]int32 `json:"ResponseStatusCodes,omitempty"`
}

// NewFiltersApiLog instantiates a new FiltersApiLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersApiLog() *FiltersApiLog {
	this := FiltersApiLog{}
	return &this
}

// NewFiltersApiLogWithDefaults instantiates a new FiltersApiLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersApiLogWithDefaults() *FiltersApiLog {
	this := FiltersApiLog{}
	return &this
}

// GetQueryAccessKeys returns the QueryAccessKeys field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryAccessKeys() []string {
	if o == nil || o.QueryAccessKeys == nil {
		var ret []string
		return ret
	}
	return *o.QueryAccessKeys
}

// GetQueryAccessKeysOk returns a tuple with the QueryAccessKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryAccessKeysOk() (*[]string, bool) {
	if o == nil || o.QueryAccessKeys == nil {
		return nil, false
	}
	return o.QueryAccessKeys, true
}

// HasQueryAccessKeys returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryAccessKeys() bool {
	if o != nil && o.QueryAccessKeys != nil {
		return true
	}

	return false
}

// SetQueryAccessKeys gets a reference to the given []string and assigns it to the QueryAccessKeys field.
func (o *FiltersApiLog) SetQueryAccessKeys(v []string) {
	o.QueryAccessKeys = &v
}

// GetQueryApiNames returns the QueryApiNames field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryApiNames() []string {
	if o == nil || o.QueryApiNames == nil {
		var ret []string
		return ret
	}
	return *o.QueryApiNames
}

// GetQueryApiNamesOk returns a tuple with the QueryApiNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryApiNamesOk() (*[]string, bool) {
	if o == nil || o.QueryApiNames == nil {
		return nil, false
	}
	return o.QueryApiNames, true
}

// HasQueryApiNames returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryApiNames() bool {
	if o != nil && o.QueryApiNames != nil {
		return true
	}

	return false
}

// SetQueryApiNames gets a reference to the given []string and assigns it to the QueryApiNames field.
func (o *FiltersApiLog) SetQueryApiNames(v []string) {
	o.QueryApiNames = &v
}

// GetQueryCallNames returns the QueryCallNames field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryCallNames() []string {
	if o == nil || o.QueryCallNames == nil {
		var ret []string
		return ret
	}
	return *o.QueryCallNames
}

// GetQueryCallNamesOk returns a tuple with the QueryCallNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryCallNamesOk() (*[]string, bool) {
	if o == nil || o.QueryCallNames == nil {
		return nil, false
	}
	return o.QueryCallNames, true
}

// HasQueryCallNames returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryCallNames() bool {
	if o != nil && o.QueryCallNames != nil {
		return true
	}

	return false
}

// SetQueryCallNames gets a reference to the given []string and assigns it to the QueryCallNames field.
func (o *FiltersApiLog) SetQueryCallNames(v []string) {
	o.QueryCallNames = &v
}

// GetQueryDateAfter returns the QueryDateAfter field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryDateAfter() string {
	if o == nil || o.QueryDateAfter == nil {
		var ret string
		return ret
	}
	return *o.QueryDateAfter
}

// GetQueryDateAfterOk returns a tuple with the QueryDateAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryDateAfterOk() (*string, bool) {
	if o == nil || o.QueryDateAfter == nil {
		return nil, false
	}
	return o.QueryDateAfter, true
}

// HasQueryDateAfter returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryDateAfter() bool {
	if o != nil && o.QueryDateAfter != nil {
		return true
	}

	return false
}

// SetQueryDateAfter gets a reference to the given string and assigns it to the QueryDateAfter field.
func (o *FiltersApiLog) SetQueryDateAfter(v string) {
	o.QueryDateAfter = &v
}

// GetQueryDateBefore returns the QueryDateBefore field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryDateBefore() string {
	if o == nil || o.QueryDateBefore == nil {
		var ret string
		return ret
	}
	return *o.QueryDateBefore
}

// GetQueryDateBeforeOk returns a tuple with the QueryDateBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryDateBeforeOk() (*string, bool) {
	if o == nil || o.QueryDateBefore == nil {
		return nil, false
	}
	return o.QueryDateBefore, true
}

// HasQueryDateBefore returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryDateBefore() bool {
	if o != nil && o.QueryDateBefore != nil {
		return true
	}

	return false
}

// SetQueryDateBefore gets a reference to the given string and assigns it to the QueryDateBefore field.
func (o *FiltersApiLog) SetQueryDateBefore(v string) {
	o.QueryDateBefore = &v
}

// GetQueryIpAddresses returns the QueryIpAddresses field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryIpAddresses() []string {
	if o == nil || o.QueryIpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.QueryIpAddresses
}

// GetQueryIpAddressesOk returns a tuple with the QueryIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryIpAddressesOk() (*[]string, bool) {
	if o == nil || o.QueryIpAddresses == nil {
		return nil, false
	}
	return o.QueryIpAddresses, true
}

// HasQueryIpAddresses returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryIpAddresses() bool {
	if o != nil && o.QueryIpAddresses != nil {
		return true
	}

	return false
}

// SetQueryIpAddresses gets a reference to the given []string and assigns it to the QueryIpAddresses field.
func (o *FiltersApiLog) SetQueryIpAddresses(v []string) {
	o.QueryIpAddresses = &v
}

// GetQueryUserAgents returns the QueryUserAgents field value if set, zero value otherwise.
func (o *FiltersApiLog) GetQueryUserAgents() []string {
	if o == nil || o.QueryUserAgents == nil {
		var ret []string
		return ret
	}
	return *o.QueryUserAgents
}

// GetQueryUserAgentsOk returns a tuple with the QueryUserAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetQueryUserAgentsOk() (*[]string, bool) {
	if o == nil || o.QueryUserAgents == nil {
		return nil, false
	}
	return o.QueryUserAgents, true
}

// HasQueryUserAgents returns a boolean if a field has been set.
func (o *FiltersApiLog) HasQueryUserAgents() bool {
	if o != nil && o.QueryUserAgents != nil {
		return true
	}

	return false
}

// SetQueryUserAgents gets a reference to the given []string and assigns it to the QueryUserAgents field.
func (o *FiltersApiLog) SetQueryUserAgents(v []string) {
	o.QueryUserAgents = &v
}

// GetRequestIds returns the RequestIds field value if set, zero value otherwise.
func (o *FiltersApiLog) GetRequestIds() []string {
	if o == nil || o.RequestIds == nil {
		var ret []string
		return ret
	}
	return *o.RequestIds
}

// GetRequestIdsOk returns a tuple with the RequestIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetRequestIdsOk() (*[]string, bool) {
	if o == nil || o.RequestIds == nil {
		return nil, false
	}
	return o.RequestIds, true
}

// HasRequestIds returns a boolean if a field has been set.
func (o *FiltersApiLog) HasRequestIds() bool {
	if o != nil && o.RequestIds != nil {
		return true
	}

	return false
}

// SetRequestIds gets a reference to the given []string and assigns it to the RequestIds field.
func (o *FiltersApiLog) SetRequestIds(v []string) {
	o.RequestIds = &v
}

// GetResponseStatusCodes returns the ResponseStatusCodes field value if set, zero value otherwise.
func (o *FiltersApiLog) GetResponseStatusCodes() []int32 {
	if o == nil || o.ResponseStatusCodes == nil {
		var ret []int32
		return ret
	}
	return *o.ResponseStatusCodes
}

// GetResponseStatusCodesOk returns a tuple with the ResponseStatusCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersApiLog) GetResponseStatusCodesOk() (*[]int32, bool) {
	if o == nil || o.ResponseStatusCodes == nil {
		return nil, false
	}
	return o.ResponseStatusCodes, true
}

// HasResponseStatusCodes returns a boolean if a field has been set.
func (o *FiltersApiLog) HasResponseStatusCodes() bool {
	if o != nil && o.ResponseStatusCodes != nil {
		return true
	}

	return false
}

// SetResponseStatusCodes gets a reference to the given []int32 and assigns it to the ResponseStatusCodes field.
func (o *FiltersApiLog) SetResponseStatusCodes(v []int32) {
	o.ResponseStatusCodes = &v
}

func (o FiltersApiLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueryAccessKeys != nil {
		toSerialize["QueryAccessKeys"] = o.QueryAccessKeys
	}
	if o.QueryApiNames != nil {
		toSerialize["QueryApiNames"] = o.QueryApiNames
	}
	if o.QueryCallNames != nil {
		toSerialize["QueryCallNames"] = o.QueryCallNames
	}
	if o.QueryDateAfter != nil {
		toSerialize["QueryDateAfter"] = o.QueryDateAfter
	}
	if o.QueryDateBefore != nil {
		toSerialize["QueryDateBefore"] = o.QueryDateBefore
	}
	if o.QueryIpAddresses != nil {
		toSerialize["QueryIpAddresses"] = o.QueryIpAddresses
	}
	if o.QueryUserAgents != nil {
		toSerialize["QueryUserAgents"] = o.QueryUserAgents
	}
	if o.RequestIds != nil {
		toSerialize["RequestIds"] = o.RequestIds
	}
	if o.ResponseStatusCodes != nil {
		toSerialize["ResponseStatusCodes"] = o.ResponseStatusCodes
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersApiLog struct {
	value *FiltersApiLog
	isSet bool
}

func (v NullableFiltersApiLog) Get() *FiltersApiLog {
	return v.value
}

func (v *NullableFiltersApiLog) Set(val *FiltersApiLog) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersApiLog) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersApiLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersApiLog(val *FiltersApiLog) *NullableFiltersApiLog {
	return &NullableFiltersApiLog{value: val, isSet: true}
}

func (v NullableFiltersApiLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersApiLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
