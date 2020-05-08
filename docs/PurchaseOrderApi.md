# \PurchaseOrderApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPurchaseOrder**](PurchaseOrderApi.md#AddPurchaseOrder) | **Post** /purchase_orders | 
[**FindPurchaseOrderById**](PurchaseOrderApi.md#FindPurchaseOrderById) | **Get** /purchase_orders/{id} | Find Purchase order by ID
[**FindPurchaseOrders**](PurchaseOrderApi.md#FindPurchaseOrders) | **Get** /purchase_orders | All purchase order


# **AddPurchaseOrder**
> PurchaseOrder AddPurchaseOrder(ctx, purchaseOrders)


Creates a new purchase order in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseOrders** | [**PurchaseOrderInput**](PurchaseOrderInput.md)| Purchase order to add to the system | 

### Return type

[**PurchaseOrder**](PurchaseOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPurchaseOrderById**
> PurchaseOrder FindPurchaseOrderById(ctx, id)
Find Purchase order by ID

Returns a single purchase order if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of purchase order to fetch | 

### Return type

[**PurchaseOrder**](PurchaseOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPurchaseOrders**
> []PurchaseOrder FindPurchaseOrders(ctx, optional)
All purchase order

Returns all purchase order from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindPurchaseOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindPurchaseOrdersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]PurchaseOrder**](PurchaseOrder.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

