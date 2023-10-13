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

// UnifiedRoleDefinition A role definition is a collection of permissions in libre graph listing the operations that can be performed and the resources against which they can performed.
type UnifiedRoleDefinition struct {
	// The description for the unifiedRoleDefinition.
	Description *string `json:"description,omitempty"`
	// The display name for the unifiedRoleDefinition. Required. Supports $filter (`eq`, `in`).
	DisplayName *string `json:"displayName,omitempty"`
	// The unique identifier for the role definition. Key, not nullable, Read-only. Inherited from entity. Supports $filter (`eq`, `in`).
	Id *string `json:"id,omitempty"`
	// List of permissions included in the role.
	RolePermissions []UnifiedRolePermission `json:"rolePermissions,omitempty"`
	// When presenting a list of roles the weight can be used to order them in a meaningful way. Lower weight gets higher precedence. So content with lower weight will come first. If set, weights should be non-zero, as 0 is interpreted as an unset weight.
	LibreGraphWeight *int32 `json:"@libre.graph.weight,omitempty"`
}

// NewUnifiedRoleDefinition instantiates a new UnifiedRoleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnifiedRoleDefinition() *UnifiedRoleDefinition {
	this := UnifiedRoleDefinition{}
	return &this
}

// NewUnifiedRoleDefinitionWithDefaults instantiates a new UnifiedRoleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnifiedRoleDefinitionWithDefaults() *UnifiedRoleDefinition {
	this := UnifiedRoleDefinition{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnifiedRoleDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UnifiedRoleDefinition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UnifiedRoleDefinition) SetId(v string) {
	o.Id = &v
}

// GetRolePermissions returns the RolePermissions field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetRolePermissions() []UnifiedRolePermission {
	if o == nil || o.RolePermissions == nil {
		var ret []UnifiedRolePermission
		return ret
	}
	return o.RolePermissions
}

// GetRolePermissionsOk returns a tuple with the RolePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetRolePermissionsOk() ([]UnifiedRolePermission, bool) {
	if o == nil || o.RolePermissions == nil {
		return nil, false
	}
	return o.RolePermissions, true
}

// HasRolePermissions returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasRolePermissions() bool {
	if o != nil && o.RolePermissions != nil {
		return true
	}

	return false
}

// SetRolePermissions gets a reference to the given []UnifiedRolePermission and assigns it to the RolePermissions field.
func (o *UnifiedRoleDefinition) SetRolePermissions(v []UnifiedRolePermission) {
	o.RolePermissions = v
}

// GetLibreGraphWeight returns the LibreGraphWeight field value if set, zero value otherwise.
func (o *UnifiedRoleDefinition) GetLibreGraphWeight() int32 {
	if o == nil || o.LibreGraphWeight == nil {
		var ret int32
		return ret
	}
	return *o.LibreGraphWeight
}

// GetLibreGraphWeightOk returns a tuple with the LibreGraphWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnifiedRoleDefinition) GetLibreGraphWeightOk() (*int32, bool) {
	if o == nil || o.LibreGraphWeight == nil {
		return nil, false
	}
	return o.LibreGraphWeight, true
}

// HasLibreGraphWeight returns a boolean if a field has been set.
func (o *UnifiedRoleDefinition) HasLibreGraphWeight() bool {
	if o != nil && o.LibreGraphWeight != nil {
		return true
	}

	return false
}

// SetLibreGraphWeight gets a reference to the given int32 and assigns it to the LibreGraphWeight field.
func (o *UnifiedRoleDefinition) SetLibreGraphWeight(v int32) {
	o.LibreGraphWeight = &v
}

func (o UnifiedRoleDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RolePermissions != nil {
		toSerialize["rolePermissions"] = o.RolePermissions
	}
	if o.LibreGraphWeight != nil {
		toSerialize["@libre.graph.weight"] = o.LibreGraphWeight
	}
	return json.Marshal(toSerialize)
}

type NullableUnifiedRoleDefinition struct {
	value *UnifiedRoleDefinition
	isSet bool
}

func (v NullableUnifiedRoleDefinition) Get() *UnifiedRoleDefinition {
	return v.value
}

func (v *NullableUnifiedRoleDefinition) Set(val *UnifiedRoleDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableUnifiedRoleDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableUnifiedRoleDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnifiedRoleDefinition(val *UnifiedRoleDefinition) *NullableUnifiedRoleDefinition {
	return &NullableUnifiedRoleDefinition{value: val, isSet: true}
}

func (v NullableUnifiedRoleDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnifiedRoleDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}