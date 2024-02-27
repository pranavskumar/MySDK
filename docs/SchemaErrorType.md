# SchemaErrorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**SchemaErrorCode**](SchemaErrorCode.md) |  | [optional] [default to EOK]
**ErrorObj** | Pointer to [**ProtobufAny**](ProtobufAny.md) |  | [optional] 
**Message** | Pointer to **string** |  A human readable string of the error  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 

## Methods

### NewSchemaErrorType

`func NewSchemaErrorType() *SchemaErrorType`

NewSchemaErrorType instantiates a new SchemaErrorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaErrorTypeWithDefaults

`func NewSchemaErrorTypeWithDefaults() *SchemaErrorType`

NewSchemaErrorTypeWithDefaults instantiates a new SchemaErrorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SchemaErrorType) GetCode() SchemaErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SchemaErrorType) GetCodeOk() (*SchemaErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SchemaErrorType) SetCode(v SchemaErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *SchemaErrorType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrorObj

`func (o *SchemaErrorType) GetErrorObj() ProtobufAny`

GetErrorObj returns the ErrorObj field if non-nil, zero value otherwise.

### GetErrorObjOk

`func (o *SchemaErrorType) GetErrorObjOk() (*ProtobufAny, bool)`

GetErrorObjOk returns a tuple with the ErrorObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorObj

`func (o *SchemaErrorType) SetErrorObj(v ProtobufAny)`

SetErrorObj sets ErrorObj field to given value.

### HasErrorObj

`func (o *SchemaErrorType) HasErrorObj() bool`

HasErrorObj returns a boolean if a field has been set.

### GetMessage

`func (o *SchemaErrorType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SchemaErrorType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SchemaErrorType) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SchemaErrorType) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


