# \PoolDataServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreate**](PoolDataServiceApi.md#BulkCreate) | **Post** /api/data/_r/{poolId} | Bulk create mobile data records in a pool
[**BulkDelete**](PoolDataServiceApi.md#BulkDelete) | **Delete** /api/data/_r/{poolId} | Bulk delete records from mobile data pool
[**BulkUpdate**](PoolDataServiceApi.md#BulkUpdate) | **Put** /api/data/_r/{poolId} | Bulk update records in mobile data pool
[**Get**](PoolDataServiceApi.md#Get) | **Get** /api/data/_r/{poolId}/{id} | Get single mobile data record
[**List**](PoolDataServiceApi.md#List) | **Get** /api/data/_r/{poolId} | List mobile data records



## BulkCreate

> DataBulkCreateResponse BulkCreate(ctx, poolId, body)

Bulk create mobile data records in a pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Pool ID (e.g. orders) | 
**body** | [**[]DataDocument**](dataDocument.md)|  | 

### Return type

[**DataBulkCreateResponse**](dataBulkCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDelete

> DataBulkDeleteResponse BulkDelete(ctx, poolId, body)

Bulk delete records from mobile data pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Pool ID (e.g. orders) | 
**body** | [**[]DataDocument**](dataDocument.md)|  | 

### Return type

[**DataBulkDeleteResponse**](dataBulkDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdate

> DataBulkUpdateResponse BulkUpdate(ctx, poolId, body)

Bulk update records in mobile data pool

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Pool ID (e.g. orders) | 
**body** | [**[]DataDocument**](dataDocument.md)|  | 

### Return type

[**DataBulkUpdateResponse**](dataBulkUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> DataDocument Get(ctx, poolId, id)

Get single mobile data record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Pool ID (e.g. orders) | 
**id** | **string**| The ID (Primary Key) of the record | 

### Return type

[**DataDocument**](dataDocument.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> DataListResponse List(ctx, poolId, limit, skip)

List mobile data records

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **string**| Pool ID (e.g. orders) | 
**limit** | **float32**| The max number of records to fetch. | 
**skip** | **float32**| The records to skip from the beggining | 

### Return type

[**DataListResponse**](dataListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

