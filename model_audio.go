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

// checks if the Audio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Audio{}

// Audio The Audio resource groups audio-related properties on an item into a single structure.  If a DriveItem has a non-null audio facet, the item represents an audio file. The properties of the Audio resource are populated by extracting metadata from the file.
type Audio struct {
	// The title of the album for this audio file.
	Album *string `json:"album,omitempty"`
	// The artist named on the album for the audio file.
	AlbumArtist *string `json:"albumArtist,omitempty"`
	// The performing artist for the audio file.
	Artist *string `json:"artist,omitempty"`
	// Bitrate expressed in kbps.
	Bitrate *int64 `json:"bitrate,omitempty"`
	// The name of the composer of the audio file.
	Composers *string `json:"composers,omitempty"`
	// Copyright information for the audio file.
	Copyright *string `json:"copyright,omitempty"`
	// The number of the disc this audio file came from.
	Disc *int32 `json:"disc,omitempty"`
	// The total number of discs in this album.
	DiscCount *int32 `json:"discCount,omitempty"`
	// Duration of the audio file, expressed in milliseconds
	Duration *int64 `json:"duration,omitempty"`
	// The genre of this audio file.
	Genre *string `json:"genre,omitempty"`
	// Indicates if the file is protected with digital rights management.
	HasDrm *bool `json:"hasDrm,omitempty"`
	// Indicates if the file is encoded with a variable bitrate.
	IsVariableBitrate *bool `json:"isVariableBitrate,omitempty"`
	// The title of the audio file.
	Title *string `json:"title,omitempty"`
	// The number of the track on the original disc for this audio file.
	Track *int32 `json:"track,omitempty"`
	// The total number of tracks on the original disc for this audio file.
	TrackCount *int32 `json:"trackCount,omitempty"`
	// The year the audio file was recorded.
	Year *int32 `json:"year,omitempty"`
}

// NewAudio instantiates a new Audio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudio() *Audio {
	this := Audio{}
	return &this
}

// NewAudioWithDefaults instantiates a new Audio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudioWithDefaults() *Audio {
	this := Audio{}
	return &this
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *Audio) GetAlbum() string {
	if o == nil || IsNil(o.Album) {
		var ret string
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetAlbumOk() (*string, bool) {
	if o == nil || IsNil(o.Album) {
		return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *Audio) HasAlbum() bool {
	if o != nil && !IsNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given string and assigns it to the Album field.
func (o *Audio) SetAlbum(v string) {
	o.Album = &v
}

// GetAlbumArtist returns the AlbumArtist field value if set, zero value otherwise.
func (o *Audio) GetAlbumArtist() string {
	if o == nil || IsNil(o.AlbumArtist) {
		var ret string
		return ret
	}
	return *o.AlbumArtist
}

// GetAlbumArtistOk returns a tuple with the AlbumArtist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetAlbumArtistOk() (*string, bool) {
	if o == nil || IsNil(o.AlbumArtist) {
		return nil, false
	}
	return o.AlbumArtist, true
}

// HasAlbumArtist returns a boolean if a field has been set.
func (o *Audio) HasAlbumArtist() bool {
	if o != nil && !IsNil(o.AlbumArtist) {
		return true
	}

	return false
}

// SetAlbumArtist gets a reference to the given string and assigns it to the AlbumArtist field.
func (o *Audio) SetAlbumArtist(v string) {
	o.AlbumArtist = &v
}

// GetArtist returns the Artist field value if set, zero value otherwise.
func (o *Audio) GetArtist() string {
	if o == nil || IsNil(o.Artist) {
		var ret string
		return ret
	}
	return *o.Artist
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetArtistOk() (*string, bool) {
	if o == nil || IsNil(o.Artist) {
		return nil, false
	}
	return o.Artist, true
}

// HasArtist returns a boolean if a field has been set.
func (o *Audio) HasArtist() bool {
	if o != nil && !IsNil(o.Artist) {
		return true
	}

	return false
}

// SetArtist gets a reference to the given string and assigns it to the Artist field.
func (o *Audio) SetArtist(v string) {
	o.Artist = &v
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise.
func (o *Audio) GetBitrate() int64 {
	if o == nil || IsNil(o.Bitrate) {
		var ret int64
		return ret
	}
	return *o.Bitrate
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetBitrateOk() (*int64, bool) {
	if o == nil || IsNil(o.Bitrate) {
		return nil, false
	}
	return o.Bitrate, true
}

// HasBitrate returns a boolean if a field has been set.
func (o *Audio) HasBitrate() bool {
	if o != nil && !IsNil(o.Bitrate) {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given int64 and assigns it to the Bitrate field.
func (o *Audio) SetBitrate(v int64) {
	o.Bitrate = &v
}

// GetComposers returns the Composers field value if set, zero value otherwise.
func (o *Audio) GetComposers() string {
	if o == nil || IsNil(o.Composers) {
		var ret string
		return ret
	}
	return *o.Composers
}

// GetComposersOk returns a tuple with the Composers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetComposersOk() (*string, bool) {
	if o == nil || IsNil(o.Composers) {
		return nil, false
	}
	return o.Composers, true
}

// HasComposers returns a boolean if a field has been set.
func (o *Audio) HasComposers() bool {
	if o != nil && !IsNil(o.Composers) {
		return true
	}

	return false
}

// SetComposers gets a reference to the given string and assigns it to the Composers field.
func (o *Audio) SetComposers(v string) {
	o.Composers = &v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *Audio) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *Audio) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *Audio) SetCopyright(v string) {
	o.Copyright = &v
}

// GetDisc returns the Disc field value if set, zero value otherwise.
func (o *Audio) GetDisc() int32 {
	if o == nil || IsNil(o.Disc) {
		var ret int32
		return ret
	}
	return *o.Disc
}

// GetDiscOk returns a tuple with the Disc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetDiscOk() (*int32, bool) {
	if o == nil || IsNil(o.Disc) {
		return nil, false
	}
	return o.Disc, true
}

// HasDisc returns a boolean if a field has been set.
func (o *Audio) HasDisc() bool {
	if o != nil && !IsNil(o.Disc) {
		return true
	}

	return false
}

// SetDisc gets a reference to the given int32 and assigns it to the Disc field.
func (o *Audio) SetDisc(v int32) {
	o.Disc = &v
}

// GetDiscCount returns the DiscCount field value if set, zero value otherwise.
func (o *Audio) GetDiscCount() int32 {
	if o == nil || IsNil(o.DiscCount) {
		var ret int32
		return ret
	}
	return *o.DiscCount
}

// GetDiscCountOk returns a tuple with the DiscCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetDiscCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscCount) {
		return nil, false
	}
	return o.DiscCount, true
}

// HasDiscCount returns a boolean if a field has been set.
func (o *Audio) HasDiscCount() bool {
	if o != nil && !IsNil(o.DiscCount) {
		return true
	}

	return false
}

// SetDiscCount gets a reference to the given int32 and assigns it to the DiscCount field.
func (o *Audio) SetDiscCount(v int32) {
	o.DiscCount = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Audio) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Audio) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *Audio) SetDuration(v int64) {
	o.Duration = &v
}

// GetGenre returns the Genre field value if set, zero value otherwise.
func (o *Audio) GetGenre() string {
	if o == nil || IsNil(o.Genre) {
		var ret string
		return ret
	}
	return *o.Genre
}

// GetGenreOk returns a tuple with the Genre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetGenreOk() (*string, bool) {
	if o == nil || IsNil(o.Genre) {
		return nil, false
	}
	return o.Genre, true
}

// HasGenre returns a boolean if a field has been set.
func (o *Audio) HasGenre() bool {
	if o != nil && !IsNil(o.Genre) {
		return true
	}

	return false
}

// SetGenre gets a reference to the given string and assigns it to the Genre field.
func (o *Audio) SetGenre(v string) {
	o.Genre = &v
}

// GetHasDrm returns the HasDrm field value if set, zero value otherwise.
func (o *Audio) GetHasDrm() bool {
	if o == nil || IsNil(o.HasDrm) {
		var ret bool
		return ret
	}
	return *o.HasDrm
}

// GetHasDrmOk returns a tuple with the HasDrm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetHasDrmOk() (*bool, bool) {
	if o == nil || IsNil(o.HasDrm) {
		return nil, false
	}
	return o.HasDrm, true
}

// HasHasDrm returns a boolean if a field has been set.
func (o *Audio) HasHasDrm() bool {
	if o != nil && !IsNil(o.HasDrm) {
		return true
	}

	return false
}

// SetHasDrm gets a reference to the given bool and assigns it to the HasDrm field.
func (o *Audio) SetHasDrm(v bool) {
	o.HasDrm = &v
}

// GetIsVariableBitrate returns the IsVariableBitrate field value if set, zero value otherwise.
func (o *Audio) GetIsVariableBitrate() bool {
	if o == nil || IsNil(o.IsVariableBitrate) {
		var ret bool
		return ret
	}
	return *o.IsVariableBitrate
}

// GetIsVariableBitrateOk returns a tuple with the IsVariableBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetIsVariableBitrateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVariableBitrate) {
		return nil, false
	}
	return o.IsVariableBitrate, true
}

// HasIsVariableBitrate returns a boolean if a field has been set.
func (o *Audio) HasIsVariableBitrate() bool {
	if o != nil && !IsNil(o.IsVariableBitrate) {
		return true
	}

	return false
}

// SetIsVariableBitrate gets a reference to the given bool and assigns it to the IsVariableBitrate field.
func (o *Audio) SetIsVariableBitrate(v bool) {
	o.IsVariableBitrate = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Audio) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Audio) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Audio) SetTitle(v string) {
	o.Title = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *Audio) GetTrack() int32 {
	if o == nil || IsNil(o.Track) {
		var ret int32
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetTrackOk() (*int32, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *Audio) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given int32 and assigns it to the Track field.
func (o *Audio) SetTrack(v int32) {
	o.Track = &v
}

// GetTrackCount returns the TrackCount field value if set, zero value otherwise.
func (o *Audio) GetTrackCount() int32 {
	if o == nil || IsNil(o.TrackCount) {
		var ret int32
		return ret
	}
	return *o.TrackCount
}

// GetTrackCountOk returns a tuple with the TrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetTrackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrackCount) {
		return nil, false
	}
	return o.TrackCount, true
}

// HasTrackCount returns a boolean if a field has been set.
func (o *Audio) HasTrackCount() bool {
	if o != nil && !IsNil(o.TrackCount) {
		return true
	}

	return false
}

// SetTrackCount gets a reference to the given int32 and assigns it to the TrackCount field.
func (o *Audio) SetTrackCount(v int32) {
	o.TrackCount = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Audio) GetYear() int32 {
	if o == nil || IsNil(o.Year) {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Audio) GetYearOk() (*int32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Audio) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *Audio) SetYear(v int32) {
	o.Year = &v
}

func (o Audio) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Audio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	if !IsNil(o.AlbumArtist) {
		toSerialize["albumArtist"] = o.AlbumArtist
	}
	if !IsNil(o.Artist) {
		toSerialize["artist"] = o.Artist
	}
	if !IsNil(o.Bitrate) {
		toSerialize["bitrate"] = o.Bitrate
	}
	if !IsNil(o.Composers) {
		toSerialize["composers"] = o.Composers
	}
	if !IsNil(o.Copyright) {
		toSerialize["copyright"] = o.Copyright
	}
	if !IsNil(o.Disc) {
		toSerialize["disc"] = o.Disc
	}
	if !IsNil(o.DiscCount) {
		toSerialize["discCount"] = o.DiscCount
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Genre) {
		toSerialize["genre"] = o.Genre
	}
	if !IsNil(o.HasDrm) {
		toSerialize["hasDrm"] = o.HasDrm
	}
	if !IsNil(o.IsVariableBitrate) {
		toSerialize["isVariableBitrate"] = o.IsVariableBitrate
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.TrackCount) {
		toSerialize["trackCount"] = o.TrackCount
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	return toSerialize, nil
}

type NullableAudio struct {
	value *Audio
	isSet bool
}

func (v NullableAudio) Get() *Audio {
	return v.value
}

func (v *NullableAudio) Set(val *Audio) {
	v.value = val
	v.isSet = true
}

func (v NullableAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudio(val *Audio) *NullableAudio {
	return &NullableAudio{value: val, isSet: true}
}

func (v NullableAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
