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

// ExportPersonalDataRequest struct for ExportPersonalDataRequest
type ExportPersonalDataRequest struct {
	// the path where the file should be created in the users personal space
	StorageLocation *string `json:"storageLocation,omitempty"`
}

// NewExportPersonalDataRequest instantiates a new ExportPersonalDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportPersonalDataRequest() *ExportPersonalDataRequest {
	this := ExportPersonalDataRequest{}
	return &this
}

// NewExportPersonalDataRequestWithDefaults instantiates a new ExportPersonalDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportPersonalDataRequestWithDefaults() *ExportPersonalDataRequest {
	this := ExportPersonalDataRequest{}
	return &this
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise.
func (o *ExportPersonalDataRequest) GetStorageLocation() string {
	if o == nil || o.StorageLocation == nil {
		var ret string
		return ret
	}
	return *o.StorageLocation
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportPersonalDataRequest) GetStorageLocationOk() (*string, bool) {
	if o == nil || o.StorageLocation == nil {
		return nil, false
	}
	return o.StorageLocation, true
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *ExportPersonalDataRequest) HasStorageLocation() bool {
	if o != nil && o.StorageLocation != nil {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given string and assigns it to the StorageLocation field.
func (o *ExportPersonalDataRequest) SetStorageLocation(v string) {
	o.StorageLocation = &v
}

func (o ExportPersonalDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StorageLocation != nil {
		toSerialize["storageLocation"] = o.StorageLocation
	}
	return json.Marshal(toSerialize)
}

type NullableExportPersonalDataRequest struct {
	value *ExportPersonalDataRequest
	isSet bool
}

func (v NullableExportPersonalDataRequest) Get() *ExportPersonalDataRequest {
	return v.value
}

func (v *NullableExportPersonalDataRequest) Set(val *ExportPersonalDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportPersonalDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportPersonalDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportPersonalDataRequest(val *ExportPersonalDataRequest) *NullableExportPersonalDataRequest {
	return &NullableExportPersonalDataRequest{value: val, isSet: true}
}

func (v NullableExportPersonalDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportPersonalDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
