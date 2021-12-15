/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// OdataError struct for OdataError
type OdataError struct {
	Error OdataErrorMain `json:"error"`
}

// NewOdataError instantiates a new OdataError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdataError(error_ OdataErrorMain) *OdataError {
	this := OdataError{}
	this.Error = error_
	return &this
}

// NewOdataErrorWithDefaults instantiates a new OdataError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdataErrorWithDefaults() *OdataError {
	this := OdataError{}
	return &this
}

// GetError returns the Error field value
func (o *OdataError) GetError() OdataErrorMain {
	if o == nil {
		var ret OdataErrorMain
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *OdataError) GetErrorOk() (*OdataErrorMain, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *OdataError) SetError(v OdataErrorMain) {
	o.Error = v
}

func (o OdataError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableOdataError struct {
	value *OdataError
	isSet bool
}

func (v NullableOdataError) Get() *OdataError {
	return v.value
}

func (v *NullableOdataError) Set(val *OdataError) {
	v.value = val
	v.isSet = true
}

func (v NullableOdataError) IsSet() bool {
	return v.isSet
}

func (v *NullableOdataError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdataError(val *OdataError) *NullableOdataError {
	return &NullableOdataError{value: val, isSet: true}
}

func (v NullableOdataError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdataError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


