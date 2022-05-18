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
	"fmt"
)

// EventsOrderBy the model 'EventsOrderBy'
type EventsOrderBy string

// List of EventsOrderBy
const (
	EVENTSORDERBY_DATE_DESCENDING EventsOrderBy = "DateDescending"
	EVENTSORDERBY_DATE_ASCENDING EventsOrderBy = "DateAscending"
)

// All allowed values of EventsOrderBy enum
var AllowedEventsOrderByEnumValues = []EventsOrderBy{
	"DateDescending",
	"DateAscending",
}

func (v *EventsOrderBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventsOrderBy(value)
	for _, existing := range AllowedEventsOrderByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventsOrderBy", value)
}

// NewEventsOrderByFromValue returns a pointer to a valid EventsOrderBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventsOrderByFromValue(v string) (*EventsOrderBy, error) {
	ev := EventsOrderBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventsOrderBy: valid values are %v", v, AllowedEventsOrderByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventsOrderBy) IsValid() bool {
	for _, existing := range AllowedEventsOrderByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventsOrderBy value
func (v EventsOrderBy) Ptr() *EventsOrderBy {
	return &v
}

type NullableEventsOrderBy struct {
	value *EventsOrderBy
	isSet bool
}

func (v NullableEventsOrderBy) Get() *EventsOrderBy {
	return v.value
}

func (v *NullableEventsOrderBy) Set(val *EventsOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsOrderBy(val *EventsOrderBy) *NullableEventsOrderBy {
	return &NullableEventsOrderBy{value: val, isSet: true}
}

func (v NullableEventsOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

