# ClusterCircuitBreaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionLimit** | Pointer to **int64** |  The maximum number of connections that loadbalancer will establish to all hosts in an upstream cluster.  In practice this is only applicable to TCP and HTTP/1.1 clusters since HTTP/2 uses a single connection to each host.  Remove endpoint out of load balancing decision, if number of connections reach connection limit.  Example: &#x60; \&quot;100\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 32768  | [optional] 
**MaxRequests** | Pointer to **int64** |  The maximum number of requests that can be outstanding to all hosts in a cluster at any given time.  In practice this is applicable to HTTP/2 clusters since HTTP/1.1 clusters are governed by the  maximum connections (connection_limit).  Remove endpoint out of load balancing decision, if requests exceed this count.  Example: &#x60; \&quot;10\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 32768  | [optional] 
**PendingRequests** | Pointer to **int64** |  The maximum number of requests that will be queued while waiting for a ready connection pool connection.  Since HTTP/2 requests are sent over a single connection, this circuit breaker only comes into play as the  initial connection is created, as requests will be multiplexed immediately afterwards. For HTTP/1.1, requests  are added to the list of pending requests whenever there aren’t enough upstream connections available to  immediately dispatch the request, so this circuit breaker will remain in play for the lifetime of the process.  Remove endpoint out of load balancing decision, if pending request reach  pending_request.  Example: &#x60; \&quot;20\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 32768  | [optional] 
**Priority** | Pointer to [**SchemaRoutingPriority**](SchemaRoutingPriority.md) |  | [optional] [default to DEFAULT]
**Retries** | Pointer to **int64** |  The maximum number of retries that can be outstanding to all hosts in a cluster at any given time.  Remove endpoint out of load balancing decision, if retries for request exceed this count.  Example: &#x60; \&quot;10\&quot;&#x60;  Validation Rules:   ves.io.schema.rules.uint32.lte: 32768  | [optional] 

## Methods

### NewClusterCircuitBreaker

`func NewClusterCircuitBreaker() *ClusterCircuitBreaker`

NewClusterCircuitBreaker instantiates a new ClusterCircuitBreaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCircuitBreakerWithDefaults

`func NewClusterCircuitBreakerWithDefaults() *ClusterCircuitBreaker`

NewClusterCircuitBreakerWithDefaults instantiates a new ClusterCircuitBreaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionLimit

`func (o *ClusterCircuitBreaker) GetConnectionLimit() int64`

GetConnectionLimit returns the ConnectionLimit field if non-nil, zero value otherwise.

### GetConnectionLimitOk

`func (o *ClusterCircuitBreaker) GetConnectionLimitOk() (*int64, bool)`

GetConnectionLimitOk returns a tuple with the ConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLimit

`func (o *ClusterCircuitBreaker) SetConnectionLimit(v int64)`

SetConnectionLimit sets ConnectionLimit field to given value.

### HasConnectionLimit

`func (o *ClusterCircuitBreaker) HasConnectionLimit() bool`

HasConnectionLimit returns a boolean if a field has been set.

### GetMaxRequests

`func (o *ClusterCircuitBreaker) GetMaxRequests() int64`

GetMaxRequests returns the MaxRequests field if non-nil, zero value otherwise.

### GetMaxRequestsOk

`func (o *ClusterCircuitBreaker) GetMaxRequestsOk() (*int64, bool)`

GetMaxRequestsOk returns a tuple with the MaxRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequests

`func (o *ClusterCircuitBreaker) SetMaxRequests(v int64)`

SetMaxRequests sets MaxRequests field to given value.

### HasMaxRequests

`func (o *ClusterCircuitBreaker) HasMaxRequests() bool`

HasMaxRequests returns a boolean if a field has been set.

### GetPendingRequests

`func (o *ClusterCircuitBreaker) GetPendingRequests() int64`

GetPendingRequests returns the PendingRequests field if non-nil, zero value otherwise.

### GetPendingRequestsOk

`func (o *ClusterCircuitBreaker) GetPendingRequestsOk() (*int64, bool)`

GetPendingRequestsOk returns a tuple with the PendingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingRequests

`func (o *ClusterCircuitBreaker) SetPendingRequests(v int64)`

SetPendingRequests sets PendingRequests field to given value.

### HasPendingRequests

`func (o *ClusterCircuitBreaker) HasPendingRequests() bool`

HasPendingRequests returns a boolean if a field has been set.

### GetPriority

`func (o *ClusterCircuitBreaker) GetPriority() SchemaRoutingPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ClusterCircuitBreaker) GetPriorityOk() (*SchemaRoutingPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ClusterCircuitBreaker) SetPriority(v SchemaRoutingPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ClusterCircuitBreaker) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetries

`func (o *ClusterCircuitBreaker) GetRetries() int64`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *ClusterCircuitBreaker) GetRetriesOk() (*int64, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *ClusterCircuitBreaker) SetRetries(v int64)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *ClusterCircuitBreaker) HasRetries() bool`

HasRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


