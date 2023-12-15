# \UtilsAPI

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UtilsErrorsErrorcodeGet**](UtilsAPI.md#UtilsErrorsErrorcodeGet) | **Get** /utils/errors/{errorcode} | 
[**UtilsTestGet**](UtilsAPI.md#UtilsTestGet) | **Get** /utils/test | 



## UtilsErrorsErrorcodeGet

> ExtendedErrorModel UtilsErrorsErrorcodeGet(ctx, errorcode).Execute()





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
    errorcode := "errorcode_example" // string | The code of the error you would like returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UtilsAPI.UtilsErrorsErrorcodeGet(context.Background(), errorcode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.UtilsErrorsErrorcodeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtilsErrorsErrorcodeGet`: ExtendedErrorModel
    fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.UtilsErrorsErrorcodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**errorcode** | **string** | The code of the error you would like returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiUtilsErrorsErrorcodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtendedErrorModel**](ExtendedErrorModel.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UtilsTestGet

> TestResponse UtilsTestGet(ctx).Execute()





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
    resp, r, err := apiClient.UtilsAPI.UtilsTestGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilsAPI.UtilsTestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UtilsTestGet`: TestResponse
    fmt.Fprintf(os.Stdout, "Response from `UtilsAPI.UtilsTestGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUtilsTestGetRequest struct via the builder pattern


### Return type

[**TestResponse**](TestResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

