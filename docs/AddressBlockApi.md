# \AddressBlockApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAddressBlock**](AddressBlockApi.md#AddAddressBlock) | **Post** /address_blocks | 
[**FindAddressBlockById**](AddressBlockApi.md#FindAddressBlockById) | **Get** /address_blocks/{id} | Find Address block by ID
[**FindAddressBlocks**](AddressBlockApi.md#FindAddressBlocks) | **Get** /address_blocks | All address block


# **AddAddressBlock**
> AddressBlock AddAddressBlock(ctx, addressBlocks)


Creates a new address block in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addressBlocks** | [**AddressBlockInput**](AddressBlockInput.md)| Address block to add to the system | 

### Return type

[**AddressBlock**](AddressBlock.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAddressBlockById**
> AddressBlock FindAddressBlockById(ctx, id)
Find Address block by ID

Returns a single address block if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of address block to fetch | 

### Return type

[**AddressBlock**](AddressBlock.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAddressBlocks**
> []AddressBlock FindAddressBlocks(ctx, optional)
All address block

Returns all address block from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindAddressBlocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindAddressBlocksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]AddressBlock**](AddressBlock.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

