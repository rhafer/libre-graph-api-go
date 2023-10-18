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

// checks if the FolderView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderView{}

// FolderView A collection of properties defining the recommended view for the folder.
type FolderView struct {
	// The method by which the folder should be sorted.
	SortBy *string `json:"sortBy,omitempty"`
	// If true, indicates that items should be sorted in descending order. Otherwise, items should be sorted ascending.
	SortOrder *string `json:"sortOrder,omitempty"`
	// The type of view that should be used to represent the folder.
	ViewType *string `json:"viewType,omitempty"`
}

// NewFolderView instantiates a new FolderView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderView() *FolderView {
	this := FolderView{}
	return &this
}

// NewFolderViewWithDefaults instantiates a new FolderView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderViewWithDefaults() *FolderView {
	this := FolderView{}
	return &this
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *FolderView) GetSortBy() string {
	if o == nil || IsNil(o.SortBy) {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderView) GetSortByOk() (*string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *FolderView) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *FolderView) SetSortBy(v string) {
	o.SortBy = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *FolderView) GetSortOrder() string {
	if o == nil || IsNil(o.SortOrder) {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderView) GetSortOrderOk() (*string, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *FolderView) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *FolderView) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *FolderView) GetViewType() string {
	if o == nil || IsNil(o.ViewType) {
		var ret string
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderView) GetViewTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ViewType) {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *FolderView) HasViewType() bool {
	if o != nil && !IsNil(o.ViewType) {
		return true
	}

	return false
}

// SetViewType gets a reference to the given string and assigns it to the ViewType field.
func (o *FolderView) SetViewType(v string) {
	o.ViewType = &v
}

func (o FolderView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SortBy) {
		toSerialize["sortBy"] = o.SortBy
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.ViewType) {
		toSerialize["viewType"] = o.ViewType
	}
	return toSerialize, nil
}

type NullableFolderView struct {
	value *FolderView
	isSet bool
}

func (v NullableFolderView) Get() *FolderView {
	return v.value
}

func (v *NullableFolderView) Set(val *FolderView) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderView) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderView(val *FolderView) *NullableFolderView {
	return &NullableFolderView{value: val, isSet: true}
}

func (v NullableFolderView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
