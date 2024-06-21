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

// checks if the VolumeState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeState{}

// VolumeState Runtime state of the volume
type VolumeState struct {
	// target exposed via a Nexus
	Target *Nexus `json:"target,omitempty"`
	// size of the volume in bytes
	Size int64 `json:"size"`
	Status VolumeStatus `json:"status"`
	// name of the volume
	Uuid string `json:"uuid"`
	// replica topology information
	ReplicaTopology map[string]ReplicaTopology `json:"replica_topology"`
	Usage *VolumeUsage `json:"usage,omitempty"`
}

type _VolumeState VolumeState

// NewVolumeState instantiates a new VolumeState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeState(size int64, status VolumeStatus, uuid string, replicaTopology map[string]ReplicaTopology) *VolumeState {
	this := VolumeState{}
	this.Size = size
	this.Status = status
	this.Uuid = uuid
	this.ReplicaTopology = replicaTopology
	return &this
}

// NewVolumeStateWithDefaults instantiates a new VolumeState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeStateWithDefaults() *VolumeState {
	this := VolumeState{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *VolumeState) GetTarget() Nexus {
	if o == nil || IsNil(o.Target) {
		var ret Nexus
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeState) GetTargetOk() (*Nexus, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *VolumeState) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given Nexus and assigns it to the Target field.
func (o *VolumeState) SetTarget(v Nexus) {
	o.Target = &v
}

// GetSize returns the Size field value
func (o *VolumeState) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VolumeState) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VolumeState) SetSize(v int64) {
	o.Size = v
}

// GetStatus returns the Status field value
func (o *VolumeState) GetStatus() VolumeStatus {
	if o == nil {
		var ret VolumeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VolumeState) GetStatusOk() (*VolumeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VolumeState) SetStatus(v VolumeStatus) {
	o.Status = v
}

// GetUuid returns the Uuid field value
func (o *VolumeState) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *VolumeState) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *VolumeState) SetUuid(v string) {
	o.Uuid = v
}

// GetReplicaTopology returns the ReplicaTopology field value
func (o *VolumeState) GetReplicaTopology() map[string]ReplicaTopology {
	if o == nil {
		var ret map[string]ReplicaTopology
		return ret
	}

	return o.ReplicaTopology
}

// GetReplicaTopologyOk returns a tuple with the ReplicaTopology field value
// and a boolean to check if the value has been set.
func (o *VolumeState) GetReplicaTopologyOk() (*map[string]ReplicaTopology, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicaTopology, true
}

// SetReplicaTopology sets field value
func (o *VolumeState) SetReplicaTopology(v map[string]ReplicaTopology) {
	o.ReplicaTopology = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *VolumeState) GetUsage() VolumeUsage {
	if o == nil || IsNil(o.Usage) {
		var ret VolumeUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeState) GetUsageOk() (*VolumeUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *VolumeState) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given VolumeUsage and assigns it to the Usage field.
func (o *VolumeState) SetUsage(v VolumeUsage) {
	o.Usage = &v
}

func (o VolumeState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	toSerialize["size"] = o.Size
	toSerialize["status"] = o.Status
	toSerialize["uuid"] = o.Uuid
	toSerialize["replica_topology"] = o.ReplicaTopology
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

func (o *VolumeState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
		"status",
		"uuid",
		"replica_topology",
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

	varVolumeState := _VolumeState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVolumeState)

	if err != nil {
		return err
	}

	*o = VolumeState(varVolumeState)

	return err
}

type NullableVolumeState struct {
	value *VolumeState
	isSet bool
}

func (v NullableVolumeState) Get() *VolumeState {
	return v.value
}

func (v *NullableVolumeState) Set(val *VolumeState) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeState) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeState(val *VolumeState) *NullableVolumeState {
	return &NullableVolumeState{value: val, isSet: true}
}

func (v NullableVolumeState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


