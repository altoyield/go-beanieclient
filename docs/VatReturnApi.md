# \VatReturnApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVatReturn**](VatReturnApi.md#AddVatReturn) | **Post** /vat_returns | 
[**FindVatReturnById**](VatReturnApi.md#FindVatReturnById) | **Get** /vat_returns/{id} | Find VAT return by ID
[**FindVatReturns**](VatReturnApi.md#FindVatReturns) | **Get** /vat_returns | All vat return


# **AddVatReturn**
> VatReturn AddVatReturn(ctx, vatReturns)


Creates a new vat return in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vatReturns** | [**VatReturnInput**](VatReturnInput.md)| VAT return to add to the system | 

### Return type

[**VatReturn**](VatReturn.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVatReturnById**
> VatReturn FindVatReturnById(ctx, id)
Find VAT return by ID

Returns a single vat return if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of vat return to fetch | 

### Return type

[**VatReturn**](VatReturn.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindVatReturns**
> []VatReturn FindVatReturns(ctx, optional)
All vat return

Returns all vat return from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindVatReturnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindVatReturnsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]VatReturn**](VatReturn.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

