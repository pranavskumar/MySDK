# ViewsTlsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSecurity** | Pointer to [**ViewsCustomCiphers**](ViewsCustomCiphers.md) |  | [optional] 
**DefaultSecurity** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**LowSecurity** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**MediumSecurity** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 

## Methods

### NewViewsTlsConfig

`func NewViewsTlsConfig() *ViewsTlsConfig`

NewViewsTlsConfig instantiates a new ViewsTlsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsTlsConfigWithDefaults

`func NewViewsTlsConfigWithDefaults() *ViewsTlsConfig`

NewViewsTlsConfigWithDefaults instantiates a new ViewsTlsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSecurity

`func (o *ViewsTlsConfig) GetCustomSecurity() ViewsCustomCiphers`

GetCustomSecurity returns the CustomSecurity field if non-nil, zero value otherwise.

### GetCustomSecurityOk

`func (o *ViewsTlsConfig) GetCustomSecurityOk() (*ViewsCustomCiphers, bool)`

GetCustomSecurityOk returns a tuple with the CustomSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSecurity

`func (o *ViewsTlsConfig) SetCustomSecurity(v ViewsCustomCiphers)`

SetCustomSecurity sets CustomSecurity field to given value.

### HasCustomSecurity

`func (o *ViewsTlsConfig) HasCustomSecurity() bool`

HasCustomSecurity returns a boolean if a field has been set.

### GetDefaultSecurity

`func (o *ViewsTlsConfig) GetDefaultSecurity() map[string]interface{}`

GetDefaultSecurity returns the DefaultSecurity field if non-nil, zero value otherwise.

### GetDefaultSecurityOk

`func (o *ViewsTlsConfig) GetDefaultSecurityOk() (*map[string]interface{}, bool)`

GetDefaultSecurityOk returns a tuple with the DefaultSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurity

`func (o *ViewsTlsConfig) SetDefaultSecurity(v map[string]interface{})`

SetDefaultSecurity sets DefaultSecurity field to given value.

### HasDefaultSecurity

`func (o *ViewsTlsConfig) HasDefaultSecurity() bool`

HasDefaultSecurity returns a boolean if a field has been set.

### GetLowSecurity

`func (o *ViewsTlsConfig) GetLowSecurity() map[string]interface{}`

GetLowSecurity returns the LowSecurity field if non-nil, zero value otherwise.

### GetLowSecurityOk

`func (o *ViewsTlsConfig) GetLowSecurityOk() (*map[string]interface{}, bool)`

GetLowSecurityOk returns a tuple with the LowSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSecurity

`func (o *ViewsTlsConfig) SetLowSecurity(v map[string]interface{})`

SetLowSecurity sets LowSecurity field to given value.

### HasLowSecurity

`func (o *ViewsTlsConfig) HasLowSecurity() bool`

HasLowSecurity returns a boolean if a field has been set.

### GetMediumSecurity

`func (o *ViewsTlsConfig) GetMediumSecurity() map[string]interface{}`

GetMediumSecurity returns the MediumSecurity field if non-nil, zero value otherwise.

### GetMediumSecurityOk

`func (o *ViewsTlsConfig) GetMediumSecurityOk() (*map[string]interface{}, bool)`

GetMediumSecurityOk returns a tuple with the MediumSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumSecurity

`func (o *ViewsTlsConfig) SetMediumSecurity(v map[string]interface{})`

SetMediumSecurity sets MediumSecurity field to given value.

### HasMediumSecurity

`func (o *ViewsTlsConfig) HasMediumSecurity() bool`

HasMediumSecurity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


