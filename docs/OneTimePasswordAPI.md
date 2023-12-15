# \OneTimePasswordAPI

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OtpMessageidGet**](OneTimePasswordAPI.md#OtpMessageidGet) | **Get** /otp/{messageid} | 
[**OtpSendPost**](OneTimePasswordAPI.md#OtpSendPost) | **Post** /otp/send | 
[**OtpVerifyPost**](OneTimePasswordAPI.md#OtpVerifyPost) | **Post** /otp/verify | 



## OtpMessageidGet

> OTPVerifyResponse OtpMessageidGet(ctx, messageid).Execute()





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
    messageid := "messageid_example" // string | The ID of the OTP you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneTimePasswordAPI.OtpMessageidGet(context.Background(), messageid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneTimePasswordAPI.OtpMessageidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OtpMessageidGet`: OTPVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `OneTimePasswordAPI.OtpMessageidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageid** | **string** | The ID of the OTP you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiOtpMessageidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OTPVerifyResponse**](OTPVerifyResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtpSendPost

> OTPResponse OtpSendPost(ctx).Otp(otp).Execute()





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
    otp := *openapiclient.NewOTP() // OTP | OTP properties

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneTimePasswordAPI.OtpSendPost(context.Background()).Otp(otp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneTimePasswordAPI.OtpSendPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OtpSendPost`: OTPResponse
    fmt.Fprintf(os.Stdout, "Response from `OneTimePasswordAPI.OtpSendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOtpSendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **otp** | [**OTP**](OTP.md) | OTP properties | 

### Return type

[**OTPResponse**](OTPResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OtpVerifyPost

> OTPVerifyResponse OtpVerifyPost(ctx).Passcode(passcode).Execute()





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
    passcode := *openapiclient.NewOTPVerify() // OTPVerify | One-Time Password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneTimePasswordAPI.OtpVerifyPost(context.Background()).Passcode(passcode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneTimePasswordAPI.OtpVerifyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OtpVerifyPost`: OTPVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `OneTimePasswordAPI.OtpVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOtpVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passcode** | [**OTPVerify**](OTPVerify.md) | One-Time Password | 

### Return type

[**OTPVerifyResponse**](OTPVerifyResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

