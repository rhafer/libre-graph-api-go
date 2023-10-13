/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// CollectionOfDriveItems1 struct for CollectionOfDriveItems1
type CollectionOfDriveItems1 struct {
	Value []DriveItem `json:"value,omitempty"`
}

// NewCollectionOfDriveItems1 instantiates a new CollectionOfDriveItems1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDriveItems1() *CollectionOfDriveItems1 {
	this := CollectionOfDriveItems1{}
	return &this
}

// NewCollectionOfDriveItems1WithDefaults instantiates a new CollectionOfDriveItems1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDriveItems1WithDefaults() *CollectionOfDriveItems1 {
	this := CollectionOfDriveItems1{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDriveItems1) GetValue() []DriveItem {
	if o == nil || o.Value == nil {
		var ret []DriveItem
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDriveItems1) GetValueOk() ([]DriveItem, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDriveItems1) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []DriveItem and assigns it to the Value field.
func (o *CollectionOfDriveItems1) SetValue(v []DriveItem) {
	o.Value = v
}

func (o CollectionOfDriveItems1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDriveItems1 struct {
	value *CollectionOfDriveItems1
	isSet bool
}

func (v NullableCollectionOfDriveItems1) Get() *CollectionOfDriveItems1 {
	return v.value
}

func (v *NullableCollectionOfDriveItems1) Set(val *CollectionOfDriveItems1) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDriveItems1) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDriveItems1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDriveItems1(val *CollectionOfDriveItems1) *NullableCollectionOfDriveItems1 {
	return &NullableCollectionOfDriveItems1{value: val, isSet: true}
}

func (v NullableCollectionOfDriveItems1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDriveItems1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}