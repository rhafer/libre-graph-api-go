/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"fmt"
)

// SharingLinkType The type of the link created.  | Value          | Display name      | Description                                                     | | -------------- | ----------------- | --------------------------------------------------------------- | | internal       | Internal          | Creates an internal link without any permissions.               | | view           | View              | Creates a read-only link to the driveItem.                      | | upload         | Upload            | Creates a read-write link to the folder driveItem.              | | edit           | Edit              | Creates a read-write link to the driveItem.                     | | createOnly     | File Drop         | Creates an upload-only link to the folder driveItem.            | | blocksDownload | Secure View       | Creates a read-only link that blocks download to the driveItem. |
type SharingLinkType string

// List of sharingLinkType
const (
	INTERNAL        SharingLinkType = "internal"
	VIEW            SharingLinkType = "view"
	UPLOAD          SharingLinkType = "upload"
	EDIT            SharingLinkType = "edit"
	CREATE_ONLY     SharingLinkType = "createOnly"
	BLOCKS_DOWNLOAD SharingLinkType = "blocksDownload"
)

// All allowed values of SharingLinkType enum
var AllowedSharingLinkTypeEnumValues = []SharingLinkType{
	"internal",
	"view",
	"upload",
	"edit",
	"createOnly",
	"blocksDownload",
}

func (v *SharingLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SharingLinkType(value)
	for _, existing := range AllowedSharingLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SharingLinkType", value)
}

// NewSharingLinkTypeFromValue returns a pointer to a valid SharingLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharingLinkTypeFromValue(v string) (*SharingLinkType, error) {
	ev := SharingLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharingLinkType: valid values are %v", v, AllowedSharingLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharingLinkType) IsValid() bool {
	for _, existing := range AllowedSharingLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sharingLinkType value
func (v SharingLinkType) Ptr() *SharingLinkType {
	return &v
}

type NullableSharingLinkType struct {
	value *SharingLinkType
	isSet bool
}

func (v NullableSharingLinkType) Get() *SharingLinkType {
	return v.value
}

func (v *NullableSharingLinkType) Set(val *SharingLinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingLinkType(val *SharingLinkType) *NullableSharingLinkType {
	return &NullableSharingLinkType{value: val, isSet: true}
}

func (v NullableSharingLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
