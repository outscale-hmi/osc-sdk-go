/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ListenerRule Information about the listener rule.
type ListenerRule struct {
	// The type of action for the rule (always `forward`).
	Action *string `json:"Action,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostNamePattern *string `json:"HostNamePattern,omitempty"`
	// The ID of the listener.
	ListenerId *int32 `json:"ListenerId,omitempty"`
	// The ID of the listener rule.
	ListenerRuleId *int32 `json:"ListenerRuleId,omitempty"`
	// A human-readable name for the listener rule.
	ListenerRuleName *string `json:"ListenerRuleName,omitempty"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&quot;'@:+?].
	PathPattern *string `json:"PathPattern,omitempty"`
	// The priority level of the listener rule, between `1` and `19999` both included. Each rule must have a unique priority level. Otherwise, an error is returned.
	Priority *int32 `json:"Priority,omitempty"`
	// The IDs of the backend VMs.
	VmIds *[]string `json:"VmIds,omitempty"`
}

// NewListenerRule instantiates a new ListenerRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenerRule() *ListenerRule {
	this := ListenerRule{}
	return &this
}

// NewListenerRuleWithDefaults instantiates a new ListenerRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerRuleWithDefaults() *ListenerRule {
	this := ListenerRule{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ListenerRule) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ListenerRule) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ListenerRule) SetAction(v string) {
	o.Action = &v
}

// GetHostNamePattern returns the HostNamePattern field value if set, zero value otherwise.
func (o *ListenerRule) GetHostNamePattern() string {
	if o == nil || o.HostNamePattern == nil {
		var ret string
		return ret
	}
	return *o.HostNamePattern
}

// GetHostNamePatternOk returns a tuple with the HostNamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetHostNamePatternOk() (*string, bool) {
	if o == nil || o.HostNamePattern == nil {
		return nil, false
	}
	return o.HostNamePattern, true
}

// HasHostNamePattern returns a boolean if a field has been set.
func (o *ListenerRule) HasHostNamePattern() bool {
	if o != nil && o.HostNamePattern != nil {
		return true
	}

	return false
}

// SetHostNamePattern gets a reference to the given string and assigns it to the HostNamePattern field.
func (o *ListenerRule) SetHostNamePattern(v string) {
	o.HostNamePattern = &v
}

// GetListenerId returns the ListenerId field value if set, zero value otherwise.
func (o *ListenerRule) GetListenerId() int32 {
	if o == nil || o.ListenerId == nil {
		var ret int32
		return ret
	}
	return *o.ListenerId
}

// GetListenerIdOk returns a tuple with the ListenerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetListenerIdOk() (*int32, bool) {
	if o == nil || o.ListenerId == nil {
		return nil, false
	}
	return o.ListenerId, true
}

// HasListenerId returns a boolean if a field has been set.
func (o *ListenerRule) HasListenerId() bool {
	if o != nil && o.ListenerId != nil {
		return true
	}

	return false
}

// SetListenerId gets a reference to the given int32 and assigns it to the ListenerId field.
func (o *ListenerRule) SetListenerId(v int32) {
	o.ListenerId = &v
}

// GetListenerRuleId returns the ListenerRuleId field value if set, zero value otherwise.
func (o *ListenerRule) GetListenerRuleId() int32 {
	if o == nil || o.ListenerRuleId == nil {
		var ret int32
		return ret
	}
	return *o.ListenerRuleId
}

// GetListenerRuleIdOk returns a tuple with the ListenerRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetListenerRuleIdOk() (*int32, bool) {
	if o == nil || o.ListenerRuleId == nil {
		return nil, false
	}
	return o.ListenerRuleId, true
}

// HasListenerRuleId returns a boolean if a field has been set.
func (o *ListenerRule) HasListenerRuleId() bool {
	if o != nil && o.ListenerRuleId != nil {
		return true
	}

	return false
}

// SetListenerRuleId gets a reference to the given int32 and assigns it to the ListenerRuleId field.
func (o *ListenerRule) SetListenerRuleId(v int32) {
	o.ListenerRuleId = &v
}

// GetListenerRuleName returns the ListenerRuleName field value if set, zero value otherwise.
func (o *ListenerRule) GetListenerRuleName() string {
	if o == nil || o.ListenerRuleName == nil {
		var ret string
		return ret
	}
	return *o.ListenerRuleName
}

// GetListenerRuleNameOk returns a tuple with the ListenerRuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetListenerRuleNameOk() (*string, bool) {
	if o == nil || o.ListenerRuleName == nil {
		return nil, false
	}
	return o.ListenerRuleName, true
}

// HasListenerRuleName returns a boolean if a field has been set.
func (o *ListenerRule) HasListenerRuleName() bool {
	if o != nil && o.ListenerRuleName != nil {
		return true
	}

	return false
}

// SetListenerRuleName gets a reference to the given string and assigns it to the ListenerRuleName field.
func (o *ListenerRule) SetListenerRuleName(v string) {
	o.ListenerRuleName = &v
}

// GetPathPattern returns the PathPattern field value if set, zero value otherwise.
func (o *ListenerRule) GetPathPattern() string {
	if o == nil || o.PathPattern == nil {
		var ret string
		return ret
	}
	return *o.PathPattern
}

// GetPathPatternOk returns a tuple with the PathPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetPathPatternOk() (*string, bool) {
	if o == nil || o.PathPattern == nil {
		return nil, false
	}
	return o.PathPattern, true
}

// HasPathPattern returns a boolean if a field has been set.
func (o *ListenerRule) HasPathPattern() bool {
	if o != nil && o.PathPattern != nil {
		return true
	}

	return false
}

// SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.
func (o *ListenerRule) SetPathPattern(v string) {
	o.PathPattern = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ListenerRule) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ListenerRule) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ListenerRule) SetPriority(v int32) {
	o.Priority = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *ListenerRule) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerRule) GetVmIdsOk() (*[]string, bool) {
	if o == nil || o.VmIds == nil {
		return nil, false
	}
	return o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *ListenerRule) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *ListenerRule) SetVmIds(v []string) {
	o.VmIds = &v
}

func (o ListenerRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.HostNamePattern != nil {
		toSerialize["HostNamePattern"] = o.HostNamePattern
	}
	if o.ListenerId != nil {
		toSerialize["ListenerId"] = o.ListenerId
	}
	if o.ListenerRuleId != nil {
		toSerialize["ListenerRuleId"] = o.ListenerRuleId
	}
	if o.ListenerRuleName != nil {
		toSerialize["ListenerRuleName"] = o.ListenerRuleName
	}
	if o.PathPattern != nil {
		toSerialize["PathPattern"] = o.PathPattern
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.VmIds != nil {
		toSerialize["VmIds"] = o.VmIds
	}
	return json.Marshal(toSerialize)
}

type NullableListenerRule struct {
	value *ListenerRule
	isSet bool
}

func (v NullableListenerRule) Get() *ListenerRule {
	return v.value
}

func (v *NullableListenerRule) Set(val *ListenerRule) {
	v.value = val
	v.isSet = true
}

func (v NullableListenerRule) IsSet() bool {
	return v.isSet
}

func (v *NullableListenerRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenerRule(val *ListenerRule) *NullableListenerRule {
	return &NullableListenerRule{value: val, isSet: true}
}

func (v NullableListenerRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenerRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
