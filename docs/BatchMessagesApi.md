# \BatchMessagesAPI

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchAnyPost**](BatchMessagesAPI.md#BatchAnyPost) | **Post** /batch/any | 
[**BatchBatchidGet**](BatchMessagesAPI.md#BatchBatchidGet) | **Get** /batch/{batchid} | 
[**BatchSchedulePost**](BatchMessagesAPI.md#BatchSchedulePost) | **Post** /batch/schedule | 
[**BatchSendPost**](BatchMessagesAPI.md#BatchSendPost) | **Post** /batch/send | 
[**BatchesScheduleBatchidDelete**](BatchMessagesAPI.md#BatchesScheduleBatchidDelete) | **Delete** /batches/schedule/{batchid} | 



## BatchAnyPost

> BatchMessageResponse BatchAnyPost(ctx).Messages(messages).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    messages := map[string]interface{}{ ... } // map[string]interface{} | An array of messages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchMessagesAPI.BatchAnyPost(context.Background()).Messages(messages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchMessagesAPI.BatchAnyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchAnyPost`: BatchMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchMessagesAPI.BatchAnyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchAnyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messages** | **map[string]interface{}** | An array of messages | 

### Return type

[**BatchMessageResponse**](BatchMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchBatchidGet

> []MessageResponse BatchBatchidGet(ctx, batchid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    batchid := "batchid_example" // string | The ID of the batch you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchMessagesAPI.BatchBatchidGet(context.Background(), batchid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchMessagesAPI.BatchBatchidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchBatchidGet`: []MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchMessagesAPI.BatchBatchidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchid** | **string** | The ID of the batch you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchBatchidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSchedulePost

> ScheduledBatchResponse BatchSchedulePost(ctx).SmsMessage(smsMessage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smsMessage := *openapiclient.NewBatchMessage("YourCompany", []string{"Destinations_example"}, "My super awesome batch message") // BatchMessage | Message properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchMessagesAPI.BatchSchedulePost(context.Background()).SmsMessage(smsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchMessagesAPI.BatchSchedulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchSchedulePost`: ScheduledBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchMessagesAPI.BatchSchedulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSchedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsMessage** | [**BatchMessage**](BatchMessage.md) | Message properties | 

### Return type

[**ScheduledBatchResponse**](ScheduledBatchResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchSendPost

> BatchMessageResponse BatchSendPost(ctx).SmsMessage(smsMessage).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    smsMessage := *openapiclient.NewBatchMessage("YourCompany", []string{"Destinations_example"}, "My super awesome batch message") // BatchMessage | Message properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchMessagesAPI.BatchSendPost(context.Background()).SmsMessage(smsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchMessagesAPI.BatchSendPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchSendPost`: BatchMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchMessagesAPI.BatchSendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsMessage** | [**BatchMessage**](BatchMessage.md) | Message properties | 

### Return type

[**BatchMessageResponse**](BatchMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchesScheduleBatchidDelete

> CancelledMessageResponse BatchesScheduleBatchidDelete(ctx, batchid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    batchid := "batchid_example" // string | The ID of the batch you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchMessagesAPI.BatchesScheduleBatchidDelete(context.Background(), batchid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchMessagesAPI.BatchesScheduleBatchidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchesScheduleBatchidDelete`: CancelledMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchMessagesAPI.BatchesScheduleBatchidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchid** | **string** | The ID of the batch you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchesScheduleBatchidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CancelledMessageResponse**](CancelledMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

