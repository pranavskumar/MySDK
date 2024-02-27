# OriginPoolTlsCertificatesType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsCertificates** | Pointer to [**[]SchemaTlsCertificateType**](SchemaTlsCertificateType.md) |  TLS Certificates  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.repeated.max_items: 16   ves.io.schema.rules.repeated.min_items: 1  | [optional] 

## Methods

### NewOriginPoolTlsCertificatesType

`func NewOriginPoolTlsCertificatesType() *OriginPoolTlsCertificatesType`

NewOriginPoolTlsCertificatesType instantiates a new OriginPoolTlsCertificatesType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolTlsCertificatesTypeWithDefaults

`func NewOriginPoolTlsCertificatesTypeWithDefaults() *OriginPoolTlsCertificatesType`

NewOriginPoolTlsCertificatesTypeWithDefaults instantiates a new OriginPoolTlsCertificatesType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsCertificates

`func (o *OriginPoolTlsCertificatesType) GetTlsCertificates() []SchemaTlsCertificateType`

GetTlsCertificates returns the TlsCertificates field if non-nil, zero value otherwise.

### GetTlsCertificatesOk

`func (o *OriginPoolTlsCertificatesType) GetTlsCertificatesOk() (*[]SchemaTlsCertificateType, bool)`

GetTlsCertificatesOk returns a tuple with the TlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertificates

`func (o *OriginPoolTlsCertificatesType) SetTlsCertificates(v []SchemaTlsCertificateType)`

SetTlsCertificates sets TlsCertificates field to given value.

### HasTlsCertificates

`func (o *OriginPoolTlsCertificatesType) HasTlsCertificates() bool`

HasTlsCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


