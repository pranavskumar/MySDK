# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VesIoSchemaViewsOriginPoolAPICreate**](DefaultAPI.md#VesIoSchemaViewsOriginPoolAPICreate) | **Post** /api/config/namespaces/{metadata.namespace}/origin_pools | Create Origin Pool
[**VesIoSchemaViewsOriginPoolAPIDelete**](DefaultAPI.md#VesIoSchemaViewsOriginPoolAPIDelete) | **Delete** /api/config/namespaces/{namespace}/origin_pools/{name} | Delete Origin Pool
[**VesIoSchemaViewsOriginPoolAPIGet**](DefaultAPI.md#VesIoSchemaViewsOriginPoolAPIGet) | **Get** /api/config/namespaces/{namespace}/origin_pools/{name} | Get Origin Pool
[**VesIoSchemaViewsOriginPoolAPIList**](DefaultAPI.md#VesIoSchemaViewsOriginPoolAPIList) | **Get** /api/config/namespaces/{namespace}/origin_pools | List Origin Pool
[**VesIoSchemaViewsOriginPoolAPIReplace**](DefaultAPI.md#VesIoSchemaViewsOriginPoolAPIReplace) | **Put** /api/config/namespaces/{metadata.namespace}/origin_pools/{metadata.name} | Replace Origin Pool



## VesIoSchemaViewsOriginPoolAPICreate

> OriginPoolCreateResponse VesIoSchemaViewsOriginPoolAPICreate(ctx, metadataNamespace).OriginPoolCreateRequest(originPoolCreateRequest).Execute()

Create Origin Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	metadataNamespace := "metadataNamespace_example" // string | namespace  x-example: \"staging\" This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"
	originPoolCreateRequest := *openapiclient.NewOriginPoolCreateRequest() // OriginPoolCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate(context.Background(), metadataNamespace).OriginPoolCreateRequest(originPoolCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VesIoSchemaViewsOriginPoolAPICreate`: OriginPoolCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPICreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataNamespace** | **string** | namespace  x-example: \&quot;staging\&quot; This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \&quot;\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiVesIoSchemaViewsOriginPoolAPICreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **originPoolCreateRequest** | [**OriginPoolCreateRequest**](OriginPoolCreateRequest.md) |  | 

### Return type

[**OriginPoolCreateResponse**](OriginPoolCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VesIoSchemaViewsOriginPoolAPIDelete

> interface{} VesIoSchemaViewsOriginPoolAPIDelete(ctx, namespace, name).OriginPoolDeleteRequest(originPoolDeleteRequest).Execute()

Delete Origin Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	namespace := "namespace_example" // string | namespace  x-example: \"ns1\" Namespace in which the configuration object is present
	name := "name_example" // string | name  x-example: \"name\" Name of the configuration object
	originPoolDeleteRequest := *openapiclient.NewOriginPoolDeleteRequest() // OriginPoolDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete(context.Background(), namespace, name).OriginPoolDeleteRequest(originPoolDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VesIoSchemaViewsOriginPoolAPIDelete`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace  x-example: \&quot;ns1\&quot; Namespace in which the configuration object is present | 
**name** | **string** | name  x-example: \&quot;name\&quot; Name of the configuration object | 

### Other Parameters

Other parameters are passed through a pointer to a apiVesIoSchemaViewsOriginPoolAPIDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **originPoolDeleteRequest** | [**OriginPoolDeleteRequest**](OriginPoolDeleteRequest.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VesIoSchemaViewsOriginPoolAPIGet

> OriginPoolGetResponse VesIoSchemaViewsOriginPoolAPIGet(ctx, namespace, name).ResponseFormat(responseFormat).Execute()

Get Origin Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	namespace := "namespace_example" // string | namespace  x-example: \"ns1\" The namespace in which the configuration object is present
	name := "name_example" // string | name  x-example: \"name\" The name of the configuration object to be fetched
	responseFormat := "responseFormat_example" // string | The format in which the configuration object is to be fetched. This could be for example     - in GetSpec form for the contents of object     - in CreateRequest form to create a new similar object     - to ReplaceRequest form to replace changeable values  Default format of returned resource Response should be in CreateRequest format Response should be in ReplaceRequest format Response should be in format of GetSpecType Response should have other objects referring to this object Response should have deleted and disabled objects referrred by this object (optional) (default to "GET_RSP_FORMAT_DEFAULT")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet(context.Background(), namespace, name).ResponseFormat(responseFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VesIoSchemaViewsOriginPoolAPIGet`: OriginPoolGetResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace  x-example: \&quot;ns1\&quot; The namespace in which the configuration object is present | 
**name** | **string** | name  x-example: \&quot;name\&quot; The name of the configuration object to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiVesIoSchemaViewsOriginPoolAPIGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **responseFormat** | **string** | The format in which the configuration object is to be fetched. This could be for example     - in GetSpec form for the contents of object     - in CreateRequest form to create a new similar object     - to ReplaceRequest form to replace changeable values  Default format of returned resource Response should be in CreateRequest format Response should be in ReplaceRequest format Response should be in format of GetSpecType Response should have other objects referring to this object Response should have deleted and disabled objects referrred by this object | [default to &quot;GET_RSP_FORMAT_DEFAULT&quot;]

### Return type

[**OriginPoolGetResponse**](OriginPoolGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VesIoSchemaViewsOriginPoolAPIList

> OriginPoolListResponse VesIoSchemaViewsOriginPoolAPIList(ctx, namespace).LabelFilter(labelFilter).ReportFields(reportFields).ReportStatusFields(reportStatusFields).Execute()

List Origin Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	namespace := "namespace_example" // string | namespace  x-example: \"ns1\" Namespace to scope the listing of origin_pool
	labelFilter := "labelFilter_example" // string | x-example: \"env in (staging, testing), tier in (web, db)\" A LabelSelectorType expression that every item in list response will satisfy (optional)
	reportFields := []string{"Inner_example"} // []string | x-example: \"\" Extra fields to return along with summary fields (optional)
	reportStatusFields := []string{"Inner_example"} // []string | x-example: \"\" Extra status fields to return along with summary fields (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIList(context.Background(), namespace).LabelFilter(labelFilter).ReportFields(reportFields).ReportStatusFields(reportStatusFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VesIoSchemaViewsOriginPoolAPIList`: OriginPoolListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | namespace  x-example: \&quot;ns1\&quot; Namespace to scope the listing of origin_pool | 

### Other Parameters

Other parameters are passed through a pointer to a apiVesIoSchemaViewsOriginPoolAPIListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelFilter** | **string** | x-example: \&quot;env in (staging, testing), tier in (web, db)\&quot; A LabelSelectorType expression that every item in list response will satisfy | 
 **reportFields** | **[]string** | x-example: \&quot;\&quot; Extra fields to return along with summary fields | 
 **reportStatusFields** | **[]string** | x-example: \&quot;\&quot; Extra status fields to return along with summary fields | 

### Return type

[**OriginPoolListResponse**](OriginPoolListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VesIoSchemaViewsOriginPoolAPIReplace

> map[string]interface{} VesIoSchemaViewsOriginPoolAPIReplace(ctx, metadataNamespace, metadataName).OriginPoolReplaceRequest(originPoolReplaceRequest).Execute()

Replace Origin Pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pranavskumar/MySDK"
)

func main() {
	metadataNamespace := "metadataNamespace_example" // string | namespace  x-example: \"staging\" This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"
	metadataName := "metadataName_example" // string | name  x-example: \"acmecorp-web\" The configuration object to be replaced will be looked up by name
	originPoolReplaceRequest := *openapiclient.NewOriginPoolReplaceRequest() // OriginPoolReplaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.VesIoSchemaViewsOriginPoolAPIReplace(context.Background(), metadataNamespace, metadataName).OriginPoolReplaceRequest(originPoolReplaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.VesIoSchemaViewsOriginPoolAPIReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VesIoSchemaViewsOriginPoolAPIReplace`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.VesIoSchemaViewsOriginPoolAPIReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataNamespace** | **string** | namespace  x-example: \&quot;staging\&quot; This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \&quot;\&quot; | 
**metadataName** | **string** | name  x-example: \&quot;acmecorp-web\&quot; The configuration object to be replaced will be looked up by name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVesIoSchemaViewsOriginPoolAPIReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **originPoolReplaceRequest** | [**OriginPoolReplaceRequest**](OriginPoolReplaceRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

