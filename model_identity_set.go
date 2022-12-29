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

// IdentitySet Optional. User account.
type IdentitySet struct {
	Application *Identity `json:"application,omitempty"`
	Device *Identity `json:"device,omitempty"`
	User *Identity `json:"user,omitempty"`
	Group *Identity `json:"group,omitempty"`
}

// NewIdentitySet instantiates a new IdentitySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySet() *IdentitySet {
	this := IdentitySet{}
	return &this
}

// NewIdentitySetWithDefaults instantiates a new IdentitySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySetWithDefaults() *IdentitySet {
	this := IdentitySet{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *IdentitySet) GetApplication() Identity {
	if o == nil || o.Application == nil {
		var ret Identity
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySet) GetApplicationOk() (*Identity, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *IdentitySet) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given Identity and assigns it to the Application field.
func (o *IdentitySet) SetApplication(v Identity) {
	o.Application = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *IdentitySet) GetDevice() Identity {
	if o == nil || o.Device == nil {
		var ret Identity
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySet) GetDeviceOk() (*Identity, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *IdentitySet) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Identity and assigns it to the Device field.
func (o *IdentitySet) SetDevice(v Identity) {
	o.Device = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IdentitySet) GetUser() Identity {
	if o == nil || o.User == nil {
		var ret Identity
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySet) GetUserOk() (*Identity, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IdentitySet) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given Identity and assigns it to the User field.
func (o *IdentitySet) SetUser(v Identity) {
	o.User = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IdentitySet) GetGroup() Identity {
	if o == nil || o.Group == nil {
		var ret Identity
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySet) GetGroupOk() (*Identity, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IdentitySet) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Identity and assigns it to the Group field.
func (o *IdentitySet) SetGroup(v Identity) {
	o.Group = &v
}

func (o IdentitySet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySet struct {
	value *IdentitySet
	isSet bool
}

func (v NullableIdentitySet) Get() *IdentitySet {
	return v.value
}

func (v *NullableIdentitySet) Set(val *IdentitySet) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySet) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySet(val *IdentitySet) *NullableIdentitySet {
	return &NullableIdentitySet{value: val, isSet: true}
}

func (v NullableIdentitySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


