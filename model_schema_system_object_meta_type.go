/*
F5 Distributed Cloud Services API for ves.io.schema.views.origin_pool

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SchemaSystemObjectMetaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaSystemObjectMetaType{}

// SchemaSystemObjectMetaType SystemObjectMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.
type SchemaSystemObjectMetaType struct {
	//  CreationTimestamp is a timestamp representing the server time when this object was  created. It is not guaranteed to be set in happens-before order across separate operations.  Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	//  A value identifying the class of the user or service which created this configuration object.  Example: ` \"value\"`
	CreatorClass *string `json:"creator_class,omitempty"`
	//  This can used by the creator of the object for later audit for e.g. by storing the  version identifying information of the object so at future it can be determined if  version present at remote end is current or stale.  Example: ` \"value\"`
	CreatorCookie *string `json:"creator_cookie,omitempty"`
	//  A value identifying the exact user or service that created this configuration object  Example: ` \"value\"`
	CreatorId *string `json:"creator_id,omitempty"`
	//  DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This  field is set by the server when a graceful deletion is requested by the user, and is not  directly settable by a client. The resource is expected to be deleted (no longer visible  from resource lists, and not reachable by name) after the time in this field, once the  finalizers list is empty. As long as the finalizers list contains items, deletion is blocked.  Once the deletionTimestamp is set, this value may not be unset or be set further into the  future, although it may be shortened or the resource may be deleted prior to this time.  For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react  by sending a graceful termination signal to the containers in the pod. After that 30 seconds,  the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup,  remove the pod from the API. In the presence of network partitions, this object may still  exist after this timestamp, until an administrator or automated process can determine the  resource is fully terminated.  If not set, graceful deletion of the object has not been requested.   Populated by the system when a graceful deletion is requested.  Read-only.
	DeletionTimestamp *time.Time `json:"deletion_timestamp,omitempty"`
	//  Must be empty before the object is deleted from the registry. Each entry  is an identifier for the responsible component that will remove the entry  from the list. If the deletionTimestamp of the object is non-nil, entries  in this list can only be removed.  Example: ` \"value\"`
	Finalizers []string `json:"finalizers,omitempty"`
	Initializers *SchemaInitializersType `json:"initializers,omitempty"`
	//  Map of string keys and values that can be used to organize and categorize  (scope and select) objects as chosen by the operator or software. Values here can be interpreted  by software(backend or frontend) to enable certain behavior e.g. things marked as soft-deleted(restorable).  Example: ` \"'ves.io/soft-deleted''true'\"`
	Labels map[string]interface{} `json:"labels,omitempty"`
	//  ModificationTimestamp is a timestamp representing the server time when this object was  last modified.
	ModificationTimestamp *time.Time `json:"modification_timestamp,omitempty"`
	//  The namespace this object belongs to. This is populated by the service based on the  metadata.namespace field when an object is created.  Validation Rules:   ves.io.schema.rules.repeated.max_items: 1 
	Namespace []IoschemaObjectRefType `json:"namespace,omitempty"`
	//  Unique index for the object. Some objects need a unique integer index to be allocated  for each object type. This field will be populated for all objects that need it and will  be zero otherwise.  Example: ` \"0\"`
	ObjectIndex *int64 `json:"object_index,omitempty"`
	OwnerView *SchemaViewRefType `json:"owner_view,omitempty"`
	//  This should be set to true If VES/SRE operator wants to suppress an object from being  presented to business-logic of a daemon(e.g. due to bad-form/issue-causing Object).  This is meant only to be used in temporary situations for operational continuity till  a fix is rolled out in business-logic.  Example: ` \"true\"`
	SreDisable *bool `json:"sre_disable,omitempty"`
	//  Tenant to which this configuration object belongs to. The value for this is found from  presented credentials.  Example: ` \"acmecorp\"`
	Tenant *string `json:"tenant,omitempty"`
	//  trace_info holds information(<trace-id>:<span-id>:<parent-span-id>) of the request doing  the object modification. This can be used on the watch side to create subsequent spans.  This information can be used to co-relate activities across services (modulo state compression)  for a synchronous API.  Example: ` \"value\"`
	TraceInfo *string `json:"trace_info,omitempty"`
	//  uid is the unique in time and space value for this object. It is generated by  the server on successful creation of an object and is not allowed to change on Replace  API. The value of is taken from uid field of ObjectMetaType, if provided.  Example: ` \"d15f1fad-4d37-48c0-8706-df1824d76d31\"`
	Uid *string `json:"uid,omitempty"`
	//  Indicate origin of this object.
	VtrpId *string `json:"vtrp_id,omitempty"`
	//  Indicate whether mars deems this object to be stale via graceful restart timer information
	VtrpStale *bool `json:"vtrp_stale,omitempty"`
}

// NewSchemaSystemObjectMetaType instantiates a new SchemaSystemObjectMetaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaSystemObjectMetaType() *SchemaSystemObjectMetaType {
	this := SchemaSystemObjectMetaType{}
	return &this
}

// NewSchemaSystemObjectMetaTypeWithDefaults instantiates a new SchemaSystemObjectMetaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaSystemObjectMetaTypeWithDefaults() *SchemaSystemObjectMetaType {
	this := SchemaSystemObjectMetaType{}
	return &this
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetCreationTimestamp() time.Time {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *SchemaSystemObjectMetaType) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

// GetCreatorClass returns the CreatorClass field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetCreatorClass() string {
	if o == nil || IsNil(o.CreatorClass) {
		var ret string
		return ret
	}
	return *o.CreatorClass
}

// GetCreatorClassOk returns a tuple with the CreatorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetCreatorClassOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorClass) {
		return nil, false
	}
	return o.CreatorClass, true
}

// HasCreatorClass returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasCreatorClass() bool {
	if o != nil && !IsNil(o.CreatorClass) {
		return true
	}

	return false
}

// SetCreatorClass gets a reference to the given string and assigns it to the CreatorClass field.
func (o *SchemaSystemObjectMetaType) SetCreatorClass(v string) {
	o.CreatorClass = &v
}

// GetCreatorCookie returns the CreatorCookie field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetCreatorCookie() string {
	if o == nil || IsNil(o.CreatorCookie) {
		var ret string
		return ret
	}
	return *o.CreatorCookie
}

// GetCreatorCookieOk returns a tuple with the CreatorCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetCreatorCookieOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorCookie) {
		return nil, false
	}
	return o.CreatorCookie, true
}

// HasCreatorCookie returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasCreatorCookie() bool {
	if o != nil && !IsNil(o.CreatorCookie) {
		return true
	}

	return false
}

// SetCreatorCookie gets a reference to the given string and assigns it to the CreatorCookie field.
func (o *SchemaSystemObjectMetaType) SetCreatorCookie(v string) {
	o.CreatorCookie = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SchemaSystemObjectMetaType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDeletionTimestamp returns the DeletionTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetDeletionTimestamp() time.Time {
	if o == nil || IsNil(o.DeletionTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTimestamp
}

// GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetDeletionTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletionTimestamp) {
		return nil, false
	}
	return o.DeletionTimestamp, true
}

// HasDeletionTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasDeletionTimestamp() bool {
	if o != nil && !IsNil(o.DeletionTimestamp) {
		return true
	}

	return false
}

// SetDeletionTimestamp gets a reference to the given time.Time and assigns it to the DeletionTimestamp field.
func (o *SchemaSystemObjectMetaType) SetDeletionTimestamp(v time.Time) {
	o.DeletionTimestamp = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetFinalizers() []string {
	if o == nil || IsNil(o.Finalizers) {
		var ret []string
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetFinalizersOk() ([]string, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []string and assigns it to the Finalizers field.
func (o *SchemaSystemObjectMetaType) SetFinalizers(v []string) {
	o.Finalizers = v
}

// GetInitializers returns the Initializers field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetInitializers() SchemaInitializersType {
	if o == nil || IsNil(o.Initializers) {
		var ret SchemaInitializersType
		return ret
	}
	return *o.Initializers
}

// GetInitializersOk returns a tuple with the Initializers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetInitializersOk() (*SchemaInitializersType, bool) {
	if o == nil || IsNil(o.Initializers) {
		return nil, false
	}
	return o.Initializers, true
}

// HasInitializers returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasInitializers() bool {
	if o != nil && !IsNil(o.Initializers) {
		return true
	}

	return false
}

// SetInitializers gets a reference to the given SchemaInitializersType and assigns it to the Initializers field.
func (o *SchemaSystemObjectMetaType) SetInitializers(v SchemaInitializersType) {
	o.Initializers = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetLabels() map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *SchemaSystemObjectMetaType) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

// GetModificationTimestamp returns the ModificationTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetModificationTimestamp() time.Time {
	if o == nil || IsNil(o.ModificationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ModificationTimestamp
}

// GetModificationTimestampOk returns a tuple with the ModificationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetModificationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationTimestamp) {
		return nil, false
	}
	return o.ModificationTimestamp, true
}

// HasModificationTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasModificationTimestamp() bool {
	if o != nil && !IsNil(o.ModificationTimestamp) {
		return true
	}

	return false
}

// SetModificationTimestamp gets a reference to the given time.Time and assigns it to the ModificationTimestamp field.
func (o *SchemaSystemObjectMetaType) SetModificationTimestamp(v time.Time) {
	o.ModificationTimestamp = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetNamespace() []IoschemaObjectRefType {
	if o == nil || IsNil(o.Namespace) {
		var ret []IoschemaObjectRefType
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetNamespaceOk() ([]IoschemaObjectRefType, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given []IoschemaObjectRefType and assigns it to the Namespace field.
func (o *SchemaSystemObjectMetaType) SetNamespace(v []IoschemaObjectRefType) {
	o.Namespace = v
}

// GetObjectIndex returns the ObjectIndex field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetObjectIndex() int64 {
	if o == nil || IsNil(o.ObjectIndex) {
		var ret int64
		return ret
	}
	return *o.ObjectIndex
}

// GetObjectIndexOk returns a tuple with the ObjectIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetObjectIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectIndex) {
		return nil, false
	}
	return o.ObjectIndex, true
}

// HasObjectIndex returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasObjectIndex() bool {
	if o != nil && !IsNil(o.ObjectIndex) {
		return true
	}

	return false
}

// SetObjectIndex gets a reference to the given int64 and assigns it to the ObjectIndex field.
func (o *SchemaSystemObjectMetaType) SetObjectIndex(v int64) {
	o.ObjectIndex = &v
}

// GetOwnerView returns the OwnerView field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetOwnerView() SchemaViewRefType {
	if o == nil || IsNil(o.OwnerView) {
		var ret SchemaViewRefType
		return ret
	}
	return *o.OwnerView
}

// GetOwnerViewOk returns a tuple with the OwnerView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetOwnerViewOk() (*SchemaViewRefType, bool) {
	if o == nil || IsNil(o.OwnerView) {
		return nil, false
	}
	return o.OwnerView, true
}

// HasOwnerView returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasOwnerView() bool {
	if o != nil && !IsNil(o.OwnerView) {
		return true
	}

	return false
}

// SetOwnerView gets a reference to the given SchemaViewRefType and assigns it to the OwnerView field.
func (o *SchemaSystemObjectMetaType) SetOwnerView(v SchemaViewRefType) {
	o.OwnerView = &v
}

// GetSreDisable returns the SreDisable field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetSreDisable() bool {
	if o == nil || IsNil(o.SreDisable) {
		var ret bool
		return ret
	}
	return *o.SreDisable
}

// GetSreDisableOk returns a tuple with the SreDisable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetSreDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.SreDisable) {
		return nil, false
	}
	return o.SreDisable, true
}

// HasSreDisable returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasSreDisable() bool {
	if o != nil && !IsNil(o.SreDisable) {
		return true
	}

	return false
}

// SetSreDisable gets a reference to the given bool and assigns it to the SreDisable field.
func (o *SchemaSystemObjectMetaType) SetSreDisable(v bool) {
	o.SreDisable = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *SchemaSystemObjectMetaType) SetTenant(v string) {
	o.Tenant = &v
}

// GetTraceInfo returns the TraceInfo field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetTraceInfo() string {
	if o == nil || IsNil(o.TraceInfo) {
		var ret string
		return ret
	}
	return *o.TraceInfo
}

// GetTraceInfoOk returns a tuple with the TraceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetTraceInfoOk() (*string, bool) {
	if o == nil || IsNil(o.TraceInfo) {
		return nil, false
	}
	return o.TraceInfo, true
}

// HasTraceInfo returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasTraceInfo() bool {
	if o != nil && !IsNil(o.TraceInfo) {
		return true
	}

	return false
}

// SetTraceInfo gets a reference to the given string and assigns it to the TraceInfo field.
func (o *SchemaSystemObjectMetaType) SetTraceInfo(v string) {
	o.TraceInfo = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *SchemaSystemObjectMetaType) SetUid(v string) {
	o.Uid = &v
}

// GetVtrpId returns the VtrpId field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetVtrpId() string {
	if o == nil || IsNil(o.VtrpId) {
		var ret string
		return ret
	}
	return *o.VtrpId
}

// GetVtrpIdOk returns a tuple with the VtrpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetVtrpIdOk() (*string, bool) {
	if o == nil || IsNil(o.VtrpId) {
		return nil, false
	}
	return o.VtrpId, true
}

// HasVtrpId returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasVtrpId() bool {
	if o != nil && !IsNil(o.VtrpId) {
		return true
	}

	return false
}

// SetVtrpId gets a reference to the given string and assigns it to the VtrpId field.
func (o *SchemaSystemObjectMetaType) SetVtrpId(v string) {
	o.VtrpId = &v
}

// GetVtrpStale returns the VtrpStale field value if set, zero value otherwise.
func (o *SchemaSystemObjectMetaType) GetVtrpStale() bool {
	if o == nil || IsNil(o.VtrpStale) {
		var ret bool
		return ret
	}
	return *o.VtrpStale
}

// GetVtrpStaleOk returns a tuple with the VtrpStale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectMetaType) GetVtrpStaleOk() (*bool, bool) {
	if o == nil || IsNil(o.VtrpStale) {
		return nil, false
	}
	return o.VtrpStale, true
}

// HasVtrpStale returns a boolean if a field has been set.
func (o *SchemaSystemObjectMetaType) HasVtrpStale() bool {
	if o != nil && !IsNil(o.VtrpStale) {
		return true
	}

	return false
}

// SetVtrpStale gets a reference to the given bool and assigns it to the VtrpStale field.
func (o *SchemaSystemObjectMetaType) SetVtrpStale(v bool) {
	o.VtrpStale = &v
}

func (o SchemaSystemObjectMetaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaSystemObjectMetaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTimestamp) {
		toSerialize["creation_timestamp"] = o.CreationTimestamp
	}
	if !IsNil(o.CreatorClass) {
		toSerialize["creator_class"] = o.CreatorClass
	}
	if !IsNil(o.CreatorCookie) {
		toSerialize["creator_cookie"] = o.CreatorCookie
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.DeletionTimestamp) {
		toSerialize["deletion_timestamp"] = o.DeletionTimestamp
	}
	if !IsNil(o.Finalizers) {
		toSerialize["finalizers"] = o.Finalizers
	}
	if !IsNil(o.Initializers) {
		toSerialize["initializers"] = o.Initializers
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.ModificationTimestamp) {
		toSerialize["modification_timestamp"] = o.ModificationTimestamp
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.ObjectIndex) {
		toSerialize["object_index"] = o.ObjectIndex
	}
	if !IsNil(o.OwnerView) {
		toSerialize["owner_view"] = o.OwnerView
	}
	if !IsNil(o.SreDisable) {
		toSerialize["sre_disable"] = o.SreDisable
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.TraceInfo) {
		toSerialize["trace_info"] = o.TraceInfo
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !IsNil(o.VtrpId) {
		toSerialize["vtrp_id"] = o.VtrpId
	}
	if !IsNil(o.VtrpStale) {
		toSerialize["vtrp_stale"] = o.VtrpStale
	}
	return toSerialize, nil
}

type NullableSchemaSystemObjectMetaType struct {
	value *SchemaSystemObjectMetaType
	isSet bool
}

func (v NullableSchemaSystemObjectMetaType) Get() *SchemaSystemObjectMetaType {
	return v.value
}

func (v *NullableSchemaSystemObjectMetaType) Set(val *SchemaSystemObjectMetaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaSystemObjectMetaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaSystemObjectMetaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaSystemObjectMetaType(val *SchemaSystemObjectMetaType) *NullableSchemaSystemObjectMetaType {
	return &NullableSchemaSystemObjectMetaType{value: val, isSet: true}
}

func (v NullableSchemaSystemObjectMetaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaSystemObjectMetaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


