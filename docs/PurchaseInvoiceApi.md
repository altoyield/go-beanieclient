# \PurchaseInvoiceApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPurchaseInvoice**](PurchaseInvoiceApi.md#AddPurchaseInvoice) | **Post** /purchase_invoices | 
[**FindPurchaseInvoiceById**](PurchaseInvoiceApi.md#FindPurchaseInvoiceById) | **Get** /purchase_invoices/{id} | Find Purchase invoice by ID
[**FindPurchaseInvoices**](PurchaseInvoiceApi.md#FindPurchaseInvoices) | **Get** /purchase_invoices | All purchase invoice


# **AddPurchaseInvoice**
> PurchaseInvoice AddPurchaseInvoice(ctx, purchaseInvoices)


Creates a new purchase invoice in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **purchaseInvoices** | [**PurchaseInvoiceInput**](PurchaseInvoiceInput.md)| Purchase invoice to add to the system | 

### Return type

[**PurchaseInvoice**](PurchaseInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPurchaseInvoiceById**
> PurchaseInvoice FindPurchaseInvoiceById(ctx, id)
Find Purchase invoice by ID

Returns a single purchase invoice if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of purchase invoice to fetch | 

### Return type

[**PurchaseInvoice**](PurchaseInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPurchaseInvoices**
> []PurchaseInvoice FindPurchaseInvoices(ctx, optional)
All purchase invoice

Returns all purchase invoice from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindPurchaseInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindPurchaseInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]PurchaseInvoice**](PurchaseInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

