# \BankAccountApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBankAccount**](BankAccountApi.md#AddBankAccount) | **Post** /bank_accounts | 
[**FindBankAccountById**](BankAccountApi.md#FindBankAccountById) | **Get** /bank_accounts/{id} | Find Bank Account by ID
[**FindBankAccounts**](BankAccountApi.md#FindBankAccounts) | **Get** /bank_accounts | All bank accounts


# **AddBankAccount**
> BankAccount AddBankAccount(ctx, bankAccount)


Creates a new bank account in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccount** | [**BankAccountInput**](BankAccountInput.md)| Bank account to add to the system | 

### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBankAccountById**
> BankAccount FindBankAccountById(ctx, id)
Find Bank Account by ID

Returns a single bank account if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of bank account to fetch | 

### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBankAccounts**
> []BankAccount FindBankAccounts(ctx, optional)
All bank accounts

Returns all bank accounts from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindBankAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindBankAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| maximum number of results to return | 

### Return type

[**[]BankAccount**](BankAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

