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

// checks if the Specs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Specs{}

// Specs Specs detailing the requested configuration of the objects.
type Specs struct {
	// Nexus Specs
	Nexuses []NexusSpec `json:"nexuses"`
	// Pool Specs
	Pools []PoolSpec `json:"pools"`
	// Replica Specs
	Replicas []ReplicaSpec `json:"replicas"`
	// Volume Specs
	Volumes []VolumeSpec `json:"volumes"`
}

type _Specs Specs

// NewSpecs instantiates a new Specs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecs(nexuses []NexusSpec, pools []PoolSpec, replicas []ReplicaSpec, volumes []VolumeSpec) *Specs {
	this := Specs{}
	this.Nexuses = nexuses
	this.Pools = pools
	this.Replicas = replicas
	this.Volumes = volumes
	return &this
}

// NewSpecsWithDefaults instantiates a new Specs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecsWithDefaults() *Specs {
	this := Specs{}
	return &this
}

// GetNexuses returns the Nexuses field value
func (o *Specs) GetNexuses() []NexusSpec {
	if o == nil {
		var ret []NexusSpec
		return ret
	}

	return o.Nexuses
}

// GetNexusesOk returns a tuple with the Nexuses field value
// and a boolean to check if the value has been set.
func (o *Specs) GetNexusesOk() ([]NexusSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nexuses, true
}

// SetNexuses sets field value
func (o *Specs) SetNexuses(v []NexusSpec) {
	o.Nexuses = v
}

// GetPools returns the Pools field value
func (o *Specs) GetPools() []PoolSpec {
	if o == nil {
		var ret []PoolSpec
		return ret
	}

	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value
// and a boolean to check if the value has been set.
func (o *Specs) GetPoolsOk() ([]PoolSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pools, true
}

// SetPools sets field value
func (o *Specs) SetPools(v []PoolSpec) {
	o.Pools = v
}

// GetReplicas returns the Replicas field value
func (o *Specs) GetReplicas() []ReplicaSpec {
	if o == nil {
		var ret []ReplicaSpec
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *Specs) GetReplicasOk() ([]ReplicaSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas, true
}

// SetReplicas sets field value
func (o *Specs) SetReplicas(v []ReplicaSpec) {
	o.Replicas = v
}

// GetVolumes returns the Volumes field value
func (o *Specs) GetVolumes() []VolumeSpec {
	if o == nil {
		var ret []VolumeSpec
		return ret
	}

	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *Specs) GetVolumesOk() ([]VolumeSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Volumes, true
}

// SetVolumes sets field value
func (o *Specs) SetVolumes(v []VolumeSpec) {
	o.Volumes = v
}

func (o Specs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Specs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nexuses"] = o.Nexuses
	toSerialize["pools"] = o.Pools
	toSerialize["replicas"] = o.Replicas
	toSerialize["volumes"] = o.Volumes
	return toSerialize, nil
}

func (o *Specs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nexuses",
		"pools",
		"replicas",
		"volumes",
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

	varSpecs := _Specs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpecs)

	if err != nil {
		return err
	}

	*o = Specs(varSpecs)

	return err
}

type NullableSpecs struct {
	value *Specs
	isSet bool
}

func (v NullableSpecs) Get() *Specs {
	return v.value
}

func (v *NullableSpecs) Set(val *Specs) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecs(val *Specs) *NullableSpecs {
	return &NullableSpecs{value: val, isSet: true}
}

func (v NullableSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


