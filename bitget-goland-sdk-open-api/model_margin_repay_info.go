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

// MarginRepayInfo struct for MarginRepayInfo
type MarginRepayInfo struct {
	Amount               *string `json:"amount,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Ctime                *string `json:"ctime,omitempty"`
	Interest             *string `json:"interest,omitempty"`
	RepayId              *string `json:"repayId,omitempty"`
	TotalAmount          *string `json:"totalAmount,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginRepayInfo MarginRepayInfo

// NewMarginRepayInfo instantiates a new MarginRepayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginRepayInfo() *MarginRepayInfo {
	this := MarginRepayInfo{}
	return &this
}

// NewMarginRepayInfoWithDefaults instantiates a new MarginRepayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginRepayInfoWithDefaults() *MarginRepayInfo {
	this := MarginRepayInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetAmount() string {
	if o == nil || isNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetAmountOk() (*string, bool) {
	if o == nil || isNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *MarginRepayInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetCoin() string {
	if o == nil || isNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetCoinOk() (*string, bool) {
	if o == nil || isNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasCoin() bool {
	if o != nil && !isNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *MarginRepayInfo) SetCoin(v string) {
	o.Coin = &v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetCtime() string {
	if o == nil || isNil(o.Ctime) {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetCtimeOk() (*string, bool) {
	if o == nil || isNil(o.Ctime) {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasCtime() bool {
	if o != nil && !isNil(o.Ctime) {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *MarginRepayInfo) SetCtime(v string) {
	o.Ctime = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetInterest() string {
	if o == nil || isNil(o.Interest) {
		var ret string
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetInterestOk() (*string, bool) {
	if o == nil || isNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasInterest() bool {
	if o != nil && !isNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given string and assigns it to the Interest field.
func (o *MarginRepayInfo) SetInterest(v string) {
	o.Interest = &v
}

// GetRepayId returns the RepayId field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetRepayId() string {
	if o == nil || isNil(o.RepayId) {
		var ret string
		return ret
	}
	return *o.RepayId
}

// GetRepayIdOk returns a tuple with the RepayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetRepayIdOk() (*string, bool) {
	if o == nil || isNil(o.RepayId) {
		return nil, false
	}
	return o.RepayId, true
}

// HasRepayId returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasRepayId() bool {
	if o != nil && !isNil(o.RepayId) {
		return true
	}

	return false
}

// SetRepayId gets a reference to the given string and assigns it to the RepayId field.
func (o *MarginRepayInfo) SetRepayId(v string) {
	o.RepayId = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetTotalAmount() string {
	if o == nil || isNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetTotalAmountOk() (*string, bool) {
	if o == nil || isNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasTotalAmount() bool {
	if o != nil && !isNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *MarginRepayInfo) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MarginRepayInfo) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginRepayInfo) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MarginRepayInfo) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MarginRepayInfo) SetType(v string) {
	o.Type = &v
}

func (o MarginRepayInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !isNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !isNil(o.Ctime) {
		toSerialize["ctime"] = o.Ctime
	}
	if !isNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !isNil(o.RepayId) {
		toSerialize["repayId"] = o.RepayId
	}
	if !isNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarginRepayInfo) UnmarshalJSON(bytes []byte) (err error) {
	varMarginRepayInfo := _MarginRepayInfo{}

	if err = json.Unmarshal(bytes, &varMarginRepayInfo); err == nil {
		*o = MarginRepayInfo(varMarginRepayInfo)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "ctime")
		delete(additionalProperties, "interest")
		delete(additionalProperties, "repayId")
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginRepayInfo struct {
	value *MarginRepayInfo
	isSet bool
}

func (v NullableMarginRepayInfo) Get() *MarginRepayInfo {
	return v.value
}

func (v *NullableMarginRepayInfo) Set(val *MarginRepayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginRepayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginRepayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginRepayInfo(val *MarginRepayInfo) *NullableMarginRepayInfo {
	return &NullableMarginRepayInfo{value: val, isSet: true}
}

func (v NullableMarginRepayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginRepayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}