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

// checks if the ReplicaSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplicaSpec{}

// ReplicaSpec User specification of a replica.
type ReplicaSpec struct {
	// Managed by our control plane
	Managed bool `json:"managed"`
	Operation *ReplicaSpecOperation `json:"operation,omitempty"`
	Owners ReplicaSpecOwners `json:"owners"`
	// The pool that the replica should live on.
	Pool string `json:"pool"`
	// storage pool unique identifier
	PoolUuid *string `json:"poolUuid,omitempty"`
	Share Protocol `json:"share"`
	// The size that the replica should be.
	Size int64 `json:"size"`
	Status SpecStatus `json:"status"`
	// Thin provisioning.
	Thin bool `json:"thin"`
	// uuid of the replica
	Uuid string `json:"uuid"`
	Kind *ReplicaKind `json:"kind,omitempty"`
}

type _ReplicaSpec ReplicaSpec

// NewReplicaSpec instantiates a new ReplicaSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplicaSpec(managed bool, owners ReplicaSpecOwners, pool string, share Protocol, size int64, status SpecStatus, thin bool, uuid string) *ReplicaSpec {
	this := ReplicaSpec{}
	this.Managed = managed
	this.Owners = owners
	this.Pool = pool
	this.Share = share
	this.Size = size
	this.Status = status
	this.Thin = thin
	this.Uuid = uuid
	var kind ReplicaKind = REGULAR
	this.Kind = &kind
	return &this
}

// NewReplicaSpecWithDefaults instantiates a new ReplicaSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplicaSpecWithDefaults() *ReplicaSpec {
	this := ReplicaSpec{}
	var kind ReplicaKind = REGULAR
	this.Kind = &kind
	return &this
}

// GetManaged returns the Managed field value
func (o *ReplicaSpec) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *ReplicaSpec) SetManaged(v bool) {
	o.Managed = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ReplicaSpec) GetOperation() ReplicaSpecOperation {
	if o == nil || IsNil(o.Operation) {
		var ret ReplicaSpecOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetOperationOk() (*ReplicaSpecOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ReplicaSpec) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given ReplicaSpecOperation and assigns it to the Operation field.
func (o *ReplicaSpec) SetOperation(v ReplicaSpecOperation) {
	o.Operation = &v
}

// GetOwners returns the Owners field value
func (o *ReplicaSpec) GetOwners() ReplicaSpecOwners {
	if o == nil {
		var ret ReplicaSpecOwners
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetOwnersOk() (*ReplicaSpecOwners, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *ReplicaSpec) SetOwners(v ReplicaSpecOwners) {
	o.Owners = v
}

// GetPool returns the Pool field value
func (o *ReplicaSpec) GetPool() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pool
}

// GetPoolOk returns a tuple with the Pool field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetPoolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pool, true
}

// SetPool sets field value
func (o *ReplicaSpec) SetPool(v string) {
	o.Pool = v
}

// GetPoolUuid returns the PoolUuid field value if set, zero value otherwise.
func (o *ReplicaSpec) GetPoolUuid() string {
	if o == nil || IsNil(o.PoolUuid) {
		var ret string
		return ret
	}
	return *o.PoolUuid
}

// GetPoolUuidOk returns a tuple with the PoolUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetPoolUuidOk() (*string, bool) {
	if o == nil || IsNil(o.PoolUuid) {
		return nil, false
	}
	return o.PoolUuid, true
}

// HasPoolUuid returns a boolean if a field has been set.
func (o *ReplicaSpec) HasPoolUuid() bool {
	if o != nil && !IsNil(o.PoolUuid) {
		return true
	}

	return false
}

// SetPoolUuid gets a reference to the given string and assigns it to the PoolUuid field.
func (o *ReplicaSpec) SetPoolUuid(v string) {
	o.PoolUuid = &v
}

// GetShare returns the Share field value
func (o *ReplicaSpec) GetShare() Protocol {
	if o == nil {
		var ret Protocol
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetShareOk() (*Protocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *ReplicaSpec) SetShare(v Protocol) {
	o.Share = v
}

// GetSize returns the Size field value
func (o *ReplicaSpec) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ReplicaSpec) SetSize(v int64) {
	o.Size = v
}

// GetStatus returns the Status field value
func (o *ReplicaSpec) GetStatus() SpecStatus {
	if o == nil {
		var ret SpecStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetStatusOk() (*SpecStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReplicaSpec) SetStatus(v SpecStatus) {
	o.Status = v
}

// GetThin returns the Thin field value
func (o *ReplicaSpec) GetThin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Thin
}

// GetThinOk returns a tuple with the Thin field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetThinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thin, true
}

// SetThin sets field value
func (o *ReplicaSpec) SetThin(v bool) {
	o.Thin = v
}

// GetUuid returns the Uuid field value
func (o *ReplicaSpec) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ReplicaSpec) SetUuid(v string) {
	o.Uuid = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ReplicaSpec) GetKind() ReplicaKind {
	if o == nil || IsNil(o.Kind) {
		var ret ReplicaKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplicaSpec) GetKindOk() (*ReplicaKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ReplicaSpec) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given ReplicaKind and assigns it to the Kind field.
func (o *ReplicaSpec) SetKind(v ReplicaKind) {
	o.Kind = &v
}

func (o ReplicaSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplicaSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["managed"] = o.Managed
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	toSerialize["owners"] = o.Owners
	toSerialize["pool"] = o.Pool
	if !IsNil(o.PoolUuid) {
		toSerialize["poolUuid"] = o.PoolUuid
	}
	toSerialize["share"] = o.Share
	toSerialize["size"] = o.Size
	toSerialize["status"] = o.Status
	toSerialize["thin"] = o.Thin
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

func (o *ReplicaSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"managed",
		"owners",
		"pool",
		"share",
		"size",
		"status",
		"thin",
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

	varReplicaSpec := _ReplicaSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplicaSpec)

	if err != nil {
		return err
	}

	*o = ReplicaSpec(varReplicaSpec)

	return err
}

type NullableReplicaSpec struct {
	value *ReplicaSpec
	isSet bool
}

func (v NullableReplicaSpec) Get() *ReplicaSpec {
	return v.value
}

func (v *NullableReplicaSpec) Set(val *ReplicaSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableReplicaSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableReplicaSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplicaSpec(val *ReplicaSpec) *NullableReplicaSpec {
	return &NullableReplicaSpec{value: val, isSet: true}
}

func (v NullableReplicaSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplicaSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


