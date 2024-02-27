# OriginPoolCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SchemaObjectGetMetaType**](SchemaObjectGetMetaType.md) |  | [optional] 
**Spec** | Pointer to [**ViewsoriginPoolGetSpecType**](ViewsoriginPoolGetSpecType.md) |  | [optional] 
**SystemMetadata** | Pointer to [**SchemaSystemObjectGetMetaType**](SchemaSystemObjectGetMetaType.md) |  | [optional] 

## Methods

### NewOriginPoolCreateResponse

`func NewOriginPoolCreateResponse() *OriginPoolCreateResponse`

NewOriginPoolCreateResponse instantiates a new OriginPoolCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolCreateResponseWithDefaults

`func NewOriginPoolCreateResponseWithDefaults() *OriginPoolCreateResponse`

NewOriginPoolCreateResponseWithDefaults instantiates a new OriginPoolCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *OriginPoolCreateResponse) GetMetadata() SchemaObjectGetMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolCreateResponse) GetMetadataOk() (*SchemaObjectGetMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolCreateResponse) SetMetadata(v SchemaObjectGetMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolCreateResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *OriginPoolCreateResponse) GetSpec() ViewsoriginPoolGetSpecType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OriginPoolCreateResponse) GetSpecOk() (*ViewsoriginPoolGetSpecType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OriginPoolCreateResponse) SetSpec(v ViewsoriginPoolGetSpecType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OriginPoolCreateResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *OriginPoolCreateResponse) GetSystemMetadata() SchemaSystemObjectGetMetaType`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *OriginPoolCreateResponse) GetSystemMetadataOk() (*SchemaSystemObjectGetMetaType, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *OriginPoolCreateResponse) SetSystemMetadata(v SchemaSystemObjectGetMetaType)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *OriginPoolCreateResponse) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


