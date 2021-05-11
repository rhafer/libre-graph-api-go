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

// OpenGraphUser struct for OpenGraphUser
type OpenGraphUser struct {
	OpenGraphDirectoryObject
	// true if the account is enabled; otherwise, false. This property is required when a user is created. Returned only on $select. Supports $filter.
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// The telephone numbers for the user. Only one number can be set for this property. Returned by default. Read-only for users synced from on-premises directory.
	BusinessPhones *[]string `json:"businessPhones,omitempty"`
	// The city in which the user is located. Returned only on $select. Supports $filter.
	City *string `json:"city,omitempty"`
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from. The maximum length of the company name is 64 characters.Returned only on $select.
	CompanyName *string `json:"companyName,omitempty"`
	// The country/region in which the user is located; for example, 'US' or 'UK'. Returned only on $select. Supports $filter.
	Country *string `json:"country,omitempty"`
	// The date and time the user was created. The value cannot be modified and is automatically populated when the entity is created. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. Property is nullable. A null value indicates that an accurate creation time couldn't be determined for the user. Returned only on $select. Read-only. Supports $filter.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The name for the department in which the user works. Returned only on $select. Supports $filter.
	Department *string `json:"department,omitempty"`
	// The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $filter and $orderby.
	DisplayName *string `json:"displayName,omitempty"`
	// The fax number of the user. Returned only on $select.
	FaxNumber *string `json:"faxNumber,omitempty"`
	// The given name (first name) of the user. Returned by default. Supports $filter.
	GivenName *string `json:"givenName,omitempty"`
	// The time when this user last changed their password. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select. Read-only.
	LastPasswordChangeDateTime *time.Time `json:"lastPasswordChangeDateTime,omitempty"`
	// Used by enterprise applications to determine the legal age group of the user. This property is read-only and calculated based on ageGroup and consentProvidedForMinor properties. Allowed values: null, minorWithOutParentalConsent, minorWithParentalConsent, minorNoParentalConsentRequired, notAdult and adult. Refer to the legal age group property definitions for further information. Returned only on $select.
	LegalAgeGroupClassification *string `json:"legalAgeGroupClassification,omitempty"`
	// The SMTP address for the user, for example, 'jeff@contoso.onowncloud.com'. Returned by default. Supports $filter and endsWith.
	Mail *string `json:"mail,omitempty"`
	// The mail alias for the user. This property must be specified when a user is created. Returned only on $select. Supports $filter.
	MailNickname *string `json:"mailNickname,omitempty"`
	// The primary cellular telephone number for the user. Returned by default. Read-only for users synced from on-premises directory.
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// The office location in the user's place of business. Returned by default.
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code. Returned only on $select.
	PostalCode *string `json:"postalCode,omitempty"`
	// The preferred language for the user. Should follow ISO 639-1 Code; for example 'en-US'. Returned by default.
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// The state or province in the user's address. Returned only on $select. Supports $filter.
	State *string `json:"state,omitempty"`
	// The street address of the user's place of business. Returned only on $select.
	StreetAddress *string `json:"streetAddress,omitempty"`
	// The user's surname (family name or last name). Returned by default. Supports $filter.
	Surname *string `json:"surname,omitempty"`
	// A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries.  Examples include: 'US', 'JP', and 'GB'. Not nullable. Returned only on $select. Supports $filter.
	UsageLocation *string `json:"usageLocation,omitempty"`
	// The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of organization. Returned by default. Supports $filter, $orderby, and endsWith.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// A string value that can be used to classify user types in your directory, such as 'Member' and 'Guest'. Returned only on $select. Supports $filter.
	UserType *string `json:"userType,omitempty"`
	// A freeform text entry field for the user to describe themselves. Returned only on $select.
	AboutMe *string `json:"aboutMe,omitempty"`
	// The birthday of the user. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Returned only on $select.
	Birthday *time.Time `json:"birthday,omitempty"`
	// The user's HomeDrive. Read-only.
	Drive *OpenGraphDrive `json:"drive,omitempty"`
	// A collection of drives available for this user. Read-only.
	Drives *[]OpenGraphDrive `json:"drives,omitempty"`
}

// NewOpenGraphUser instantiates a new OpenGraphUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenGraphUser() *OpenGraphUser {
	this := OpenGraphUser{}
	return &this
}

// NewOpenGraphUserWithDefaults instantiates a new OpenGraphUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenGraphUserWithDefaults() *OpenGraphUser {
	this := OpenGraphUser{}
	return &this
}

// GetAccountEnabled returns the AccountEnabled field value if set, zero value otherwise.
func (o *OpenGraphUser) GetAccountEnabled() bool {
	if o == nil || o.AccountEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AccountEnabled
}

// GetAccountEnabledOk returns a tuple with the AccountEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetAccountEnabledOk() (*bool, bool) {
	if o == nil || o.AccountEnabled == nil {
		return nil, false
	}
	return o.AccountEnabled, true
}

// HasAccountEnabled returns a boolean if a field has been set.
func (o *OpenGraphUser) HasAccountEnabled() bool {
	if o != nil && o.AccountEnabled != nil {
		return true
	}

	return false
}

// SetAccountEnabled gets a reference to the given bool and assigns it to the AccountEnabled field.
func (o *OpenGraphUser) SetAccountEnabled(v bool) {
	o.AccountEnabled = &v
}

// GetBusinessPhones returns the BusinessPhones field value if set, zero value otherwise.
func (o *OpenGraphUser) GetBusinessPhones() []string {
	if o == nil || o.BusinessPhones == nil {
		var ret []string
		return ret
	}
	return *o.BusinessPhones
}

// GetBusinessPhonesOk returns a tuple with the BusinessPhones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetBusinessPhonesOk() (*[]string, bool) {
	if o == nil || o.BusinessPhones == nil {
		return nil, false
	}
	return o.BusinessPhones, true
}

// HasBusinessPhones returns a boolean if a field has been set.
func (o *OpenGraphUser) HasBusinessPhones() bool {
	if o != nil && o.BusinessPhones != nil {
		return true
	}

	return false
}

// SetBusinessPhones gets a reference to the given []string and assigns it to the BusinessPhones field.
func (o *OpenGraphUser) SetBusinessPhones(v []string) {
	o.BusinessPhones = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OpenGraphUser) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OpenGraphUser) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OpenGraphUser) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OpenGraphUser) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OpenGraphUser) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OpenGraphUser) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OpenGraphUser) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OpenGraphUser) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *OpenGraphUser) SetCountry(v string) {
	o.Country = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *OpenGraphUser) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *OpenGraphUser) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *OpenGraphUser) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *OpenGraphUser) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *OpenGraphUser) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *OpenGraphUser) SetDepartment(v string) {
	o.Department = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OpenGraphUser) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OpenGraphUser) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OpenGraphUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFaxNumber returns the FaxNumber field value if set, zero value otherwise.
func (o *OpenGraphUser) GetFaxNumber() string {
	if o == nil || o.FaxNumber == nil {
		var ret string
		return ret
	}
	return *o.FaxNumber
}

// GetFaxNumberOk returns a tuple with the FaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetFaxNumberOk() (*string, bool) {
	if o == nil || o.FaxNumber == nil {
		return nil, false
	}
	return o.FaxNumber, true
}

// HasFaxNumber returns a boolean if a field has been set.
func (o *OpenGraphUser) HasFaxNumber() bool {
	if o != nil && o.FaxNumber != nil {
		return true
	}

	return false
}

// SetFaxNumber gets a reference to the given string and assigns it to the FaxNumber field.
func (o *OpenGraphUser) SetFaxNumber(v string) {
	o.FaxNumber = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *OpenGraphUser) GetGivenName() string {
	if o == nil || o.GivenName == nil {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetGivenNameOk() (*string, bool) {
	if o == nil || o.GivenName == nil {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *OpenGraphUser) HasGivenName() bool {
	if o != nil && o.GivenName != nil {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *OpenGraphUser) SetGivenName(v string) {
	o.GivenName = &v
}

// GetLastPasswordChangeDateTime returns the LastPasswordChangeDateTime field value if set, zero value otherwise.
func (o *OpenGraphUser) GetLastPasswordChangeDateTime() time.Time {
	if o == nil || o.LastPasswordChangeDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordChangeDateTime
}

// GetLastPasswordChangeDateTimeOk returns a tuple with the LastPasswordChangeDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetLastPasswordChangeDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastPasswordChangeDateTime == nil {
		return nil, false
	}
	return o.LastPasswordChangeDateTime, true
}

// HasLastPasswordChangeDateTime returns a boolean if a field has been set.
func (o *OpenGraphUser) HasLastPasswordChangeDateTime() bool {
	if o != nil && o.LastPasswordChangeDateTime != nil {
		return true
	}

	return false
}

// SetLastPasswordChangeDateTime gets a reference to the given time.Time and assigns it to the LastPasswordChangeDateTime field.
func (o *OpenGraphUser) SetLastPasswordChangeDateTime(v time.Time) {
	o.LastPasswordChangeDateTime = &v
}

// GetLegalAgeGroupClassification returns the LegalAgeGroupClassification field value if set, zero value otherwise.
func (o *OpenGraphUser) GetLegalAgeGroupClassification() string {
	if o == nil || o.LegalAgeGroupClassification == nil {
		var ret string
		return ret
	}
	return *o.LegalAgeGroupClassification
}

// GetLegalAgeGroupClassificationOk returns a tuple with the LegalAgeGroupClassification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetLegalAgeGroupClassificationOk() (*string, bool) {
	if o == nil || o.LegalAgeGroupClassification == nil {
		return nil, false
	}
	return o.LegalAgeGroupClassification, true
}

// HasLegalAgeGroupClassification returns a boolean if a field has been set.
func (o *OpenGraphUser) HasLegalAgeGroupClassification() bool {
	if o != nil && o.LegalAgeGroupClassification != nil {
		return true
	}

	return false
}

// SetLegalAgeGroupClassification gets a reference to the given string and assigns it to the LegalAgeGroupClassification field.
func (o *OpenGraphUser) SetLegalAgeGroupClassification(v string) {
	o.LegalAgeGroupClassification = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *OpenGraphUser) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *OpenGraphUser) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *OpenGraphUser) SetMail(v string) {
	o.Mail = &v
}

// GetMailNickname returns the MailNickname field value if set, zero value otherwise.
func (o *OpenGraphUser) GetMailNickname() string {
	if o == nil || o.MailNickname == nil {
		var ret string
		return ret
	}
	return *o.MailNickname
}

// GetMailNicknameOk returns a tuple with the MailNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetMailNicknameOk() (*string, bool) {
	if o == nil || o.MailNickname == nil {
		return nil, false
	}
	return o.MailNickname, true
}

// HasMailNickname returns a boolean if a field has been set.
func (o *OpenGraphUser) HasMailNickname() bool {
	if o != nil && o.MailNickname != nil {
		return true
	}

	return false
}

// SetMailNickname gets a reference to the given string and assigns it to the MailNickname field.
func (o *OpenGraphUser) SetMailNickname(v string) {
	o.MailNickname = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *OpenGraphUser) GetMobilePhone() string {
	if o == nil || o.MobilePhone == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetMobilePhoneOk() (*string, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *OpenGraphUser) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *OpenGraphUser) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetOfficeLocation returns the OfficeLocation field value if set, zero value otherwise.
func (o *OpenGraphUser) GetOfficeLocation() string {
	if o == nil || o.OfficeLocation == nil {
		var ret string
		return ret
	}
	return *o.OfficeLocation
}

// GetOfficeLocationOk returns a tuple with the OfficeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetOfficeLocationOk() (*string, bool) {
	if o == nil || o.OfficeLocation == nil {
		return nil, false
	}
	return o.OfficeLocation, true
}

// HasOfficeLocation returns a boolean if a field has been set.
func (o *OpenGraphUser) HasOfficeLocation() bool {
	if o != nil && o.OfficeLocation != nil {
		return true
	}

	return false
}

// SetOfficeLocation gets a reference to the given string and assigns it to the OfficeLocation field.
func (o *OpenGraphUser) SetOfficeLocation(v string) {
	o.OfficeLocation = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OpenGraphUser) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OpenGraphUser) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OpenGraphUser) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise.
func (o *OpenGraphUser) GetPreferredLanguage() string {
	if o == nil || o.PreferredLanguage == nil {
		var ret string
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetPreferredLanguageOk() (*string, bool) {
	if o == nil || o.PreferredLanguage == nil {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *OpenGraphUser) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage != nil {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.
func (o *OpenGraphUser) SetPreferredLanguage(v string) {
	o.PreferredLanguage = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OpenGraphUser) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OpenGraphUser) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OpenGraphUser) SetState(v string) {
	o.State = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *OpenGraphUser) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *OpenGraphUser) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *OpenGraphUser) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *OpenGraphUser) GetSurname() string {
	if o == nil || o.Surname == nil {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetSurnameOk() (*string, bool) {
	if o == nil || o.Surname == nil {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *OpenGraphUser) HasSurname() bool {
	if o != nil && o.Surname != nil {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *OpenGraphUser) SetSurname(v string) {
	o.Surname = &v
}

// GetUsageLocation returns the UsageLocation field value if set, zero value otherwise.
func (o *OpenGraphUser) GetUsageLocation() string {
	if o == nil || o.UsageLocation == nil {
		var ret string
		return ret
	}
	return *o.UsageLocation
}

// GetUsageLocationOk returns a tuple with the UsageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetUsageLocationOk() (*string, bool) {
	if o == nil || o.UsageLocation == nil {
		return nil, false
	}
	return o.UsageLocation, true
}

// HasUsageLocation returns a boolean if a field has been set.
func (o *OpenGraphUser) HasUsageLocation() bool {
	if o != nil && o.UsageLocation != nil {
		return true
	}

	return false
}

// SetUsageLocation gets a reference to the given string and assigns it to the UsageLocation field.
func (o *OpenGraphUser) SetUsageLocation(v string) {
	o.UsageLocation = &v
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise.
func (o *OpenGraphUser) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil || o.UserPrincipalName == nil {
		return nil, false
	}
	return o.UserPrincipalName, true
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *OpenGraphUser) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName != nil {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given string and assigns it to the UserPrincipalName field.
func (o *OpenGraphUser) SetUserPrincipalName(v string) {
	o.UserPrincipalName = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *OpenGraphUser) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *OpenGraphUser) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *OpenGraphUser) SetUserType(v string) {
	o.UserType = &v
}

// GetAboutMe returns the AboutMe field value if set, zero value otherwise.
func (o *OpenGraphUser) GetAboutMe() string {
	if o == nil || o.AboutMe == nil {
		var ret string
		return ret
	}
	return *o.AboutMe
}

// GetAboutMeOk returns a tuple with the AboutMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetAboutMeOk() (*string, bool) {
	if o == nil || o.AboutMe == nil {
		return nil, false
	}
	return o.AboutMe, true
}

// HasAboutMe returns a boolean if a field has been set.
func (o *OpenGraphUser) HasAboutMe() bool {
	if o != nil && o.AboutMe != nil {
		return true
	}

	return false
}

// SetAboutMe gets a reference to the given string and assigns it to the AboutMe field.
func (o *OpenGraphUser) SetAboutMe(v string) {
	o.AboutMe = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *OpenGraphUser) GetBirthday() time.Time {
	if o == nil || o.Birthday == nil {
		var ret time.Time
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetBirthdayOk() (*time.Time, bool) {
	if o == nil || o.Birthday == nil {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *OpenGraphUser) HasBirthday() bool {
	if o != nil && o.Birthday != nil {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.
func (o *OpenGraphUser) SetBirthday(v time.Time) {
	o.Birthday = &v
}

// GetDrive returns the Drive field value if set, zero value otherwise.
func (o *OpenGraphUser) GetDrive() OpenGraphDrive {
	if o == nil || o.Drive == nil {
		var ret OpenGraphDrive
		return ret
	}
	return *o.Drive
}

// GetDriveOk returns a tuple with the Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetDriveOk() (*OpenGraphDrive, bool) {
	if o == nil || o.Drive == nil {
		return nil, false
	}
	return o.Drive, true
}

// HasDrive returns a boolean if a field has been set.
func (o *OpenGraphUser) HasDrive() bool {
	if o != nil && o.Drive != nil {
		return true
	}

	return false
}

// SetDrive gets a reference to the given OpenGraphDrive and assigns it to the Drive field.
func (o *OpenGraphUser) SetDrive(v OpenGraphDrive) {
	o.Drive = &v
}

// GetDrives returns the Drives field value if set, zero value otherwise.
func (o *OpenGraphUser) GetDrives() []OpenGraphDrive {
	if o == nil || o.Drives == nil {
		var ret []OpenGraphDrive
		return ret
	}
	return *o.Drives
}

// GetDrivesOk returns a tuple with the Drives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenGraphUser) GetDrivesOk() (*[]OpenGraphDrive, bool) {
	if o == nil || o.Drives == nil {
		return nil, false
	}
	return o.Drives, true
}

// HasDrives returns a boolean if a field has been set.
func (o *OpenGraphUser) HasDrives() bool {
	if o != nil && o.Drives != nil {
		return true
	}

	return false
}

// SetDrives gets a reference to the given []OpenGraphDrive and assigns it to the Drives field.
func (o *OpenGraphUser) SetDrives(v []OpenGraphDrive) {
	o.Drives = &v
}

func (o OpenGraphUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOpenGraphDirectoryObject, errOpenGraphDirectoryObject := json.Marshal(o.OpenGraphDirectoryObject)
	if errOpenGraphDirectoryObject != nil {
		return []byte{}, errOpenGraphDirectoryObject
	}
	errOpenGraphDirectoryObject = json.Unmarshal([]byte(serializedOpenGraphDirectoryObject), &toSerialize)
	if errOpenGraphDirectoryObject != nil {
		return []byte{}, errOpenGraphDirectoryObject
	}
	if o.AccountEnabled != nil {
		toSerialize["accountEnabled"] = o.AccountEnabled
	}
	if o.BusinessPhones != nil {
		toSerialize["businessPhones"] = o.BusinessPhones
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CompanyName != nil {
		toSerialize["companyName"] = o.CompanyName
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.FaxNumber != nil {
		toSerialize["faxNumber"] = o.FaxNumber
	}
	if o.GivenName != nil {
		toSerialize["givenName"] = o.GivenName
	}
	if o.LastPasswordChangeDateTime != nil {
		toSerialize["lastPasswordChangeDateTime"] = o.LastPasswordChangeDateTime
	}
	if o.LegalAgeGroupClassification != nil {
		toSerialize["legalAgeGroupClassification"] = o.LegalAgeGroupClassification
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.MailNickname != nil {
		toSerialize["mailNickname"] = o.MailNickname
	}
	if o.MobilePhone != nil {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if o.OfficeLocation != nil {
		toSerialize["officeLocation"] = o.OfficeLocation
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.PreferredLanguage != nil {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StreetAddress != nil {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if o.Surname != nil {
		toSerialize["surname"] = o.Surname
	}
	if o.UsageLocation != nil {
		toSerialize["usageLocation"] = o.UsageLocation
	}
	if o.UserPrincipalName != nil {
		toSerialize["userPrincipalName"] = o.UserPrincipalName
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}
	if o.AboutMe != nil {
		toSerialize["aboutMe"] = o.AboutMe
	}
	if o.Birthday != nil {
		toSerialize["birthday"] = o.Birthday
	}
	if o.Drive != nil {
		toSerialize["drive"] = o.Drive
	}
	if o.Drives != nil {
		toSerialize["drives"] = o.Drives
	}
	return json.Marshal(toSerialize)
}

type NullableOpenGraphUser struct {
	value *OpenGraphUser
	isSet bool
}

func (v NullableOpenGraphUser) Get() *OpenGraphUser {
	return v.value
}

func (v *NullableOpenGraphUser) Set(val *OpenGraphUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenGraphUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenGraphUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenGraphUser(val *OpenGraphUser) *NullableOpenGraphUser {
	return &NullableOpenGraphUser{value: val, isSet: true}
}

func (v NullableOpenGraphUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenGraphUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


