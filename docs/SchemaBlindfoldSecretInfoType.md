# SchemaBlindfoldSecretInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecryptionProvider** | Pointer to **string** |  Name of the Secret Management Access object that contains information about the backend Secret Management service.  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 
**Location** | Pointer to **string** |  Location is the uri_ref. It could be in url format for string:///  Or it could be a path if the store provider is an http/https location  Example: &#x60; \&quot;string:///U2VjcmV0SW5mb3JtYXRpb24&#x3D;\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.uri_ref: true  | [optional] 
**StoreProvider** | Pointer to **string** |  Name of the Secret Management Access object that contains information about the store to get encrypted bytes  This field needs to be provided only if the url scheme is not string:///  Example: &#x60; \&quot;value\&quot;&#x60; | [optional] 

## Methods

### NewSchemaBlindfoldSecretInfoType

`func NewSchemaBlindfoldSecretInfoType() *SchemaBlindfoldSecretInfoType`

NewSchemaBlindfoldSecretInfoType instantiates a new SchemaBlindfoldSecretInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBlindfoldSecretInfoTypeWithDefaults

`func NewSchemaBlindfoldSecretInfoTypeWithDefaults() *SchemaBlindfoldSecretInfoType`

NewSchemaBlindfoldSecretInfoTypeWithDefaults instantiates a new SchemaBlindfoldSecretInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecryptionProvider

`func (o *SchemaBlindfoldSecretInfoType) GetDecryptionProvider() string`

GetDecryptionProvider returns the DecryptionProvider field if non-nil, zero value otherwise.

### GetDecryptionProviderOk

`func (o *SchemaBlindfoldSecretInfoType) GetDecryptionProviderOk() (*string, bool)`

GetDecryptionProviderOk returns a tuple with the DecryptionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptionProvider

`func (o *SchemaBlindfoldSecretInfoType) SetDecryptionProvider(v string)`

SetDecryptionProvider sets DecryptionProvider field to given value.

### HasDecryptionProvider

`func (o *SchemaBlindfoldSecretInfoType) HasDecryptionProvider() bool`

HasDecryptionProvider returns a boolean if a field has been set.

### GetLocation

`func (o *SchemaBlindfoldSecretInfoType) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SchemaBlindfoldSecretInfoType) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SchemaBlindfoldSecretInfoType) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SchemaBlindfoldSecretInfoType) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStoreProvider

`func (o *SchemaBlindfoldSecretInfoType) GetStoreProvider() string`

GetStoreProvider returns the StoreProvider field if non-nil, zero value otherwise.

### GetStoreProviderOk

`func (o *SchemaBlindfoldSecretInfoType) GetStoreProviderOk() (*string, bool)`

GetStoreProviderOk returns a tuple with the StoreProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreProvider

`func (o *SchemaBlindfoldSecretInfoType) SetStoreProvider(v string)`

SetStoreProvider sets StoreProvider field to given value.

### HasStoreProvider

`func (o *SchemaBlindfoldSecretInfoType) HasStoreProvider() bool`

HasStoreProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


