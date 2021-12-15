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

// OpenGraphFile File metadata, if the item is a file. Read-only.
type OpenGraphFile struct {
	Hashes *Hashes `json:"hashes,omitempty"`
	// The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
	MimeType *string `json:"mimeType,omitempty"`
	ProcessingMetadata *bool `json:"processingMetadata,omitempty"`
}

// NewOpenGraphFile instantiates a new OpenGraphFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphFile() *OpenGraphFile {
	this := OpenGraphFile{}
	return &this
}

// NewOpenGraphFileWithDefaults instantiates a new OpenGraphFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphFileWithDefaults() *OpenGraphFile {
	this := OpenGraphFile{}
	return &this
}

// GetHashes returns the Hashes field value if set, zero value otherwise.
func (o *OpenGraphFile) GetHashes() Hashes {
	if o == nil || o.Hashes == nil {
		var ret Hashes
		return ret
	}
	return *o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphFile) GetHashesOk() (*Hashes, bool) {
	if o == nil || o.Hashes == nil {
		return nil, false
	}
	return o.Hashes, true
}

// HasHashes returns a boolean if a field has been set.
func (o *OpenGraphFile) HasHashes() bool {
	if o != nil && o.Hashes != nil {
		return true
	}

	return false
}

// SetHashes gets a reference to the given Hashes and assigns it to the Hashes field.
func (o *OpenGraphFile) SetHashes(v Hashes) {
	o.Hashes = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *OpenGraphFile) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphFile) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *OpenGraphFile) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *OpenGraphFile) SetMimeType(v string) {
	o.MimeType = &v
}

// GetProcessingMetadata returns the ProcessingMetadata field value if set, zero value otherwise.
func (o *OpenGraphFile) GetProcessingMetadata() bool {
	if o == nil || o.ProcessingMetadata == nil {
		var ret bool
		return ret
	}
	return *o.ProcessingMetadata
}

// GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphFile) GetProcessingMetadataOk() (*bool, bool) {
	if o == nil || o.ProcessingMetadata == nil {
		return nil, false
	}
	return o.ProcessingMetadata, true
}

// HasProcessingMetadata returns a boolean if a field has been set.
func (o *OpenGraphFile) HasProcessingMetadata() bool {
	if o != nil && o.ProcessingMetadata != nil {
		return true
	}

	return false
}

// SetProcessingMetadata gets a reference to the given bool and assigns it to the ProcessingMetadata field.
func (o *OpenGraphFile) SetProcessingMetadata(v bool) {
	o.ProcessingMetadata = &v
}

func (o OpenGraphFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hashes != nil {
		toSerialize["hashes"] = o.Hashes
	}
	if o.MimeType != nil {
		toSerialize["mimeType"] = o.MimeType
	}
	if o.ProcessingMetadata != nil {
		toSerialize["processingMetadata"] = o.ProcessingMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphFile struct {
	value *OpenGraphFile
	isSet bool
}

func (v NullableOpenGraphFile) Get() *OpenGraphFile {
	return v.value
}

func (v *NullableOpenGraphFile) Set(val *OpenGraphFile) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphFile) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphFile(val *OpenGraphFile) *NullableOpenGraphFile {
	return &NullableOpenGraphFile{value: val, isSet: true}
}

func (v NullableOpenGraphFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


