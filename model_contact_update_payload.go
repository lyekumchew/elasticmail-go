/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ContactUpdatePayload struct for ContactUpdatePayload
type ContactUpdatePayload struct {
	// First name.
	FirstName *string `json:"FirstName,omitempty"`
	// Last name.
	LastName *string `json:"LastName,omitempty"`
	// A key-value collection of custom contact fields which can be used in the system.
	CustomFields *map[string]string `json:"CustomFields,omitempty"`
}

// NewContactUpdatePayload instantiates a new ContactUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactUpdatePayload() *ContactUpdatePayload {
	this := ContactUpdatePayload{}
	return &this
}

// NewContactUpdatePayloadWithDefaults instantiates a new ContactUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactUpdatePayloadWithDefaults() *ContactUpdatePayload {
	this := ContactUpdatePayload{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ContactUpdatePayload) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactUpdatePayload) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ContactUpdatePayload) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ContactUpdatePayload) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ContactUpdatePayload) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactUpdatePayload) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ContactUpdatePayload) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ContactUpdatePayload) SetLastName(v string) {
	o.LastName = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ContactUpdatePayload) GetCustomFields() map[string]string {
	if o == nil || o.CustomFields == nil {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactUpdatePayload) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ContactUpdatePayload) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *ContactUpdatePayload) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

func (o ContactUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["FirstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["LastName"] = o.LastName
	}
	if o.CustomFields != nil {
		toSerialize["CustomFields"] = o.CustomFields
	}
	return json.Marshal(toSerialize)
}

type NullableContactUpdatePayload struct {
	value *ContactUpdatePayload
	isSet bool
}

func (v NullableContactUpdatePayload) Get() *ContactUpdatePayload {
	return v.value
}

func (v *NullableContactUpdatePayload) Set(val *ContactUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableContactUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableContactUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactUpdatePayload(val *ContactUpdatePayload) *NullableContactUpdatePayload {
	return &NullableContactUpdatePayload{value: val, isSet: true}
}

func (v NullableContactUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


