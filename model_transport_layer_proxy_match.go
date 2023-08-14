/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.104.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
	"fmt"
)

// TransportLayerProxyMatch - struct for TransportLayerProxyMatch
type TransportLayerProxyMatch struct {
	ArrayOfString *[]string
	String *string
}

// []stringAsTransportLayerProxyMatch is a convenience function that returns []string wrapped in TransportLayerProxyMatch
func ArrayOfStringAsTransportLayerProxyMatch(v *[]string) TransportLayerProxyMatch {
	return TransportLayerProxyMatch{
		ArrayOfString: v,
	}
}

// stringAsTransportLayerProxyMatch is a convenience function that returns string wrapped in TransportLayerProxyMatch
func StringAsTransportLayerProxyMatch(v *string) TransportLayerProxyMatch {
	return TransportLayerProxyMatch{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TransportLayerProxyMatch) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TransportLayerProxyMatch)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TransportLayerProxyMatch)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TransportLayerProxyMatch) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TransportLayerProxyMatch) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableTransportLayerProxyMatch struct {
	value *TransportLayerProxyMatch
	isSet bool
}

func (v NullableTransportLayerProxyMatch) Get() *TransportLayerProxyMatch {
	return v.value
}

func (v *NullableTransportLayerProxyMatch) Set(val *TransportLayerProxyMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyMatch(val *TransportLayerProxyMatch) *NullableTransportLayerProxyMatch {
	return &NullableTransportLayerProxyMatch{value: val, isSet: true}
}

func (v NullableTransportLayerProxyMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


