# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://zestlabs.io](https://zestlabs.io)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthServiceApi* | [**AddUsersToRole**](docs/AuthServiceApi.md#adduserstorole) | **Post** /api/auth/v1/role/users/add | 
*AuthServiceApi* | [**CheckHMACAuth**](docs/AuthServiceApi.md#checkhmacauth) | **Post** /api/auth/v1/check-hmac | Auth API
*AuthServiceApi* | [**CreateAccessKey**](docs/AuthServiceApi.md#createaccesskey) | **Post** /api/auth/v1/accesskey | User Access Keys API
*AuthServiceApi* | [**CreateAccount**](docs/AuthServiceApi.md#createaccount) | **Post** /api/auth/v1/account | Account API
*AuthServiceApi* | [**CreatePolicy**](docs/AuthServiceApi.md#createpolicy) | **Post** /api/auth/v1/policy | 
*AuthServiceApi* | [**CreateRole**](docs/AuthServiceApi.md#createrole) | **Post** /api/auth/v1/role | 
*AuthServiceApi* | [**CreateUser**](docs/AuthServiceApi.md#createuser) | **Post** /api/auth/v1/user | 
*AuthServiceApi* | [**DeleteAccessKey**](docs/AuthServiceApi.md#deleteaccesskey) | **Delete** /api/auth/v1/accesskey/{accessKeyID} | 
*AuthServiceApi* | [**DeleteAccount**](docs/AuthServiceApi.md#deleteaccount) | **Delete** /api/auth/v1/account/{accountID} | 
*AuthServiceApi* | [**DeletePolicy**](docs/AuthServiceApi.md#deletepolicy) | **Delete** /api/auth/v1/policy/{policyID} | 
*AuthServiceApi* | [**DeleteRole**](docs/AuthServiceApi.md#deleterole) | **Delete** /api/auth/v1/role/{roleID} | 
*AuthServiceApi* | [**DeleteUser**](docs/AuthServiceApi.md#deleteuser) | **Delete** /api/auth/v1/user/{userID} | 
*AuthServiceApi* | [**GetAccount**](docs/AuthServiceApi.md#getaccount) | **Get** /api/auth/v1/account/{accountID} | 
*AuthServiceApi* | [**GetDefaultPolicies**](docs/AuthServiceApi.md#getdefaultpolicies) | **Get** /api/auth/v1/default-policies | 
*AuthServiceApi* | [**GetOwnAccount**](docs/AuthServiceApi.md#getownaccount) | **Get** /api/auth/v1/account | 
*AuthServiceApi* | [**GetPasswordPolicy**](docs/AuthServiceApi.md#getpasswordpolicy) | **Get** /api/auth/v1/password-policy | 
*AuthServiceApi* | [**GetPolicies**](docs/AuthServiceApi.md#getpolicies) | **Get** /api/auth/v1/policies | 
*AuthServiceApi* | [**GetPolicy**](docs/AuthServiceApi.md#getpolicy) | **Get** /api/auth/v1/policy/{policyID} | RBAC API
*AuthServiceApi* | [**GetRole**](docs/AuthServiceApi.md#getrole) | **Get** /api/auth/v1/role/{roleID} | 
*AuthServiceApi* | [**GetRoles**](docs/AuthServiceApi.md#getroles) | **Get** /api/auth/v1/roles | 
*AuthServiceApi* | [**GetUser**](docs/AuthServiceApi.md#getuser) | **Get** /api/auth/v1/user/{userID} | 
*AuthServiceApi* | [**GetUserAccessKeys**](docs/AuthServiceApi.md#getuseraccesskeys) | **Get** /api/auth/v1/user/accesskeys/{userID} | 
*AuthServiceApi* | [**GetUserIDByEmail**](docs/AuthServiceApi.md#getuseridbyemail) | **Get** /api/auth/v1/user-id-by-email/{email} | 
*AuthServiceApi* | [**GetUserInfo**](docs/AuthServiceApi.md#getuserinfo) | **Get** /api/auth/v1/userinfo | User API
*AuthServiceApi* | [**GetUsers**](docs/AuthServiceApi.md#getusers) | **Get** /api/auth/v1/users | 
*AuthServiceApi* | [**GetUsersForRole**](docs/AuthServiceApi.md#getusersforrole) | **Get** /api/auth/v1/role/users/{roleID} | 
*AuthServiceApi* | [**RemoveUsersFromRole**](docs/AuthServiceApi.md#removeusersfromrole) | **Post** /api/auth/v1/role/users/remove | 
*AuthServiceApi* | [**ResetPassword**](docs/AuthServiceApi.md#resetpassword) | **Post** /api/auth/v1/reset-password | 
*AuthServiceApi* | [**SetPassword**](docs/AuthServiceApi.md#setpassword) | **Post** /api/auth/v1/set-password | Password API
*AuthServiceApi* | [**UpdateAccount**](docs/AuthServiceApi.md#updateaccount) | **Put** /api/auth/v1/account/{account.accountID} | 
*AuthServiceApi* | [**UpdatePolicy**](docs/AuthServiceApi.md#updatepolicy) | **Put** /api/auth/v1/policy | 
*AuthServiceApi* | [**UpdateRole**](docs/AuthServiceApi.md#updaterole) | **Put** /api/auth/v1/role | 
*AuthServiceApi* | [**UpdateUser**](docs/AuthServiceApi.md#updateuser) | **Put** /api/auth/v1/user/{user.userID} | 
*DistrConfigServiceApi* | [**ActivatePool**](docs/DistrConfigServiceApi.md#activatepool) | **Post** /api/distribution/v1/activate-pool | ActivatePool data pool task
*DistrConfigServiceApi* | [**AssignAppPools**](docs/DistrConfigServiceApi.md#assignapppools) | **Post** /api/distribution/v1/assign-app-pools | Assign data pools to mobile application
*DistrConfigServiceApi* | [**AssignAppUsers**](docs/DistrConfigServiceApi.md#assignappusers) | **Post** /api/distribution/v1/assign-app-users | Assign mobile application to users
*DistrConfigServiceApi* | [**AssignTagToUser**](docs/DistrConfigServiceApi.md#assigntagtouser) | **Post** /api/distribution/v1/assign-tag-to-user | Assign filtering tag to user
*DistrConfigServiceApi* | [**CreateApp**](docs/DistrConfigServiceApi.md#createapp) | **Post** /api/distribution/v1/app | Create mobile application
*DistrConfigServiceApi* | [**CreatePool**](docs/DistrConfigServiceApi.md#createpool) | **Post** /api/distribution/v1/pool | CreatePool new data pool task
*DistrConfigServiceApi* | [**CreatePools**](docs/DistrConfigServiceApi.md#createpools) | **Post** /api/distribution/v1/pools | CreatePool new data pool task
*DistrConfigServiceApi* | [**CreateUser**](docs/DistrConfigServiceApi.md#createuser) | **Post** /api/distribution/v1/user | Registers user in distribution database
*DistrConfigServiceApi* | [**DeleteApp**](docs/DistrConfigServiceApi.md#deleteapp) | **Delete** /api/distribution/v1/app/{id} | Delete mobile application
*DistrConfigServiceApi* | [**DeletePool**](docs/DistrConfigServiceApi.md#deletepool) | **Delete** /api/distribution/v1/pool/{id} | DeletePool data pool task
*DistrConfigServiceApi* | [**DeleteUser**](docs/DistrConfigServiceApi.md#deleteuser) | **Delete** /api/distribution/v1/user/{id} | Deletes user from distribution db
*DistrConfigServiceApi* | [**GetApps**](docs/DistrConfigServiceApi.md#getapps) | **Get** /api/distribution/v1/apps | Get all mobile applications
*DistrConfigServiceApi* | [**GetPool**](docs/DistrConfigServiceApi.md#getpool) | **Get** /api/distribution/v1/pool/{id} | Get data pool task
*DistrConfigServiceApi* | [**GetPoolDistribution**](docs/DistrConfigServiceApi.md#getpooldistribution) | **Get** /api/distribution/v1/get-pool-distribution/{id} | ActivatePool data pool task
*DistrConfigServiceApi* | [**GetPools**](docs/DistrConfigServiceApi.md#getpools) | **Get** /api/distribution/v1/pools | Get all data pool tasks
*DistrConfigServiceApi* | [**GetUsers**](docs/DistrConfigServiceApi.md#getusers) | **Get** /api/distribution/v1/users | Fetch all users in account distribution db
*DistrConfigServiceApi* | [**SetAppBundle**](docs/DistrConfigServiceApi.md#setappbundle) | **Post** /api/distribution/v1/set-app-bundle | Set the mobile application bundle
*DistrConfigServiceApi* | [**UnassignAppPools**](docs/DistrConfigServiceApi.md#unassignapppools) | **Post** /api/distribution/v1/unassign-app-pools | Unassign data pool from mobile application
*DistrConfigServiceApi* | [**UnassignAppUsers**](docs/DistrConfigServiceApi.md#unassignappusers) | **Post** /api/distribution/v1/unassign-app-users | Unassign mobile application from users
*DistrConfigServiceApi* | [**UnassignTagFromUser**](docs/DistrConfigServiceApi.md#unassigntagfromuser) | **Post** /api/distribution/v1/unassign-tag-from-user | Unassign filtering tag from user
*DistrConfigServiceApi* | [**UpdatePool**](docs/DistrConfigServiceApi.md#updatepool) | **Put** /api/distribution/v1/pool | UpdatePool data pool task
*DistrConfigServiceApi* | [**UpdatePool2**](docs/DistrConfigServiceApi.md#updatepool2) | **Patch** /api/distribution/v1/pool/{dataPool.id} | UpdatePool data pool task
*DistrConfigServiceApi* | [**ValidatePoolData**](docs/DistrConfigServiceApi.md#validatepooldata) | **Post** /api/distribution/v1/validate-pool-data | Validate pool data against pool specification
*FunctionsServiceApi* | [**CallAsyncFunction**](docs/FunctionsServiceApi.md#callasyncfunction) | **Post** /api/func/v1/post-call/{reason}/{call.id} | Post function for execution in async way
*FunctionsServiceApi* | [**CallSyncFunction**](docs/FunctionsServiceApi.md#callsyncfunction) | **Post** /api/func/v1/call-sync/{id} | Call a function synchroniously
*FunctionsServiceApi* | [**CallTestFunction**](docs/FunctionsServiceApi.md#calltestfunction) | **Post** /api/func/v1/call-test/{id} | Test a function call
*FunctionsServiceApi* | [**CreateFunction**](docs/FunctionsServiceApi.md#createfunction) | **Post** /api/func/v1/function | Create function
*FunctionsServiceApi* | [**DeleteFunction**](docs/FunctionsServiceApi.md#deletefunction) | **Delete** /api/func/v1/function/{id} | Delete function
*FunctionsServiceApi* | [**GetFunction**](docs/FunctionsServiceApi.md#getfunction) | **Get** /api/func/v1/function/{id} | Get function
*FunctionsServiceApi* | [**GetFunctions**](docs/FunctionsServiceApi.md#getfunctions) | **Get** /api/func/v1/functions | Fetch all functions
*FunctionsServiceApi* | [**UpdateFunction**](docs/FunctionsServiceApi.md#updatefunction) | **Put** /api/func/v1/function | Update function
*FunctionsServiceApi* | [**UpdateFunction2**](docs/FunctionsServiceApi.md#updatefunction2) | **Patch** /api/func/v1/function/{function.id} | Update function
*PoolDataServiceApi* | [**BulkCreate**](docs/PoolDataServiceApi.md#bulkcreate) | **Post** /api/data/_r/{poolId} | Bulk create mobile data records in a pool
*PoolDataServiceApi* | [**BulkDelete**](docs/PoolDataServiceApi.md#bulkdelete) | **Delete** /api/data/_r/{poolId} | Bulk delete records from mobile data pool
*PoolDataServiceApi* | [**BulkUpdate**](docs/PoolDataServiceApi.md#bulkupdate) | **Put** /api/data/_r/{poolId} | Bulk update records in mobile data pool
*PoolDataServiceApi* | [**Get**](docs/PoolDataServiceApi.md#get) | **Get** /api/data/_r/{poolId}/{id} | Get single mobile data record
*PoolDataServiceApi* | [**List**](docs/PoolDataServiceApi.md#list) | **Get** /api/data/_r/{poolId} | List mobile data records


## Documentation For Models

 - [AccountAccountStatus](docs/AccountAccountStatus.md)
 - [DataBulkCreateResponse](docs/DataBulkCreateResponse.md)
 - [DataBulkDeleteResponse](docs/DataBulkDeleteResponse.md)
 - [DataBulkUpdateResponse](docs/DataBulkUpdateResponse.md)
 - [DataDocument](docs/DataDocument.md)
 - [DataListResponse](docs/DataListResponse.md)
 - [DataPersistResponse](docs/DataPersistResponse.md)
 - [DistrconfigActivatePoolRequest](docs/DistrconfigActivatePoolRequest.md)
 - [DistrconfigAssignAppPoolsRequest](docs/DistrconfigAssignAppPoolsRequest.md)
 - [DistrconfigAssignAppUsersRequest](docs/DistrconfigAssignAppUsersRequest.md)
 - [DistrconfigAssignTagToUserRequest](docs/DistrconfigAssignTagToUserRequest.md)
 - [DistrconfigCreatePoolsRequest](docs/DistrconfigCreatePoolsRequest.md)
 - [DistrconfigCreateUserRequest](docs/DistrconfigCreateUserRequest.md)
 - [DistrconfigDataPool](docs/DistrconfigDataPool.md)
 - [DistrconfigDistributionUser](docs/DistrconfigDistributionUser.md)
 - [DistrconfigGetAppsResponse](docs/DistrconfigGetAppsResponse.md)
 - [DistrconfigGetPoolDistributionResponse](docs/DistrconfigGetPoolDistributionResponse.md)
 - [DistrconfigGetPoolResponse](docs/DistrconfigGetPoolResponse.md)
 - [DistrconfigGetPoolsResponse](docs/DistrconfigGetPoolsResponse.md)
 - [DistrconfigGetUsersResponse](docs/DistrconfigGetUsersResponse.md)
 - [DistrconfigMobileApp](docs/DistrconfigMobileApp.md)
 - [DistrconfigMobileAppWithPools](docs/DistrconfigMobileAppWithPools.md)
 - [DistrconfigPoolSize](docs/DistrconfigPoolSize.md)
 - [DistrconfigPoolType](docs/DistrconfigPoolType.md)
 - [DistrconfigSetAppBundleRequest](docs/DistrconfigSetAppBundleRequest.md)
 - [DistrconfigUnassignAppPoolsRequest](docs/DistrconfigUnassignAppPoolsRequest.md)
 - [DistrconfigUnassignAppUsersRequest](docs/DistrconfigUnassignAppUsersRequest.md)
 - [DistrconfigUnassignTagFromUserRequest](docs/DistrconfigUnassignTagFromUserRequest.md)
 - [DistrconfigUserTagAssignment](docs/DistrconfigUserTagAssignment.md)
 - [DistrconfigValidatePoolDataRequest](docs/DistrconfigValidatePoolDataRequest.md)
 - [DistrconfigValidatePoolDataResponse](docs/DistrconfigValidatePoolDataResponse.md)
 - [FunctionsCallAsyncFunctionRequest](docs/FunctionsCallAsyncFunctionRequest.md)
 - [FunctionsCallFunctionRequest](docs/FunctionsCallFunctionRequest.md)
 - [FunctionsCallFunctionResponse](docs/FunctionsCallFunctionResponse.md)
 - [FunctionsCaller](docs/FunctionsCaller.md)
 - [FunctionsCreateFunctionRequest](docs/FunctionsCreateFunctionRequest.md)
 - [FunctionsCreateFunctionResponse](docs/FunctionsCreateFunctionResponse.md)
 - [FunctionsFunction](docs/FunctionsFunction.md)
 - [FunctionsGetFunctionResponse](docs/FunctionsGetFunctionResponse.md)
 - [FunctionsGetFunctionsResponse](docs/FunctionsGetFunctionsResponse.md)
 - [PaymentDetailsPaymentType](docs/PaymentDetailsPaymentType.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [ProtobufNullValue](docs/ProtobufNullValue.md)
 - [RuntimeError](docs/RuntimeError.md)
 - [V1AccessKey](docs/V1AccessKey.md)
 - [V1Account](docs/V1Account.md)
 - [V1Action](docs/V1Action.md)
 - [V1AddUsersToRoleRequest](docs/V1AddUsersToRoleRequest.md)
 - [V1AddUsersToRoleResponse](docs/V1AddUsersToRoleResponse.md)
 - [V1CheckHmacAuthRequest](docs/V1CheckHmacAuthRequest.md)
 - [V1CheckHmacAuthResponse](docs/V1CheckHmacAuthResponse.md)
 - [V1ContactDetails](docs/V1ContactDetails.md)
 - [V1CreateAccessKeyRequest](docs/V1CreateAccessKeyRequest.md)
 - [V1CreateAccessKeyResponse](docs/V1CreateAccessKeyResponse.md)
 - [V1CreateAccountRequest](docs/V1CreateAccountRequest.md)
 - [V1CreateAccountResponse](docs/V1CreateAccountResponse.md)
 - [V1CreatePolicyRequest](docs/V1CreatePolicyRequest.md)
 - [V1CreatePolicyResponse](docs/V1CreatePolicyResponse.md)
 - [V1CreateRoleRequest](docs/V1CreateRoleRequest.md)
 - [V1CreateRoleResponse](docs/V1CreateRoleResponse.md)
 - [V1CreateUserRequest](docs/V1CreateUserRequest.md)
 - [V1CreateUserResponse](docs/V1CreateUserResponse.md)
 - [V1GetAccountResponse](docs/V1GetAccountResponse.md)
 - [V1GetDefaultPoliciesResponse](docs/V1GetDefaultPoliciesResponse.md)
 - [V1GetPasswordPolicyResponse](docs/V1GetPasswordPolicyResponse.md)
 - [V1GetPoliciesResponse](docs/V1GetPoliciesResponse.md)
 - [V1GetPolicyResponse](docs/V1GetPolicyResponse.md)
 - [V1GetRoleResponse](docs/V1GetRoleResponse.md)
 - [V1GetRolesResponse](docs/V1GetRolesResponse.md)
 - [V1GetUserAccessKeysResponse](docs/V1GetUserAccessKeysResponse.md)
 - [V1GetUserIdByEmailResponse](docs/V1GetUserIdByEmailResponse.md)
 - [V1GetUserInfoResponse](docs/V1GetUserInfoResponse.md)
 - [V1GetUserResponse](docs/V1GetUserResponse.md)
 - [V1GetUsersForRoleResponse](docs/V1GetUsersForRoleResponse.md)
 - [V1GetUsersResponse](docs/V1GetUsersResponse.md)
 - [V1PasswordPolicy](docs/V1PasswordPolicy.md)
 - [V1PaymentDetails](docs/V1PaymentDetails.md)
 - [V1Permission](docs/V1Permission.md)
 - [V1Policy](docs/V1Policy.md)
 - [V1RemoveUsersFromRoleRequest](docs/V1RemoveUsersFromRoleRequest.md)
 - [V1RemoveUsersFromRoleResponse](docs/V1RemoveUsersFromRoleResponse.md)
 - [V1ResetPasswordRequest](docs/V1ResetPasswordRequest.md)
 - [V1Role](docs/V1Role.md)
 - [V1SetPasswordRequest](docs/V1SetPasswordRequest.md)
 - [V1UpdateAccountRequest](docs/V1UpdateAccountRequest.md)
 - [V1UpdatePolicyRequest](docs/V1UpdatePolicyRequest.md)
 - [V1UpdateRoleRequest](docs/V1UpdateRoleRequest.md)
 - [V1UpdateUserRequest](docs/V1UpdateUserRequest.md)
 - [V1User](docs/V1User.md)
 - [V1UserInfo](docs/V1UserInfo.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author

contact@zestlabs.io
