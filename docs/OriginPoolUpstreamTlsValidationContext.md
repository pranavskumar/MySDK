# OriginPoolUpstreamTlsValidationContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustedCaUrl** | Pointer to **string** | Exclusive with []  Inline Root CA certificate for verification of Server&#39;s certificate  Validation Rules:   ves.io.schema.rules.string.max_bytes: 131072   ves.io.schema.rules.string.min_bytes: 1   ves.io.schema.rules.string.truststore_url: true  | [optional] 

## Methods

### NewOriginPoolUpstreamTlsValidationContext

`func NewOriginPoolUpstreamTlsValidationContext() *OriginPoolUpstreamTlsValidationContext`

NewOriginPoolUpstreamTlsValidationContext instantiates a new OriginPoolUpstreamTlsValidationContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolUpstreamTlsValidationContextWithDefaults

`func NewOriginPoolUpstreamTlsValidationContextWithDefaults() *OriginPoolUpstreamTlsValidationContext`

NewOriginPoolUpstreamTlsValidationContextWithDefaults instantiates a new OriginPoolUpstreamTlsValidationContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrustedCaUrl

`func (o *OriginPoolUpstreamTlsValidationContext) GetTrustedCaUrl() string`

GetTrustedCaUrl returns the TrustedCaUrl field if non-nil, zero value otherwise.

### GetTrustedCaUrlOk

`func (o *OriginPoolUpstreamTlsValidationContext) GetTrustedCaUrlOk() (*string, bool)`

GetTrustedCaUrlOk returns a tuple with the TrustedCaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCaUrl

`func (o *OriginPoolUpstreamTlsValidationContext) SetTrustedCaUrl(v string)`

SetTrustedCaUrl sets TrustedCaUrl field to given value.

### HasTrustedCaUrl

`func (o *OriginPoolUpstreamTlsValidationContext) HasTrustedCaUrl() bool`

HasTrustedCaUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


