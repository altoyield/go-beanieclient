# \ShippingCentreApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddShippingCentre**](ShippingCentreApi.md#AddShippingCentre) | **Post** /shipping_centres | 
[**FindShippingCentreById**](ShippingCentreApi.md#FindShippingCentreById) | **Get** /shipping_centres/{id} | Find Shipping centre by ID
[**FindShippingCentres**](ShippingCentreApi.md#FindShippingCentres) | **Get** /shipping_centres | All shipping centre


# **AddShippingCentre**
> ShippingCentre AddShippingCentre(ctx, shippingCentres)


Creates a new shipping centre in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shippingCentres** | [**ShippingCentreInput**](ShippingCentreInput.md)| Shipping centre to add to the system | 

### Return type

[**ShippingCentre**](ShippingCentre.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindShippingCentreById**
> ShippingCentre FindShippingCentreById(ctx, id)
Find Shipping centre by ID

Returns a single shipping centre if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of shipping centre to fetch | 

### Return type

[**ShippingCentre**](ShippingCentre.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindShippingCentres**
> []ShippingCentre FindShippingCentres(ctx, optional)
All shipping centre

Returns all shipping centre from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindShippingCentresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindShippingCentresOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]ShippingCentre**](ShippingCentre.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

