# \BillOfMaterialsApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBillOfMaterial**](BillOfMaterialsApi.md#AddBillOfMaterial) | **Post** /bill_of_materials | 
[**FindBillOfMaterialById**](BillOfMaterialsApi.md#FindBillOfMaterialById) | **Get** /bill_of_materials/{id} | Find Bill of Materials by ID
[**FindBillOfMaterials**](BillOfMaterialsApi.md#FindBillOfMaterials) | **Get** /bill_of_materials | All bill of materials


# **AddBillOfMaterial**
> BillOfMaterial AddBillOfMaterial(ctx, billOfMaterials)


Creates a new bill of materials in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **billOfMaterials** | [**BillOfMaterialInput**](BillOfMaterialInput.md)| Bill of Materials to add to the system | 

### Return type

[**BillOfMaterial**](BillOfMaterial.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBillOfMaterialById**
> BillOfMaterial FindBillOfMaterialById(ctx, id)
Find Bill of Materials by ID

Returns a single bill of materials if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of bill of materials to fetch | 

### Return type

[**BillOfMaterial**](BillOfMaterial.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindBillOfMaterials**
> []BillOfMaterial FindBillOfMaterials(ctx, optional)
All bill of materials

Returns all bill of materials from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindBillOfMaterialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindBillOfMaterialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]BillOfMaterial**](BillOfMaterial.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

