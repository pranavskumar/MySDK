# OriginPoolOriginServerPrivateName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **string** |  DNS Name  Example: &#x60; \&quot;value\&quot;&#x60;  Required: YES  Validation Rules:   ves.io.schema.rules.message.required: true  | [optional] 
**InsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**OutsideNetwork** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**RefreshInterval** | Pointer to **int64** |  Interval for DNS refresh in seconds.  Max value is 7 days as per https://datatracker.ietf.org/doc/html/rfc8767  Example: &#x60; \&quot;20\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 604800  | [optional] 
**SiteLocator** | Pointer to [**ViewsSiteLocator**](ViewsSiteLocator.md) |  | [optional] 

## Methods

### NewOriginPoolOriginServerPrivateName

`func NewOriginPoolOriginServerPrivateName() *OriginPoolOriginServerPrivateName`

NewOriginPoolOriginServerPrivateName instantiates a new OriginPoolOriginServerPrivateName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginServerPrivateNameWithDefaults

`func NewOriginPoolOriginServerPrivateNameWithDefaults() *OriginPoolOriginServerPrivateName`

NewOriginPoolOriginServerPrivateNameWithDefaults instantiates a new OriginPoolOriginServerPrivateName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *OriginPoolOriginServerPrivateName) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *OriginPoolOriginServerPrivateName) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *OriginPoolOriginServerPrivateName) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *OriginPoolOriginServerPrivateName) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetInsideNetwork

`func (o *OriginPoolOriginServerPrivateName) GetInsideNetwork() map[string]interface{}`

GetInsideNetwork returns the InsideNetwork field if non-nil, zero value otherwise.

### GetInsideNetworkOk

`func (o *OriginPoolOriginServerPrivateName) GetInsideNetworkOk() (*map[string]interface{}, bool)`

GetInsideNetworkOk returns a tuple with the InsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsideNetwork

`func (o *OriginPoolOriginServerPrivateName) SetInsideNetwork(v map[string]interface{})`

SetInsideNetwork sets InsideNetwork field to given value.

### HasInsideNetwork

`func (o *OriginPoolOriginServerPrivateName) HasInsideNetwork() bool`

HasInsideNetwork returns a boolean if a field has been set.

### GetOutsideNetwork

`func (o *OriginPoolOriginServerPrivateName) GetOutsideNetwork() map[string]interface{}`

GetOutsideNetwork returns the OutsideNetwork field if non-nil, zero value otherwise.

### GetOutsideNetworkOk

`func (o *OriginPoolOriginServerPrivateName) GetOutsideNetworkOk() (*map[string]interface{}, bool)`

GetOutsideNetworkOk returns a tuple with the OutsideNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideNetwork

`func (o *OriginPoolOriginServerPrivateName) SetOutsideNetwork(v map[string]interface{})`

SetOutsideNetwork sets OutsideNetwork field to given value.

### HasOutsideNetwork

`func (o *OriginPoolOriginServerPrivateName) HasOutsideNetwork() bool`

HasOutsideNetwork returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *OriginPoolOriginServerPrivateName) GetRefreshInterval() int64`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *OriginPoolOriginServerPrivateName) GetRefreshIntervalOk() (*int64, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *OriginPoolOriginServerPrivateName) SetRefreshInterval(v int64)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *OriginPoolOriginServerPrivateName) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetSiteLocator

`func (o *OriginPoolOriginServerPrivateName) GetSiteLocator() ViewsSiteLocator`

GetSiteLocator returns the SiteLocator field if non-nil, zero value otherwise.

### GetSiteLocatorOk

`func (o *OriginPoolOriginServerPrivateName) GetSiteLocatorOk() (*ViewsSiteLocator, bool)`

GetSiteLocatorOk returns a tuple with the SiteLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteLocator

`func (o *OriginPoolOriginServerPrivateName) SetSiteLocator(v ViewsSiteLocator)`

SetSiteLocator sets SiteLocator field to given value.

### HasSiteLocator

`func (o *OriginPoolOriginServerPrivateName) HasSiteLocator() bool`

HasSiteLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


