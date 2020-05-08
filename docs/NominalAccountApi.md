# \NominalAccountApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNominalAccount**](NominalAccountApi.md#AddNominalAccount) | **Post** /nominal_accounts | 
[**FindNominalAccountById**](NominalAccountApi.md#FindNominalAccountById) | **Get** /nominal_accounts/{id} | Find Nominal account by ID
[**FindNominalAccounts**](NominalAccountApi.md#FindNominalAccounts) | **Get** /nominal_accounts | All nominal account


# **AddNominalAccount**
> NominalAccount AddNominalAccount(ctx, nominalAccounts)


Creates a new nominal account in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nominalAccounts** | [**NominalAccountInput**](NominalAccountInput.md)| Nominal account to add to the system | 

### Return type

[**NominalAccount**](NominalAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNominalAccountById**
> NominalAccount FindNominalAccountById(ctx, id)
Find Nominal account by ID

Returns a single nominal account if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of nominal account to fetch | 

### Return type

[**NominalAccount**](NominalAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNominalAccounts**
> []NominalAccount FindNominalAccounts(ctx, optional)
All nominal account

Returns all nominal account from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindNominalAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindNominalAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]NominalAccount**](NominalAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

