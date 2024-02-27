# ClusterOutlierDetectionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseEjectionTime** | Pointer to **int64** |  The base time that a host is ejected for. The real time is equal to the  base time multiplied by the number of times the host has been ejected.  This causes hosts to get ejected for longer periods if they continue to fail.  Defaults to 30000ms or 30s. Specified in milliseconds.  Example: &#x60; \&quot;20000\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 1800000  | [optional] 
**Consecutive5xx** | Pointer to **int64** |  If an upstream endpoint returns some number of consecutive 5xx, it will be ejected.  Note that in this case a 5xx means an actual 5xx respond code, or an event that would  cause the HTTP router to return one on the upstream’s behalf(reset, connection failure, etc.)  consecutive_5xx indicates the number of consecutive 5xx responses required before  a consecutive 5xx ejection occurs. Defaults to 5.  Example: &#x60; \&quot;3\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 1024  | [optional] 
**ConsecutiveGatewayFailure** | Pointer to **int64** |  If an upstream endpoint returns some number of consecutive “gateway errors”  (502, 503 or 504 status code), it will be ejected. Note that this includes events  that would cause the HTTP router to return one of these status codes on the  upstream’s behalf (reset, connection failure, etc.).  consecutive_gateway_failure indicates the number of consecutive gateway failures  before a consecutive gateway failure ejection occurs. Defaults to 5.  Example: &#x60; \&quot;5\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 1024  | [optional] 
**Interval** | Pointer to **int64** |  The time interval between ejection analysis sweeps. This can result in  both new ejections as well as endpoints being returned to service. Defaults  to 10000ms or 10s. Specified in milliseconds.  Example: &#x60; \&quot;5000\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 600000  | [optional] 
**MaxEjectionPercent** | Pointer to **int64** |  The maximum % of an upstream cluster that can be ejected due to outlier  detection. Defaults to 10% but will eject at least one host regardless of the value.  Example: &#x60; \&quot;20\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 100  | [optional] 

## Methods

### NewClusterOutlierDetectionType

`func NewClusterOutlierDetectionType() *ClusterOutlierDetectionType`

NewClusterOutlierDetectionType instantiates a new ClusterOutlierDetectionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOutlierDetectionTypeWithDefaults

`func NewClusterOutlierDetectionTypeWithDefaults() *ClusterOutlierDetectionType`

NewClusterOutlierDetectionTypeWithDefaults instantiates a new ClusterOutlierDetectionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseEjectionTime

`func (o *ClusterOutlierDetectionType) GetBaseEjectionTime() int64`

GetBaseEjectionTime returns the BaseEjectionTime field if non-nil, zero value otherwise.

### GetBaseEjectionTimeOk

`func (o *ClusterOutlierDetectionType) GetBaseEjectionTimeOk() (*int64, bool)`

GetBaseEjectionTimeOk returns a tuple with the BaseEjectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEjectionTime

`func (o *ClusterOutlierDetectionType) SetBaseEjectionTime(v int64)`

SetBaseEjectionTime sets BaseEjectionTime field to given value.

### HasBaseEjectionTime

`func (o *ClusterOutlierDetectionType) HasBaseEjectionTime() bool`

HasBaseEjectionTime returns a boolean if a field has been set.

### GetConsecutive5xx

`func (o *ClusterOutlierDetectionType) GetConsecutive5xx() int64`

GetConsecutive5xx returns the Consecutive5xx field if non-nil, zero value otherwise.

### GetConsecutive5xxOk

`func (o *ClusterOutlierDetectionType) GetConsecutive5xxOk() (*int64, bool)`

GetConsecutive5xxOk returns a tuple with the Consecutive5xx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutive5xx

`func (o *ClusterOutlierDetectionType) SetConsecutive5xx(v int64)`

SetConsecutive5xx sets Consecutive5xx field to given value.

### HasConsecutive5xx

`func (o *ClusterOutlierDetectionType) HasConsecutive5xx() bool`

HasConsecutive5xx returns a boolean if a field has been set.

### GetConsecutiveGatewayFailure

`func (o *ClusterOutlierDetectionType) GetConsecutiveGatewayFailure() int64`

GetConsecutiveGatewayFailure returns the ConsecutiveGatewayFailure field if non-nil, zero value otherwise.

### GetConsecutiveGatewayFailureOk

`func (o *ClusterOutlierDetectionType) GetConsecutiveGatewayFailureOk() (*int64, bool)`

GetConsecutiveGatewayFailureOk returns a tuple with the ConsecutiveGatewayFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveGatewayFailure

`func (o *ClusterOutlierDetectionType) SetConsecutiveGatewayFailure(v int64)`

SetConsecutiveGatewayFailure sets ConsecutiveGatewayFailure field to given value.

### HasConsecutiveGatewayFailure

`func (o *ClusterOutlierDetectionType) HasConsecutiveGatewayFailure() bool`

HasConsecutiveGatewayFailure returns a boolean if a field has been set.

### GetInterval

`func (o *ClusterOutlierDetectionType) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ClusterOutlierDetectionType) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ClusterOutlierDetectionType) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ClusterOutlierDetectionType) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxEjectionPercent

`func (o *ClusterOutlierDetectionType) GetMaxEjectionPercent() int64`

GetMaxEjectionPercent returns the MaxEjectionPercent field if non-nil, zero value otherwise.

### GetMaxEjectionPercentOk

`func (o *ClusterOutlierDetectionType) GetMaxEjectionPercentOk() (*int64, bool)`

GetMaxEjectionPercentOk returns a tuple with the MaxEjectionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEjectionPercent

`func (o *ClusterOutlierDetectionType) SetMaxEjectionPercent(v int64)`

SetMaxEjectionPercent sets MaxEjectionPercent field to given value.

### HasMaxEjectionPercent

`func (o *ClusterOutlierDetectionType) HasMaxEjectionPercent() bool`

HasMaxEjectionPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


