# OriginPoolReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**SchemaObjectReplaceMetaType**](SchemaObjectReplaceMetaType.md) |  | [optional] 
**Spec** | Pointer to [**ViewsoriginPoolReplaceSpecType**](ViewsoriginPoolReplaceSpecType.md) |  | [optional] 

## Methods

### NewOriginPoolReplaceRequest

`func NewOriginPoolReplaceRequest() *OriginPoolReplaceRequest`

NewOriginPoolReplaceRequest instantiates a new OriginPoolReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolReplaceRequestWithDefaults

`func NewOriginPoolReplaceRequestWithDefaults() *OriginPoolReplaceRequest`

NewOriginPoolReplaceRequestWithDefaults instantiates a new OriginPoolReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *OriginPoolReplaceRequest) GetMetadata() SchemaObjectReplaceMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolReplaceRequest) GetMetadataOk() (*SchemaObjectReplaceMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolReplaceRequest) SetMetadata(v SchemaObjectReplaceMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolReplaceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *OriginPoolReplaceRequest) GetSpec() ViewsoriginPoolReplaceSpecType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OriginPoolReplaceRequest) GetSpecOk() (*ViewsoriginPoolReplaceSpecType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OriginPoolReplaceRequest) SetSpec(v ViewsoriginPoolReplaceSpecType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OriginPoolReplaceRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


