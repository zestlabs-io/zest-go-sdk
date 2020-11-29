# \DistrConfigServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePool**](DistrConfigServiceApi.md#ActivatePool) | **Post** /api/distribution/v1/pool/activate | Activate data pool, which will enable the distribution of its data to  mobile users.
[**AssignAppToUsers**](DistrConfigServiceApi.md#AssignAppToUsers) | **Post** /api/distribution/v1/app/users/assign | Assign mobile application to users
[**AssignPoolsToApp**](DistrConfigServiceApi.md#AssignPoolsToApp) | **Post** /api/distribution/v1/app/pools/assign | Assign data pools to mobile application
[**AssignTagToUser**](DistrConfigServiceApi.md#AssignTagToUser) | **Post** /api/distribution/v1/user/tags/assign | Assign filtering tag to user
[**CreateApp**](DistrConfigServiceApi.md#CreateApp) | **Post** /api/distribution/v1/app | Create mobile application
[**CreatePool**](DistrConfigServiceApi.md#CreatePool) | **Post** /api/distribution/v1/pool | Create a data pool that can be used to distribute data to mobile users.
[**CreatePools**](DistrConfigServiceApi.md#CreatePools) | **Post** /api/distribution/v1/pools | Create several data pools, which can be used to distribute data to mobile users.
[**CreateUser**](DistrConfigServiceApi.md#CreateUser) | **Post** /api/distribution/v1/user | Registers user in distribution database
[**DeleteApp**](DistrConfigServiceApi.md#DeleteApp) | **Delete** /api/distribution/v1/app/{id} | Delete mobile application
[**DeletePool**](DistrConfigServiceApi.md#DeletePool) | **Delete** /api/distribution/v1/pool/{id} | Delete data pool with all its data from the system and the mobile users&#39;  devices.
[**DeleteUser**](DistrConfigServiceApi.md#DeleteUser) | **Delete** /api/distribution/v1/user/{id} | Deletes user from distribution db
[**GetApp**](DistrConfigServiceApi.md#GetApp) | **Get** /api/distribution/v1/app/{id} | Get a mobile application
[**GetApps**](DistrConfigServiceApi.md#GetApps) | **Get** /api/distribution/v1/apps | Get all mobile applications
[**GetPool**](DistrConfigServiceApi.md#GetPool) | **Get** /api/distribution/v1/pool/{id} | Get the data pool with the given ID.
[**GetPoolDistribution**](DistrConfigServiceApi.md#GetPoolDistribution) | **Get** /api/distribution/v1/pool/distribution/{id} | Get the data distribution details for the pool.
[**GetPools**](DistrConfigServiceApi.md#GetPools) | **Get** /api/distribution/v1/pools | Get the list of all data pool that have been created in the system.
[**GetUser**](DistrConfigServiceApi.md#GetUser) | **Get** /api/distribution/v1/user/{userID} | Fetch the user with the given ID
[**GetUsers**](DistrConfigServiceApi.md#GetUsers) | **Get** /api/distribution/v1/users | Fetch all users in account distribution db
[**SetAppBundle**](DistrConfigServiceApi.md#SetAppBundle) | **Post** /api/distribution/v1/app/{id}/bundle | Set the mobile application bundle
[**UnassignAppFromUsers**](DistrConfigServiceApi.md#UnassignAppFromUsers) | **Post** /api/distribution/v1/app/users/unassign | Unassign mobile application from users
[**UnassignPoolsFromApp**](DistrConfigServiceApi.md#UnassignPoolsFromApp) | **Post** /api/distribution/v1/app/pools/unassign | Unassign data pool from mobile application
[**UnassignTagFromUser**](DistrConfigServiceApi.md#UnassignTagFromUser) | **Post** /api/distribution/v1/user/tags/unassign | Unassign filtering tag from user
[**UpdatePool**](DistrConfigServiceApi.md#UpdatePool) | **Put** /api/distribution/v1/pool | Update an already existing data pool.
[**UpdatePool2**](DistrConfigServiceApi.md#UpdatePool2) | **Patch** /api/distribution/v1/pool/{dataPool.id} | Update an already existing data pool.
[**ValidatePoolData**](DistrConfigServiceApi.md#ValidatePoolData) | **Post** /api/distribution/v1/pool/validate | Validate pool data against pool specification



## ActivatePool

> map[string]interface{} ActivatePool(ctx, body)

Activate data pool, which will enable the distribution of its data to  mobile users.

Errors: - 400    Returned when no valid pool ID is provided in the request, or           when the pool is already active - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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


## AssignAppToUsers

> DistrconfigAssignAppToUsersResponse AssignAppToUsers(ctx, body)

Assign mobile application to users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigAssignAppToUsersRequest**](DistrconfigAssignAppToUsersRequest.md)|  | 

### Return type

[**DistrconfigAssignAppToUsersResponse**](distrconfigAssignAppToUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignPoolsToApp

> DistrconfigAssignPoolsToAppResponse AssignPoolsToApp(ctx, body)

Assign data pools to mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigAssignPoolsToAppRequest**](DistrconfigAssignPoolsToAppRequest.md)|  | 

### Return type

[**DistrconfigAssignPoolsToAppResponse**](distrconfigAssignPoolsToAppResponse.md)

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

Create a data pool that can be used to distribute data to mobile users.

Errors: - 400    Returned when no valid pool definition is provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Create several data pools, which can be used to distribute data to mobile users.

Errors: - 400    Returned when no valid pool definitions are provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Delete data pool with all its data from the system and the mobile users'  devices.

Errors: - 400    Returned when no valid pool ID is provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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


## GetApp

> DistrconfigGetAppResponse GetApp(ctx, id)

Get a mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DistrconfigGetAppResponse**](distrconfigGetAppResponse.md)

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

Get the data pool with the given ID.

Errors: - 400    Returned when no valid pool ID is provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Get the data distribution details for the pool.

Errors: - 400    Returned when no valid pool ID is provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Get the list of all data pool that have been created in the system.

Errors: - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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


## GetUser

> DistrconfigGetUserResponse GetUser(ctx, userID)

Fetch the user with the given ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string**|  | 

### Return type

[**DistrconfigGetUserResponse**](distrconfigGetUserResponse.md)

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

> map[string]interface{} SetAppBundle(ctx, id, body)

Set the mobile application bundle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
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


## UnassignAppFromUsers

> DistrconfigUnassignAppFromUsersResponse UnassignAppFromUsers(ctx, body)

Unassign mobile application from users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigUnassignAppFromUsersRequest**](DistrconfigUnassignAppFromUsersRequest.md)|  | 

### Return type

[**DistrconfigUnassignAppFromUsersResponse**](distrconfigUnassignAppFromUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignPoolsFromApp

> DistrconfigUnassignPoolsFromAppResponse UnassignPoolsFromApp(ctx, body)

Unassign data pool from mobile application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**DistrconfigUnassignPoolsFromAppRequest**](DistrconfigUnassignPoolsFromAppRequest.md)|  | 

### Return type

[**DistrconfigUnassignPoolsFromAppResponse**](distrconfigUnassignPoolsFromAppResponse.md)

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

Update an already existing data pool.

Errors: - 400    Returned when no valid pool definition is provided in the request - 404    Returned when a pool with the given ID does not exist - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Update an already existing data pool.

Errors: - 400    Returned when no valid pool definition is provided in the request - 404    Returned when a pool with the given ID does not exist - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

Errors: - 400    Returned when no valid pool definition is provided in the request - 403    Returned when the caller is not allowed to perform this call - 500    Returned whenever an internall error occurs

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

