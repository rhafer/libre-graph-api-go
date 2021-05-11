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
	"time"
)

// OpenGraphBaseItemAllOf struct for OpenGraphBaseItemAllOf
type OpenGraphBaseItemAllOf struct {
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *OpenGraphIdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *OpenGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *OpenGraphItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *OpenGraphUser `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *OpenGraphUser `json:"lastModifiedByUser,omitempty"`
}

// NewOpenGraphBaseItemAllOf instantiates a new OpenGraphBaseItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphBaseItemAllOf() *OpenGraphBaseItemAllOf {
	this := OpenGraphBaseItemAllOf{}
	return &this
}

// NewOpenGraphBaseItemAllOfWithDefaults instantiates a new OpenGraphBaseItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphBaseItemAllOfWithDefaults() *OpenGraphBaseItemAllOf {
	this := OpenGraphBaseItemAllOf{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetCreatedBy() OpenGraphIdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret OpenGraphIdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetCreatedByOk() (*OpenGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given OpenGraphIdentitySet and assigns it to the CreatedBy field.
func (o *OpenGraphBaseItemAllOf) SetCreatedBy(v OpenGraphIdentitySet) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *OpenGraphBaseItemAllOf) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenGraphBaseItemAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetETagOk() (*string, bool) {
	if o == nil || o.ETag == nil {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *OpenGraphBaseItemAllOf) SetETag(v string) {
	o.ETag = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedBy() OpenGraphIdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret OpenGraphIdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByOk() (*OpenGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given OpenGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedBy(v OpenGraphIdentitySet) {
	o.LastModifiedBy = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenGraphBaseItemAllOf) SetName(v string) {
	o.Name = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetParentReference() OpenGraphItemReference {
	if o == nil || o.ParentReference == nil {
		var ret OpenGraphItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetParentReferenceOk() (*OpenGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given OpenGraphItemReference and assigns it to the ParentReference field.
func (o *OpenGraphBaseItemAllOf) SetParentReference(v OpenGraphItemReference) {
	o.ParentReference = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetWebUrlOk() (*string, bool) {
	if o == nil || o.WebUrl == nil {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *OpenGraphBaseItemAllOf) SetWebUrl(v string) {
	o.WebUrl = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetCreatedByUser() OpenGraphUser {
	if o == nil || o.CreatedByUser == nil {
		var ret OpenGraphUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetCreatedByUserOk() (*OpenGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given OpenGraphUser and assigns it to the CreatedByUser field.
func (o *OpenGraphBaseItemAllOf) SetCreatedByUser(v OpenGraphUser) {
	o.CreatedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByUser() OpenGraphUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret OpenGraphUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphBaseItemAllOf) GetLastModifiedByUserOk() (*OpenGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *OpenGraphBaseItemAllOf) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given OpenGraphUser and assigns it to the LastModifiedByUser field.
func (o *OpenGraphBaseItemAllOf) SetLastModifiedByUser(v OpenGraphUser) {
	o.LastModifiedByUser = &v
}

func (o OpenGraphBaseItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ETag != nil {
		toSerialize["eTag"] = o.ETag
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.WebUrl != nil {
		toSerialize["webUrl"] = o.WebUrl
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.LastModifiedByUser != nil {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphBaseItemAllOf struct {
	value *OpenGraphBaseItemAllOf
	isSet bool
}

func (v NullableOpenGraphBaseItemAllOf) Get() *OpenGraphBaseItemAllOf {
	return v.value
}

func (v *NullableOpenGraphBaseItemAllOf) Set(val *OpenGraphBaseItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphBaseItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphBaseItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphBaseItemAllOf(val *OpenGraphBaseItemAllOf) *NullableOpenGraphBaseItemAllOf {
	return &NullableOpenGraphBaseItemAllOf{value: val, isSet: true}
}

func (v NullableOpenGraphBaseItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphBaseItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


