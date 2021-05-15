# {{classname}}

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelScheduledBatchJob**](BatchMessagesApi.md#CancelScheduledBatchJob) | **Delete** /batches/schedule/{batchid} | 
[**GetBatchById**](BatchMessagesApi.md#GetBatchById) | **Get** /batch/{batchid} | 
[**ScheduleBatch**](BatchMessagesApi.md#ScheduleBatch) | **Post** /batch/schedule | 
[**SendBatch**](BatchMessagesApi.md#SendBatch) | **Post** /batch/send | 

# **CancelScheduledBatchJob**
> CancelledMessageResponse CancelScheduledBatchJob(ctx, batchid)


Cancels a scheduled SMS message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchid** | **string**| The ID of the batch you would like returned | 

### Return type

[**CancelledMessageResponse**](CancelledMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchById**
> []MessageResponse GetBatchById(ctx, batchid)


Retrieve all messages in a batch with the given batch ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchid** | **string**| The ID of the batch you would like returned | 

### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleBatch**
> ScheduledBatchResponse ScheduleBatch(ctx, body)


Schedules a batch of SMS messages to be sent at the date time you specify

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchMessage**](BatchMessage.md)| Message properties | 

### Return type

[**ScheduledBatchResponse**](ScheduledBatchResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendBatch**
> BatchMessageResponse SendBatch(ctx, body)


Send a single SMS message to multiple recipients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchMessage**](BatchMessage.md)| Message properties | 

### Return type

[**BatchMessageResponse**](BatchMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

