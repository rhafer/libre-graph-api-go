/*
 * Open Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OpenGraphDeleted struct for OpenGraphDeleted
type OpenGraphDeleted struct {
	// Represents the state of the deleted item.
	State *string `json:"state,omitempty"`
}

// NewOpenGraphDeleted instantiates a new OpenGraphDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphDeleted() *OpenGraphDeleted {
	this := OpenGraphDeleted{}
	return &this
}

// NewOpenGraphDeletedWithDefaults instantiates a new OpenGraphDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphDeletedWithDefaults() *OpenGraphDeleted {
	this := OpenGraphDeleted{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OpenGraphDeleted) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphDeleted) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OpenGraphDeleted) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OpenGraphDeleted) SetState(v string) {
	o.State = &v
}

func (o OpenGraphDeleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphDeleted struct {
	value *OpenGraphDeleted
	isSet bool
}

func (v NullableOpenGraphDeleted) Get() *OpenGraphDeleted {
	return v.value
}

func (v *NullableOpenGraphDeleted) Set(val *OpenGraphDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphDeleted(val *OpenGraphDeleted) *NullableOpenGraphDeleted {
	return &NullableOpenGraphDeleted{value: val, isSet: true}
}

func (v NullableOpenGraphDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


