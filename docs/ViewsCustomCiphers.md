# ViewsCustomCiphers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherSuites** | Pointer to **[]string** |  The TLS listener will only support the specified cipher list.  Example: &#x60; \&quot;TLS_AES_128_GCM_SHA256\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.items.string.in: [\\\&quot;TLS_AES_128_GCM_SHA256\\\&quot;,\\\&quot;TLS_AES_256_GCM_SHA384\\\&quot;,\\\&quot;TLS_CHACHA20_POLY1305_SHA256\\\&quot;,\\\&quot;TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\&quot;,\\\&quot;TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\&quot;,\\\&quot;TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\&quot;,\\\&quot;TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\&quot;,\\\&quot;TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\&quot;,\\\&quot;TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\&quot;,\\\&quot;TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\&quot;,\\\&quot;TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\&quot;,\\\&quot;TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\&quot;,\\\&quot;TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\&quot;,\\\&quot;TLS_RSA_WITH_AES_128_CBC_SHA\\\&quot;,\\\&quot;TLS_RSA_WITH_AES_128_GCM_SHA256\\\&quot;,\\\&quot;TLS_RSA_WITH_AES_256_CBC_SHA\\\&quot;,\\\&quot;TLS_RSA_WITH_AES_256_GCM_SHA384\\\&quot;]   ves.io.schema.rules.repeated.unique: true  | [optional] 
**MaxVersion** | Pointer to [**SchemaTlsProtocol**](SchemaTlsProtocol.md) |  | [optional] [default to TLS_AUTO]
**MinVersion** | Pointer to [**SchemaTlsProtocol**](SchemaTlsProtocol.md) |  | [optional] [default to TLS_AUTO]

## Methods

### NewViewsCustomCiphers

`func NewViewsCustomCiphers() *ViewsCustomCiphers`

NewViewsCustomCiphers instantiates a new ViewsCustomCiphers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewsCustomCiphersWithDefaults

`func NewViewsCustomCiphersWithDefaults() *ViewsCustomCiphers`

NewViewsCustomCiphersWithDefaults instantiates a new ViewsCustomCiphers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherSuites

`func (o *ViewsCustomCiphers) GetCipherSuites() []string`

GetCipherSuites returns the CipherSuites field if non-nil, zero value otherwise.

### GetCipherSuitesOk

`func (o *ViewsCustomCiphers) GetCipherSuitesOk() (*[]string, bool)`

GetCipherSuitesOk returns a tuple with the CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuites

`func (o *ViewsCustomCiphers) SetCipherSuites(v []string)`

SetCipherSuites sets CipherSuites field to given value.

### HasCipherSuites

`func (o *ViewsCustomCiphers) HasCipherSuites() bool`

HasCipherSuites returns a boolean if a field has been set.

### GetMaxVersion

`func (o *ViewsCustomCiphers) GetMaxVersion() SchemaTlsProtocol`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *ViewsCustomCiphers) GetMaxVersionOk() (*SchemaTlsProtocol, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *ViewsCustomCiphers) SetMaxVersion(v SchemaTlsProtocol)`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *ViewsCustomCiphers) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### GetMinVersion

`func (o *ViewsCustomCiphers) GetMinVersion() SchemaTlsProtocol`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *ViewsCustomCiphers) GetMinVersionOk() (*SchemaTlsProtocol, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *ViewsCustomCiphers) SetMinVersion(v SchemaTlsProtocol)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *ViewsCustomCiphers) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


