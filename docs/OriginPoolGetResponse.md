# OriginPoolGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateForm** | Pointer to [**OriginPoolCreateRequest**](OriginPoolCreateRequest.md) |  | [optional] 
**DeletedReferredObjects** | Pointer to [**[]IoschemaObjectRefType**](IoschemaObjectRefType.md) | The set of deleted objects that are referred by this object | [optional] 
**DisabledReferredObjects** | Pointer to [**[]IoschemaObjectRefType**](IoschemaObjectRefType.md) | The set of deleted objects that are referred by this object | [optional] 
**Metadata** | Pointer to [**SchemaObjectGetMetaType**](SchemaObjectGetMetaType.md) |  | [optional] 
**Object** | Pointer to [**OriginPoolObject**](OriginPoolObject.md) |  | [optional] 
**ReferringObjects** | Pointer to [**[]IoschemaObjectRefType**](IoschemaObjectRefType.md) | The set of objects that are referring to this object in their spec | [optional] 
**ReplaceForm** | Pointer to [**OriginPoolReplaceRequest**](OriginPoolReplaceRequest.md) |  | [optional] 
**Spec** | Pointer to [**ViewsoriginPoolGetSpecType**](ViewsoriginPoolGetSpecType.md) |  | [optional] 
**SystemMetadata** | Pointer to [**SchemaSystemObjectGetMetaType**](SchemaSystemObjectGetMetaType.md) |  | [optional] 

## Methods

### NewOriginPoolGetResponse

`func NewOriginPoolGetResponse() *OriginPoolGetResponse`

NewOriginPoolGetResponse instantiates a new OriginPoolGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolGetResponseWithDefaults

`func NewOriginPoolGetResponseWithDefaults() *OriginPoolGetResponse`

NewOriginPoolGetResponseWithDefaults instantiates a new OriginPoolGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateForm

`func (o *OriginPoolGetResponse) GetCreateForm() OriginPoolCreateRequest`

GetCreateForm returns the CreateForm field if non-nil, zero value otherwise.

### GetCreateFormOk

`func (o *OriginPoolGetResponse) GetCreateFormOk() (*OriginPoolCreateRequest, bool)`

GetCreateFormOk returns a tuple with the CreateForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateForm

`func (o *OriginPoolGetResponse) SetCreateForm(v OriginPoolCreateRequest)`

SetCreateForm sets CreateForm field to given value.

### HasCreateForm

`func (o *OriginPoolGetResponse) HasCreateForm() bool`

HasCreateForm returns a boolean if a field has been set.

### GetDeletedReferredObjects

`func (o *OriginPoolGetResponse) GetDeletedReferredObjects() []IoschemaObjectRefType`

GetDeletedReferredObjects returns the DeletedReferredObjects field if non-nil, zero value otherwise.

### GetDeletedReferredObjectsOk

`func (o *OriginPoolGetResponse) GetDeletedReferredObjectsOk() (*[]IoschemaObjectRefType, bool)`

GetDeletedReferredObjectsOk returns a tuple with the DeletedReferredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedReferredObjects

`func (o *OriginPoolGetResponse) SetDeletedReferredObjects(v []IoschemaObjectRefType)`

SetDeletedReferredObjects sets DeletedReferredObjects field to given value.

### HasDeletedReferredObjects

`func (o *OriginPoolGetResponse) HasDeletedReferredObjects() bool`

HasDeletedReferredObjects returns a boolean if a field has been set.

### GetDisabledReferredObjects

`func (o *OriginPoolGetResponse) GetDisabledReferredObjects() []IoschemaObjectRefType`

GetDisabledReferredObjects returns the DisabledReferredObjects field if non-nil, zero value otherwise.

### GetDisabledReferredObjectsOk

`func (o *OriginPoolGetResponse) GetDisabledReferredObjectsOk() (*[]IoschemaObjectRefType, bool)`

GetDisabledReferredObjectsOk returns a tuple with the DisabledReferredObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledReferredObjects

`func (o *OriginPoolGetResponse) SetDisabledReferredObjects(v []IoschemaObjectRefType)`

SetDisabledReferredObjects sets DisabledReferredObjects field to given value.

### HasDisabledReferredObjects

`func (o *OriginPoolGetResponse) HasDisabledReferredObjects() bool`

HasDisabledReferredObjects returns a boolean if a field has been set.

### GetMetadata

`func (o *OriginPoolGetResponse) GetMetadata() SchemaObjectGetMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolGetResponse) GetMetadataOk() (*SchemaObjectGetMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolGetResponse) SetMetadata(v SchemaObjectGetMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolGetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetObject

`func (o *OriginPoolGetResponse) GetObject() OriginPoolObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OriginPoolGetResponse) GetObjectOk() (*OriginPoolObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OriginPoolGetResponse) SetObject(v OriginPoolObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *OriginPoolGetResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetReferringObjects

`func (o *OriginPoolGetResponse) GetReferringObjects() []IoschemaObjectRefType`

GetReferringObjects returns the ReferringObjects field if non-nil, zero value otherwise.

### GetReferringObjectsOk

`func (o *OriginPoolGetResponse) GetReferringObjectsOk() (*[]IoschemaObjectRefType, bool)`

GetReferringObjectsOk returns a tuple with the ReferringObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferringObjects

`func (o *OriginPoolGetResponse) SetReferringObjects(v []IoschemaObjectRefType)`

SetReferringObjects sets ReferringObjects field to given value.

### HasReferringObjects

`func (o *OriginPoolGetResponse) HasReferringObjects() bool`

HasReferringObjects returns a boolean if a field has been set.

### GetReplaceForm

`func (o *OriginPoolGetResponse) GetReplaceForm() OriginPoolReplaceRequest`

GetReplaceForm returns the ReplaceForm field if non-nil, zero value otherwise.

### GetReplaceFormOk

`func (o *OriginPoolGetResponse) GetReplaceFormOk() (*OriginPoolReplaceRequest, bool)`

GetReplaceFormOk returns a tuple with the ReplaceForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceForm

`func (o *OriginPoolGetResponse) SetReplaceForm(v OriginPoolReplaceRequest)`

SetReplaceForm sets ReplaceForm field to given value.

### HasReplaceForm

`func (o *OriginPoolGetResponse) HasReplaceForm() bool`

HasReplaceForm returns a boolean if a field has been set.

### GetSpec

`func (o *OriginPoolGetResponse) GetSpec() ViewsoriginPoolGetSpecType`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OriginPoolGetResponse) GetSpecOk() (*ViewsoriginPoolGetSpecType, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OriginPoolGetResponse) SetSpec(v ViewsoriginPoolGetSpecType)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OriginPoolGetResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *OriginPoolGetResponse) GetSystemMetadata() SchemaSystemObjectGetMetaType`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *OriginPoolGetResponse) GetSystemMetadataOk() (*SchemaSystemObjectGetMetaType, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *OriginPoolGetResponse) SetSystemMetadata(v SchemaSystemObjectGetMetaType)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *OriginPoolGetResponse) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


