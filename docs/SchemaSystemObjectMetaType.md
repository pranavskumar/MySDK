# SchemaSystemObjectMetaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimestamp** | Pointer to **time.Time** |  CreationTimestamp is a timestamp representing the server time when this object was  created. It is not guaranteed to be set in happens-before order across separate operations.  Clients may not set this value. It is represented in RFC3339 form and is in UTC. | [optional] 
**CreatorClass** | Pointer to **string** |  A value identifying the class of the user or service which created this configuration object.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**CreatorCookie** | Pointer to **string** |  This can used by the creator of the object for later audit for e.g. by storing the  version identifying information of the object so at future it can be determined if  version present at remote end is current or stale.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**CreatorId** | Pointer to **string** |  A value identifying the exact user or service that created this configuration object  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**DeletionTimestamp** | Pointer to **time.Time** |  DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This  field is set by the server when a graceful deletion is requested by the user, and is not  directly settable by a client. The resource is expected to be deleted (no longer visible  from resource lists, and not reachable by name) after the time in this field, once the  finalizers list is empty. As long as the finalizers list contains items, deletion is blocked.  Once the deletionTimestamp is set, this value may not be unset or be set further into the  future, although it may be shortened or the resource may be deleted prior to this time.  For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react  by sending a graceful termination signal to the containers in the pod. After that 30 seconds,  the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup,  remove the pod from the API. In the presence of network partitions, this object may still  exist after this timestamp, until an administrator or automated process can determine the  resource is fully terminated.  If not set, graceful deletion of the object has not been requested.   Populated by the system when a graceful deletion is requested.  Read-only. | [optional] 
**Finalizers** | Pointer to **[]string** |  Must be empty before the object is deleted from the registry. Each entry  is an identifier for the responsible component that will remove the entry  from the list. If the deletionTimestamp of the object is non-nil, entries  in this list can only be removed.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**Initializers** | Pointer to [**SchemaInitializersType**](SchemaInitializersType.md) |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  Map of string keys and values that can be used to organize and categorize  (scope and select) objects as chosen by the operator or software. Values here can be interpreted  by software(backend or frontend) to enable certain behavior e.g. things marked as soft-deleted(restorable).  Example: &#x60; \&quot;&#39;ves.io/soft-deleted&#39;&#39;true&#39;\&quot;&#x60; | [optional] 
**ModificationTimestamp** | Pointer to **time.Time** |  ModificationTimestamp is a timestamp representing the server time when this object was  last modified. | [optional] 
**Namespace** | Pointer to [**[]IoschemaObjectRefType**](IoschemaObjectRefType.md) |  The namespace this object belongs to. This is populated by the service based on the  metadata.namespace field when an object is created.  Validation Rules:   ves.io.schema.rules.repeated.max_items: 1  | [optional] 
**ObjectIndex** | Pointer to **int64** |  Unique index for the object. Some objects need a unique integer index to be allocated  for each object type. This field will be populated for all objects that need it and will  be zero otherwise.  Example: &#x60; \&quot;0\&quot;&#x60; | [optional] 
**OwnerView** | Pointer to [**SchemaViewRefType**](SchemaViewRefType.md) |  | [optional] 
**SreDisable** | Pointer to **bool** |  This should be set to true If VES/SRE operator wants to suppress an object from being  presented to business-logic of a daemon(e.g. due to bad-form/issue-causing Object).  This is meant only to be used in temporary situations for operational continuity till  a fix is rolled out in business-logic.  Example: &#x60; \&quot;true\&quot;&#x60; | [optional] 
**Tenant** | Pointer to **string** |  Tenant to which this configuration object belongs to. The value for this is found from  presented credentials.  Example: &#x60; \&quot;acmecorp\&quot;&#x60; | [optional] 
**TraceInfo** | Pointer to **string** |  trace_info holds information(&lt;trace-id&gt;:&lt;span-id&gt;:&lt;parent-span-id&gt;) of the request doing  the object modification. This can be used on the watch side to create subsequent spans.  This information can be used to co-relate activities across services (modulo state compression)  for a synchronous API.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**Uid** | Pointer to **string** |  uid is the unique in time and space value for this object. It is generated by  the server on successful creation of an object and is not allowed to change on Replace  API. The value of is taken from uid field of ObjectMetaType, if provided.  Example: &#x60; \&quot;d15f1fad-4d37-48c0-8706-df1824d76d31\&quot;&#x60; | [optional] 
**VtrpId** | Pointer to **string** |  Indicate origin of this object. | [optional] 
**VtrpStale** | Pointer to **bool** |  Indicate whether mars deems this object to be stale via graceful restart timer information | [optional] 

## Methods

### NewSchemaSystemObjectMetaType

`func NewSchemaSystemObjectMetaType() *SchemaSystemObjectMetaType`

NewSchemaSystemObjectMetaType instantiates a new SchemaSystemObjectMetaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaSystemObjectMetaTypeWithDefaults

`func NewSchemaSystemObjectMetaTypeWithDefaults() *SchemaSystemObjectMetaType`

NewSchemaSystemObjectMetaTypeWithDefaults instantiates a new SchemaSystemObjectMetaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimestamp

`func (o *SchemaSystemObjectMetaType) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *SchemaSystemObjectMetaType) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *SchemaSystemObjectMetaType) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *SchemaSystemObjectMetaType) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetCreatorClass

`func (o *SchemaSystemObjectMetaType) GetCreatorClass() string`

GetCreatorClass returns the CreatorClass field if non-nil, zero value otherwise.

### GetCreatorClassOk

`func (o *SchemaSystemObjectMetaType) GetCreatorClassOk() (*string, bool)`

GetCreatorClassOk returns a tuple with the CreatorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorClass

`func (o *SchemaSystemObjectMetaType) SetCreatorClass(v string)`

SetCreatorClass sets CreatorClass field to given value.

### HasCreatorClass

`func (o *SchemaSystemObjectMetaType) HasCreatorClass() bool`

HasCreatorClass returns a boolean if a field has been set.

### GetCreatorCookie

`func (o *SchemaSystemObjectMetaType) GetCreatorCookie() string`

GetCreatorCookie returns the CreatorCookie field if non-nil, zero value otherwise.

### GetCreatorCookieOk

`func (o *SchemaSystemObjectMetaType) GetCreatorCookieOk() (*string, bool)`

GetCreatorCookieOk returns a tuple with the CreatorCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCookie

`func (o *SchemaSystemObjectMetaType) SetCreatorCookie(v string)`

SetCreatorCookie sets CreatorCookie field to given value.

### HasCreatorCookie

`func (o *SchemaSystemObjectMetaType) HasCreatorCookie() bool`

HasCreatorCookie returns a boolean if a field has been set.

### GetCreatorId

`func (o *SchemaSystemObjectMetaType) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SchemaSystemObjectMetaType) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SchemaSystemObjectMetaType) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *SchemaSystemObjectMetaType) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *SchemaSystemObjectMetaType) GetDeletionTimestamp() time.Time`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *SchemaSystemObjectMetaType) GetDeletionTimestampOk() (*time.Time, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *SchemaSystemObjectMetaType) SetDeletionTimestamp(v time.Time)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *SchemaSystemObjectMetaType) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *SchemaSystemObjectMetaType) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *SchemaSystemObjectMetaType) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *SchemaSystemObjectMetaType) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *SchemaSystemObjectMetaType) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetInitializers

`func (o *SchemaSystemObjectMetaType) GetInitializers() SchemaInitializersType`

GetInitializers returns the Initializers field if non-nil, zero value otherwise.

### GetInitializersOk

`func (o *SchemaSystemObjectMetaType) GetInitializersOk() (*SchemaInitializersType, bool)`

GetInitializersOk returns a tuple with the Initializers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializers

`func (o *SchemaSystemObjectMetaType) SetInitializers(v SchemaInitializersType)`

SetInitializers sets Initializers field to given value.

### HasInitializers

`func (o *SchemaSystemObjectMetaType) HasInitializers() bool`

HasInitializers returns a boolean if a field has been set.

### GetLabels

`func (o *SchemaSystemObjectMetaType) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SchemaSystemObjectMetaType) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SchemaSystemObjectMetaType) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SchemaSystemObjectMetaType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetModificationTimestamp

`func (o *SchemaSystemObjectMetaType) GetModificationTimestamp() time.Time`

GetModificationTimestamp returns the ModificationTimestamp field if non-nil, zero value otherwise.

### GetModificationTimestampOk

`func (o *SchemaSystemObjectMetaType) GetModificationTimestampOk() (*time.Time, bool)`

GetModificationTimestampOk returns a tuple with the ModificationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationTimestamp

`func (o *SchemaSystemObjectMetaType) SetModificationTimestamp(v time.Time)`

SetModificationTimestamp sets ModificationTimestamp field to given value.

### HasModificationTimestamp

`func (o *SchemaSystemObjectMetaType) HasModificationTimestamp() bool`

HasModificationTimestamp returns a boolean if a field has been set.

### GetNamespace

`func (o *SchemaSystemObjectMetaType) GetNamespace() []IoschemaObjectRefType`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SchemaSystemObjectMetaType) GetNamespaceOk() (*[]IoschemaObjectRefType, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SchemaSystemObjectMetaType) SetNamespace(v []IoschemaObjectRefType)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SchemaSystemObjectMetaType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObjectIndex

`func (o *SchemaSystemObjectMetaType) GetObjectIndex() int64`

GetObjectIndex returns the ObjectIndex field if non-nil, zero value otherwise.

### GetObjectIndexOk

`func (o *SchemaSystemObjectMetaType) GetObjectIndexOk() (*int64, bool)`

GetObjectIndexOk returns a tuple with the ObjectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIndex

`func (o *SchemaSystemObjectMetaType) SetObjectIndex(v int64)`

SetObjectIndex sets ObjectIndex field to given value.

### HasObjectIndex

`func (o *SchemaSystemObjectMetaType) HasObjectIndex() bool`

HasObjectIndex returns a boolean if a field has been set.

### GetOwnerView

`func (o *SchemaSystemObjectMetaType) GetOwnerView() SchemaViewRefType`

GetOwnerView returns the OwnerView field if non-nil, zero value otherwise.

### GetOwnerViewOk

`func (o *SchemaSystemObjectMetaType) GetOwnerViewOk() (*SchemaViewRefType, bool)`

GetOwnerViewOk returns a tuple with the OwnerView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerView

`func (o *SchemaSystemObjectMetaType) SetOwnerView(v SchemaViewRefType)`

SetOwnerView sets OwnerView field to given value.

### HasOwnerView

`func (o *SchemaSystemObjectMetaType) HasOwnerView() bool`

HasOwnerView returns a boolean if a field has been set.

### GetSreDisable

`func (o *SchemaSystemObjectMetaType) GetSreDisable() bool`

GetSreDisable returns the SreDisable field if non-nil, zero value otherwise.

### GetSreDisableOk

`func (o *SchemaSystemObjectMetaType) GetSreDisableOk() (*bool, bool)`

GetSreDisableOk returns a tuple with the SreDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSreDisable

`func (o *SchemaSystemObjectMetaType) SetSreDisable(v bool)`

SetSreDisable sets SreDisable field to given value.

### HasSreDisable

`func (o *SchemaSystemObjectMetaType) HasSreDisable() bool`

HasSreDisable returns a boolean if a field has been set.

### GetTenant

`func (o *SchemaSystemObjectMetaType) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SchemaSystemObjectMetaType) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SchemaSystemObjectMetaType) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *SchemaSystemObjectMetaType) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTraceInfo

`func (o *SchemaSystemObjectMetaType) GetTraceInfo() string`

GetTraceInfo returns the TraceInfo field if non-nil, zero value otherwise.

### GetTraceInfoOk

`func (o *SchemaSystemObjectMetaType) GetTraceInfoOk() (*string, bool)`

GetTraceInfoOk returns a tuple with the TraceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceInfo

`func (o *SchemaSystemObjectMetaType) SetTraceInfo(v string)`

SetTraceInfo sets TraceInfo field to given value.

### HasTraceInfo

`func (o *SchemaSystemObjectMetaType) HasTraceInfo() bool`

HasTraceInfo returns a boolean if a field has been set.

### GetUid

`func (o *SchemaSystemObjectMetaType) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SchemaSystemObjectMetaType) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SchemaSystemObjectMetaType) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SchemaSystemObjectMetaType) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVtrpId

`func (o *SchemaSystemObjectMetaType) GetVtrpId() string`

GetVtrpId returns the VtrpId field if non-nil, zero value otherwise.

### GetVtrpIdOk

`func (o *SchemaSystemObjectMetaType) GetVtrpIdOk() (*string, bool)`

GetVtrpIdOk returns a tuple with the VtrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtrpId

`func (o *SchemaSystemObjectMetaType) SetVtrpId(v string)`

SetVtrpId sets VtrpId field to given value.

### HasVtrpId

`func (o *SchemaSystemObjectMetaType) HasVtrpId() bool`

HasVtrpId returns a boolean if a field has been set.

### GetVtrpStale

`func (o *SchemaSystemObjectMetaType) GetVtrpStale() bool`

GetVtrpStale returns the VtrpStale field if non-nil, zero value otherwise.

### GetVtrpStaleOk

`func (o *SchemaSystemObjectMetaType) GetVtrpStaleOk() (*bool, bool)`

GetVtrpStaleOk returns a tuple with the VtrpStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtrpStale

`func (o *SchemaSystemObjectMetaType) SetVtrpStale(v bool)`

SetVtrpStale sets VtrpStale field to given value.

### HasVtrpStale

`func (o *SchemaSystemObjectMetaType) HasVtrpStale() bool`

HasVtrpStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


