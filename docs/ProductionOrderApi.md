# \ProductionOrderApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProductionOrder**](ProductionOrderApi.md#AddProductionOrder) | **Post** /production_orders | 
[**FindProductionOrderById**](ProductionOrderApi.md#FindProductionOrderById) | **Get** /production_orders/{id} | Find Production order by ID
[**FindProductionOrders**](ProductionOrderApi.md#FindProductionOrders) | **Get** /production_orders | All production order


# **AddProductionOrder**
> ProductionOrder AddProductionOrder(ctx, productionOrders)


Creates a new production order in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productionOrders** | [**ProductionOrderInput**](ProductionOrderInput.md)| Production order to add to the system | 

### Return type

[**ProductionOrder**](ProductionOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductionOrderById**
> ProductionOrder FindProductionOrderById(ctx, id)
Find Production order by ID

Returns a single production order if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of production order to fetch | 

### Return type

[**ProductionOrder**](ProductionOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProductionOrders**
> []ProductionOrder FindProductionOrders(ctx, optional)
All production order

Returns all production order from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindProductionOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindProductionOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]ProductionOrder**](ProductionOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

