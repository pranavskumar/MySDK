# OriginPoolOriginPoolAdvancedOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoHttpConfig** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**CircuitBreaker** | Pointer to [**ClusterCircuitBreaker**](ClusterCircuitBreaker.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  The timeout for new network connections to endpoints in the cluster.  This is specified in milliseconds. The default value is 2 seconds  Example: &#x60; \&quot;4000\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 1800000  | [optional] 
**DefaultCircuitBreaker** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**DisableCircuitBreaker** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**DisableOutlierDetection** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**DisableSubsets** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**EnableSubsets** | Pointer to [**OriginPoolOriginPoolSubsets**](OriginPoolOriginPoolSubsets.md) |  | [optional] 
**HeaderTransformationType** | Pointer to [**SchemaHeaderTransformationType**](SchemaHeaderTransformationType.md) |  | [optional] 
**Http1Config** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**Http2Options** | Pointer to [**ClusterHttp2ProtocolOptions**](ClusterHttp2ProtocolOptions.md) |  | [optional] 
**HttpIdleTimeout** | Pointer to **int64** |  The idle timeout for upstream connection pool connections. The idle timeout is defined as the  period in which there are no active requests. When the idle timeout is reached the connection  will be closed. Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.  This is specified in milliseconds. The default value is 5 minutes.  Example: &#x60; \&quot;60000\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 600000  | [optional] 
**NoPanicThreshold** | Pointer to **map[string]interface{}** | This can be used for messages where no values are needed | [optional] 
**OutlierDetection** | Pointer to [**ClusterOutlierDetectionType**](ClusterOutlierDetectionType.md) |  | [optional] 
**PanicThreshold** | Pointer to **int64** | Exclusive with [no_panic_threshold]  x-example:\&quot;25\&quot;  Configure a threshold (percentage of unhealthy endpoints) below which  all endpoints will be considered for load balancing ignoring its health status.  Validation Rules:   ves.io.schema.rules.uint32.lte: 100  | [optional] 

## Methods

### NewOriginPoolOriginPoolAdvancedOptions

`func NewOriginPoolOriginPoolAdvancedOptions() *OriginPoolOriginPoolAdvancedOptions`

NewOriginPoolOriginPoolAdvancedOptions instantiates a new OriginPoolOriginPoolAdvancedOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolOriginPoolAdvancedOptionsWithDefaults

`func NewOriginPoolOriginPoolAdvancedOptionsWithDefaults() *OriginPoolOriginPoolAdvancedOptions`

NewOriginPoolOriginPoolAdvancedOptionsWithDefaults instantiates a new OriginPoolOriginPoolAdvancedOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoHttpConfig

`func (o *OriginPoolOriginPoolAdvancedOptions) GetAutoHttpConfig() map[string]interface{}`

GetAutoHttpConfig returns the AutoHttpConfig field if non-nil, zero value otherwise.

### GetAutoHttpConfigOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetAutoHttpConfigOk() (*map[string]interface{}, bool)`

GetAutoHttpConfigOk returns a tuple with the AutoHttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHttpConfig

`func (o *OriginPoolOriginPoolAdvancedOptions) SetAutoHttpConfig(v map[string]interface{})`

SetAutoHttpConfig sets AutoHttpConfig field to given value.

### HasAutoHttpConfig

`func (o *OriginPoolOriginPoolAdvancedOptions) HasAutoHttpConfig() bool`

HasAutoHttpConfig returns a boolean if a field has been set.

### GetCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) GetCircuitBreaker() ClusterCircuitBreaker`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetCircuitBreakerOk() (*ClusterCircuitBreaker, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) SetCircuitBreaker(v ClusterCircuitBreaker)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetDefaultCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDefaultCircuitBreaker() map[string]interface{}`

GetDefaultCircuitBreaker returns the DefaultCircuitBreaker field if non-nil, zero value otherwise.

### GetDefaultCircuitBreakerOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDefaultCircuitBreakerOk() (*map[string]interface{}, bool)`

GetDefaultCircuitBreakerOk returns a tuple with the DefaultCircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) SetDefaultCircuitBreaker(v map[string]interface{})`

SetDefaultCircuitBreaker sets DefaultCircuitBreaker field to given value.

### HasDefaultCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) HasDefaultCircuitBreaker() bool`

HasDefaultCircuitBreaker returns a boolean if a field has been set.

### GetDisableCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableCircuitBreaker() map[string]interface{}`

GetDisableCircuitBreaker returns the DisableCircuitBreaker field if non-nil, zero value otherwise.

### GetDisableCircuitBreakerOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableCircuitBreakerOk() (*map[string]interface{}, bool)`

GetDisableCircuitBreakerOk returns a tuple with the DisableCircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) SetDisableCircuitBreaker(v map[string]interface{})`

SetDisableCircuitBreaker sets DisableCircuitBreaker field to given value.

### HasDisableCircuitBreaker

`func (o *OriginPoolOriginPoolAdvancedOptions) HasDisableCircuitBreaker() bool`

HasDisableCircuitBreaker returns a boolean if a field has been set.

### GetDisableOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableOutlierDetection() map[string]interface{}`

GetDisableOutlierDetection returns the DisableOutlierDetection field if non-nil, zero value otherwise.

### GetDisableOutlierDetectionOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableOutlierDetectionOk() (*map[string]interface{}, bool)`

GetDisableOutlierDetectionOk returns a tuple with the DisableOutlierDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) SetDisableOutlierDetection(v map[string]interface{})`

SetDisableOutlierDetection sets DisableOutlierDetection field to given value.

### HasDisableOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) HasDisableOutlierDetection() bool`

HasDisableOutlierDetection returns a boolean if a field has been set.

### GetDisableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableSubsets() map[string]interface{}`

GetDisableSubsets returns the DisableSubsets field if non-nil, zero value otherwise.

### GetDisableSubsetsOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetDisableSubsetsOk() (*map[string]interface{}, bool)`

GetDisableSubsetsOk returns a tuple with the DisableSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) SetDisableSubsets(v map[string]interface{})`

SetDisableSubsets sets DisableSubsets field to given value.

### HasDisableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) HasDisableSubsets() bool`

HasDisableSubsets returns a boolean if a field has been set.

### GetEnableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) GetEnableSubsets() OriginPoolOriginPoolSubsets`

GetEnableSubsets returns the EnableSubsets field if non-nil, zero value otherwise.

### GetEnableSubsetsOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetEnableSubsetsOk() (*OriginPoolOriginPoolSubsets, bool)`

GetEnableSubsetsOk returns a tuple with the EnableSubsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) SetEnableSubsets(v OriginPoolOriginPoolSubsets)`

SetEnableSubsets sets EnableSubsets field to given value.

### HasEnableSubsets

`func (o *OriginPoolOriginPoolAdvancedOptions) HasEnableSubsets() bool`

HasEnableSubsets returns a boolean if a field has been set.

### GetHeaderTransformationType

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHeaderTransformationType() SchemaHeaderTransformationType`

GetHeaderTransformationType returns the HeaderTransformationType field if non-nil, zero value otherwise.

### GetHeaderTransformationTypeOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHeaderTransformationTypeOk() (*SchemaHeaderTransformationType, bool)`

GetHeaderTransformationTypeOk returns a tuple with the HeaderTransformationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderTransformationType

`func (o *OriginPoolOriginPoolAdvancedOptions) SetHeaderTransformationType(v SchemaHeaderTransformationType)`

SetHeaderTransformationType sets HeaderTransformationType field to given value.

### HasHeaderTransformationType

`func (o *OriginPoolOriginPoolAdvancedOptions) HasHeaderTransformationType() bool`

HasHeaderTransformationType returns a boolean if a field has been set.

### GetHttp1Config

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttp1Config() map[string]interface{}`

GetHttp1Config returns the Http1Config field if non-nil, zero value otherwise.

### GetHttp1ConfigOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttp1ConfigOk() (*map[string]interface{}, bool)`

GetHttp1ConfigOk returns a tuple with the Http1Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp1Config

`func (o *OriginPoolOriginPoolAdvancedOptions) SetHttp1Config(v map[string]interface{})`

SetHttp1Config sets Http1Config field to given value.

### HasHttp1Config

`func (o *OriginPoolOriginPoolAdvancedOptions) HasHttp1Config() bool`

HasHttp1Config returns a boolean if a field has been set.

### GetHttp2Options

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttp2Options() ClusterHttp2ProtocolOptions`

GetHttp2Options returns the Http2Options field if non-nil, zero value otherwise.

### GetHttp2OptionsOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttp2OptionsOk() (*ClusterHttp2ProtocolOptions, bool)`

GetHttp2OptionsOk returns a tuple with the Http2Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp2Options

`func (o *OriginPoolOriginPoolAdvancedOptions) SetHttp2Options(v ClusterHttp2ProtocolOptions)`

SetHttp2Options sets Http2Options field to given value.

### HasHttp2Options

`func (o *OriginPoolOriginPoolAdvancedOptions) HasHttp2Options() bool`

HasHttp2Options returns a boolean if a field has been set.

### GetHttpIdleTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttpIdleTimeout() int64`

GetHttpIdleTimeout returns the HttpIdleTimeout field if non-nil, zero value otherwise.

### GetHttpIdleTimeoutOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetHttpIdleTimeoutOk() (*int64, bool)`

GetHttpIdleTimeoutOk returns a tuple with the HttpIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpIdleTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) SetHttpIdleTimeout(v int64)`

SetHttpIdleTimeout sets HttpIdleTimeout field to given value.

### HasHttpIdleTimeout

`func (o *OriginPoolOriginPoolAdvancedOptions) HasHttpIdleTimeout() bool`

HasHttpIdleTimeout returns a boolean if a field has been set.

### GetNoPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) GetNoPanicThreshold() map[string]interface{}`

GetNoPanicThreshold returns the NoPanicThreshold field if non-nil, zero value otherwise.

### GetNoPanicThresholdOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetNoPanicThresholdOk() (*map[string]interface{}, bool)`

GetNoPanicThresholdOk returns a tuple with the NoPanicThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) SetNoPanicThreshold(v map[string]interface{})`

SetNoPanicThreshold sets NoPanicThreshold field to given value.

### HasNoPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) HasNoPanicThreshold() bool`

HasNoPanicThreshold returns a boolean if a field has been set.

### GetOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) GetOutlierDetection() ClusterOutlierDetectionType`

GetOutlierDetection returns the OutlierDetection field if non-nil, zero value otherwise.

### GetOutlierDetectionOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetOutlierDetectionOk() (*ClusterOutlierDetectionType, bool)`

GetOutlierDetectionOk returns a tuple with the OutlierDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) SetOutlierDetection(v ClusterOutlierDetectionType)`

SetOutlierDetection sets OutlierDetection field to given value.

### HasOutlierDetection

`func (o *OriginPoolOriginPoolAdvancedOptions) HasOutlierDetection() bool`

HasOutlierDetection returns a boolean if a field has been set.

### GetPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) GetPanicThreshold() int64`

GetPanicThreshold returns the PanicThreshold field if non-nil, zero value otherwise.

### GetPanicThresholdOk

`func (o *OriginPoolOriginPoolAdvancedOptions) GetPanicThresholdOk() (*int64, bool)`

GetPanicThresholdOk returns a tuple with the PanicThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) SetPanicThreshold(v int64)`

SetPanicThreshold sets PanicThreshold field to given value.

### HasPanicThreshold

`func (o *OriginPoolOriginPoolAdvancedOptions) HasPanicThreshold() bool`

HasPanicThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


