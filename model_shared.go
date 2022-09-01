/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.17.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// Shared struct for Shared
type Shared struct {
	Owner *IdentitySet `json:"owner,omitempty"`
	// Indicates the scope of how the item is shared: anonymous, organization, or users. Read-only.
	Scope *string `json:"scope,omitempty"`
	SharedBy *IdentitySet `json:"sharedBy,omitempty"`
	// The UTC date and time when the item was shared. Read-only.
	SharedDateTime *time.Time `json:"sharedDateTime,omitempty"`
}

// NewShared instantiates a new Shared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShared() *Shared {
	this := Shared{}
	return &this
}

// NewSharedWithDefaults instantiates a new Shared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedWithDefaults() *Shared {
	this := Shared{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Shared) GetOwner() IdentitySet {
	if o == nil || o.Owner == nil {
		var ret IdentitySet
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shared) GetOwnerOk() (*IdentitySet, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Shared) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given IdentitySet and assigns it to the Owner field.
func (o *Shared) SetOwner(v IdentitySet) {
	o.Owner = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Shared) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shared) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Shared) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *Shared) SetScope(v string) {
	o.Scope = &v
}

// GetSharedBy returns the SharedBy field value if set, zero value otherwise.
func (o *Shared) GetSharedBy() IdentitySet {
	if o == nil || o.SharedBy == nil {
		var ret IdentitySet
		return ret
	}
	return *o.SharedBy
}

// GetSharedByOk returns a tuple with the SharedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shared) GetSharedByOk() (*IdentitySet, bool) {
	if o == nil || o.SharedBy == nil {
		return nil, false
	}
	return o.SharedBy, true
}

// HasSharedBy returns a boolean if a field has been set.
func (o *Shared) HasSharedBy() bool {
	if o != nil && o.SharedBy != nil {
		return true
	}

	return false
}

// SetSharedBy gets a reference to the given IdentitySet and assigns it to the SharedBy field.
func (o *Shared) SetSharedBy(v IdentitySet) {
	o.SharedBy = &v
}

// GetSharedDateTime returns the SharedDateTime field value if set, zero value otherwise.
func (o *Shared) GetSharedDateTime() time.Time {
	if o == nil || o.SharedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.SharedDateTime
}

// GetSharedDateTimeOk returns a tuple with the SharedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Shared) GetSharedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.SharedDateTime == nil {
		return nil, false
	}
	return o.SharedDateTime, true
}

// HasSharedDateTime returns a boolean if a field has been set.
func (o *Shared) HasSharedDateTime() bool {
	if o != nil && o.SharedDateTime != nil {
		return true
	}

	return false
}

// SetSharedDateTime gets a reference to the given time.Time and assigns it to the SharedDateTime field.
func (o *Shared) SetSharedDateTime(v time.Time) {
	o.SharedDateTime = &v
}

func (o Shared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.SharedBy != nil {
		toSerialize["sharedBy"] = o.SharedBy
	}
	if o.SharedDateTime != nil {
		toSerialize["sharedDateTime"] = o.SharedDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableShared struct {
	value *Shared
	isSet bool
}

func (v NullableShared) Get() *Shared {
	return v.value
}

func (v *NullableShared) Set(val *Shared) {
	v.value = val
	v.isSet = true
}

func (v NullableShared) IsSet() bool {
	return v.isSet
}

func (v *NullableShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShared(val *Shared) *NullableShared {
	return &NullableShared{value: val, isSet: true}
}

func (v NullableShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


