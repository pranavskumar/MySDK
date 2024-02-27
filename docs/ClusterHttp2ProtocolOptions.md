# ClusterHttp2ProtocolOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  Enable/disable HTTP2 Protocol for upstream connections | [optional] 

## Methods

### NewClusterHttp2ProtocolOptions

`func NewClusterHttp2ProtocolOptions() *ClusterHttp2ProtocolOptions`

NewClusterHttp2ProtocolOptions instantiates a new ClusterHttp2ProtocolOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHttp2ProtocolOptionsWithDefaults

`func NewClusterHttp2ProtocolOptionsWithDefaults() *ClusterHttp2ProtocolOptions`

NewClusterHttp2ProtocolOptionsWithDefaults instantiates a new ClusterHttp2ProtocolOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterHttp2ProtocolOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterHttp2ProtocolOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterHttp2ProtocolOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterHttp2ProtocolOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


