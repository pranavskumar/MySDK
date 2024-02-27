# OriginPoolUpstreamTlsParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableSni** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**NoMtls** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**SkipServerVerification** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**Sni** | Pointer to **string** | Exclusive with [disable_sni use_host_header_as_sni]  SNI value to be used.  Validation Rules:   ves.io.schema.rules.string.hostname: true   ves.io.schema.rules.string.max_len: 256  | [optional] 
**TlsConfig** | Pointer to [**ViewsTlsConfig**](ViewsTlsConfig.md) |  | [optional] 
**UseHostHeaderAsSni** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**UseMtls** | Pointer to [**OriginPoolTlsCertificatesType**](OriginPoolTlsCertificatesType.md) |  | [optional] 
**UseServerVerification** | Pointer to [**OriginPoolUpstreamTlsValidationContext**](OriginPoolUpstreamTlsValidationContext.md) |  | [optional] 
**VolterraTrustedCa** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 

## Methods

### NewOriginPoolUpstreamTlsParameters

`func NewOriginPoolUpstreamTlsParameters() *OriginPoolUpstreamTlsParameters`

NewOriginPoolUpstreamTlsParameters instantiates a new OriginPoolUpstreamTlsParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolUpstreamTlsParametersWithDefaults

`func NewOriginPoolUpstreamTlsParametersWithDefaults() *OriginPoolUpstreamTlsParameters`

NewOriginPoolUpstreamTlsParametersWithDefaults instantiates a new OriginPoolUpstreamTlsParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableSni

`func (o *OriginPoolUpstreamTlsParameters) GetDisableSni() map[string]interface{}`

GetDisableSni returns the DisableSni field if non-nil, zero value otherwise.

### GetDisableSniOk

`func (o *OriginPoolUpstreamTlsParameters) GetDisableSniOk() (*map[string]interface{}, bool)`

GetDisableSniOk returns a tuple with the DisableSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSni

`func (o *OriginPoolUpstreamTlsParameters) SetDisableSni(v map[string]interface{})`

SetDisableSni sets DisableSni field to given value.

### HasDisableSni

`func (o *OriginPoolUpstreamTlsParameters) HasDisableSni() bool`

HasDisableSni returns a boolean if a field has been set.

### GetNoMtls

`func (o *OriginPoolUpstreamTlsParameters) GetNoMtls() map[string]interface{}`

GetNoMtls returns the NoMtls field if non-nil, zero value otherwise.

### GetNoMtlsOk

`func (o *OriginPoolUpstreamTlsParameters) GetNoMtlsOk() (*map[string]interface{}, bool)`

GetNoMtlsOk returns a tuple with the NoMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMtls

`func (o *OriginPoolUpstreamTlsParameters) SetNoMtls(v map[string]interface{})`

SetNoMtls sets NoMtls field to given value.

### HasNoMtls

`func (o *OriginPoolUpstreamTlsParameters) HasNoMtls() bool`

HasNoMtls returns a boolean if a field has been set.

### GetSkipServerVerification

`func (o *OriginPoolUpstreamTlsParameters) GetSkipServerVerification() map[string]interface{}`

GetSkipServerVerification returns the SkipServerVerification field if non-nil, zero value otherwise.

### GetSkipServerVerificationOk

`func (o *OriginPoolUpstreamTlsParameters) GetSkipServerVerificationOk() (*map[string]interface{}, bool)`

GetSkipServerVerificationOk returns a tuple with the SkipServerVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerVerification

`func (o *OriginPoolUpstreamTlsParameters) SetSkipServerVerification(v map[string]interface{})`

SetSkipServerVerification sets SkipServerVerification field to given value.

### HasSkipServerVerification

`func (o *OriginPoolUpstreamTlsParameters) HasSkipServerVerification() bool`

HasSkipServerVerification returns a boolean if a field has been set.

### GetSni

`func (o *OriginPoolUpstreamTlsParameters) GetSni() string`

GetSni returns the Sni field if non-nil, zero value otherwise.

### GetSniOk

`func (o *OriginPoolUpstreamTlsParameters) GetSniOk() (*string, bool)`

GetSniOk returns a tuple with the Sni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSni

`func (o *OriginPoolUpstreamTlsParameters) SetSni(v string)`

SetSni sets Sni field to given value.

### HasSni

`func (o *OriginPoolUpstreamTlsParameters) HasSni() bool`

HasSni returns a boolean if a field has been set.

### GetTlsConfig

`func (o *OriginPoolUpstreamTlsParameters) GetTlsConfig() ViewsTlsConfig`

GetTlsConfig returns the TlsConfig field if non-nil, zero value otherwise.

### GetTlsConfigOk

`func (o *OriginPoolUpstreamTlsParameters) GetTlsConfigOk() (*ViewsTlsConfig, bool)`

GetTlsConfigOk returns a tuple with the TlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsConfig

`func (o *OriginPoolUpstreamTlsParameters) SetTlsConfig(v ViewsTlsConfig)`

SetTlsConfig sets TlsConfig field to given value.

### HasTlsConfig

`func (o *OriginPoolUpstreamTlsParameters) HasTlsConfig() bool`

HasTlsConfig returns a boolean if a field has been set.

### GetUseHostHeaderAsSni

`func (o *OriginPoolUpstreamTlsParameters) GetUseHostHeaderAsSni() map[string]interface{}`

GetUseHostHeaderAsSni returns the UseHostHeaderAsSni field if non-nil, zero value otherwise.

### GetUseHostHeaderAsSniOk

`func (o *OriginPoolUpstreamTlsParameters) GetUseHostHeaderAsSniOk() (*map[string]interface{}, bool)`

GetUseHostHeaderAsSniOk returns a tuple with the UseHostHeaderAsSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostHeaderAsSni

`func (o *OriginPoolUpstreamTlsParameters) SetUseHostHeaderAsSni(v map[string]interface{})`

SetUseHostHeaderAsSni sets UseHostHeaderAsSni field to given value.

### HasUseHostHeaderAsSni

`func (o *OriginPoolUpstreamTlsParameters) HasUseHostHeaderAsSni() bool`

HasUseHostHeaderAsSni returns a boolean if a field has been set.

### GetUseMtls

`func (o *OriginPoolUpstreamTlsParameters) GetUseMtls() OriginPoolTlsCertificatesType`

GetUseMtls returns the UseMtls field if non-nil, zero value otherwise.

### GetUseMtlsOk

`func (o *OriginPoolUpstreamTlsParameters) GetUseMtlsOk() (*OriginPoolTlsCertificatesType, bool)`

GetUseMtlsOk returns a tuple with the UseMtls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMtls

`func (o *OriginPoolUpstreamTlsParameters) SetUseMtls(v OriginPoolTlsCertificatesType)`

SetUseMtls sets UseMtls field to given value.

### HasUseMtls

`func (o *OriginPoolUpstreamTlsParameters) HasUseMtls() bool`

HasUseMtls returns a boolean if a field has been set.

### GetUseServerVerification

`func (o *OriginPoolUpstreamTlsParameters) GetUseServerVerification() OriginPoolUpstreamTlsValidationContext`

GetUseServerVerification returns the UseServerVerification field if non-nil, zero value otherwise.

### GetUseServerVerificationOk

`func (o *OriginPoolUpstreamTlsParameters) GetUseServerVerificationOk() (*OriginPoolUpstreamTlsValidationContext, bool)`

GetUseServerVerificationOk returns a tuple with the UseServerVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseServerVerification

`func (o *OriginPoolUpstreamTlsParameters) SetUseServerVerification(v OriginPoolUpstreamTlsValidationContext)`

SetUseServerVerification sets UseServerVerification field to given value.

### HasUseServerVerification

`func (o *OriginPoolUpstreamTlsParameters) HasUseServerVerification() bool`

HasUseServerVerification returns a boolean if a field has been set.

### GetVolterraTrustedCa

`func (o *OriginPoolUpstreamTlsParameters) GetVolterraTrustedCa() map[string]interface{}`

GetVolterraTrustedCa returns the VolterraTrustedCa field if non-nil, zero value otherwise.

### GetVolterraTrustedCaOk

`func (o *OriginPoolUpstreamTlsParameters) GetVolterraTrustedCaOk() (*map[string]interface{}, bool)`

GetVolterraTrustedCaOk returns a tuple with the VolterraTrustedCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolterraTrustedCa

`func (o *OriginPoolUpstreamTlsParameters) SetVolterraTrustedCa(v map[string]interface{})`

SetVolterraTrustedCa sets VolterraTrustedCa field to given value.

### HasVolterraTrustedCa

`func (o *OriginPoolUpstreamTlsParameters) HasVolterraTrustedCa() bool`

HasVolterraTrustedCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


