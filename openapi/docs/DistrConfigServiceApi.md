# \DistrConfigServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePool**](DistrConfigServiceApi.md#ActivatePool) | **Post** /api/distribution/v1/activate-pool | ActivatePool data pool task
[**AssignAppPools**](DistrConfigServiceApi.md#AssignAppPools) | **Post** /api/distribution/v1/assign-app-pools | Assign data pools to mobile application
[**AssignAppUsers**](DistrConfigServiceApi.md#AssignAppUsers) | **Post** /api/distribution/v1/assign-app-users | Assign mobile application to users
[**AssignTagToUser**](DistrConfigServiceApi.md#AssignTagToUser) | **Post** /api/distribution/v1/assign-tag-to-user | Assign filtering tag to user
[**CreateApp**](DistrConfigServiceApi.md#CreateApp) | **Post** /api/distribution/v1/app | Create mobile application
[**CreatePool**](DistrConfigServiceApi.md#CreatePool) | **Post** /api/distribution/v1/pool | CreatePool new data pool task
[**CreatePools**](DistrConfigServiceApi.md#CreatePools) | **Post** /api/distribution/v1/pools | CreatePool new data pool task
[**CreateUser**](DistrConfigServiceApi.md#CreateUser) | **Post** /api/distribution/v1/user | Registers user in distribution database
[**DeleteApp**](DistrConfigServiceApi.md#DeleteApp) | **Delete** /api/distribution/v1/app/{id} | Delete mobile application
[**DeletePool**](DistrConfigServiceApi.md#DeletePool) | **Delete** /api/distribution/v1/pool/{id} | DeletePool data pool task
[**DeleteUser**](DistrConfigServiceApi.md#DeleteUser) | **Delete** /api/distribution/v1/user/{id} | Deletes user from distribution db
[**GetApps**](DistrConfigServiceApi.md#GetApps) | **Get** /api/distribution/v1/apps | Get all mobile applications
[**GetPool**](DistrConfigServiceApi.md#GetPool) | **Get** /api/distribution/v1/pool/{id} | Get data pool task
[**GetPoolDistribution**](DistrConfigServiceApi.md#GetPoolDistribution) | **Get** /api/distribution/v1/get-pool-distribution/{id} | ActivatePool data pool task
[**GetPools**](DistrConfigServiceApi.md#GetPools) | **Get** /api/distribution/v1/pools | Get all data pool tasks
[**GetUsers**](DistrConfigServiceApi.md#GetUsers) | **Get** /api/distribution/v1/users | Fetch all users in account distribution db
[**SetAppBundle**](DistrConfigServiceApi.md#SetAppBundle) | **Post** /api/distribution/v1/set-app-bundle | Set the mobile application bundle
[**UnassignAppPools**](DistrConfigServiceApi.md#UnassignAppPools) | **Post** /api/distribution/v1/unassign-app-pools | Unassign data pool from mobile application
[**UnassignAppUsers**](DistrConfigServiceApi.md#UnassignAppUsers) | **Post** /api/distribution/v1/unassign-app-users | Unassign mobile application from users
[**UnassignTagFromUser**](DistrConfigServiceApi.md#UnassignTagFromUser) | **Post** /api/distribution/v1/unassign-tag-from-user | Unassign filtering tag from user
[**UpdatePool**](DistrConfigServiceApi.md#UpdatePool) | **Put** /api/distribution/v1/pool | UpdatePool data pool task
[**UpdatePool2**](DistrConfigServiceApi.md#UpdatePool2) | **Patch** /api/distribution/v1/pool/{dataPool.id} | UpdatePool data pool task
[**ValidatePoolData**](DistrConfigServiceApi.md#ValidatePoolData) | **Post** /api/distribution/v1/validate-pool-data | Validate pool data against pool specification



## ActivatePool

> map[string]interface{} ActivatePool(ctx, body)

ActivatePool data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigActivatePoolRequest**](DistrconfigActivatePoolRequest.md)|  | 

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


## AssignAppPools

> map[string]interface{} AssignAppPools(ctx, body)

Assign data pools to mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigAssignAppPoolsRequest**](DistrconfigAssignAppPoolsRequest.md)|  | 

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


## AssignAppUsers

> map[string]interface{} AssignAppUsers(ctx, body)

Assign mobile application to users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigAssignAppUsersRequest**](DistrconfigAssignAppUsersRequest.md)|  | 

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


## AssignTagToUser

> map[string]interface{} AssignTagToUser(ctx, body)

Assign filtering tag to user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigAssignTagToUserRequest**](DistrconfigAssignTagToUserRequest.md)|  | 

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


## CreateApp

> map[string]interface{} CreateApp(ctx, body)

Create mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigMobileApp**](DistrconfigMobileApp.md)|  | 

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


## CreatePool

> map[string]interface{} CreatePool(ctx, body)

CreatePool new data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigDataPool**](DistrconfigDataPool.md)|  | 

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


## CreatePools

> map[string]interface{} CreatePools(ctx, body)

CreatePool new data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigCreatePoolsRequest**](DistrconfigCreatePoolsRequest.md)|  | 

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


## CreateUser

> map[string]interface{} CreateUser(ctx, body)

Registers user in distribution database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigCreateUserRequest**](DistrconfigCreateUserRequest.md)|  | 

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


## DeleteApp

> map[string]interface{} DeleteApp(ctx, id)

Delete mobile application

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


## DeletePool

> map[string]interface{} DeletePool(ctx, id)

DeletePool data pool task

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


## DeleteUser

> map[string]interface{} DeleteUser(ctx, id)

Deletes user from distribution db

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


## GetApps

> DistrconfigGetAppsResponse GetApps(ctx, )

Get all mobile applications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DistrconfigGetAppsResponse**](distrconfigGetAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPool

> DistrconfigGetPoolResponse GetPool(ctx, id)

Get data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DistrconfigGetPoolResponse**](distrconfigGetPoolResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPoolDistribution

> DistrconfigGetPoolDistributionResponse GetPoolDistribution(ctx, id, optional)

ActivatePool data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetPoolDistributionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPoolDistributionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**|  | 

### Return type

[**DistrconfigGetPoolDistributionResponse**](distrconfigGetPoolDistributionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPools

> DistrconfigGetPoolsResponse GetPools(ctx, )

Get all data pool tasks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DistrconfigGetPoolsResponse**](distrconfigGetPoolsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> DistrconfigGetUsersResponse GetUsers(ctx, )

Fetch all users in account distribution db

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DistrconfigGetUsersResponse**](distrconfigGetUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAppBundle

> map[string]interface{} SetAppBundle(ctx, body)

Set the mobile application bundle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigSetAppBundleRequest**](DistrconfigSetAppBundleRequest.md)|  | 

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


## UnassignAppPools

> map[string]interface{} UnassignAppPools(ctx, body)

Unassign data pool from mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigUnassignAppPoolsRequest**](DistrconfigUnassignAppPoolsRequest.md)|  | 

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


## UnassignAppUsers

> map[string]interface{} UnassignAppUsers(ctx, body)

Unassign mobile application from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigUnassignAppUsersRequest**](DistrconfigUnassignAppUsersRequest.md)|  | 

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


## UnassignTagFromUser

> map[string]interface{} UnassignTagFromUser(ctx, body)

Unassign filtering tag from user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigUnassignTagFromUserRequest**](DistrconfigUnassignTagFromUserRequest.md)|  | 

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


## UpdatePool

> map[string]interface{} UpdatePool(ctx, body)

UpdatePool data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigDataPool**](DistrconfigDataPool.md)|  | 

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


## UpdatePool2

> map[string]interface{} UpdatePool2(ctx, dataPoolId, body)

UpdatePool data pool task

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataPoolId** | **string**| Primary key together with accountId - should be unique by client - used also as name of the pool | 
**body** | [**DistrconfigDataPool**](DistrconfigDataPool.md)|  | 

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


## ValidatePoolData

> DistrconfigValidatePoolDataResponse ValidatePoolData(ctx, body)

Validate pool data against pool specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigValidatePoolDataRequest**](DistrconfigValidatePoolDataRequest.md)|  | 

### Return type

[**DistrconfigValidatePoolDataResponse**](distrconfigValidatePoolDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

