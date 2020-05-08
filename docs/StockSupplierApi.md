# \StockSupplierApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddStockSupplier**](StockSupplierApi.md#AddStockSupplier) | **Post** /stock_suppliers | 
[**FindStockSupplierById**](StockSupplierApi.md#FindStockSupplierById) | **Get** /stock_suppliers/{id} | Find Stock supplier by ID
[**FindStockSuppliers**](StockSupplierApi.md#FindStockSuppliers) | **Get** /stock_suppliers | All stock supplier


# **AddStockSupplier**
> StockSupplier AddStockSupplier(ctx, stockSuppliers)


Creates a new stock supplier in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stockSuppliers** | [**StockSupplierInput**](StockSupplierInput.md)| Stock supplier to add to the system | 

### Return type

[**StockSupplier**](StockSupplier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindStockSupplierById**
> StockSupplier FindStockSupplierById(ctx, id)
Find Stock supplier by ID

Returns a single stock supplier if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of stock supplier to fetch | 

### Return type

[**StockSupplier**](StockSupplier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindStockSuppliers**
> []StockSupplier FindStockSuppliers(ctx, optional)
All stock supplier

Returns all stock supplier from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindStockSuppliersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindStockSuppliersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]StockSupplier**](StockSupplier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

