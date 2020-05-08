# \NominalAccountCategoryApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNominalAccountCategory**](NominalAccountCategoryApi.md#AddNominalAccountCategory) | **Post** /nominal_account_categories | 
[**FindNominalAccountCategories**](NominalAccountCategoryApi.md#FindNominalAccountCategories) | **Get** /nominal_account_categories | All nominal account category
[**FindNominalAccountCategoryById**](NominalAccountCategoryApi.md#FindNominalAccountCategoryById) | **Get** /nominal_account_categories/{id} | Find Nominal account category by ID


# **AddNominalAccountCategory**
> NominalAccountCategory AddNominalAccountCategory(ctx, nominalAccountCategories)


Creates a new nominal account category in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nominalAccountCategories** | [**NominalAccountCategoryInput**](NominalAccountCategoryInput.md)| Nominal account category to add to the system | 

### Return type

[**NominalAccountCategory**](NominalAccountCategory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNominalAccountCategories**
> []NominalAccountCategory FindNominalAccountCategories(ctx, optional)
All nominal account category

Returns all nominal account category from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindNominalAccountCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindNominalAccountCategoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]NominalAccountCategory**](NominalAccountCategory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindNominalAccountCategoryById**
> NominalAccountCategory FindNominalAccountCategoryById(ctx, id)
Find Nominal account category by ID

Returns a single nominal account category if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of nominal account category to fetch | 

### Return type

[**NominalAccountCategory**](NominalAccountCategory.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

