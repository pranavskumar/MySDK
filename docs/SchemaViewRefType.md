# SchemaViewRefType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  Kind of the view object  Example: &#x60; \&quot;http_proxy\&quot;&#x60; | [optional] 
**Name** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then name will hold the referred object&#39;s(e.g. route&#39;s) name.  Example: &#x60; \&quot;contactus-route\&quot;&#x60; | [optional] 
**Namespace** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then namespace will hold the referred object&#39;s(e.g. route&#39;s) namespace.  Example: &#x60; \&quot;ns1\&quot;&#x60; | [optional] 
**Uid** | Pointer to **string** |  UID of the view object  Example: &#x60; \&quot;f3744323-1adf-4aaa-a5dc-0707c1d1bd82\&quot;&#x60; | [optional] 

## Methods

### NewSchemaViewRefType

`func NewSchemaViewRefType() *SchemaViewRefType`

NewSchemaViewRefType instantiates a new SchemaViewRefType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaViewRefTypeWithDefaults

`func NewSchemaViewRefTypeWithDefaults() *SchemaViewRefType`

NewSchemaViewRefTypeWithDefaults instantiates a new SchemaViewRefType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SchemaViewRefType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SchemaViewRefType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SchemaViewRefType) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SchemaViewRefType) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *SchemaViewRefType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaViewRefType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaViewRefType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaViewRefType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SchemaViewRefType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SchemaViewRefType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SchemaViewRefType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SchemaViewRefType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUid

`func (o *SchemaViewRefType) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SchemaViewRefType) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SchemaViewRefType) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SchemaViewRefType) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


