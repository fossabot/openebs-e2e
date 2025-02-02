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

// checks if the SnapshotAsSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotAsSource{}

// SnapshotAsSource The snapshot source for the volume content.
type SnapshotAsSource struct {
	Snapshot string `json:"snapshot"`
	Volume string `json:"volume"`
}

type _SnapshotAsSource SnapshotAsSource

// NewSnapshotAsSource instantiates a new SnapshotAsSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotAsSource(snapshot string, volume string) *SnapshotAsSource {
	this := SnapshotAsSource{}
	this.Snapshot = snapshot
	this.Volume = volume
	return &this
}

// NewSnapshotAsSourceWithDefaults instantiates a new SnapshotAsSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotAsSourceWithDefaults() *SnapshotAsSource {
	this := SnapshotAsSource{}
	return &this
}

// GetSnapshot returns the Snapshot field value
func (o *SnapshotAsSource) GetSnapshot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value
// and a boolean to check if the value has been set.
func (o *SnapshotAsSource) GetSnapshotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snapshot, true
}

// SetSnapshot sets field value
func (o *SnapshotAsSource) SetSnapshot(v string) {
	o.Snapshot = v
}

// GetVolume returns the Volume field value
func (o *SnapshotAsSource) GetVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *SnapshotAsSource) GetVolumeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *SnapshotAsSource) SetVolume(v string) {
	o.Volume = v
}

func (o SnapshotAsSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotAsSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snapshot"] = o.Snapshot
	toSerialize["volume"] = o.Volume
	return toSerialize, nil
}

func (o *SnapshotAsSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"snapshot",
		"volume",
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

	varSnapshotAsSource := _SnapshotAsSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnapshotAsSource)

	if err != nil {
		return err
	}

	*o = SnapshotAsSource(varSnapshotAsSource)

	return err
}

type NullableSnapshotAsSource struct {
	value *SnapshotAsSource
	isSet bool
}

func (v NullableSnapshotAsSource) Get() *SnapshotAsSource {
	return v.value
}

func (v *NullableSnapshotAsSource) Set(val *SnapshotAsSource) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotAsSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotAsSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotAsSource(val *SnapshotAsSource) *NullableSnapshotAsSource {
	return &NullableSnapshotAsSource{value: val, isSet: true}
}

func (v NullableSnapshotAsSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotAsSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


