# SchemaObjectCreateMetaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]interface{}** |  Annotations is an unstructured key value map stored with a resource that may be  set by external tools to store and retrieve arbitrary metadata. They are not  queryable and should be preserved when modifying objects.  Example: &#x60; \&quot;value\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.map.keys.string.max_len: 64   ves.io.schema.rules.map.keys.string.min_len: 1   ves.io.schema.rules.map.values.string.max_len: 1024   ves.io.schema.rules.map.values.string.min_len: 1  | [optional] 
**Description** | Pointer to **string** |  Human readable description for the object  Example: &#x60; \&quot;Virtual Host for acmecorp website\&quot;&#x60; | [optional] 
**Disable** | Pointer to **bool** |  A value of true will administratively disable the object  Example: &#x60; \&quot;true\&quot;&#x60; | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  Map of string keys and values that can be used to organize and categorize  (scope and select) objects as chosen by the user. Values specified here will be used  by selector expression  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**Name** | Pointer to **string** |  This is the name of configuration object. It has to be unique within the namespace.  It can only be specified during create API and cannot be changed during replace API.  The value of name has to follow DNS-1035 format.  Example: &#x60; \&quot;acmecorp-web\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true  | [optional] 
**Namespace** | Pointer to **string** |  This defines the workspace within which each the configuration object is to be created.  Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \&quot;\&quot;  Example: &#x60; \&quot;staging\&quot;&#x60; | [optional] 

## Methods

### NewSchemaObjectCreateMetaType

`func NewSchemaObjectCreateMetaType() *SchemaObjectCreateMetaType`

NewSchemaObjectCreateMetaType instantiates a new SchemaObjectCreateMetaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaObjectCreateMetaTypeWithDefaults

`func NewSchemaObjectCreateMetaTypeWithDefaults() *SchemaObjectCreateMetaType`

NewSchemaObjectCreateMetaTypeWithDefaults instantiates a new SchemaObjectCreateMetaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *SchemaObjectCreateMetaType) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *SchemaObjectCreateMetaType) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *SchemaObjectCreateMetaType) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *SchemaObjectCreateMetaType) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaObjectCreateMetaType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaObjectCreateMetaType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaObjectCreateMetaType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaObjectCreateMetaType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisable

`func (o *SchemaObjectCreateMetaType) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *SchemaObjectCreateMetaType) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *SchemaObjectCreateMetaType) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *SchemaObjectCreateMetaType) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetLabels

`func (o *SchemaObjectCreateMetaType) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SchemaObjectCreateMetaType) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SchemaObjectCreateMetaType) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SchemaObjectCreateMetaType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *SchemaObjectCreateMetaType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaObjectCreateMetaType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaObjectCreateMetaType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaObjectCreateMetaType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SchemaObjectCreateMetaType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SchemaObjectCreateMetaType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SchemaObjectCreateMetaType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SchemaObjectCreateMetaType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


