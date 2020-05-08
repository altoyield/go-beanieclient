# \CompanyApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindCompanyById**](CompanyApi.md#FindCompanyById) | **Get** /companies/{id} | Find Company details by ID


# **FindCompanyById**
> Company FindCompanyById(ctx, id)
Find Company details by ID

Returns a single company details if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of company details to fetch | 

### Return type

[**Company**](Company.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

