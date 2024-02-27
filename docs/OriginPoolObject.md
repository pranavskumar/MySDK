# OriginPoolObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SchemaObjectMetaType**](SchemaObjectMetaType.md) |  | [optional] 
**Spec** | Pointer to [**OriginPoolSpecType**](OriginPoolSpecType.md) |  | [optional] 
**SystemMetadata** | Pointer to [**SchemaSystemObjectMetaType**](SchemaSystemObjectMetaType.md) |  | [optional] 

## Methods

### NewOriginPoolObject

`func NewOriginPoolObject() *OriginPoolObject`

NewOriginPoolObject instantiates a new OriginPoolObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolObjectWithDefaults

`func NewOriginPoolObjectWithDefaults() *OriginPoolObject`

NewOriginPoolObjectWithDefaults instantiates a new OriginPoolObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *OriginPoolObject) GetMetadata() SchemaObjectMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolObject) GetMetadataOk() (*SchemaObjectMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolObject) SetMetadata(v SchemaObjectMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *OriginPoolObject) GetSpec() OriginPoolSpecType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OriginPoolObject) GetSpecOk() (*OriginPoolSpecType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OriginPoolObject) SetSpec(v OriginPoolSpecType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OriginPoolObject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *OriginPoolObject) GetSystemMetadata() SchemaSystemObjectMetaType`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *OriginPoolObject) GetSystemMetadataOk() (*SchemaSystemObjectMetaType, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *OriginPoolObject) SetSystemMetadata(v SchemaSystemObjectMetaType)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *OriginPoolObject) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


