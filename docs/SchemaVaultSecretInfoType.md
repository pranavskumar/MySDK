# SchemaVaultSecretInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | x-displayName: \&quot;Key\&quot; x-example: \&quot;key_pem\&quot; Key of the individual secret. Vault Secrets are stored as key-value pair. If user is only interested in one value from the map, this field should be set to the corresponding key. If not provided entire secret will be returned. | [optional] 
**Location** | Pointer to **string** | x-displayName: \&quot;Location\&quot; x-required x-example: \&quot;v1/data/vhost_key\&quot; Path to secret in Vault. | [optional] 
**Provider** | Pointer to **string** | x-displayName: \&quot;Provider\&quot; x-required x-example: \&quot;vault-vh-provider\&quot; Name of the Secret Management Access object that contains information about the backend Vault. | [optional] 
**SecretEncoding** | Pointer to [**SchemaSecretEncodingType**](SchemaSecretEncodingType.md) |  | [optional] [default to ENCODING_NONE]
**Version** | Pointer to **int64** | x-displayName: \&quot;Version\&quot; x-example: \&quot;1\&quot; Version of the secret to be fetched. As vault secrets are versioned, user can specify this field to fetch specific version. If not provided latest version will be returned. | [optional] 

## Methods

### NewSchemaVaultSecretInfoType

`func NewSchemaVaultSecretInfoType() *SchemaVaultSecretInfoType`

NewSchemaVaultSecretInfoType instantiates a new SchemaVaultSecretInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaVaultSecretInfoTypeWithDefaults

`func NewSchemaVaultSecretInfoTypeWithDefaults() *SchemaVaultSecretInfoType`

NewSchemaVaultSecretInfoTypeWithDefaults instantiates a new SchemaVaultSecretInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SchemaVaultSecretInfoType) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaVaultSecretInfoType) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaVaultSecretInfoType) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaVaultSecretInfoType) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocation

`func (o *SchemaVaultSecretInfoType) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SchemaVaultSecretInfoType) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SchemaVaultSecretInfoType) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SchemaVaultSecretInfoType) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProvider

`func (o *SchemaVaultSecretInfoType) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SchemaVaultSecretInfoType) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SchemaVaultSecretInfoType) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SchemaVaultSecretInfoType) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSecretEncoding

`func (o *SchemaVaultSecretInfoType) GetSecretEncoding() SchemaSecretEncodingType`

GetSecretEncoding returns the SecretEncoding field if non-nil, zero value otherwise.

### GetSecretEncodingOk

`func (o *SchemaVaultSecretInfoType) GetSecretEncodingOk() (*SchemaSecretEncodingType, bool)`

GetSecretEncodingOk returns a tuple with the SecretEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretEncoding

`func (o *SchemaVaultSecretInfoType) SetSecretEncoding(v SchemaSecretEncodingType)`

SetSecretEncoding sets SecretEncoding field to given value.

### HasSecretEncoding

`func (o *SchemaVaultSecretInfoType) HasSecretEncoding() bool`

HasSecretEncoding returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaVaultSecretInfoType) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaVaultSecretInfoType) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaVaultSecretInfoType) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SchemaVaultSecretInfoType) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


