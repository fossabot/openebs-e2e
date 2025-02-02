# AppNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | App node identifier. | 
**Spec** | [**AppNodeSpec**](AppNodeSpec.md) |  | 
**State** | Pointer to [**AppNodeState**](AppNodeState.md) |  | [optional] 

## Methods

### NewAppNode

`func NewAppNode(id string, spec AppNodeSpec, ) *AppNode`

NewAppNode instantiates a new AppNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppNodeWithDefaults

`func NewAppNodeWithDefaults() *AppNode`

NewAppNodeWithDefaults instantiates a new AppNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppNode) SetId(v string)`

SetId sets Id field to given value.


### GetSpec

`func (o *AppNode) GetSpec() AppNodeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AppNode) GetSpecOk() (*AppNodeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AppNode) SetSpec(v AppNodeSpec)`

SetSpec sets Spec field to given value.


### GetState

`func (o *AppNode) GetState() AppNodeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppNode) GetStateOk() (*AppNodeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppNode) SetState(v AppNodeState)`

SetState sets State field to given value.

### HasState

`func (o *AppNode) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


