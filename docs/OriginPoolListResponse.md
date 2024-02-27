# OriginPoolListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]SchemaErrorType**](SchemaErrorType.md) |  Errors(if any) while listing items from collection | [optional] 
**Items** | Pointer to [**[]OriginPoolListResponseItem**](OriginPoolListResponseItem.md) |  items represents the collection in response | [optional] 

## Methods

### NewOriginPoolListResponse

`func NewOriginPoolListResponse() *OriginPoolListResponse`

NewOriginPoolListResponse instantiates a new OriginPoolListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginPoolListResponseWithDefaults

`func NewOriginPoolListResponseWithDefaults() *OriginPoolListResponse`

NewOriginPoolListResponseWithDefaults instantiates a new OriginPoolListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *OriginPoolListResponse) GetErrors() []SchemaErrorType`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OriginPoolListResponse) GetErrorsOk() (*[]SchemaErrorType, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OriginPoolListResponse) SetErrors(v []SchemaErrorType)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OriginPoolListResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetItems

`func (o *OriginPoolListResponse) GetItems() []OriginPoolListResponseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OriginPoolListResponse) GetItemsOk() (*[]OriginPoolListResponseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OriginPoolListResponse) SetItems(v []OriginPoolListResponseItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *OriginPoolListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


