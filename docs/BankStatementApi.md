# \BankStatementApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBankStatement**](BankStatementApi.md#AddBankStatement) | **Post** /bank_statements | 
[**FindBankStatementById**](BankStatementApi.md#FindBankStatementById) | **Get** /bank_statements/{id} | Find Bank statement by ID
[**FindBankStatements**](BankStatementApi.md#FindBankStatements) | **Get** /bank_statements | All bank statement


# **AddBankStatement**
> BankStatement AddBankStatement(ctx, bankStatements)


Creates a new bank statement in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankStatements** | [**BankStatementInput**](BankStatementInput.md)| Bank statement to add to the system | 

### Return type

[**BankStatement**](BankStatement.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBankStatementById**
> BankStatement FindBankStatementById(ctx, id)
Find Bank statement by ID

Returns a single bank statement if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of bank statement to fetch | 

### Return type

[**BankStatement**](BankStatement.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBankStatements**
> []BankStatement FindBankStatements(ctx, optional)
All bank statement

Returns all bank statement from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindBankStatementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindBankStatementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]BankStatement**](BankStatement.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

