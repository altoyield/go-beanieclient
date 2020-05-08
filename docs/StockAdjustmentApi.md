# \StockAdjustmentApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStockAdjustment**](StockAdjustmentApi.md#AddStockAdjustment) | **Post** /stock_adjustments | 
[**FindStockAdjustmentById**](StockAdjustmentApi.md#FindStockAdjustmentById) | **Get** /stock_adjustments/{id} | Find Stock adjustment by ID
[**FindStockAdjustments**](StockAdjustmentApi.md#FindStockAdjustments) | **Get** /stock_adjustments | All stock adjustment


# **AddStockAdjustment**
> StockAdjustment AddStockAdjustment(ctx, stockAdjustments)


Creates a new stock adjustment in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stockAdjustments** | [**StockAdjustmentInput**](StockAdjustmentInput.md)| Stock adjustment to add to the system | 

### Return type

[**StockAdjustment**](StockAdjustment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindStockAdjustmentById**
> StockAdjustment FindStockAdjustmentById(ctx, id)
Find Stock adjustment by ID

Returns a single stock adjustment if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of stock adjustment to fetch | 

### Return type

[**StockAdjustment**](StockAdjustment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindStockAdjustments**
> []StockAdjustment FindStockAdjustments(ctx, optional)
All stock adjustment

Returns all stock adjustment from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindStockAdjustmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindStockAdjustmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]StockAdjustment**](StockAdjustment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

