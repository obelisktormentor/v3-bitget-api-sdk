/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EndOrderRequest struct for EndOrderRequest
type EndOrderRequest struct {
	// trackingOrderNos
	TrackingOrderNos     []string `json:"trackingOrderNos"`
	AdditionalProperties map[string]interface{}
}

type _EndOrderRequest EndOrderRequest

// NewEndOrderRequest instantiates a new EndOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndOrderRequest(trackingOrderNos []string) *EndOrderRequest {
	this := EndOrderRequest{}
	this.TrackingOrderNos = trackingOrderNos
	return &this
}

// NewEndOrderRequestWithDefaults instantiates a new EndOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndOrderRequestWithDefaults() *EndOrderRequest {
	this := EndOrderRequest{}
	return &this
}

// GetTrackingOrderNos returns the TrackingOrderNos field value
func (o *EndOrderRequest) GetTrackingOrderNos() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrackingOrderNos
}

// GetTrackingOrderNosOk returns a tuple with the TrackingOrderNos field value
// and a boolean to check if the value has been set.
func (o *EndOrderRequest) GetTrackingOrderNosOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingOrderNos, true
}

// SetTrackingOrderNos sets field value
func (o *EndOrderRequest) SetTrackingOrderNos(v []string) {
	o.TrackingOrderNos = v
}

func (o EndOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trackingOrderNos"] = o.TrackingOrderNos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EndOrderRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEndOrderRequest := _EndOrderRequest{}

	if err = json.Unmarshal(bytes, &varEndOrderRequest); err == nil {
		*o = EndOrderRequest(varEndOrderRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "trackingOrderNos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndOrderRequest struct {
	value *EndOrderRequest
	isSet bool
}

func (v NullableEndOrderRequest) Get() *EndOrderRequest {
	return v.value
}

func (v *NullableEndOrderRequest) Set(val *EndOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndOrderRequest(val *EndOrderRequest) *NullableEndOrderRequest {
	return &NullableEndOrderRequest{value: val, isSet: true}
}

func (v NullableEndOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}