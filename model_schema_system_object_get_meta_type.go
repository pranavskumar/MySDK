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

// checks if the SchemaSystemObjectGetMetaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaSystemObjectGetMetaType{}

// SchemaSystemObjectGetMetaType SystemObjectGetMetaType is metadata generated or populated by the system for all persisted objects and cannot be updated directly by users.
type SchemaSystemObjectGetMetaType struct {
	//  CreationTimestamp is a timestamp representing the server time when this object was  created. It is not guaranteed to be set in happens-before order across separate operations.  Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp *time.Time `json:"creation_timestamp,omitempty"`
	//  A value identifying the class of the user or service which created this configuration object.  Example: ` \"value\"`
	CreatorClass *string `json:"creator_class,omitempty"`
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
	//  Unique index for the object. Some objects need a unique integer index to be allocated  for each object type. This field will be populated for all objects that need it and will  be zero otherwise.  Example: ` \"0\"`
	ObjectIndex *int64 `json:"object_index,omitempty"`
	OwnerView *SchemaViewRefType `json:"owner_view,omitempty"`
	//  Tenant to which this configuration object belongs to. The value for this is found from  presented credentials.  Example: ` \"acmecorp\"`
	Tenant *string `json:"tenant,omitempty"`
	//  uid is the unique in time and space value for this object. It is generated by  the server on successful creation of an object and is not allowed to change on Replace  API. The value of is taken from uid field of ObjectMetaType, if provided.  Example: ` \"d15f1fad-4d37-48c0-8706-df1824d76d31\"`
	Uid *string `json:"uid,omitempty"`
}

// NewSchemaSystemObjectGetMetaType instantiates a new SchemaSystemObjectGetMetaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaSystemObjectGetMetaType() *SchemaSystemObjectGetMetaType {
	this := SchemaSystemObjectGetMetaType{}
	return &this
}

// NewSchemaSystemObjectGetMetaTypeWithDefaults instantiates a new SchemaSystemObjectGetMetaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaSystemObjectGetMetaTypeWithDefaults() *SchemaSystemObjectGetMetaType {
	this := SchemaSystemObjectGetMetaType{}
	return &this
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetCreationTimestamp() time.Time {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *SchemaSystemObjectGetMetaType) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

// GetCreatorClass returns the CreatorClass field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetCreatorClass() string {
	if o == nil || IsNil(o.CreatorClass) {
		var ret string
		return ret
	}
	return *o.CreatorClass
}

// GetCreatorClassOk returns a tuple with the CreatorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetCreatorClassOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorClass) {
		return nil, false
	}
	return o.CreatorClass, true
}

// HasCreatorClass returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasCreatorClass() bool {
	if o != nil && !IsNil(o.CreatorClass) {
		return true
	}

	return false
}

// SetCreatorClass gets a reference to the given string and assigns it to the CreatorClass field.
func (o *SchemaSystemObjectGetMetaType) SetCreatorClass(v string) {
	o.CreatorClass = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *SchemaSystemObjectGetMetaType) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDeletionTimestamp returns the DeletionTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetDeletionTimestamp() time.Time {
	if o == nil || IsNil(o.DeletionTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.DeletionTimestamp
}

// GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetDeletionTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletionTimestamp) {
		return nil, false
	}
	return o.DeletionTimestamp, true
}

// HasDeletionTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasDeletionTimestamp() bool {
	if o != nil && !IsNil(o.DeletionTimestamp) {
		return true
	}

	return false
}

// SetDeletionTimestamp gets a reference to the given time.Time and assigns it to the DeletionTimestamp field.
func (o *SchemaSystemObjectGetMetaType) SetDeletionTimestamp(v time.Time) {
	o.DeletionTimestamp = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetFinalizers() []string {
	if o == nil || IsNil(o.Finalizers) {
		var ret []string
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetFinalizersOk() ([]string, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []string and assigns it to the Finalizers field.
func (o *SchemaSystemObjectGetMetaType) SetFinalizers(v []string) {
	o.Finalizers = v
}

// GetInitializers returns the Initializers field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetInitializers() SchemaInitializersType {
	if o == nil || IsNil(o.Initializers) {
		var ret SchemaInitializersType
		return ret
	}
	return *o.Initializers
}

// GetInitializersOk returns a tuple with the Initializers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetInitializersOk() (*SchemaInitializersType, bool) {
	if o == nil || IsNil(o.Initializers) {
		return nil, false
	}
	return o.Initializers, true
}

// HasInitializers returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasInitializers() bool {
	if o != nil && !IsNil(o.Initializers) {
		return true
	}

	return false
}

// SetInitializers gets a reference to the given SchemaInitializersType and assigns it to the Initializers field.
func (o *SchemaSystemObjectGetMetaType) SetInitializers(v SchemaInitializersType) {
	o.Initializers = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetLabels() map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetLabelsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *SchemaSystemObjectGetMetaType) SetLabels(v map[string]interface{}) {
	o.Labels = v
}

// GetModificationTimestamp returns the ModificationTimestamp field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetModificationTimestamp() time.Time {
	if o == nil || IsNil(o.ModificationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ModificationTimestamp
}

// GetModificationTimestampOk returns a tuple with the ModificationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetModificationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationTimestamp) {
		return nil, false
	}
	return o.ModificationTimestamp, true
}

// HasModificationTimestamp returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasModificationTimestamp() bool {
	if o != nil && !IsNil(o.ModificationTimestamp) {
		return true
	}

	return false
}

// SetModificationTimestamp gets a reference to the given time.Time and assigns it to the ModificationTimestamp field.
func (o *SchemaSystemObjectGetMetaType) SetModificationTimestamp(v time.Time) {
	o.ModificationTimestamp = &v
}

// GetObjectIndex returns the ObjectIndex field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetObjectIndex() int64 {
	if o == nil || IsNil(o.ObjectIndex) {
		var ret int64
		return ret
	}
	return *o.ObjectIndex
}

// GetObjectIndexOk returns a tuple with the ObjectIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetObjectIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectIndex) {
		return nil, false
	}
	return o.ObjectIndex, true
}

// HasObjectIndex returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasObjectIndex() bool {
	if o != nil && !IsNil(o.ObjectIndex) {
		return true
	}

	return false
}

// SetObjectIndex gets a reference to the given int64 and assigns it to the ObjectIndex field.
func (o *SchemaSystemObjectGetMetaType) SetObjectIndex(v int64) {
	o.ObjectIndex = &v
}

// GetOwnerView returns the OwnerView field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetOwnerView() SchemaViewRefType {
	if o == nil || IsNil(o.OwnerView) {
		var ret SchemaViewRefType
		return ret
	}
	return *o.OwnerView
}

// GetOwnerViewOk returns a tuple with the OwnerView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetOwnerViewOk() (*SchemaViewRefType, bool) {
	if o == nil || IsNil(o.OwnerView) {
		return nil, false
	}
	return o.OwnerView, true
}

// HasOwnerView returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasOwnerView() bool {
	if o != nil && !IsNil(o.OwnerView) {
		return true
	}

	return false
}

// SetOwnerView gets a reference to the given SchemaViewRefType and assigns it to the OwnerView field.
func (o *SchemaSystemObjectGetMetaType) SetOwnerView(v SchemaViewRefType) {
	o.OwnerView = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *SchemaSystemObjectGetMetaType) SetTenant(v string) {
	o.Tenant = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SchemaSystemObjectGetMetaType) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSystemObjectGetMetaType) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SchemaSystemObjectGetMetaType) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *SchemaSystemObjectGetMetaType) SetUid(v string) {
	o.Uid = &v
}

func (o SchemaSystemObjectGetMetaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaSystemObjectGetMetaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationTimestamp) {
		toSerialize["creation_timestamp"] = o.CreationTimestamp
	}
	if !IsNil(o.CreatorClass) {
		toSerialize["creator_class"] = o.CreatorClass
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
	if !IsNil(o.ObjectIndex) {
		toSerialize["object_index"] = o.ObjectIndex
	}
	if !IsNil(o.OwnerView) {
		toSerialize["owner_view"] = o.OwnerView
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableSchemaSystemObjectGetMetaType struct {
	value *SchemaSystemObjectGetMetaType
	isSet bool
}

func (v NullableSchemaSystemObjectGetMetaType) Get() *SchemaSystemObjectGetMetaType {
	return v.value
}

func (v *NullableSchemaSystemObjectGetMetaType) Set(val *SchemaSystemObjectGetMetaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaSystemObjectGetMetaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaSystemObjectGetMetaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaSystemObjectGetMetaType(val *SchemaSystemObjectGetMetaType) *NullableSchemaSystemObjectGetMetaType {
	return &NullableSchemaSystemObjectGetMetaType{value: val, isSet: true}
}

func (v NullableSchemaSystemObjectGetMetaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaSystemObjectGetMetaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


