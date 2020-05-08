# \ProductVariantApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProductVariant**](ProductVariantApi.md#AddProductVariant) | **Post** /product_variants | 
[**FindProductVariantById**](ProductVariantApi.md#FindProductVariantById) | **Get** /product_variants/{id} | Find Product variant by ID
[**FindProductVariants**](ProductVariantApi.md#FindProductVariants) | **Get** /product_variants | All product variant


# **AddProductVariant**
> ProductVariant AddProductVariant(ctx, productVariants)


Creates a new product variant in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productVariants** | [**ProductVariantInput**](ProductVariantInput.md)| Product variant to add to the system | 

### Return type

[**ProductVariant**](ProductVariant.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductVariantById**
> ProductVariant FindProductVariantById(ctx, id)
Find Product variant by ID

Returns a single product variant if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of product variant to fetch | 

### Return type

[**ProductVariant**](ProductVariant.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductVariants**
> []ProductVariant FindProductVariants(ctx, optional)
All product variant

Returns all product variant from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindProductVariantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindProductVariantsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]ProductVariant**](ProductVariant.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

