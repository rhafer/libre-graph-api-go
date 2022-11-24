/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// PasswordChange struct for PasswordChange
type PasswordChange struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
}

// NewPasswordChange instantiates a new PasswordChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordChange() *PasswordChange {
	this := PasswordChange{}
	return &this
}

// NewPasswordChangeWithDefaults instantiates a new PasswordChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordChangeWithDefaults() *PasswordChange {
	this := PasswordChange{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *PasswordChange) GetCurrentPassword() string {
	if o == nil || o.CurrentPassword == nil {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChange) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || o.CurrentPassword == nil {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *PasswordChange) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword != nil {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *PasswordChange) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *PasswordChange) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordChange) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *PasswordChange) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *PasswordChange) SetNewPassword(v string) {
	o.NewPassword = &v
}

func (o PasswordChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPassword != nil {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	if o.NewPassword != nil {
		toSerialize["newPassword"] = o.NewPassword
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordChange struct {
	value *PasswordChange
	isSet bool
}

func (v NullablePasswordChange) Get() *PasswordChange {
	return v.value
}

func (v *NullablePasswordChange) Set(val *PasswordChange) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordChange) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordChange(val *PasswordChange) *NullablePasswordChange {
	return &NullablePasswordChange{value: val, isSet: true}
}

func (v NullablePasswordChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


