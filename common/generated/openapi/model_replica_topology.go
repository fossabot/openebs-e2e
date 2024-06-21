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

// checks if the ReplicaTopology type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicaTopology{}

// ReplicaTopology Volume Replica information.
type ReplicaTopology struct {
	// storage node identifier
	Node *string `json:"node,omitempty"`
	// storage pool identifier
	Pool *string `json:"pool,omitempty"`
	State ReplicaState `json:"state"`
	ChildStatus *ChildState `json:"child-status,omitempty"`
	ChildStatusReason *ChildStateReason `json:"child-status-reason,omitempty"`
	Usage *ReplicaUsage `json:"usage,omitempty"`
	// current rebuild progress (%)
	RebuildProgress *int32 `json:"rebuild-progress,omitempty"`
}

type _ReplicaTopology ReplicaTopology

// NewReplicaTopology instantiates a new ReplicaTopology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaTopology(state ReplicaState) *ReplicaTopology {
	this := ReplicaTopology{}
	this.State = state
	return &this
}

// NewReplicaTopologyWithDefaults instantiates a new ReplicaTopology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaTopologyWithDefaults() *ReplicaTopology {
	this := ReplicaTopology{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *ReplicaTopology) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *ReplicaTopology) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *ReplicaTopology) SetNode(v string) {
	o.Node = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ReplicaTopology) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *ReplicaTopology) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *ReplicaTopology) SetPool(v string) {
	o.Pool = &v
}

// GetState returns the State field value
func (o *ReplicaTopology) GetState() ReplicaState {
	if o == nil {
		var ret ReplicaState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetStateOk() (*ReplicaState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReplicaTopology) SetState(v ReplicaState) {
	o.State = v
}

// GetChildStatus returns the ChildStatus field value if set, zero value otherwise.
func (o *ReplicaTopology) GetChildStatus() ChildState {
	if o == nil || IsNil(o.ChildStatus) {
		var ret ChildState
		return ret
	}
	return *o.ChildStatus
}

// GetChildStatusOk returns a tuple with the ChildStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetChildStatusOk() (*ChildState, bool) {
	if o == nil || IsNil(o.ChildStatus) {
		return nil, false
	}
	return o.ChildStatus, true
}

// HasChildStatus returns a boolean if a field has been set.
func (o *ReplicaTopology) HasChildStatus() bool {
	if o != nil && !IsNil(o.ChildStatus) {
		return true
	}

	return false
}

// SetChildStatus gets a reference to the given ChildState and assigns it to the ChildStatus field.
func (o *ReplicaTopology) SetChildStatus(v ChildState) {
	o.ChildStatus = &v
}

// GetChildStatusReason returns the ChildStatusReason field value if set, zero value otherwise.
func (o *ReplicaTopology) GetChildStatusReason() ChildStateReason {
	if o == nil || IsNil(o.ChildStatusReason) {
		var ret ChildStateReason
		return ret
	}
	return *o.ChildStatusReason
}

// GetChildStatusReasonOk returns a tuple with the ChildStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetChildStatusReasonOk() (*ChildStateReason, bool) {
	if o == nil || IsNil(o.ChildStatusReason) {
		return nil, false
	}
	return o.ChildStatusReason, true
}

// HasChildStatusReason returns a boolean if a field has been set.
func (o *ReplicaTopology) HasChildStatusReason() bool {
	if o != nil && !IsNil(o.ChildStatusReason) {
		return true
	}

	return false
}

// SetChildStatusReason gets a reference to the given ChildStateReason and assigns it to the ChildStatusReason field.
func (o *ReplicaTopology) SetChildStatusReason(v ChildStateReason) {
	o.ChildStatusReason = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ReplicaTopology) GetUsage() ReplicaUsage {
	if o == nil || IsNil(o.Usage) {
		var ret ReplicaUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetUsageOk() (*ReplicaUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ReplicaTopology) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given ReplicaUsage and assigns it to the Usage field.
func (o *ReplicaTopology) SetUsage(v ReplicaUsage) {
	o.Usage = &v
}

// GetRebuildProgress returns the RebuildProgress field value if set, zero value otherwise.
func (o *ReplicaTopology) GetRebuildProgress() int32 {
	if o == nil || IsNil(o.RebuildProgress) {
		var ret int32
		return ret
	}
	return *o.RebuildProgress
}

// GetRebuildProgressOk returns a tuple with the RebuildProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaTopology) GetRebuildProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.RebuildProgress) {
		return nil, false
	}
	return o.RebuildProgress, true
}

// HasRebuildProgress returns a boolean if a field has been set.
func (o *ReplicaTopology) HasRebuildProgress() bool {
	if o != nil && !IsNil(o.RebuildProgress) {
		return true
	}

	return false
}

// SetRebuildProgress gets a reference to the given int32 and assigns it to the RebuildProgress field.
func (o *ReplicaTopology) SetRebuildProgress(v int32) {
	o.RebuildProgress = &v
}

func (o ReplicaTopology) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicaTopology) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	toSerialize["state"] = o.State
	if !IsNil(o.ChildStatus) {
		toSerialize["child-status"] = o.ChildStatus
	}
	if !IsNil(o.ChildStatusReason) {
		toSerialize["child-status-reason"] = o.ChildStatusReason
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.RebuildProgress) {
		toSerialize["rebuild-progress"] = o.RebuildProgress
	}
	return toSerialize, nil
}

func (o *ReplicaTopology) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
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

	varReplicaTopology := _ReplicaTopology{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplicaTopology)

	if err != nil {
		return err
	}

	*o = ReplicaTopology(varReplicaTopology)

	return err
}

type NullableReplicaTopology struct {
	value *ReplicaTopology
	isSet bool
}

func (v NullableReplicaTopology) Get() *ReplicaTopology {
	return v.value
}

func (v *NullableReplicaTopology) Set(val *ReplicaTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaTopology(val *ReplicaTopology) *NullableReplicaTopology {
	return &NullableReplicaTopology{value: val, isSet: true}
}

func (v NullableReplicaTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


