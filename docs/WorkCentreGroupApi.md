# \WorkCentreGroupApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindWorkCentreGroupById**](WorkCentreGroupApi.md#FindWorkCentreGroupById) | **Get** /work_centre_groups/{id} | Find Work centre group by ID
[**FindWorkCentreGroups**](WorkCentreGroupApi.md#FindWorkCentreGroups) | **Get** /work_centre_groups | All work centre group


# **FindWorkCentreGroupById**
> WorkCentreGroup FindWorkCentreGroupById(ctx, id)
Find Work centre group by ID

Returns a single work centre group if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of work centre group to fetch | 

### Return type

[**WorkCentreGroup**](WorkCentreGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindWorkCentreGroups**
> []WorkCentreGroup FindWorkCentreGroups(ctx, optional)
All work centre group

Returns all work centre group from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindWorkCentreGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindWorkCentreGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]WorkCentreGroup**](WorkCentreGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

