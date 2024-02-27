# IoschemaObjectRefType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then kind will hold the referred object&#39;s kind (e.g. \&quot;route\&quot;)  Example: &#x60; \&quot;virtual_site\&quot;&#x60; | [optional] 
**Name** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then name will hold the referred object&#39;s(e.g. route&#39;s) name.  Example: &#x60; \&quot;contactus-route\&quot;&#x60; | [optional] 
**Namespace** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then namespace will hold the referred object&#39;s(e.g. route&#39;s) namespace.  Example: &#x60; \&quot;ns1\&quot;&#x60; | [optional] 
**Tenant** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then tenant will hold the referred object&#39;s(e.g. route&#39;s) tenant.  Example: &#x60; \&quot;acmecorp\&quot;&#x60; | [optional] 
**Uid** | Pointer to **string** |  When a configuration object(e.g. virtual_host) refers to another(e.g route)  then uid will hold the referred object&#39;s(e.g. route&#39;s) uid.  Example: &#x60; \&quot;d15f1fad-4d37-48c0-8706-df1824d76d31\&quot;&#x60; | [optional] 

## Methods

### NewIoschemaObjectRefType

`func NewIoschemaObjectRefType() *IoschemaObjectRefType`

NewIoschemaObjectRefType instantiates a new IoschemaObjectRefType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoschemaObjectRefTypeWithDefaults

`func NewIoschemaObjectRefTypeWithDefaults() *IoschemaObjectRefType`

NewIoschemaObjectRefTypeWithDefaults instantiates a new IoschemaObjectRefType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *IoschemaObjectRefType) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoschemaObjectRefType) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoschemaObjectRefType) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoschemaObjectRefType) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *IoschemaObjectRefType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoschemaObjectRefType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoschemaObjectRefType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoschemaObjectRefType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *IoschemaObjectRefType) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoschemaObjectRefType) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoschemaObjectRefType) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoschemaObjectRefType) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *IoschemaObjectRefType) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IoschemaObjectRefType) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IoschemaObjectRefType) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IoschemaObjectRefType) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUid

`func (o *IoschemaObjectRefType) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IoschemaObjectRefType) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IoschemaObjectRefType) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IoschemaObjectRefType) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


