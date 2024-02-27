# SchemaTlsCertificateType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateUrl** | Pointer to **string** |  TLS certificate.  Certificate or certificate chain in PEM format including the PEM headers.  Example: &#x60; \&quot;value\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.certificate_url: true   ves.io.schema.rules.string.max_bytes: 131072   ves.io.schema.rules.string.min_bytes: 1  | [optional] 
**CustomHashAlgorithms** | Pointer to [**SchemaHashAlgorithms**](SchemaHashAlgorithms.md) |  | [optional] 
**Description** | Pointer to **string** |  Description for the certificate  Example: &#x60; \&quot;Certificate used in production environment\&quot;&#x60; | [optional] 
**DisableOcspStapling** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**PrivateKey** | Pointer to [**SchemaSecretType**](SchemaSecretType.md) |  | [optional] 
**UseSystemDefaults** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 

## Methods

### NewSchemaTlsCertificateType

`func NewSchemaTlsCertificateType() *SchemaTlsCertificateType`

NewSchemaTlsCertificateType instantiates a new SchemaTlsCertificateType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTlsCertificateTypeWithDefaults

`func NewSchemaTlsCertificateTypeWithDefaults() *SchemaTlsCertificateType`

NewSchemaTlsCertificateTypeWithDefaults instantiates a new SchemaTlsCertificateType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateUrl

`func (o *SchemaTlsCertificateType) GetCertificateUrl() string`

GetCertificateUrl returns the CertificateUrl field if non-nil, zero value otherwise.

### GetCertificateUrlOk

`func (o *SchemaTlsCertificateType) GetCertificateUrlOk() (*string, bool)`

GetCertificateUrlOk returns a tuple with the CertificateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUrl

`func (o *SchemaTlsCertificateType) SetCertificateUrl(v string)`

SetCertificateUrl sets CertificateUrl field to given value.

### HasCertificateUrl

`func (o *SchemaTlsCertificateType) HasCertificateUrl() bool`

HasCertificateUrl returns a boolean if a field has been set.

### GetCustomHashAlgorithms

`func (o *SchemaTlsCertificateType) GetCustomHashAlgorithms() SchemaHashAlgorithms`

GetCustomHashAlgorithms returns the CustomHashAlgorithms field if non-nil, zero value otherwise.

### GetCustomHashAlgorithmsOk

`func (o *SchemaTlsCertificateType) GetCustomHashAlgorithmsOk() (*SchemaHashAlgorithms, bool)`

GetCustomHashAlgorithmsOk returns a tuple with the CustomHashAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHashAlgorithms

`func (o *SchemaTlsCertificateType) SetCustomHashAlgorithms(v SchemaHashAlgorithms)`

SetCustomHashAlgorithms sets CustomHashAlgorithms field to given value.

### HasCustomHashAlgorithms

`func (o *SchemaTlsCertificateType) HasCustomHashAlgorithms() bool`

HasCustomHashAlgorithms returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaTlsCertificateType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaTlsCertificateType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaTlsCertificateType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaTlsCertificateType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableOcspStapling

`func (o *SchemaTlsCertificateType) GetDisableOcspStapling() map[string]interface{}`

GetDisableOcspStapling returns the DisableOcspStapling field if non-nil, zero value otherwise.

### GetDisableOcspStaplingOk

`func (o *SchemaTlsCertificateType) GetDisableOcspStaplingOk() (*map[string]interface{}, bool)`

GetDisableOcspStaplingOk returns a tuple with the DisableOcspStapling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOcspStapling

`func (o *SchemaTlsCertificateType) SetDisableOcspStapling(v map[string]interface{})`

SetDisableOcspStapling sets DisableOcspStapling field to given value.

### HasDisableOcspStapling

`func (o *SchemaTlsCertificateType) HasDisableOcspStapling() bool`

HasDisableOcspStapling returns a boolean if a field has been set.

### GetPrivateKey

`func (o *SchemaTlsCertificateType) GetPrivateKey() SchemaSecretType`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SchemaTlsCertificateType) GetPrivateKeyOk() (*SchemaSecretType, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SchemaTlsCertificateType) SetPrivateKey(v SchemaSecretType)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *SchemaTlsCertificateType) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetUseSystemDefaults

`func (o *SchemaTlsCertificateType) GetUseSystemDefaults() map[string]interface{}`

GetUseSystemDefaults returns the UseSystemDefaults field if non-nil, zero value otherwise.

### GetUseSystemDefaultsOk

`func (o *SchemaTlsCertificateType) GetUseSystemDefaultsOk() (*map[string]interface{}, bool)`

GetUseSystemDefaultsOk returns a tuple with the UseSystemDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSystemDefaults

`func (o *SchemaTlsCertificateType) SetUseSystemDefaults(v map[string]interface{})`

SetUseSystemDefaults sets UseSystemDefaults field to given value.

### HasUseSystemDefaults

`func (o *SchemaTlsCertificateType) HasUseSystemDefaults() bool`

HasUseSystemDefaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


