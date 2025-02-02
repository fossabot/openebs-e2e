/*
IoEngine RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BlockDevicePartition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockDevicePartition{}

// BlockDevicePartition partition information in case where device represents a partition
type BlockDevicePartition struct {
	// partition name
	Name string `json:"name"`
	// partition number
	Number int32 `json:"number"`
	// devname of parent device to which this partition belongs
	Parent string `json:"parent"`
	// partition scheme: gpt, dos, ...
	Scheme string `json:"scheme"`
	// partition type identifier
	Typeid string `json:"typeid"`
	// UUID identifying partition
	Uuid string `json:"uuid"`
}

type _BlockDevicePartition BlockDevicePartition

// NewBlockDevicePartition instantiates a new BlockDevicePartition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockDevicePartition(name string, number int32, parent string, scheme string, typeid string, uuid string) *BlockDevicePartition {
	this := BlockDevicePartition{}
	this.Name = name
	this.Number = number
	this.Parent = parent
	this.Scheme = scheme
	this.Typeid = typeid
	this.Uuid = uuid
	return &this
}

// NewBlockDevicePartitionWithDefaults instantiates a new BlockDevicePartition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockDevicePartitionWithDefaults() *BlockDevicePartition {
	this := BlockDevicePartition{}
	return &this
}

// GetName returns the Name field value
func (o *BlockDevicePartition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlockDevicePartition) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *BlockDevicePartition) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *BlockDevicePartition) SetNumber(v int32) {
	o.Number = v
}

// GetParent returns the Parent field value
func (o *BlockDevicePartition) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *BlockDevicePartition) SetParent(v string) {
	o.Parent = v
}

// GetScheme returns the Scheme field value
func (o *BlockDevicePartition) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *BlockDevicePartition) SetScheme(v string) {
	o.Scheme = v
}

// GetTypeid returns the Typeid field value
func (o *BlockDevicePartition) GetTypeid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typeid
}

// GetTypeidOk returns a tuple with the Typeid field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetTypeidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typeid, true
}

// SetTypeid sets field value
func (o *BlockDevicePartition) SetTypeid(v string) {
	o.Typeid = v
}

// GetUuid returns the Uuid field value
func (o *BlockDevicePartition) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *BlockDevicePartition) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *BlockDevicePartition) SetUuid(v string) {
	o.Uuid = v
}

func (o BlockDevicePartition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockDevicePartition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["number"] = o.Number
	toSerialize["parent"] = o.Parent
	toSerialize["scheme"] = o.Scheme
	toSerialize["typeid"] = o.Typeid
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *BlockDevicePartition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"number",
		"parent",
		"scheme",
		"typeid",
		"uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBlockDevicePartition := _BlockDevicePartition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockDevicePartition)

	if err != nil {
		return err
	}

	*o = BlockDevicePartition(varBlockDevicePartition)

	return err
}

type NullableBlockDevicePartition struct {
	value *BlockDevicePartition
	isSet bool
}

func (v NullableBlockDevicePartition) Get() *BlockDevicePartition {
	return v.value
}

func (v *NullableBlockDevicePartition) Set(val *BlockDevicePartition) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockDevicePartition) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockDevicePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockDevicePartition(val *BlockDevicePartition) *NullableBlockDevicePartition {
	return &NullableBlockDevicePartition{value: val, isSet: true}
}

func (v NullableBlockDevicePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockDevicePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


