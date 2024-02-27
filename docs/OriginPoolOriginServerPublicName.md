# OriginPoolOriginServerPublicName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** |  DNS Name  Example: &#x60; \&quot;value\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true   ves.io.schema.rules.string.hostname: true   ves.io.schema.rules.string.max_len: 256   ves.io.schema.rules.string.min_len: 1  | [optional] 
**RefreshInterval** | Pointer to **int64** |  Interval for DNS refresh in seconds.  Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767  Example: &#x60; \&quot;20\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 604800  | [optional] 

## Methods

### NewOriginPoolOriginServerPublicName

`func NewOriginPoolOriginServerPublicName() *OriginPoolOriginServerPublicName`

NewOriginPoolOriginServerPublicName instantiates a new OriginPoolOriginServerPublicName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerPublicNameWithDefaults

`func NewOriginPoolOriginServerPublicNameWithDefaults() *OriginPoolOriginServerPublicName`

NewOriginPoolOriginServerPublicNameWithDefaults instantiates a new OriginPoolOriginServerPublicName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *OriginPoolOriginServerPublicName) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *OriginPoolOriginServerPublicName) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *OriginPoolOriginServerPublicName) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *OriginPoolOriginServerPublicName) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *OriginPoolOriginServerPublicName) GetRefreshInterval() int64`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *OriginPoolOriginServerPublicName) GetRefreshIntervalOk() (*int64, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *OriginPoolOriginServerPublicName) SetRefreshInterval(v int64)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *OriginPoolOriginServerPublicName) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


