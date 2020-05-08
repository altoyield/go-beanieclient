# \ProductPriceApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProductPrice**](ProductPriceApi.md#AddProductPrice) | **Post** /product_prices | 
[**FindProductPriceById**](ProductPriceApi.md#FindProductPriceById) | **Get** /product_prices/{id} | Find Product price by ID
[**FindProductPrices**](ProductPriceApi.md#FindProductPrices) | **Get** /product_prices | All product price


# **AddProductPrice**
> ProductPrice AddProductPrice(ctx, productPrices)


Creates a new product price in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productPrices** | [**ProductPriceInput**](ProductPriceInput.md)| Product price to add to the system | 

### Return type

[**ProductPrice**](ProductPrice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductPriceById**
> ProductPrice FindProductPriceById(ctx, id)
Find Product price by ID

Returns a single product price if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of product price to fetch | 

### Return type

[**ProductPrice**](ProductPrice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductPrices**
> []ProductPrice FindProductPrices(ctx, optional)
All product price

Returns all product price from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindProductPricesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindProductPricesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]ProductPrice**](ProductPrice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

