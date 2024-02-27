# SchemaSecretType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlindfoldSecretInfo** | Pointer to [**SchemaBlindfoldSecretInfoType**](SchemaBlindfoldSecretInfoType.md) |  | [optional] 
**ClearSecretInfo** | Pointer to [**SchemaClearSecretInfoType**](SchemaClearSecretInfoType.md) |  | [optional] 

## Methods

### NewSchemaSecretType

`func NewSchemaSecretType() *SchemaSecretType`

NewSchemaSecretType instantiates a new SchemaSecretType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaSecretTypeWithDefaults

`func NewSchemaSecretTypeWithDefaults() *SchemaSecretType`

NewSchemaSecretTypeWithDefaults instantiates a new SchemaSecretType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlindfoldSecretInfo

`func (o *SchemaSecretType) GetBlindfoldSecretInfo() SchemaBlindfoldSecretInfoType`

GetBlindfoldSecretInfo returns the BlindfoldSecretInfo field if non-nil, zero value otherwise.

### GetBlindfoldSecretInfoOk

`func (o *SchemaSecretType) GetBlindfoldSecretInfoOk() (*SchemaBlindfoldSecretInfoType, bool)`

GetBlindfoldSecretInfoOk returns a tuple with the BlindfoldSecretInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlindfoldSecretInfo

`func (o *SchemaSecretType) SetBlindfoldSecretInfo(v SchemaBlindfoldSecretInfoType)`

SetBlindfoldSecretInfo sets BlindfoldSecretInfo field to given value.

### HasBlindfoldSecretInfo

`func (o *SchemaSecretType) HasBlindfoldSecretInfo() bool`

HasBlindfoldSecretInfo returns a boolean if a field has been set.

### GetClearSecretInfo

`func (o *SchemaSecretType) GetClearSecretInfo() SchemaClearSecretInfoType`

GetClearSecretInfo returns the ClearSecretInfo field if non-nil, zero value otherwise.

### GetClearSecretInfoOk

`func (o *SchemaSecretType) GetClearSecretInfoOk() (*SchemaClearSecretInfoType, bool)`

GetClearSecretInfoOk returns a tuple with the ClearSecretInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSecretInfo

`func (o *SchemaSecretType) SetClearSecretInfo(v SchemaClearSecretInfoType)`

SetClearSecretInfo sets ClearSecretInfo field to given value.

### HasClearSecretInfo

`func (o *SchemaSecretType) HasClearSecretInfo() bool`

HasClearSecretInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


