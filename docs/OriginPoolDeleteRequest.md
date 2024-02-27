# OriginPoolDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailIfReferred** | Pointer to **bool** |  Fail the delete operation if this object is being referred by other objects | [optional] 
**Name** | Pointer to **string** |  Name of the configuration object  Example: &#x60; \&quot;name\&quot;&#x60; | [optional] 
**Namespace** | Pointer to **string** |  Namespace in which the configuration object is present  Example: &#x60; \&quot;ns1\&quot;&#x60; | [optional] 

## Methods

### NewOriginPoolDeleteRequest

`func NewOriginPoolDeleteRequest() *OriginPoolDeleteRequest`

NewOriginPoolDeleteRequest instantiates a new OriginPoolDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolDeleteRequestWithDefaults

`func NewOriginPoolDeleteRequestWithDefaults() *OriginPoolDeleteRequest`

NewOriginPoolDeleteRequestWithDefaults instantiates a new OriginPoolDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailIfReferred

`func (o *OriginPoolDeleteRequest) GetFailIfReferred() bool`

GetFailIfReferred returns the FailIfReferred field if non-nil, zero value otherwise.

### GetFailIfReferredOk

`func (o *OriginPoolDeleteRequest) GetFailIfReferredOk() (*bool, bool)`

GetFailIfReferredOk returns a tuple with the FailIfReferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIfReferred

`func (o *OriginPoolDeleteRequest) SetFailIfReferred(v bool)`

SetFailIfReferred sets FailIfReferred field to given value.

### HasFailIfReferred

`func (o *OriginPoolDeleteRequest) HasFailIfReferred() bool`

HasFailIfReferred returns a boolean if a field has been set.

### GetName

`func (o *OriginPoolDeleteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OriginPoolDeleteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OriginPoolDeleteRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OriginPoolDeleteRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *OriginPoolDeleteRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OriginPoolDeleteRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OriginPoolDeleteRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OriginPoolDeleteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


