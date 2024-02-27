# SchemaClearSecretInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  Name of the Secret Management Access object that contains information about the store to get encrypted bytes  This field needs to be provided only if the url scheme is not string:///  Example: &#x60; \&quot;box-provider\&quot;&#x60; | [optional] 
**Url** | Pointer to **string** |  URL of the secret. Currently supported URL schemes is string:///.  For string:/// scheme, Secret needs to be encoded Base64 format.  When asked for this secret, caller will get Secret bytes after Base64 decoding.  Example: &#x60; \&quot;string:///U2VjcmV0SW5mb3JtYXRpb24&#x3D;\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.max_bytes: 131072   ves.io.schema.rules.string.uri_ref: true  | [optional] 

## Methods

### NewSchemaClearSecretInfoType

`func NewSchemaClearSecretInfoType() *SchemaClearSecretInfoType`

NewSchemaClearSecretInfoType instantiates a new SchemaClearSecretInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaClearSecretInfoTypeWithDefaults

`func NewSchemaClearSecretInfoTypeWithDefaults() *SchemaClearSecretInfoType`

NewSchemaClearSecretInfoTypeWithDefaults instantiates a new SchemaClearSecretInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *SchemaClearSecretInfoType) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SchemaClearSecretInfoType) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SchemaClearSecretInfoType) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SchemaClearSecretInfoType) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUrl

`func (o *SchemaClearSecretInfoType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SchemaClearSecretInfoType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SchemaClearSecretInfoType) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SchemaClearSecretInfoType) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


