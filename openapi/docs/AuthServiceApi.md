# \AuthServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUsersToRole**](AuthServiceApi.md#AddUsersToRole) | **Post** /api/auth/v1/role/users/add | 
[**CheckHMACAuth**](AuthServiceApi.md#CheckHMACAuth) | **Post** /api/auth/v1/check-hmac | Auth API
[**CreateAccessKey**](AuthServiceApi.md#CreateAccessKey) | **Post** /api/auth/v1/accesskey | User Access Keys API
[**CreatePolicy**](AuthServiceApi.md#CreatePolicy) | **Post** /api/auth/v1/policy | 
[**CreateRole**](AuthServiceApi.md#CreateRole) | **Post** /api/auth/v1/role | 
[**CreateUser**](AuthServiceApi.md#CreateUser) | **Post** /api/auth/v1/user | 
[**DeleteAccessKey**](AuthServiceApi.md#DeleteAccessKey) | **Delete** /api/auth/v1/accesskey/{accessKeyID} | 
[**DeletePolicy**](AuthServiceApi.md#DeletePolicy) | **Delete** /api/auth/v1/policy/{policyID} | 
[**DeleteRole**](AuthServiceApi.md#DeleteRole) | **Delete** /api/auth/v1/role/{roleID} | 
[**DeleteUser**](AuthServiceApi.md#DeleteUser) | **Delete** /api/auth/v1/user/{userID} | 
[**GetDefaultPolicies**](AuthServiceApi.md#GetDefaultPolicies) | **Get** /api/auth/v1/default-policies | 
[**GetOwnAccount**](AuthServiceApi.md#GetOwnAccount) | **Get** /api/auth/v1/account | 
[**GetPasswordPolicy**](AuthServiceApi.md#GetPasswordPolicy) | **Get** /api/auth/v1/password-policy | 
[**GetPolicies**](AuthServiceApi.md#GetPolicies) | **Get** /api/auth/v1/policies | 
[**GetPolicy**](AuthServiceApi.md#GetPolicy) | **Get** /api/auth/v1/policy/{policyID} | RBAC API
[**GetRole**](AuthServiceApi.md#GetRole) | **Get** /api/auth/v1/role/{roleID} | 
[**GetRoles**](AuthServiceApi.md#GetRoles) | **Get** /api/auth/v1/roles | 
[**GetUser**](AuthServiceApi.md#GetUser) | **Get** /api/auth/v1/user/{userID} | 
[**GetUserAccessKeys**](AuthServiceApi.md#GetUserAccessKeys) | **Get** /api/auth/v1/user/accesskeys/{userID} | 
[**GetUserIDByEmail**](AuthServiceApi.md#GetUserIDByEmail) | **Get** /api/auth/v1/user-id-by-email/{email} | 
[**GetUserInfo**](AuthServiceApi.md#GetUserInfo) | **Get** /api/auth/v1/userinfo | User API
[**GetUsers**](AuthServiceApi.md#GetUsers) | **Get** /api/auth/v1/users | 
[**GetUsersForRole**](AuthServiceApi.md#GetUsersForRole) | **Get** /api/auth/v1/role/users/{roleID} | 
[**RemoveUsersFromRole**](AuthServiceApi.md#RemoveUsersFromRole) | **Post** /api/auth/v1/role/users/remove | 
[**ResetPassword**](AuthServiceApi.md#ResetPassword) | **Post** /api/auth/v1/reset-password | 
[**SetPassword**](AuthServiceApi.md#SetPassword) | **Post** /api/auth/v1/set-password | Password API
[**UpdateAccount**](AuthServiceApi.md#UpdateAccount) | **Put** /api/auth/v1/account/{account.accountID} | 
[**UpdatePolicy**](AuthServiceApi.md#UpdatePolicy) | **Put** /api/auth/v1/policy | 
[**UpdateRole**](AuthServiceApi.md#UpdateRole) | **Put** /api/auth/v1/role | 
[**UpdateUser**](AuthServiceApi.md#UpdateUser) | **Put** /api/auth/v1/user/{user.userID} | 



## AddUsersToRole

> V1AddUsersToRoleResponse AddUsersToRole(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1AddUsersToRoleRequest**](V1AddUsersToRoleRequest.md)|  | 

### Return type

[**V1AddUsersToRoleResponse**](v1AddUsersToRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckHMACAuth

> V1CheckHmacAuthResponse CheckHMACAuth(ctx, body)

Auth API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1CheckHmacAuthRequest**](V1CheckHmacAuthRequest.md)|  | 

### Return type

[**V1CheckHmacAuthResponse**](v1CheckHMACAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccessKey

> V1CreateAccessKeyResponse CreateAccessKey(ctx, body)

User Access Keys API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1CreateAccessKeyRequest**](V1CreateAccessKeyRequest.md)|  | 

### Return type

[**V1CreateAccessKeyResponse**](v1CreateAccessKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> V1CreatePolicyResponse CreatePolicy(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1CreatePolicyRequest**](V1CreatePolicyRequest.md)|  | 

### Return type

[**V1CreatePolicyResponse**](v1CreatePolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRole

> V1CreateRoleResponse CreateRole(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1CreateRoleRequest**](V1CreateRoleRequest.md)|  | 

### Return type

[**V1CreateRoleResponse**](v1CreateRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> V1CreateUserResponse CreateUser(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1CreateUserRequest**](V1CreateUserRequest.md)|  | 

### Return type

[**V1CreateUserResponse**](v1CreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKey

> map[string]interface{} DeleteAccessKey(ctx, accessKeyID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessKeyID** | **string**|  | 

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


## DeletePolicy

> map[string]interface{} DeletePolicy(ctx, policyID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyID** | **string**|  | 

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


## DeleteRole

> map[string]interface{} DeleteRole(ctx, roleID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **string**|  | 

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

> map[string]interface{} DeleteUser(ctx, userID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string**|  | 

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


## GetDefaultPolicies

> V1GetDefaultPoliciesResponse GetDefaultPolicies(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDefaultPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDefaultPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.String**|  | 
 **limit** | **optional.String**|  | 

### Return type

[**V1GetDefaultPoliciesResponse**](v1GetDefaultPoliciesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnAccount

> V1GetAccountResponse GetOwnAccount(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetAccountResponse**](v1GetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordPolicy

> V1GetPasswordPolicyResponse GetPasswordPolicy(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetPasswordPolicyResponse**](v1GetPasswordPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicies

> V1GetPoliciesResponse GetPolicies(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetPoliciesResponse**](v1GetPoliciesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> V1GetPolicyResponse GetPolicy(ctx, policyID)

RBAC API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyID** | **string**|  | 

### Return type

[**V1GetPolicyResponse**](v1GetPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> V1GetRoleResponse GetRole(ctx, roleID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **string**|  | 

### Return type

[**V1GetRoleResponse**](v1GetRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> V1GetRolesResponse GetRoles(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetRolesResponse**](v1GetRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> V1GetUserResponse GetUser(ctx, userID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string**|  | 

### Return type

[**V1GetUserResponse**](v1GetUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAccessKeys

> V1GetUserAccessKeysResponse GetUserAccessKeys(ctx, userID)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string**|  | 

### Return type

[**V1GetUserAccessKeysResponse**](v1GetUserAccessKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserIDByEmail

> V1GetUserIdByEmailResponse GetUserIDByEmail(ctx, email)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**|  | 

### Return type

[**V1GetUserIdByEmailResponse**](v1GetUserIDByEmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> V1GetUserInfoResponse GetUserInfo(ctx, )

User API

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**V1GetUserInfoResponse**](v1GetUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> V1GetUsersResponse GetUsers(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**V1GetUsersResponse**](v1GetUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersForRole

> V1GetUsersForRoleResponse GetUsersForRole(ctx, roleID, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **string**|  | 
 **optional** | ***GetUsersForRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersForRoleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.String**|  | 
 **limit** | **optional.String**|  | 

### Return type

[**V1GetUsersForRoleResponse**](v1GetUsersForRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUsersFromRole

> V1RemoveUsersFromRoleResponse RemoveUsersFromRole(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1RemoveUsersFromRoleRequest**](V1RemoveUsersFromRoleRequest.md)|  | 

### Return type

[**V1RemoveUsersFromRoleResponse**](v1RemoveUsersFromRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> map[string]interface{} ResetPassword(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1ResetPasswordRequest**](V1ResetPasswordRequest.md)|  | 

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


## SetPassword

> map[string]interface{} SetPassword(ctx, body)

Password API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1SetPasswordRequest**](V1SetPasswordRequest.md)|  | 

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


## UpdateAccount

> map[string]interface{} UpdateAccount(ctx, accountAccountID, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountAccountID** | **string**|  | 
**body** | [**V1UpdateAccountRequest**](V1UpdateAccountRequest.md)|  | 

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


## UpdatePolicy

> map[string]interface{} UpdatePolicy(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1UpdatePolicyRequest**](V1UpdatePolicyRequest.md)|  | 

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


## UpdateRole

> map[string]interface{} UpdateRole(ctx, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1UpdateRoleRequest**](V1UpdateRoleRequest.md)|  | 

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


## UpdateUser

> map[string]interface{} UpdateUser(ctx, userUserID, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUserID** | **string**|  | 
**body** | [**V1UpdateUserRequest**](V1UpdateUserRequest.md)|  | 

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

