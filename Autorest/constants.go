//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.10.2, generator: @autorest/go@4.0.0-preview.63)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package generated

const host = "http://localhost:8080"

// ClusterEndpointSelectionPolicy - Policy for selection of endpoints from local site/remote site/both
// Consider both remote and local endpoints for load balancing LOCAL_ONLY: Consider only local endpoints for load balancing
// Enable this policy to load balance ONLY among locally discovered endpoints
// Prefer the local endpoints for load balancing. If local endpoints are not present remote endpoints will be considered.
type ClusterEndpointSelectionPolicy string

const (
	ClusterEndpointSelectionPolicyDISTRIBUTED ClusterEndpointSelectionPolicy = "DISTRIBUTED"
	ClusterEndpointSelectionPolicyLOCALONLY ClusterEndpointSelectionPolicy = "LOCAL_ONLY"
	ClusterEndpointSelectionPolicyLOCALPREFERRED ClusterEndpointSelectionPolicy = "LOCAL_PREFERRED"
)

// PossibleClusterEndpointSelectionPolicyValues returns the possible values for the ClusterEndpointSelectionPolicy const type.
func PossibleClusterEndpointSelectionPolicyValues() []ClusterEndpointSelectionPolicy {
	return []ClusterEndpointSelectionPolicy{	
		ClusterEndpointSelectionPolicyDISTRIBUTED,
		ClusterEndpointSelectionPolicyLOCALONLY,
		ClusterEndpointSelectionPolicyLOCALPREFERRED,
	}
}

// ClusterLoadbalancerAlgorithm - Different load balancing algorithms supported When a connection to a endpoint in an upstream
// cluster is required, the load balancer uses loadbalancer_algorithm to determine which host is selected.
// * ROUNDROBIN: ROUNDROBIN
// Policy in which each healthy/available upstream endpoint is selected in round robin order.
// * LEASTREQUEST: LEASTREQUEST
// Policy in which loadbalancer picks the upstream endpoint which has the fewest active requests
// * RINGHASH: RINGHASH
// Policy implements consistent hashing to upstream endpoints using ring hash of endpoint names Hash of the incoming request
// is calculated using request hash policy. The ring/modulo hash load balancer
// implements consistent hashing to upstream hosts. The algorithm is based on mapping all hosts onto a circle such that the
// addition or removal of a host from the host set changes only affect 1/N
// requests. This technique is also commonly known as “ketama” hashing. A consistent hashing load balancer is only effective
// when protocol routing is used that specifies a value to hash on. The minimum
// ring size governs the replication factor for each host in the ring. For example, if the minimum ring size is 1024 and there
// are 16 hosts, each host will be replicated 64 times.
// * RANDOM: RANDOM
// Policy in which each available upstream endpoint is selected in random order. The random load balancer selects a random
// healthy host. The random load balancer generally performs better than round
// robin if no health checking policy is configured. Random selection avoids bias towards the host in the set that comes after
// a failed host.
// * LB_OVERRIDE: Load Balancer Override
// Hash policy is taken from from the load balancer which is using this origin pool
type ClusterLoadbalancerAlgorithm string

const (
	ClusterLoadbalancerAlgorithmLBOVERRIDE ClusterLoadbalancerAlgorithm = "LB_OVERRIDE"
	ClusterLoadbalancerAlgorithmLEASTREQUEST ClusterLoadbalancerAlgorithm = "LEAST_REQUEST"
	ClusterLoadbalancerAlgorithmRANDOM ClusterLoadbalancerAlgorithm = "RANDOM"
	ClusterLoadbalancerAlgorithmRINGHASH ClusterLoadbalancerAlgorithm = "RING_HASH"
	ClusterLoadbalancerAlgorithmROUNDROBIN ClusterLoadbalancerAlgorithm = "ROUND_ROBIN"
)

// PossibleClusterLoadbalancerAlgorithmValues returns the possible values for the ClusterLoadbalancerAlgorithm const type.
func PossibleClusterLoadbalancerAlgorithmValues() []ClusterLoadbalancerAlgorithm {
	return []ClusterLoadbalancerAlgorithm{	
		ClusterLoadbalancerAlgorithmLBOVERRIDE,
		ClusterLoadbalancerAlgorithmLEASTREQUEST,
		ClusterLoadbalancerAlgorithmRANDOM,
		ClusterLoadbalancerAlgorithmRINGHASH,
		ClusterLoadbalancerAlgorithmROUNDROBIN,
	}
}

type Enum4 string

const (
	Enum4GETRSPFORMATBROKENREFERENCES Enum4 = "GET_RSP_FORMAT_BROKEN_REFERENCES"
	Enum4GETRSPFORMATDEFAULT Enum4 = "GET_RSP_FORMAT_DEFAULT"
	Enum4GETRSPFORMATFORCREATE Enum4 = "GET_RSP_FORMAT_FOR_CREATE"
	Enum4GETRSPFORMATFORREPLACE Enum4 = "GET_RSP_FORMAT_FOR_REPLACE"
	Enum4GETRSPFORMATREAD Enum4 = "GET_RSP_FORMAT_READ"
	Enum4GETRSPFORMATREFERRINGOBJECTS Enum4 = "GET_RSP_FORMAT_REFERRING_OBJECTS"
)

// PossibleEnum4Values returns the possible values for the Enum4 const type.
func PossibleEnum4Values() []Enum4 {
	return []Enum4{	
		Enum4GETRSPFORMATBROKENREFERENCES,
		Enum4GETRSPFORMATDEFAULT,
		Enum4GETRSPFORMATFORCREATE,
		Enum4GETRSPFORMATFORREPLACE,
		Enum4GETRSPFORMATREAD,
		Enum4GETRSPFORMATREFERRINGOBJECTS,
	}
}

// OriginPoolGetResponseFormatCode - x-displayName: "Get Response Format" This is the various forms that can be requested
// to be sent in the GetResponse
// * GETRSPFORMAT_DEFAULT: x-displayName: "Default Format" Default format of returned resource
// * GETRSPFORMATFORCREATE: x-displayName: "Create request Format" Response should be in CreateRequest format
// * GETRSPFORMATFORREPLACE: x-displayName: "Replace request format" Response should be in ReplaceRequest format
// * GETRSPFORMAT_READ: x-displayName: "GetSpecType format" Response should be in format of GetSpecType
// * GETRSPFORMATREFERRINGOBJECTS: x-displayName: "Referring Objects" Response should have other objects referring to this
// object
// * GETRSPFORMATBROKENREFERENCES: x-displayName: "Broken Referred Objects" Response should have deleted and disabled objects
// referrred by this object
type OriginPoolGetResponseFormatCode string

const (
	OriginPoolGetResponseFormatCodeGETRSPFORMATBROKENREFERENCES OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_BROKEN_REFERENCES"
	OriginPoolGetResponseFormatCodeGETRSPFORMATDEFAULT OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_DEFAULT"
	OriginPoolGetResponseFormatCodeGETRSPFORMATFORCREATE OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_FOR_CREATE"
	OriginPoolGetResponseFormatCodeGETRSPFORMATFORREPLACE OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_FOR_REPLACE"
	OriginPoolGetResponseFormatCodeGETRSPFORMATREAD OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_READ"
	OriginPoolGetResponseFormatCodeGETRSPFORMATREFERRINGOBJECTS OriginPoolGetResponseFormatCode = "GET_RSP_FORMAT_REFERRING_OBJECTS"
)

// PossibleOriginPoolGetResponseFormatCodeValues returns the possible values for the OriginPoolGetResponseFormatCode const type.
func PossibleOriginPoolGetResponseFormatCodeValues() []OriginPoolGetResponseFormatCode {
	return []OriginPoolGetResponseFormatCode{	
		OriginPoolGetResponseFormatCodeGETRSPFORMATBROKENREFERENCES,
		OriginPoolGetResponseFormatCodeGETRSPFORMATDEFAULT,
		OriginPoolGetResponseFormatCodeGETRSPFORMATFORCREATE,
		OriginPoolGetResponseFormatCodeGETRSPFORMATFORREPLACE,
		OriginPoolGetResponseFormatCodeGETRSPFORMATREAD,
		OriginPoolGetResponseFormatCodeGETRSPFORMATREFERRINGOBJECTS,
	}
}

// SchemaErrorCode - Union of all possible error-codes from system
// * EOK: No error
// * EPERMS: Permissions error
// * EBADINPUT: Input is not correct
// * ENOTFOUND: Not found
// * EEXISTS: Already exists
// * EUNKNOWN: Unknown/catchall error
// * ESERIALIZE: Error in serializing/de-serializing
// * EINTERNAL: Server error
// * EPARTIAL: Partial error
type SchemaErrorCode string

const (
	SchemaErrorCodeEBADINPUT SchemaErrorCode = "EBADINPUT"
	SchemaErrorCodeEEXISTS SchemaErrorCode = "EEXISTS"
	SchemaErrorCodeEINTERNAL SchemaErrorCode = "EINTERNAL"
	SchemaErrorCodeENOTFOUND SchemaErrorCode = "ENOTFOUND"
	SchemaErrorCodeEOK SchemaErrorCode = "EOK"
	SchemaErrorCodeEPARTIAL SchemaErrorCode = "EPARTIAL"
	SchemaErrorCodeEPERMS SchemaErrorCode = "EPERMS"
	SchemaErrorCodeESERIALIZE SchemaErrorCode = "ESERIALIZE"
	SchemaErrorCodeEUNKNOWN SchemaErrorCode = "EUNKNOWN"
)

// PossibleSchemaErrorCodeValues returns the possible values for the SchemaErrorCode const type.
func PossibleSchemaErrorCodeValues() []SchemaErrorCode {
	return []SchemaErrorCode{	
		SchemaErrorCodeEBADINPUT,
		SchemaErrorCodeEEXISTS,
		SchemaErrorCodeEINTERNAL,
		SchemaErrorCodeENOTFOUND,
		SchemaErrorCodeEOK,
		SchemaErrorCodeEPARTIAL,
		SchemaErrorCodeEPERMS,
		SchemaErrorCodeESERIALIZE,
		SchemaErrorCodeEUNKNOWN,
	}
}

// SchemaHashAlgorithm - Specifies the Hash Algorithm to be used
// Invalid hash algorithm sha256 hash algorithm sha1 hash algorithm
type SchemaHashAlgorithm string

const (
	SchemaHashAlgorithmINVALIDHASHALGORITHM SchemaHashAlgorithm = "INVALID_HASH_ALGORITHM"
	SchemaHashAlgorithmSHA1 SchemaHashAlgorithm = "SHA1"
	SchemaHashAlgorithmSHA256 SchemaHashAlgorithm = "SHA256"
)

// PossibleSchemaHashAlgorithmValues returns the possible values for the SchemaHashAlgorithm const type.
func PossibleSchemaHashAlgorithmValues() []SchemaHashAlgorithm {
	return []SchemaHashAlgorithm{	
		SchemaHashAlgorithmINVALIDHASHALGORITHM,
		SchemaHashAlgorithmSHA1,
		SchemaHashAlgorithmSHA256,
	}
}

// SchemaRoutingPriority - Priority routing for each request. Different connection pools are used based on the priority selected
// for the request. Also, circuit-breaker configuration at destination cluster is chosen based on
// selected priority.
// Default routing mechanism High-Priority routing mechanism
type SchemaRoutingPriority string

const (
	SchemaRoutingPriorityDEFAULT SchemaRoutingPriority = "DEFAULT"
	SchemaRoutingPriorityHIGH SchemaRoutingPriority = "HIGH"
)

// PossibleSchemaRoutingPriorityValues returns the possible values for the SchemaRoutingPriority const type.
func PossibleSchemaRoutingPriorityValues() []SchemaRoutingPriority {
	return []SchemaRoutingPriority{	
		SchemaRoutingPriorityDEFAULT,
		SchemaRoutingPriorityHIGH,
	}
}

// SchemaSecretEncodingType - x-displayName: "Secret Encoding" SecretEncodingType defines the encoding type of the secret
// before handled by the Secret Management Service.
// * EncodingNone: x-displayName: "None" No Encoding
// * EncodingBase64: Base64
// x-displayName: "Base64" Base64 encoding
type SchemaSecretEncodingType string

const (
	SchemaSecretEncodingTypeEncodingBase64 SchemaSecretEncodingType = "EncodingBase64"
	SchemaSecretEncodingTypeEncodingNone SchemaSecretEncodingType = "EncodingNone"
)

// PossibleSchemaSecretEncodingTypeValues returns the possible values for the SchemaSecretEncodingType const type.
func PossibleSchemaSecretEncodingTypeValues() []SchemaSecretEncodingType {
	return []SchemaSecretEncodingType{	
		SchemaSecretEncodingTypeEncodingBase64,
		SchemaSecretEncodingTypeEncodingNone,
	}
}

// SchemaTLSProtocol - TlsProtocol is enumeration of supported TLS versions
// F5 Distributed Cloud will choose the optimal TLS version.
type SchemaTLSProtocol string

const (
	SchemaTLSProtocolTLSAUTO SchemaTLSProtocol = "TLS_AUTO"
	SchemaTLSProtocolTLSv10 SchemaTLSProtocol = "TLSv1_0"
	SchemaTLSProtocolTLSv11 SchemaTLSProtocol = "TLSv1_1"
	SchemaTLSProtocolTLSv12 SchemaTLSProtocol = "TLSv1_2"
	SchemaTLSProtocolTLSv13 SchemaTLSProtocol = "TLSv1_3"
)

// PossibleSchemaTLSProtocolValues returns the possible values for the SchemaTLSProtocol const type.
func PossibleSchemaTLSProtocolValues() []SchemaTLSProtocol {
	return []SchemaTLSProtocol{	
		SchemaTLSProtocolTLSAUTO,
		SchemaTLSProtocolTLSv10,
		SchemaTLSProtocolTLSv11,
		SchemaTLSProtocolTLSv12,
		SchemaTLSProtocolTLSv13,
	}
}

