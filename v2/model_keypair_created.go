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

// KeypairCreated Information about the created keypair.
type KeypairCreated struct {
	// The MD5 public key fingerprint as specified in section 4 of RFC 4716.
	KeypairFingerprint *string `json:"KeypairFingerprint,omitempty"`
	// The name of the keypair.
	KeypairName *string `json:"KeypairName,omitempty"`
	// The private key. When saving the private key in a .rsa file, replace the `\\n` escape sequences with line breaks.
	PrivateKey *string `json:"PrivateKey,omitempty"`
}

// NewKeypairCreated instantiates a new KeypairCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeypairCreated() *KeypairCreated {
	this := KeypairCreated{}
	return &this
}

// NewKeypairCreatedWithDefaults instantiates a new KeypairCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeypairCreatedWithDefaults() *KeypairCreated {
	this := KeypairCreated{}
	return &this
}

// GetKeypairFingerprint returns the KeypairFingerprint field value if set, zero value otherwise.
func (o *KeypairCreated) GetKeypairFingerprint() string {
	if o == nil || o.KeypairFingerprint == nil {
		var ret string
		return ret
	}
	return *o.KeypairFingerprint
}

// GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetKeypairFingerprintOk() (*string, bool) {
	if o == nil || o.KeypairFingerprint == nil {
		return nil, false
	}
	return o.KeypairFingerprint, true
}

// HasKeypairFingerprint returns a boolean if a field has been set.
func (o *KeypairCreated) HasKeypairFingerprint() bool {
	if o != nil && o.KeypairFingerprint != nil {
		return true
	}

	return false
}

// SetKeypairFingerprint gets a reference to the given string and assigns it to the KeypairFingerprint field.
func (o *KeypairCreated) SetKeypairFingerprint(v string) {
	o.KeypairFingerprint = &v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *KeypairCreated) GetKeypairName() string {
	if o == nil || o.KeypairName == nil {
		var ret string
		return ret
	}
	return *o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetKeypairNameOk() (*string, bool) {
	if o == nil || o.KeypairName == nil {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *KeypairCreated) HasKeypairName() bool {
	if o != nil && o.KeypairName != nil {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *KeypairCreated) SetKeypairName(v string) {
	o.KeypairName = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *KeypairCreated) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeypairCreated) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *KeypairCreated) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *KeypairCreated) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

func (o KeypairCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeypairFingerprint != nil {
		toSerialize["KeypairFingerprint"] = o.KeypairFingerprint
	}
	if o.KeypairName != nil {
		toSerialize["KeypairName"] = o.KeypairName
	}
	if o.PrivateKey != nil {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	return json.Marshal(toSerialize)
}

type NullableKeypairCreated struct {
	value *KeypairCreated
	isSet bool
}

func (v NullableKeypairCreated) Get() *KeypairCreated {
	return v.value
}

func (v *NullableKeypairCreated) Set(val *KeypairCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableKeypairCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableKeypairCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeypairCreated(val *KeypairCreated) *NullableKeypairCreated {
	return &NullableKeypairCreated{value: val, isSet: true}
}

func (v NullableKeypairCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeypairCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
