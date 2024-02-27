# SchemaStatusType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  Suggested HTTP return code for this status, 0 if not set.  Example: &#x60; \&quot;0\&quot;&#x60; | [optional] 
**Reason** | Pointer to **string** |  A human-readable description of why this operation is in the  \&quot;Failure\&quot; status. If this value is empty there  is no information available.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**Status** | Pointer to **string** |  Status of the operation.  One of: \&quot;Success\&quot; or \&quot;Failure\&quot;.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 

## Methods

### NewSchemaStatusType

`func NewSchemaStatusType() *SchemaStatusType`

NewSchemaStatusType instantiates a new SchemaStatusType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaStatusTypeWithDefaults

`func NewSchemaStatusTypeWithDefaults() *SchemaStatusType`

NewSchemaStatusTypeWithDefaults instantiates a new SchemaStatusType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SchemaStatusType) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SchemaStatusType) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SchemaStatusType) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *SchemaStatusType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetReason

`func (o *SchemaStatusType) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SchemaStatusType) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SchemaStatusType) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SchemaStatusType) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *SchemaStatusType) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchemaStatusType) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchemaStatusType) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchemaStatusType) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


