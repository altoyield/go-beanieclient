# \WorkCentreApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindWorkCentreById**](WorkCentreApi.md#FindWorkCentreById) | **Get** /work_centres/{id} | Find Work centre by ID
[**FindWorkCentres**](WorkCentreApi.md#FindWorkCentres) | **Get** /work_centre_groups/{work_centre_group_id}/work_centres | All work centre


# **FindWorkCentreById**
> WorkCentre FindWorkCentreById(ctx, id)
Find Work centre by ID

Returns a single work centre if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of work centre to fetch | 

### Return type

[**WorkCentre**](WorkCentre.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindWorkCentres**
> []WorkCentre FindWorkCentres(ctx, workCentreGroupId, optional)
All work centre

Returns all work centre from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workCentreGroupId** | **int64**| ID of Work Centre Group for list of Work Centres | 
 **optional** | ***FindWorkCentresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindWorkCentresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]WorkCentre**](WorkCentre.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

