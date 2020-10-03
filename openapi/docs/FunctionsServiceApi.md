# \FunctionsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallAsyncFunction**](FunctionsServiceApi.md#CallAsyncFunction) | **Post** /api/func/v1/post-call/{reason}/{call.id} | Post function for execution in async way
[**CallSyncFunction**](FunctionsServiceApi.md#CallSyncFunction) | **Post** /api/func/v1/call-sync/{id} | Call a function synchroniously
[**CallTestFunction**](FunctionsServiceApi.md#CallTestFunction) | **Post** /api/func/v1/call-test/{id} | Test a function call
[**CreateFunction**](FunctionsServiceApi.md#CreateFunction) | **Post** /api/func/v1/function | Create function
[**DeleteFunction**](FunctionsServiceApi.md#DeleteFunction) | **Delete** /api/func/v1/function/{id} | Delete function
[**GetFunction**](FunctionsServiceApi.md#GetFunction) | **Get** /api/func/v1/function/{id} | Get function
[**GetFunctions**](FunctionsServiceApi.md#GetFunctions) | **Get** /api/func/v1/functions | Fetch all functions
[**UpdateFunction**](FunctionsServiceApi.md#UpdateFunction) | **Put** /api/func/v1/function | Update function
[**UpdateFunction2**](FunctionsServiceApi.md#UpdateFunction2) | **Patch** /api/func/v1/function/{function.id} | Update function



## CallAsyncFunction

> map[string]interface{} CallAsyncFunction(ctx, reason, callId, body)

Post function for execution in async way

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reason** | **string**|  | 
**callId** | **string**|  | 
**body** | **map[string]interface{}**|  | 

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


## CallSyncFunction

> FunctionsCallFunctionResponse CallSyncFunction(ctx, id, body)

Call a function synchroniously

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**body** | **map[string]interface{}**|  | 

### Return type

[**FunctionsCallFunctionResponse**](functionsCallFunctionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallTestFunction

> FunctionsCallFunctionResponse CallTestFunction(ctx, id, body)

Test a function call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**body** | **map[string]interface{}**|  | 

### Return type

[**FunctionsCallFunctionResponse**](functionsCallFunctionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFunction

> FunctionsCreateFunctionResponse CreateFunction(ctx, body)

Create function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**FunctionsCreateFunctionRequest**](FunctionsCreateFunctionRequest.md)|  | 

### Return type

[**FunctionsCreateFunctionResponse**](functionsCreateFunctionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFunction

> map[string]interface{} DeleteFunction(ctx, id)

Delete function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> FunctionsGetFunctionResponse GetFunction(ctx, id)

Get function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**FunctionsGetFunctionResponse**](functionsGetFunctionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctions

> FunctionsGetFunctionsResponse GetFunctions(ctx, )

Fetch all functions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FunctionsGetFunctionsResponse**](functionsGetFunctionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFunction

> map[string]interface{} UpdateFunction(ctx, body)

Update function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**FunctionsFunction**](FunctionsFunction.md)|  | 

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


## UpdateFunction2

> map[string]interface{} UpdateFunction2(ctx, functionId, body)

Update function

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**functionId** | **string**|  | 
**body** | [**FunctionsFunction**](FunctionsFunction.md)|  | 

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

