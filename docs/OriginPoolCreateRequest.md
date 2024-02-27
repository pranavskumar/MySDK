# OriginPoolCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SchemaObjectCreateMetaType**](SchemaObjectCreateMetaType.md) |  | [optional] 
**Spec** | Pointer to [**ViewsoriginPoolCreateSpecType**](ViewsoriginPoolCreateSpecType.md) |  | [optional] 

## Methods

### NewOriginPoolCreateRequest

`func NewOriginPoolCreateRequest() *OriginPoolCreateRequest`

NewOriginPoolCreateRequest instantiates a new OriginPoolCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolCreateRequestWithDefaults

`func NewOriginPoolCreateRequestWithDefaults() *OriginPoolCreateRequest`

NewOriginPoolCreateRequestWithDefaults instantiates a new OriginPoolCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *OriginPoolCreateRequest) GetMetadata() SchemaObjectCreateMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolCreateRequest) GetMetadataOk() (*SchemaObjectCreateMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolCreateRequest) SetMetadata(v SchemaObjectCreateMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *OriginPoolCreateRequest) GetSpec() ViewsoriginPoolCreateSpecType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OriginPoolCreateRequest) GetSpecOk() (*ViewsoriginPoolCreateSpecType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OriginPoolCreateRequest) SetSpec(v ViewsoriginPoolCreateSpecType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OriginPoolCreateRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


