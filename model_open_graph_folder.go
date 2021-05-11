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

// OpenGraphFolder struct for OpenGraphFolder
type OpenGraphFolder struct {
	// Number of children contained immediately within this container.
	ChildCount *int32 `json:"childCount,omitempty"`
	// A collection of properties defining the recommended view for the folder.
	View *OpenGraphFolderView `json:"view,omitempty"`
}

// NewOpenGraphFolder instantiates a new OpenGraphFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphFolder() *OpenGraphFolder {
	this := OpenGraphFolder{}
	return &this
}

// NewOpenGraphFolderWithDefaults instantiates a new OpenGraphFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphFolderWithDefaults() *OpenGraphFolder {
	this := OpenGraphFolder{}
	return &this
}

// GetChildCount returns the ChildCount field value if set, zero value otherwise.
func (o *OpenGraphFolder) GetChildCount() int32 {
	if o == nil || o.ChildCount == nil {
		var ret int32
		return ret
	}
	return *o.ChildCount
}

// GetChildCountOk returns a tuple with the ChildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphFolder) GetChildCountOk() (*int32, bool) {
	if o == nil || o.ChildCount == nil {
		return nil, false
	}
	return o.ChildCount, true
}

// HasChildCount returns a boolean if a field has been set.
func (o *OpenGraphFolder) HasChildCount() bool {
	if o != nil && o.ChildCount != nil {
		return true
	}

	return false
}

// SetChildCount gets a reference to the given int32 and assigns it to the ChildCount field.
func (o *OpenGraphFolder) SetChildCount(v int32) {
	o.ChildCount = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *OpenGraphFolder) GetView() OpenGraphFolderView {
	if o == nil || o.View == nil {
		var ret OpenGraphFolderView
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphFolder) GetViewOk() (*OpenGraphFolderView, bool) {
	if o == nil || o.View == nil {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *OpenGraphFolder) HasView() bool {
	if o != nil && o.View != nil {
		return true
	}

	return false
}

// SetView gets a reference to the given OpenGraphFolderView and assigns it to the View field.
func (o *OpenGraphFolder) SetView(v OpenGraphFolderView) {
	o.View = &v
}

func (o OpenGraphFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChildCount != nil {
		toSerialize["childCount"] = o.ChildCount
	}
	if o.View != nil {
		toSerialize["view"] = o.View
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphFolder struct {
	value *OpenGraphFolder
	isSet bool
}

func (v NullableOpenGraphFolder) Get() *OpenGraphFolder {
	return v.value
}

func (v *NullableOpenGraphFolder) Set(val *OpenGraphFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphFolder(val *OpenGraphFolder) *NullableOpenGraphFolder {
	return &NullableOpenGraphFolder{value: val, isSet: true}
}

func (v NullableOpenGraphFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


