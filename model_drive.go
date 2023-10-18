/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// checks if the Drive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Drive{}

// Drive The drive represents a space on the storage.
type Drive struct {
	// The unique idenfier for this drive.
	Id        *string      `json:"id,omitempty"`
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag           *string      `json:"eTag,omitempty"`
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name            string         `json:"name"`
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Describes the type of drive represented by this resource. Values are \"personal\" for users home spaces, \"project\", \"virtual\" or \"share\". Read-only.
	DriveType *string `json:"driveType,omitempty"`
	// The drive alias can be used in clients to make the urls user friendly. Example: 'personal/einstein'. This will be used to resolve to the correct driveID.
	DriveAlias *string      `json:"driveAlias,omitempty"`
	Owner      *IdentitySet `json:"owner,omitempty"`
	Quota      *Quota       `json:"quota,omitempty"`
	// All items contained in the drive. Read-only. Nullable.
	Items []DriveItem `json:"items,omitempty"`
	Root  *DriveItem  `json:"root,omitempty"`
	// A collection of special drive resources.
	Special []DriveItem `json:"special,omitempty"`
}

// NewDrive instantiates a new Drive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrive(name string) *Drive {
	this := Drive{}
	this.Name = name
	return &this
}

// NewDriveWithDefaults instantiates a new Drive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveWithDefaults() *Drive {
	this := Drive{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Drive) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Drive) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Drive) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Drive) GetCreatedBy() IdentitySet {
	if o == nil || IsNil(o.CreatedBy) {
		var ret IdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetCreatedByOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Drive) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given IdentitySet and assigns it to the CreatedBy field.
func (o *Drive) SetCreatedBy(v IdentitySet) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *Drive) GetCreatedDateTime() time.Time {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Drive) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *Drive) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Drive) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Drive) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Drive) SetDescription(v string) {
	o.Description = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *Drive) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *Drive) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *Drive) SetETag(v string) {
	o.ETag = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *Drive) GetLastModifiedBy() IdentitySet {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret IdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetLastModifiedByOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *Drive) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given IdentitySet and assigns it to the LastModifiedBy field.
func (o *Drive) SetLastModifiedBy(v IdentitySet) {
	o.LastModifiedBy = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *Drive) GetLastModifiedDateTime() time.Time {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *Drive) HasLastModifiedDateTime() bool {
	if o != nil && !IsNil(o.LastModifiedDateTime) {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *Drive) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value
func (o *Drive) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Drive) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Drive) SetName(v string) {
	o.Name = v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *Drive) GetParentReference() ItemReference {
	if o == nil || IsNil(o.ParentReference) {
		var ret ItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetParentReferenceOk() (*ItemReference, bool) {
	if o == nil || IsNil(o.ParentReference) {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *Drive) HasParentReference() bool {
	if o != nil && !IsNil(o.ParentReference) {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given ItemReference and assigns it to the ParentReference field.
func (o *Drive) SetParentReference(v ItemReference) {
	o.ParentReference = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *Drive) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *Drive) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *Drive) SetWebUrl(v string) {
	o.WebUrl = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *Drive) GetDriveType() string {
	if o == nil || IsNil(o.DriveType) {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetDriveTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DriveType) {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *Drive) HasDriveType() bool {
	if o != nil && !IsNil(o.DriveType) {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *Drive) SetDriveType(v string) {
	o.DriveType = &v
}

// GetDriveAlias returns the DriveAlias field value if set, zero value otherwise.
func (o *Drive) GetDriveAlias() string {
	if o == nil || IsNil(o.DriveAlias) {
		var ret string
		return ret
	}
	return *o.DriveAlias
}

// GetDriveAliasOk returns a tuple with the DriveAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetDriveAliasOk() (*string, bool) {
	if o == nil || IsNil(o.DriveAlias) {
		return nil, false
	}
	return o.DriveAlias, true
}

// HasDriveAlias returns a boolean if a field has been set.
func (o *Drive) HasDriveAlias() bool {
	if o != nil && !IsNil(o.DriveAlias) {
		return true
	}

	return false
}

// SetDriveAlias gets a reference to the given string and assigns it to the DriveAlias field.
func (o *Drive) SetDriveAlias(v string) {
	o.DriveAlias = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Drive) GetOwner() IdentitySet {
	if o == nil || IsNil(o.Owner) {
		var ret IdentitySet
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetOwnerOk() (*IdentitySet, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Drive) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given IdentitySet and assigns it to the Owner field.
func (o *Drive) SetOwner(v IdentitySet) {
	o.Owner = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *Drive) GetQuota() Quota {
	if o == nil || IsNil(o.Quota) {
		var ret Quota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetQuotaOk() (*Quota, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *Drive) HasQuota() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given Quota and assigns it to the Quota field.
func (o *Drive) SetQuota(v Quota) {
	o.Quota = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Drive) GetItems() []DriveItem {
	if o == nil || IsNil(o.Items) {
		var ret []DriveItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetItemsOk() ([]DriveItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Drive) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DriveItem and assigns it to the Items field.
func (o *Drive) SetItems(v []DriveItem) {
	o.Items = v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *Drive) GetRoot() DriveItem {
	if o == nil || IsNil(o.Root) {
		var ret DriveItem
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetRootOk() (*DriveItem, bool) {
	if o == nil || IsNil(o.Root) {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *Drive) HasRoot() bool {
	if o != nil && !IsNil(o.Root) {
		return true
	}

	return false
}

// SetRoot gets a reference to the given DriveItem and assigns it to the Root field.
func (o *Drive) SetRoot(v DriveItem) {
	o.Root = &v
}

// GetSpecial returns the Special field value if set, zero value otherwise.
func (o *Drive) GetSpecial() []DriveItem {
	if o == nil || IsNil(o.Special) {
		var ret []DriveItem
		return ret
	}
	return o.Special
}

// GetSpecialOk returns a tuple with the Special field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetSpecialOk() ([]DriveItem, bool) {
	if o == nil || IsNil(o.Special) {
		return nil, false
	}
	return o.Special, true
}

// HasSpecial returns a boolean if a field has been set.
func (o *Drive) HasSpecial() bool {
	if o != nil && !IsNil(o.Special) {
		return true
	}

	return false
}

// SetSpecial gets a reference to the given []DriveItem and assigns it to the Special field.
func (o *Drive) SetSpecial(v []DriveItem) {
	o.Special = v
}

func (o Drive) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Drive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastModifiedDateTime) {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentReference) {
		toSerialize["parentReference"] = o.ParentReference
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	if !IsNil(o.DriveType) {
		toSerialize["driveType"] = o.DriveType
	}
	if !IsNil(o.DriveAlias) {
		toSerialize["driveAlias"] = o.DriveAlias
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Root) {
		toSerialize["root"] = o.Root
	}
	if !IsNil(o.Special) {
		toSerialize["special"] = o.Special
	}
	return toSerialize, nil
}

type NullableDrive struct {
	value *Drive
	isSet bool
}

func (v NullableDrive) Get() *Drive {
	return v.value
}

func (v *NullableDrive) Set(val *Drive) {
	v.value = val
	v.isSet = true
}

func (v NullableDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrive(val *Drive) *NullableDrive {
	return &NullableDrive{value: val, isSet: true}
}

func (v NullableDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
