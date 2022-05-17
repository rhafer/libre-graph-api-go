/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.14.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// RemoteItem Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
type RemoteItem struct {
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	File *OpenGraphFile `json:"file,omitempty"`
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`
	Folder *Folder `json:"folder,omitempty"`
	// Unique identifier for the remote item in its drive. Read-only.
	Id *string `json:"id,omitempty"`
	Image *Image `json:"image,omitempty"`
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Optional. Filename of the remote item. Read-only.
	Name *string `json:"name,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only.
	CTag *string `json:"cTag,omitempty"`
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	Shared *Shared `json:"shared,omitempty"`
	// Size of the remote item. Read-only.
	Size *int64 `json:"size,omitempty"`
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`
	// DAV compatible URL for the item.
	WebDavUrl *string `json:"webDavUrl,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
}

// NewRemoteItem instantiates a new RemoteItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteItem() *RemoteItem {
	this := RemoteItem{}
	return &this
}

// NewRemoteItemWithDefaults instantiates a new RemoteItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteItemWithDefaults() *RemoteItem {
	this := RemoteItem{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RemoteItem) GetCreatedBy() IdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret IdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCreatedByOk() (*IdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RemoteItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given IdentitySet and assigns it to the CreatedBy field.
func (o *RemoteItem) SetCreatedBy(v IdentitySet) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *RemoteItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *RemoteItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *RemoteItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *RemoteItem) GetFile() OpenGraphFile {
	if o == nil || o.File == nil {
		var ret OpenGraphFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFileOk() (*OpenGraphFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *RemoteItem) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given OpenGraphFile and assigns it to the File field.
func (o *RemoteItem) SetFile(v OpenGraphFile) {
	o.File = &v
}

// GetFileSystemInfo returns the FileSystemInfo field value if set, zero value otherwise.
func (o *RemoteItem) GetFileSystemInfo() FileSystemInfo {
	if o == nil || o.FileSystemInfo == nil {
		var ret FileSystemInfo
		return ret
	}
	return *o.FileSystemInfo
}

// GetFileSystemInfoOk returns a tuple with the FileSystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFileSystemInfoOk() (*FileSystemInfo, bool) {
	if o == nil || o.FileSystemInfo == nil {
		return nil, false
	}
	return o.FileSystemInfo, true
}

// HasFileSystemInfo returns a boolean if a field has been set.
func (o *RemoteItem) HasFileSystemInfo() bool {
	if o != nil && o.FileSystemInfo != nil {
		return true
	}

	return false
}

// SetFileSystemInfo gets a reference to the given FileSystemInfo and assigns it to the FileSystemInfo field.
func (o *RemoteItem) SetFileSystemInfo(v FileSystemInfo) {
	o.FileSystemInfo = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *RemoteItem) GetFolder() Folder {
	if o == nil || o.Folder == nil {
		var ret Folder
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetFolderOk() (*Folder, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *RemoteItem) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given Folder and assigns it to the Folder field.
func (o *RemoteItem) SetFolder(v Folder) {
	o.Folder = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteItem) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *RemoteItem) GetImage() Image {
	if o == nil || o.Image == nil {
		var ret Image
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetImageOk() (*Image, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *RemoteItem) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given Image and assigns it to the Image field.
func (o *RemoteItem) SetImage(v Image) {
	o.Image = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *RemoteItem) GetLastModifiedBy() IdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret IdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetLastModifiedByOk() (*IdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *RemoteItem) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given IdentitySet and assigns it to the LastModifiedBy field.
func (o *RemoteItem) SetLastModifiedBy(v IdentitySet) {
	o.LastModifiedBy = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *RemoteItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *RemoteItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *RemoteItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteItem) SetName(v string) {
	o.Name = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *RemoteItem) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetETagOk() (*string, bool) {
	if o == nil || o.ETag == nil {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *RemoteItem) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *RemoteItem) SetETag(v string) {
	o.ETag = &v
}

// GetCTag returns the CTag field value if set, zero value otherwise.
func (o *RemoteItem) GetCTag() string {
	if o == nil || o.CTag == nil {
		var ret string
		return ret
	}
	return *o.CTag
}

// GetCTagOk returns a tuple with the CTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetCTagOk() (*string, bool) {
	if o == nil || o.CTag == nil {
		return nil, false
	}
	return o.CTag, true
}

// HasCTag returns a boolean if a field has been set.
func (o *RemoteItem) HasCTag() bool {
	if o != nil && o.CTag != nil {
		return true
	}

	return false
}

// SetCTag gets a reference to the given string and assigns it to the CTag field.
func (o *RemoteItem) SetCTag(v string) {
	o.CTag = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *RemoteItem) GetParentReference() ItemReference {
	if o == nil || o.ParentReference == nil {
		var ret ItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetParentReferenceOk() (*ItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *RemoteItem) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given ItemReference and assigns it to the ParentReference field.
func (o *RemoteItem) SetParentReference(v ItemReference) {
	o.ParentReference = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *RemoteItem) GetShared() Shared {
	if o == nil || o.Shared == nil {
		var ret Shared
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSharedOk() (*Shared, bool) {
	if o == nil || o.Shared == nil {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *RemoteItem) HasShared() bool {
	if o != nil && o.Shared != nil {
		return true
	}

	return false
}

// SetShared gets a reference to the given Shared and assigns it to the Shared field.
func (o *RemoteItem) SetShared(v Shared) {
	o.Shared = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RemoteItem) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RemoteItem) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *RemoteItem) SetSize(v int64) {
	o.Size = &v
}

// GetSpecialFolder returns the SpecialFolder field value if set, zero value otherwise.
func (o *RemoteItem) GetSpecialFolder() SpecialFolder {
	if o == nil || o.SpecialFolder == nil {
		var ret SpecialFolder
		return ret
	}
	return *o.SpecialFolder
}

// GetSpecialFolderOk returns a tuple with the SpecialFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetSpecialFolderOk() (*SpecialFolder, bool) {
	if o == nil || o.SpecialFolder == nil {
		return nil, false
	}
	return o.SpecialFolder, true
}

// HasSpecialFolder returns a boolean if a field has been set.
func (o *RemoteItem) HasSpecialFolder() bool {
	if o != nil && o.SpecialFolder != nil {
		return true
	}

	return false
}

// SetSpecialFolder gets a reference to the given SpecialFolder and assigns it to the SpecialFolder field.
func (o *RemoteItem) SetSpecialFolder(v SpecialFolder) {
	o.SpecialFolder = &v
}

// GetWebDavUrl returns the WebDavUrl field value if set, zero value otherwise.
func (o *RemoteItem) GetWebDavUrl() string {
	if o == nil || o.WebDavUrl == nil {
		var ret string
		return ret
	}
	return *o.WebDavUrl
}

// GetWebDavUrlOk returns a tuple with the WebDavUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetWebDavUrlOk() (*string, bool) {
	if o == nil || o.WebDavUrl == nil {
		return nil, false
	}
	return o.WebDavUrl, true
}

// HasWebDavUrl returns a boolean if a field has been set.
func (o *RemoteItem) HasWebDavUrl() bool {
	if o != nil && o.WebDavUrl != nil {
		return true
	}

	return false
}

// SetWebDavUrl gets a reference to the given string and assigns it to the WebDavUrl field.
func (o *RemoteItem) SetWebDavUrl(v string) {
	o.WebDavUrl = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *RemoteItem) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteItem) GetWebUrlOk() (*string, bool) {
	if o == nil || o.WebUrl == nil {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *RemoteItem) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *RemoteItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o RemoteItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.FileSystemInfo != nil {
		toSerialize["fileSystemInfo"] = o.FileSystemInfo
	}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
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
	if o.ETag != nil {
		toSerialize["eTag"] = o.ETag
	}
	if o.CTag != nil {
		toSerialize["cTag"] = o.CTag
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.Shared != nil {
		toSerialize["shared"] = o.Shared
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.SpecialFolder != nil {
		toSerialize["specialFolder"] = o.SpecialFolder
	}
	if o.WebDavUrl != nil {
		toSerialize["webDavUrl"] = o.WebDavUrl
	}
	if o.WebUrl != nil {
		toSerialize["webUrl"] = o.WebUrl
	}
	return json.Marshal(toSerialize)
}

type NullableRemoteItem struct {
	value *RemoteItem
	isSet bool
}

func (v NullableRemoteItem) Get() *RemoteItem {
	return v.value
}

func (v *NullableRemoteItem) Set(val *RemoteItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteItem(val *RemoteItem) *NullableRemoteItem {
	return &NullableRemoteItem{value: val, isSet: true}
}

func (v NullableRemoteItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


