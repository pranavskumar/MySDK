# OriginPoolListResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]interface{}** |  The set of annotations present on this origin_pool | [optional] 
**Description** | Pointer to **string** |  The description set for this origin_pool | [optional] 
**Disabled** | Pointer to **bool** |  A value of true indicates origin_pool is administratively disabled | [optional] 
**GetSpec** | Pointer to [**ViewsoriginPoolGetSpecType**](ViewsoriginPoolGetSpecType.md) |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  The set of labels present on this origin_pool | [optional] 
**Metadata** | Pointer to [**SchemaObjectGetMetaType**](SchemaObjectGetMetaType.md) |  | [optional] 
**Name** | Pointer to **string** |  The name of this origin_pool  Example: &#x60; \&quot;name\&quot;&#x60; | [optional] 
**Namespace** | Pointer to **string** |  The namespace this item belongs to  Example: &#x60; \&quot;ns1\&quot;&#x60; | [optional] 
**Object** | Pointer to [**OriginPoolObject**](OriginPoolObject.md) |  | [optional] 
**OwnerView** | Pointer to [**SchemaViewRefType**](SchemaViewRefType.md) |  | [optional] 
**SystemMetadata** | Pointer to [**SchemaSystemObjectGetMetaType**](SchemaSystemObjectGetMetaType.md) |  | [optional] 
**Tenant** | Pointer to **string** |  The tenant this item belongs to  Example: &#x60; \&quot;acmecorp\&quot;&#x60; | [optional] 
**Uid** | Pointer to **string** |  The unique uid of this origin_pool  Example: &#x60; \&quot;d27938ba-967e-40a7-9709-57b8627f9f75\&quot;&#x60; | [optional] 

## Methods

### NewOriginPoolListResponseItem

`func NewOriginPoolListResponseItem() *OriginPoolListResponseItem`

NewOriginPoolListResponseItem instantiates a new OriginPoolListResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolListResponseItemWithDefaults

`func NewOriginPoolListResponseItemWithDefaults() *OriginPoolListResponseItem`

NewOriginPoolListResponseItemWithDefaults instantiates a new OriginPoolListResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *OriginPoolListResponseItem) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *OriginPoolListResponseItem) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *OriginPoolListResponseItem) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *OriginPoolListResponseItem) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDescription

`func (o *OriginPoolListResponseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OriginPoolListResponseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OriginPoolListResponseItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OriginPoolListResponseItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *OriginPoolListResponseItem) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *OriginPoolListResponseItem) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *OriginPoolListResponseItem) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *OriginPoolListResponseItem) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGetSpec

`func (o *OriginPoolListResponseItem) GetGetSpec() ViewsoriginPoolGetSpecType`

GetGetSpec returns the GetSpec field if non-nil, zero value otherwise.

### GetGetSpecOk

`func (o *OriginPoolListResponseItem) GetGetSpecOk() (*ViewsoriginPoolGetSpecType, bool)`

GetGetSpecOk returns a tuple with the GetSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetSpec

`func (o *OriginPoolListResponseItem) SetGetSpec(v ViewsoriginPoolGetSpecType)`

SetGetSpec sets GetSpec field to given value.

### HasGetSpec

`func (o *OriginPoolListResponseItem) HasGetSpec() bool`

HasGetSpec returns a boolean if a field has been set.

### GetLabels

`func (o *OriginPoolListResponseItem) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OriginPoolListResponseItem) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OriginPoolListResponseItem) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OriginPoolListResponseItem) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMetadata

`func (o *OriginPoolListResponseItem) GetMetadata() SchemaObjectGetMetaType`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OriginPoolListResponseItem) GetMetadataOk() (*SchemaObjectGetMetaType, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OriginPoolListResponseItem) SetMetadata(v SchemaObjectGetMetaType)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OriginPoolListResponseItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OriginPoolListResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OriginPoolListResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OriginPoolListResponseItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OriginPoolListResponseItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *OriginPoolListResponseItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OriginPoolListResponseItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OriginPoolListResponseItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OriginPoolListResponseItem) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObject

`func (o *OriginPoolListResponseItem) GetObject() OriginPoolObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OriginPoolListResponseItem) GetObjectOk() (*OriginPoolObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OriginPoolListResponseItem) SetObject(v OriginPoolObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *OriginPoolListResponseItem) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOwnerView

`func (o *OriginPoolListResponseItem) GetOwnerView() SchemaViewRefType`

GetOwnerView returns the OwnerView field if non-nil, zero value otherwise.

### GetOwnerViewOk

`func (o *OriginPoolListResponseItem) GetOwnerViewOk() (*SchemaViewRefType, bool)`

GetOwnerViewOk returns a tuple with the OwnerView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerView

`func (o *OriginPoolListResponseItem) SetOwnerView(v SchemaViewRefType)`

SetOwnerView sets OwnerView field to given value.

### HasOwnerView

`func (o *OriginPoolListResponseItem) HasOwnerView() bool`

HasOwnerView returns a boolean if a field has been set.

### GetSystemMetadata

`func (o *OriginPoolListResponseItem) GetSystemMetadata() SchemaSystemObjectGetMetaType`

GetSystemMetadata returns the SystemMetadata field if non-nil, zero value otherwise.

### GetSystemMetadataOk

`func (o *OriginPoolListResponseItem) GetSystemMetadataOk() (*SchemaSystemObjectGetMetaType, bool)`

GetSystemMetadataOk returns a tuple with the SystemMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMetadata

`func (o *OriginPoolListResponseItem) SetSystemMetadata(v SchemaSystemObjectGetMetaType)`

SetSystemMetadata sets SystemMetadata field to given value.

### HasSystemMetadata

`func (o *OriginPoolListResponseItem) HasSystemMetadata() bool`

HasSystemMetadata returns a boolean if a field has been set.

### GetTenant

`func (o *OriginPoolListResponseItem) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *OriginPoolListResponseItem) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *OriginPoolListResponseItem) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *OriginPoolListResponseItem) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUid

`func (o *OriginPoolListResponseItem) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OriginPoolListResponseItem) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OriginPoolListResponseItem) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *OriginPoolListResponseItem) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


