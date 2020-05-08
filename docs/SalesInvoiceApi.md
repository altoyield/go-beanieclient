# \SalesInvoiceApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSalesInvoice**](SalesInvoiceApi.md#AddSalesInvoice) | **Post** /sales_invoices | 
[**FindSalesInvoiceById**](SalesInvoiceApi.md#FindSalesInvoiceById) | **Get** /sales_invoices/{id} | Find Sales invoice by ID
[**FindSalesInvoices**](SalesInvoiceApi.md#FindSalesInvoices) | **Get** /sales_invoices | All sales invoice


# **AddSalesInvoice**
> SalesInvoice AddSalesInvoice(ctx, salesInvoices)


Creates a new sales invoice in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **salesInvoices** | [**SalesInvoiceInput**](SalesInvoiceInput.md)| Sales invoice to add to the system | 

### Return type

[**SalesInvoice**](SalesInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSalesInvoiceById**
> SalesInvoice FindSalesInvoiceById(ctx, id)
Find Sales invoice by ID

Returns a single sales invoice if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of sales invoice to fetch | 

### Return type

[**SalesInvoice**](SalesInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindSalesInvoices**
> []SalesInvoice FindSalesInvoices(ctx, optional)
All sales invoice

Returns all sales invoice from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindSalesInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindSalesInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]SalesInvoice**](SalesInvoice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

