# \PartnerAddressApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartnerAddress**](PartnerAddressApi.md#AddPartnerAddress) | **Post** /partner_addresses | 
[**FindPartnerAddressById**](PartnerAddressApi.md#FindPartnerAddressById) | **Get** /partner_addresses/{id} | Find Partner address by ID
[**FindPartnerAddresses**](PartnerAddressApi.md#FindPartnerAddresses) | **Get** /partner_addresses | All partner address


# **AddPartnerAddress**
> PartnerAddress AddPartnerAddress(ctx, partnerAddresses)


Creates a new partner address in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerAddresses** | [**PartnerAddressInput**](PartnerAddressInput.md)| Partner address to add to the system | 

### Return type

[**PartnerAddress**](PartnerAddress.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPartnerAddressById**
> PartnerAddress FindPartnerAddressById(ctx, id)
Find Partner address by ID

Returns a single partner address if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of partner address to fetch | 

### Return type

[**PartnerAddress**](PartnerAddress.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindPartnerAddresses**
> []PartnerAddress FindPartnerAddresses(ctx, optional)
All partner address

Returns all partner address from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindPartnerAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindPartnerAddressesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]PartnerAddress**](PartnerAddress.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

