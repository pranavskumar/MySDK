# SchemaInitializersType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to [**[]SchemaInitializerType**](SchemaInitializerType.md) |  Pending is a list of initializers that must execute in order before this object is initialized.  When the last pending initializer is removed, and no failing result is set, the initializers  struct will be set to nil and the object is considered as initialized and visible to all  clients. | [optional] 
**Result** | Pointer to [**SchemaStatusType**](SchemaStatusType.md) |  | [optional] 

## Methods

### NewSchemaInitializersType

`func NewSchemaInitializersType() *SchemaInitializersType`

NewSchemaInitializersType instantiates a new SchemaInitializersType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaInitializersTypeWithDefaults

`func NewSchemaInitializersTypeWithDefaults() *SchemaInitializersType`

NewSchemaInitializersTypeWithDefaults instantiates a new SchemaInitializersType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *SchemaInitializersType) GetPending() []SchemaInitializerType`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *SchemaInitializersType) GetPendingOk() (*[]SchemaInitializerType, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *SchemaInitializersType) SetPending(v []SchemaInitializerType)`

SetPending sets Pending field to given value.

### HasPending

`func (o *SchemaInitializersType) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetResult

`func (o *SchemaInitializersType) GetResult() SchemaStatusType`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SchemaInitializersType) GetResultOk() (*SchemaStatusType, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SchemaInitializersType) SetResult(v SchemaStatusType)`

SetResult sets Result field to given value.

### HasResult

`func (o *SchemaInitializersType) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


