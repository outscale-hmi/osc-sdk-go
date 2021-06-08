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

// VpnOptions Information about the VPN options.
type VpnOptions struct {
	Phase1Options *Phase1Options `json:"Phase1Options,omitempty"`
	Phase2Options *Phase2Options `json:"Phase2Options,omitempty"`
	// The IP range for the tunnel in your VPN connection.
	TunnelInsideIpRange *string `json:"TunnelInsideIpRange,omitempty"`
}

// NewVpnOptions instantiates a new VpnOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVpnOptions() *VpnOptions {
	this := VpnOptions{}
	return &this
}

// NewVpnOptionsWithDefaults instantiates a new VpnOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVpnOptionsWithDefaults() *VpnOptions {
	this := VpnOptions{}
	return &this
}

// GetPhase1Options returns the Phase1Options field value if set, zero value otherwise.
func (o *VpnOptions) GetPhase1Options() Phase1Options {
	if o == nil || o.Phase1Options == nil {
		var ret Phase1Options
		return ret
	}
	return *o.Phase1Options
}

// GetPhase1OptionsOk returns a tuple with the Phase1Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnOptions) GetPhase1OptionsOk() (*Phase1Options, bool) {
	if o == nil || o.Phase1Options == nil {
		return nil, false
	}
	return o.Phase1Options, true
}

// HasPhase1Options returns a boolean if a field has been set.
func (o *VpnOptions) HasPhase1Options() bool {
	if o != nil && o.Phase1Options != nil {
		return true
	}

	return false
}

// SetPhase1Options gets a reference to the given Phase1Options and assigns it to the Phase1Options field.
func (o *VpnOptions) SetPhase1Options(v Phase1Options) {
	o.Phase1Options = &v
}

// GetPhase2Options returns the Phase2Options field value if set, zero value otherwise.
func (o *VpnOptions) GetPhase2Options() Phase2Options {
	if o == nil || o.Phase2Options == nil {
		var ret Phase2Options
		return ret
	}
	return *o.Phase2Options
}

// GetPhase2OptionsOk returns a tuple with the Phase2Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnOptions) GetPhase2OptionsOk() (*Phase2Options, bool) {
	if o == nil || o.Phase2Options == nil {
		return nil, false
	}
	return o.Phase2Options, true
}

// HasPhase2Options returns a boolean if a field has been set.
func (o *VpnOptions) HasPhase2Options() bool {
	if o != nil && o.Phase2Options != nil {
		return true
	}

	return false
}

// SetPhase2Options gets a reference to the given Phase2Options and assigns it to the Phase2Options field.
func (o *VpnOptions) SetPhase2Options(v Phase2Options) {
	o.Phase2Options = &v
}

// GetTunnelInsideIpRange returns the TunnelInsideIpRange field value if set, zero value otherwise.
func (o *VpnOptions) GetTunnelInsideIpRange() string {
	if o == nil || o.TunnelInsideIpRange == nil {
		var ret string
		return ret
	}
	return *o.TunnelInsideIpRange
}

// GetTunnelInsideIpRangeOk returns a tuple with the TunnelInsideIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VpnOptions) GetTunnelInsideIpRangeOk() (*string, bool) {
	if o == nil || o.TunnelInsideIpRange == nil {
		return nil, false
	}
	return o.TunnelInsideIpRange, true
}

// HasTunnelInsideIpRange returns a boolean if a field has been set.
func (o *VpnOptions) HasTunnelInsideIpRange() bool {
	if o != nil && o.TunnelInsideIpRange != nil {
		return true
	}

	return false
}

// SetTunnelInsideIpRange gets a reference to the given string and assigns it to the TunnelInsideIpRange field.
func (o *VpnOptions) SetTunnelInsideIpRange(v string) {
	o.TunnelInsideIpRange = &v
}

func (o VpnOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Phase1Options != nil {
		toSerialize["Phase1Options"] = o.Phase1Options
	}
	if o.Phase2Options != nil {
		toSerialize["Phase2Options"] = o.Phase2Options
	}
	if o.TunnelInsideIpRange != nil {
		toSerialize["TunnelInsideIpRange"] = o.TunnelInsideIpRange
	}
	return json.Marshal(toSerialize)
}

type NullableVpnOptions struct {
	value *VpnOptions
	isSet bool
}

func (v NullableVpnOptions) Get() *VpnOptions {
	return v.value
}

func (v *NullableVpnOptions) Set(val *VpnOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVpnOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVpnOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVpnOptions(val *VpnOptions) *NullableVpnOptions {
	return &NullableVpnOptions{value: val, isSet: true}
}

func (v NullableVpnOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVpnOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
