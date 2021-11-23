/*
Open Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
	"time"
)

// FileSystemInfo File system information on client. Read-write.
type FileSystemInfo struct {
	// The UTC date and time the file was created on a client.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The UTC date and time the file was last accessed. Available for the recent file list only.
	LastAccessedDateTime *time.Time `json:"lastAccessedDateTime,omitempty"`
	// The UTC date and time the file was last modified on a client.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// NewFileSystemInfo instantiates a new FileSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSystemInfo() *FileSystemInfo {
	this := FileSystemInfo{}
	return &this
}

// NewFileSystemInfoWithDefaults instantiates a new FileSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSystemInfoWithDefaults() *FileSystemInfo {
	this := FileSystemInfo{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *FileSystemInfo) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemInfo) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *FileSystemInfo) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *FileSystemInfo) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetLastAccessedDateTime returns the LastAccessedDateTime field value if set, zero value otherwise.
func (o *FileSystemInfo) GetLastAccessedDateTime() time.Time {
	if o == nil || o.LastAccessedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedDateTime
}

// GetLastAccessedDateTimeOk returns a tuple with the LastAccessedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemInfo) GetLastAccessedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessedDateTime == nil {
		return nil, false
	}
	return o.LastAccessedDateTime, true
}

// HasLastAccessedDateTime returns a boolean if a field has been set.
func (o *FileSystemInfo) HasLastAccessedDateTime() bool {
	if o != nil && o.LastAccessedDateTime != nil {
		return true
	}

	return false
}

// SetLastAccessedDateTime gets a reference to the given time.Time and assigns it to the LastAccessedDateTime field.
func (o *FileSystemInfo) SetLastAccessedDateTime(v time.Time) {
	o.LastAccessedDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *FileSystemInfo) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSystemInfo) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *FileSystemInfo) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *FileSystemInfo) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

func (o FileSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.LastAccessedDateTime != nil {
		toSerialize["lastAccessedDateTime"] = o.LastAccessedDateTime
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableFileSystemInfo struct {
	value *FileSystemInfo
	isSet bool
}

func (v NullableFileSystemInfo) Get() *FileSystemInfo {
	return v.value
}

func (v *NullableFileSystemInfo) Set(val *FileSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSystemInfo(val *FileSystemInfo) *NullableFileSystemInfo {
	return &NullableFileSystemInfo{value: val, isSet: true}
}

func (v NullableFileSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


