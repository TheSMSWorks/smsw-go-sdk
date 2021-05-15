# {{classname}}

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetError**](UtilsApi.md#GetError) | **Get** /utils/errors/{errorcode} | 
[**Test**](UtilsApi.md#Test) | **Get** /utils/test | 

# **GetError**
> ExtendedErrorModel GetError(ctx, errorcode)


Returns a sample error object for the given error code. Useful for designing code to react to errors when they occur for real.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **errorcode** | **string**| The code of the error you would like returned | 

### Return type

[**ExtendedErrorModel**](ExtendedErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Test**
> TestResponse Test(ctx, )


Returns the customer ID to the caller

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TestResponse**](TestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

