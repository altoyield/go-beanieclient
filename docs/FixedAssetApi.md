# \FixedAssetApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFixedAsset**](FixedAssetApi.md#AddFixedAsset) | **Post** /fixed_assets | 
[**FindFixedAssetById**](FixedAssetApi.md#FindFixedAssetById) | **Get** /fixed_assets/{id} | Find Fixed asset by ID
[**FindFixedAssets**](FixedAssetApi.md#FindFixedAssets) | **Get** /fixed_assets | All fixed asset


# **AddFixedAsset**
> FixedAsset AddFixedAsset(ctx, fixedAssets)


Creates a new fixed asset in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fixedAssets** | [**FixedAssetInput**](FixedAssetInput.md)| Fixed asset to add to the system | 

### Return type

[**FixedAsset**](FixedAsset.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFixedAssetById**
> FixedAsset FindFixedAssetById(ctx, id)
Find Fixed asset by ID

Returns a single fixed asset if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of fixed asset to fetch | 

### Return type

[**FixedAsset**](FixedAsset.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFixedAssets**
> []FixedAsset FindFixedAssets(ctx, optional)
All fixed asset

Returns all fixed asset from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindFixedAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindFixedAssetsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]FixedAsset**](FixedAsset.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

