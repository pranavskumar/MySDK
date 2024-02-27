# SchemaviewsObjectRefType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then name will hold the referred object&#39;s(e.g. route&#39;s) name.  Example: &#x60; \&quot;contacts-route\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.max_bytes: 128   ves.io.schema.rules.string.min_bytes: 1  | [optional] 
**Namespace** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then namespace will hold the referred object&#39;s(e.g. route&#39;s) namespace.  Example: &#x60; \&quot;ns1\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.string.max_bytes: 64  | [optional] 
**Tenant** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then tenant will hold the referred object&#39;s(e.g. route&#39;s) tenant.  Example: &#x60; \&quot;acmecorp\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.string.max_bytes: 64  | [optional] 

## Methods

### NewSchemaviewsObjectRefType

`func NewSchemaviewsObjectRefType() *SchemaviewsObjectRefType`

NewSchemaviewsObjectRefType instantiates a new SchemaviewsObjectRefType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaviewsObjectRefTypeWithDefaults

`func NewSchemaviewsObjectRefTypeWithDefaults() *SchemaviewsObjectRefType`

NewSchemaviewsObjectRefTypeWithDefaults instantiates a new SchemaviewsObjectRefType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchemaviewsObjectRefType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaviewsObjectRefType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaviewsObjectRefType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaviewsObjectRefType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SchemaviewsObjectRefType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SchemaviewsObjectRefType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SchemaviewsObjectRefType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SchemaviewsObjectRefType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *SchemaviewsObjectRefType) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SchemaviewsObjectRefType) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SchemaviewsObjectRefType) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *SchemaviewsObjectRefType) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


