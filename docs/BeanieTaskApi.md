# \BeanieTaskApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBeanieTask**](BeanieTaskApi.md#AddBeanieTask) | **Post** /beanie_tasks | 
[**FindBeanieTaskById**](BeanieTaskApi.md#FindBeanieTaskById) | **Get** /beanie_tasks/{id} | Find Beanie task by ID
[**FindBeanieTasks**](BeanieTaskApi.md#FindBeanieTasks) | **Get** /beanie_tasks | All beanie task


# **AddBeanieTask**
> BeanieTask AddBeanieTask(ctx, beanieTasks)


Creates a new beanie task in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **beanieTasks** | [**BeanieTaskInput**](BeanieTaskInput.md)| Beanie task to add to the system | 

### Return type

[**BeanieTask**](BeanieTask.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBeanieTaskById**
> BeanieTask FindBeanieTaskById(ctx, id)
Find Beanie task by ID

Returns a single beanie task if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of beanie task to fetch | 

### Return type

[**BeanieTask**](BeanieTask.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBeanieTasks**
> []BeanieTask FindBeanieTasks(ctx, optional)
All beanie task

Returns all beanie task from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindBeanieTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindBeanieTasksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]BeanieTask**](BeanieTask.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

