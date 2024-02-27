/*
F5 Distributed Cloud Services API for ves.io.schema.views.origin_pool

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ClusterLoadbalancerAlgorithm Different load balancing algorithms supported When a connection to a endpoint in an upstream cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected.   - ROUND_ROBIN: ROUND_ROBIN  Policy in which each healthy/available upstream endpoint is selected in round robin order.  - LEAST_REQUEST: LEAST_REQUEST  Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests  - RING_HASH: RING_HASH  Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request is calculated using request hash policy. The ring/modulo hash load balancer implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the addition or removal of a host from the host set changes only affect 1/N requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective when protocol routing is used that specifies a value to hash on. The minimum ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there are 16 hosts, each host will be replicated 64 times.  - RANDOM: RANDOM  Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random healthy host. The random load balancer generally performs better than round robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after a failed host.  - LB_OVERRIDE: Load Balancer Override  Hash policy is taken from from the load balancer which is using this origin pool
type ClusterLoadbalancerAlgorithm string

// List of clusterLoadbalancerAlgorithm
const (
	ROUND_ROBIN ClusterLoadbalancerAlgorithm = "ROUND_ROBIN"
	LEAST_REQUEST ClusterLoadbalancerAlgorithm = "LEAST_REQUEST"
	RING_HASH ClusterLoadbalancerAlgorithm = "RING_HASH"
	RANDOM ClusterLoadbalancerAlgorithm = "RANDOM"
	LB_OVERRIDE ClusterLoadbalancerAlgorithm = "LB_OVERRIDE"
)

// All allowed values of ClusterLoadbalancerAlgorithm enum
var AllowedClusterLoadbalancerAlgorithmEnumValues = []ClusterLoadbalancerAlgorithm{
	"ROUND_ROBIN",
	"LEAST_REQUEST",
	"RING_HASH",
	"RANDOM",
	"LB_OVERRIDE",
}

func (v *ClusterLoadbalancerAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterLoadbalancerAlgorithm(value)
	for _, existing := range AllowedClusterLoadbalancerAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterLoadbalancerAlgorithm", value)
}

// NewClusterLoadbalancerAlgorithmFromValue returns a pointer to a valid ClusterLoadbalancerAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterLoadbalancerAlgorithmFromValue(v string) (*ClusterLoadbalancerAlgorithm, error) {
	ev := ClusterLoadbalancerAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterLoadbalancerAlgorithm: valid values are %v", v, AllowedClusterLoadbalancerAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterLoadbalancerAlgorithm) IsValid() bool {
	for _, existing := range AllowedClusterLoadbalancerAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clusterLoadbalancerAlgorithm value
func (v ClusterLoadbalancerAlgorithm) Ptr() *ClusterLoadbalancerAlgorithm {
	return &v
}

type NullableClusterLoadbalancerAlgorithm struct {
	value *ClusterLoadbalancerAlgorithm
	isSet bool
}

func (v NullableClusterLoadbalancerAlgorithm) Get() *ClusterLoadbalancerAlgorithm {
	return v.value
}

func (v *NullableClusterLoadbalancerAlgorithm) Set(val *ClusterLoadbalancerAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLoadbalancerAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLoadbalancerAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLoadbalancerAlgorithm(val *ClusterLoadbalancerAlgorithm) *NullableClusterLoadbalancerAlgorithm {
	return &NullableClusterLoadbalancerAlgorithm{value: val, isSet: true}
}

func (v NullableClusterLoadbalancerAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLoadbalancerAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

