/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// AppRole struct for AppRole
type AppRole struct {
	// Specifies whether this app role can be assigned to users and groups (by setting to ['User']), to other application's (by setting to ['Application'], or both (by setting to ['User', 'Application']). App roles supporting assignment to other applications' service principals are also known as application permissions. The 'Application' value is only supported for app roles defined on application entities.
	AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
	// The description for the app role. This is displayed when the app role is being assigned and, if the app role functions as an application permission, during  consent experiences.
	Description NullableString `json:"description,omitempty"`
	// Display name for the permission that appears in the app role assignment and consent experiences.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique role identifier inside the appRoles collection. When creating a new app role, a new GUID identifier must be provided.
	Id string `json:"id"`
}

// NewAppRole instantiates a new AppRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRole(id string) *AppRole {
	this := AppRole{}
	this.Id = id
	return &this
}

// NewAppRoleWithDefaults instantiates a new AppRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWithDefaults() *AppRole {
	this := AppRole{}
	return &this
}

// GetAllowedMemberTypes returns the AllowedMemberTypes field value if set, zero value otherwise.
func (o *AppRole) GetAllowedMemberTypes() []string {
	if o == nil || o.AllowedMemberTypes == nil {
		var ret []string
		return ret
	}
	return o.AllowedMemberTypes
}

// GetAllowedMemberTypesOk returns a tuple with the AllowedMemberTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRole) GetAllowedMemberTypesOk() ([]string, bool) {
	if o == nil || o.AllowedMemberTypes == nil {
		return nil, false
	}
	return o.AllowedMemberTypes, true
}

// HasAllowedMemberTypes returns a boolean if a field has been set.
func (o *AppRole) HasAllowedMemberTypes() bool {
	if o != nil && o.AllowedMemberTypes != nil {
		return true
	}

	return false
}

// SetAllowedMemberTypes gets a reference to the given []string and assigns it to the AllowedMemberTypes field.
func (o *AppRole) SetAllowedMemberTypes(v []string) {
	o.AllowedMemberTypes = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRole) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRole) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppRole) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRole) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRole) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppRole) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AppRole) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AppRole) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AppRole) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value
func (o *AppRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRole) SetId(v string) {
	o.Id = v
}

func (o AppRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedMemberTypes != nil {
		toSerialize["allowedMemberTypes"] = o.AllowedMemberTypes
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAppRole struct {
	value *AppRole
	isSet bool
}

func (v NullableAppRole) Get() *AppRole {
	return v.value
}

func (v *NullableAppRole) Set(val *AppRole) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRole) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRole(val *AppRole) *NullableAppRole {
	return &NullableAppRole{value: val, isSet: true}
}

func (v NullableAppRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

