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

// UpdateLoadBalancerRequest struct for UpdateLoadBalancerRequest
type UpdateLoadBalancerRequest struct {
	AccessLog *AccessLog `json:"AccessLog,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun      *bool        `json:"DryRun,omitempty"`
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort *int32 `json:"LoadBalancerPort,omitempty"`
	// The list of policy names (must contain all the policies to be enabled).
	PolicyNames *[]string `json:"PolicyNames,omitempty"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).
	ServerCertificateId *string `json:"ServerCertificateId,omitempty"`
}

// NewUpdateLoadBalancerRequest instantiates a new UpdateLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoadBalancerRequest(loadBalancerName string) *UpdateLoadBalancerRequest {
	this := UpdateLoadBalancerRequest{}
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewUpdateLoadBalancerRequestWithDefaults instantiates a new UpdateLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoadBalancerRequestWithDefaults() *UpdateLoadBalancerRequest {
	this := UpdateLoadBalancerRequest{}
	return &this
}

// GetAccessLog returns the AccessLog field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetAccessLog() AccessLog {
	if o == nil || o.AccessLog == nil {
		var ret AccessLog
		return ret
	}
	return *o.AccessLog
}

// GetAccessLogOk returns a tuple with the AccessLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetAccessLogOk() (*AccessLog, bool) {
	if o == nil || o.AccessLog == nil {
		return nil, false
	}
	return o.AccessLog, true
}

// HasAccessLog returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasAccessLog() bool {
	if o != nil && o.AccessLog != nil {
		return true
	}

	return false
}

// SetAccessLog gets a reference to the given AccessLog and assigns it to the AccessLog field.
func (o *UpdateLoadBalancerRequest) SetAccessLog(v AccessLog) {
	o.AccessLog = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateLoadBalancerRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetHealthCheck() HealthCheck {
	if o == nil || o.HealthCheck == nil {
		var ret HealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetHealthCheckOk() (*HealthCheck, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given HealthCheck and assigns it to the HealthCheck field.
func (o *UpdateLoadBalancerRequest) SetHealthCheck(v HealthCheck) {
	o.HealthCheck = &v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *UpdateLoadBalancerRequest) GetLoadBalancerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *UpdateLoadBalancerRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetLoadBalancerPort() int32 {
	if o == nil || o.LoadBalancerPort == nil {
		var ret int32
		return ret
	}
	return *o.LoadBalancerPort
}

// GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetLoadBalancerPortOk() (*int32, bool) {
	if o == nil || o.LoadBalancerPort == nil {
		return nil, false
	}
	return o.LoadBalancerPort, true
}

// HasLoadBalancerPort returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasLoadBalancerPort() bool {
	if o != nil && o.LoadBalancerPort != nil {
		return true
	}

	return false
}

// SetLoadBalancerPort gets a reference to the given int32 and assigns it to the LoadBalancerPort field.
func (o *UpdateLoadBalancerRequest) SetLoadBalancerPort(v int32) {
	o.LoadBalancerPort = &v
}

// GetPolicyNames returns the PolicyNames field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetPolicyNames() []string {
	if o == nil || o.PolicyNames == nil {
		var ret []string
		return ret
	}
	return *o.PolicyNames
}

// GetPolicyNamesOk returns a tuple with the PolicyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetPolicyNamesOk() (*[]string, bool) {
	if o == nil || o.PolicyNames == nil {
		return nil, false
	}
	return o.PolicyNames, true
}

// HasPolicyNames returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasPolicyNames() bool {
	if o != nil && o.PolicyNames != nil {
		return true
	}

	return false
}

// SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.
func (o *UpdateLoadBalancerRequest) SetPolicyNames(v []string) {
	o.PolicyNames = &v
}

// GetServerCertificateId returns the ServerCertificateId field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetServerCertificateId() string {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificateId
}

// GetServerCertificateIdOk returns a tuple with the ServerCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetServerCertificateIdOk() (*string, bool) {
	if o == nil || o.ServerCertificateId == nil {
		return nil, false
	}
	return o.ServerCertificateId, true
}

// HasServerCertificateId returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasServerCertificateId() bool {
	if o != nil && o.ServerCertificateId != nil {
		return true
	}

	return false
}

// SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.
func (o *UpdateLoadBalancerRequest) SetServerCertificateId(v string) {
	o.ServerCertificateId = &v
}

func (o UpdateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessLog != nil {
		toSerialize["AccessLog"] = o.AccessLog
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.HealthCheck != nil {
		toSerialize["HealthCheck"] = o.HealthCheck
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if o.LoadBalancerPort != nil {
		toSerialize["LoadBalancerPort"] = o.LoadBalancerPort
	}
	if o.PolicyNames != nil {
		toSerialize["PolicyNames"] = o.PolicyNames
	}
	if o.ServerCertificateId != nil {
		toSerialize["ServerCertificateId"] = o.ServerCertificateId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateLoadBalancerRequest struct {
	value *UpdateLoadBalancerRequest
	isSet bool
}

func (v NullableUpdateLoadBalancerRequest) Get() *UpdateLoadBalancerRequest {
	return v.value
}

func (v *NullableUpdateLoadBalancerRequest) Set(val *UpdateLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoadBalancerRequest(val *UpdateLoadBalancerRequest) *NullableUpdateLoadBalancerRequest {
	return &NullableUpdateLoadBalancerRequest{value: val, isSet: true}
}

func (v NullableUpdateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
