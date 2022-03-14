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

// AccessKeySecretKey Information about the access key.
type AccessKeySecretKey struct {
	// The ID of the access key.
	AccessKeyId *string `json:"AccessKeyId,omitempty"`
	// The date and time (UTC) of creation of the access key.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The date and time (UTC) at which the access key expires.
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// The date and time (UTC) of the last modification of the access key.
	LastModificationDate *string `json:"LastModificationDate,omitempty"`
	// The access key that enables you to send requests.
	SecretKey *string `json:"SecretKey,omitempty"`
	// The state of the access key (`ACTIVE` if the key is valid for API calls, or `INACTIVE` if not).
	State *string `json:"State,omitempty"`
}

// NewAccessKeySecretKey instantiates a new AccessKeySecretKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeySecretKey() *AccessKeySecretKey {
	this := AccessKeySecretKey{}
	return &this
}

// NewAccessKeySecretKeyWithDefaults instantiates a new AccessKeySecretKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeySecretKeyWithDefaults() *AccessKeySecretKey {
	this := AccessKeySecretKey{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasAccessKeyId() bool {
	if o != nil && o.AccessKeyId != nil {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AccessKeySecretKey) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *AccessKeySecretKey) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *AccessKeySecretKey) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetLastModificationDate returns the LastModificationDate field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetLastModificationDate() string {
	if o == nil || o.LastModificationDate == nil {
		var ret string
		return ret
	}
	return *o.LastModificationDate
}

// GetLastModificationDateOk returns a tuple with the LastModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetLastModificationDateOk() (*string, bool) {
	if o == nil || o.LastModificationDate == nil {
		return nil, false
	}
	return o.LastModificationDate, true
}

// HasLastModificationDate returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasLastModificationDate() bool {
	if o != nil && o.LastModificationDate != nil {
		return true
	}

	return false
}

// SetLastModificationDate gets a reference to the given string and assigns it to the LastModificationDate field.
func (o *AccessKeySecretKey) SetLastModificationDate(v string) {
	o.LastModificationDate = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AccessKeySecretKey) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccessKeySecretKey) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeySecretKey) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccessKeySecretKey) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccessKeySecretKey) SetState(v string) {
	o.State = &v
}

func (o AccessKeySecretKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeyId != nil {
		toSerialize["AccessKeyId"] = o.AccessKeyId
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if o.LastModificationDate != nil {
		toSerialize["LastModificationDate"] = o.LastModificationDate
	}
	if o.SecretKey != nil {
		toSerialize["SecretKey"] = o.SecretKey
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAccessKeySecretKey struct {
	value *AccessKeySecretKey
	isSet bool
}

func (v NullableAccessKeySecretKey) Get() *AccessKeySecretKey {
	return v.value
}

func (v *NullableAccessKeySecretKey) Set(val *AccessKeySecretKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeySecretKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeySecretKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeySecretKey(val *AccessKeySecretKey) *NullableAccessKeySecretKey {
	return &NullableAccessKeySecretKey{value: val, isSet: true}
}

func (v NullableAccessKeySecretKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeySecretKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
