# \SalesOrderApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSalesOrder**](SalesOrderApi.md#AddSalesOrder) | **Post** /sales_orders | 
[**FindSalesOrder**](SalesOrderApi.md#FindSalesOrder) | **Get** /sales_orders | All sales order
[**FindSalesOrderById**](SalesOrderApi.md#FindSalesOrderById) | **Get** /sales_orders/{id} | Find Sales order by ID


# **AddSalesOrder**
> SalesOrder AddSalesOrder(ctx, salesOrders)


Creates a new sales order in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **salesOrders** | [**SalesOrderInput**](SalesOrderInput.md)| Sales order to add to the system | 

### Return type

[**SalesOrder**](SalesOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSalesOrder**
> []SalesOrder FindSalesOrder(ctx, optional)
All sales order

Returns all sales order from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindSalesOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindSalesOrderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]SalesOrder**](SalesOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSalesOrderById**
> SalesOrder FindSalesOrderById(ctx, id)
Find Sales order by ID

Returns a single sales order if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of sales order to fetch | 

### Return type

[**SalesOrder**](SalesOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

