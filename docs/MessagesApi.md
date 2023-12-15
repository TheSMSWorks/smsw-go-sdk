# \MessagesAPI

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessageSchedulePost**](MessagesAPI.md#MessageSchedulePost) | **Post** /message/schedule | 
[**MessageSendPost**](MessagesAPI.md#MessageSendPost) | **Post** /message/send | 
[**MessagesFailedPost**](MessagesAPI.md#MessagesFailedPost) | **Post** /messages/failed | 
[**MessagesInboxPost**](MessagesAPI.md#MessagesInboxPost) | **Post** /messages/inbox | 
[**MessagesMessageidDelete**](MessagesAPI.md#MessagesMessageidDelete) | **Delete** /messages/{messageid} | 
[**MessagesMessageidGet**](MessagesAPI.md#MessagesMessageidGet) | **Get** /messages/{messageid} | 
[**MessagesPost**](MessagesAPI.md#MessagesPost) | **Post** /messages | 
[**MessagesScheduleGet**](MessagesAPI.md#MessagesScheduleGet) | **Get** /messages/schedule | 
[**MessagesScheduleMessageidDelete**](MessagesAPI.md#MessagesScheduleMessageidDelete) | **Delete** /messages/schedule/{messageid} | 
[**SendFlashMessage**](MessagesAPI.md#SendFlashMessage) | **Post** /message/flash | 



## MessageSchedulePost

> []ScheduledMessageResponse MessageSchedulePost(ctx).SmsMessage(smsMessage).Execute()





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
    smsMessage := *openapiclient.NewMessage("YourCompany", "447777777777", "Your super awesome message") // Message | Message properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessageSchedulePost(context.Background()).SmsMessage(smsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessageSchedulePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessageSchedulePost`: []ScheduledMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessageSchedulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessageSchedulePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsMessage** | [**Message**](Message.md) | Message properties | 

### Return type

[**[]ScheduledMessageResponse**](ScheduledMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessageSendPost

> SendMessageResponse MessageSendPost(ctx).SmsMessage(smsMessage).Execute()





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
    smsMessage := *openapiclient.NewMessage("YourCompany", "447777777777", "Your super awesome message") // Message | Message properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessageSendPost(context.Background()).SmsMessage(smsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessageSendPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessageSendPost`: SendMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessageSendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessageSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsMessage** | [**Message**](Message.md) | Message properties | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesFailedPost

> []MessageResponse MessagesFailedPost(ctx).Query(query).Execute()





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
    query := *openapiclient.NewQuery() // Query | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesFailedPost(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesFailedPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesFailedPost`: []MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesFailedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesFailedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesInboxPost

> []MessageResponse MessagesInboxPost(ctx).Query(query).Execute()





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
    query := *openapiclient.NewQuery() // Query | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesInboxPost(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesInboxPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesInboxPost`: []MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesInboxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesInboxPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesMessageidDelete

> DeletedMessageResponse MessagesMessageidDelete(ctx, messageid).Execute()





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
    messageid := "messageid_example" // string | The ID of the message you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesMessageidDelete(context.Background(), messageid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesMessageidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesMessageidDelete`: DeletedMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesMessageidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageid** | **string** | The ID of the message you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiMessagesMessageidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedMessageResponse**](DeletedMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesMessageidGet

> MessageResponse MessagesMessageidGet(ctx, messageid).Execute()





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
    messageid := "messageid_example" // string | The ID of the message you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesMessageidGet(context.Background(), messageid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesMessageidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesMessageidGet`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesMessageidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageid** | **string** | The ID of the message you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiMessagesMessageidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesPost

> []MessageResponse MessagesPost(ctx).Query(query).Execute()





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
    query := *openapiclient.NewQuery() // Query | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesPost(context.Background()).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesPost`: []MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**Query**](Query.md) |  | 

### Return type

[**[]MessageResponse**](MessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesScheduleGet

> ScheduledMessagesResponse MessagesScheduleGet(ctx).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesScheduleGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesScheduleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesScheduleGet`: ScheduledMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesScheduleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMessagesScheduleGetRequest struct via the builder pattern


### Return type

[**ScheduledMessagesResponse**](ScheduledMessagesResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesScheduleMessageidDelete

> CancelledMessageResponse MessagesScheduleMessageidDelete(ctx, messageid).Execute()





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
    messageid := "messageid_example" // string | The ID of the message you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.MessagesScheduleMessageidDelete(context.Background(), messageid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.MessagesScheduleMessageidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesScheduleMessageidDelete`: CancelledMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.MessagesScheduleMessageidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageid** | **string** | The ID of the message you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiMessagesScheduleMessageidDeleteRequest struct via the builder pattern


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


## SendFlashMessage

> SendMessageResponse SendFlashMessage(ctx).SmsMessage(smsMessage).Execute()





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
    smsMessage := *openapiclient.NewMessage("YourCompany", "447777777777", "Your super awesome message") // Message | Message properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.SendFlashMessage(context.Background()).SmsMessage(smsMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendFlashMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendFlashMessage`: SendMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendFlashMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendFlashMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsMessage** | [**Message**](Message.md) | Message properties | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

