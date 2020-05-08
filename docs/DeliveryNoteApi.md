# \DeliveryNoteApi

All URIs are relative to *https://bean.ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDeliveryNote**](DeliveryNoteApi.md#AddDeliveryNote) | **Post** /delivery_notes | 
[**FindDeliveryNoteById**](DeliveryNoteApi.md#FindDeliveryNoteById) | **Get** /delivery_notes/{id} | Find Delivery note by ID
[**FindDeliveryNotes**](DeliveryNoteApi.md#FindDeliveryNotes) | **Get** /delivery_notes | All delivery note


# **AddDeliveryNote**
> DeliveryNote AddDeliveryNote(ctx, deliveryNotes)


Creates a new delivery note in the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deliveryNotes** | [**DeliveryNoteInput**](DeliveryNoteInput.md)| Delivery note to add to the system | 

### Return type

[**DeliveryNote**](DeliveryNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDeliveryNoteById**
> DeliveryNote FindDeliveryNoteById(ctx, id)
Find Delivery note by ID

Returns a single delivery note if the user has access

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of delivery note to fetch | 

### Return type

[**DeliveryNote**](DeliveryNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDeliveryNotes**
> []DeliveryNote FindDeliveryNotes(ctx, optional)
All delivery note

Returns all delivery note from the system that the user has access to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindDeliveryNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FindDeliveryNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| tags to filter by | 
 **limit** | **optional.Int32**| Maximum number of results to return | 

### Return type

[**[]DeliveryNote**](DeliveryNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

