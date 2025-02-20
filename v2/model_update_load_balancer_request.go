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

// UpdateLoadBalancerRequest struct for UpdateLoadBalancerRequest
type UpdateLoadBalancerRequest struct {
	AccessLog *AccessLog `json:"AccessLog,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun      *bool        `json:"DryRun,omitempty"`
	HealthCheck *HealthCheck `json:"HealthCheck,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included). This parameter is required if you want to update the server certificate.
	LoadBalancerPort *int32 `json:"LoadBalancerPort,omitempty"`
	// The name of the policy you want to enable for the listener.
	PolicyNames *[]string `json:"PolicyNames,omitempty"`
	// (internet-facing only) The public IP you want to associate with the load balancer. The former public IP of the load balancer is then disassociated. If you specify an empty string and the former public IP belonged to you, it is disassociated and replaced by a public IP owned by 3DS OUTSCALE.
	PublicIp *string `json:"PublicIp,omitempty"`
	// If true, secure cookies are enabled for the load balancer.
	SecuredCookies *bool `json:"SecuredCookies,omitempty"`
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. You need to specify the already assigned security groups that you want to keep along with the new ones you are assigning. If the list is empty, the default security group of the Net is assigned to the load balancer.
	SecurityGroups *[]string `json:"SecurityGroups,omitempty"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://docs.outscale.com/en/userguide/Resource-Identifiers.html#_outscale_resource_names_orns). If this parameter is specified, you must also specify the `LoadBalancerPort` parameter.
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

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *UpdateLoadBalancerRequest) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetSecuredCookies returns the SecuredCookies field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetSecuredCookies() bool {
	if o == nil || o.SecuredCookies == nil {
		var ret bool
		return ret
	}
	return *o.SecuredCookies
}

// GetSecuredCookiesOk returns a tuple with the SecuredCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetSecuredCookiesOk() (*bool, bool) {
	if o == nil || o.SecuredCookies == nil {
		return nil, false
	}
	return o.SecuredCookies, true
}

// HasSecuredCookies returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasSecuredCookies() bool {
	if o != nil && o.SecuredCookies != nil {
		return true
	}

	return false
}

// SetSecuredCookies gets a reference to the given bool and assigns it to the SecuredCookies field.
func (o *UpdateLoadBalancerRequest) SetSecuredCookies(v bool) {
	o.SecuredCookies = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *UpdateLoadBalancerRequest) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerRequest) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *UpdateLoadBalancerRequest) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *UpdateLoadBalancerRequest) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
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
	if o.PublicIp != nil {
		toSerialize["PublicIp"] = o.PublicIp
	}
	if o.SecuredCookies != nil {
		toSerialize["SecuredCookies"] = o.SecuredCookies
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
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
