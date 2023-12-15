# Go API client for openapi

The SMS Works provides a low-cost, reliable SMS API for developers. Pay only for delivered texts, all failed UK messages are refunded.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.9.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://thesmsworks.co.uk/contact](https://thesmsworks.co.uk/contact)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.thesmsworks.co.uk/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BatchMessagesAPI* | [**BatchAnyPost**](docs/BatchMessagesAPI.md#batchanypost) | **Post** /batch/any | 
*BatchMessagesAPI* | [**BatchBatchidGet**](docs/BatchMessagesAPI.md#batchbatchidget) | **Get** /batch/{batchid} | 
*BatchMessagesAPI* | [**BatchSchedulePost**](docs/BatchMessagesAPI.md#batchschedulepost) | **Post** /batch/schedule | 
*BatchMessagesAPI* | [**BatchSendPost**](docs/BatchMessagesAPI.md#batchsendpost) | **Post** /batch/send | 
*BatchMessagesAPI* | [**BatchesScheduleBatchidDelete**](docs/BatchMessagesAPI.md#batchesschedulebatchiddelete) | **Delete** /batches/schedule/{batchid} | 
*CreditsAPI* | [**CreditsBalanceGet**](docs/CreditsAPI.md#creditsbalanceget) | **Get** /credits/balance | 
*MessagesAPI* | [**MessageSchedulePost**](docs/MessagesAPI.md#messageschedulepost) | **Post** /message/schedule | 
*MessagesAPI* | [**MessageSendPost**](docs/MessagesAPI.md#messagesendpost) | **Post** /message/send | 
*MessagesAPI* | [**MessagesFailedPost**](docs/MessagesAPI.md#messagesfailedpost) | **Post** /messages/failed | 
*MessagesAPI* | [**MessagesInboxPost**](docs/MessagesAPI.md#messagesinboxpost) | **Post** /messages/inbox | 
*MessagesAPI* | [**MessagesMessageidDelete**](docs/MessagesAPI.md#messagesmessageiddelete) | **Delete** /messages/{messageid} | 
*MessagesAPI* | [**MessagesMessageidGet**](docs/MessagesAPI.md#messagesmessageidget) | **Get** /messages/{messageid} | 
*MessagesAPI* | [**MessagesPost**](docs/MessagesAPI.md#messagespost) | **Post** /messages | 
*MessagesAPI* | [**MessagesScheduleGet**](docs/MessagesAPI.md#messagesscheduleget) | **Get** /messages/schedule | 
*MessagesAPI* | [**MessagesScheduleMessageidDelete**](docs/MessagesAPI.md#messagesschedulemessageiddelete) | **Delete** /messages/schedule/{messageid} | 
*MessagesAPI* | [**SendFlashMessage**](docs/MessagesAPI.md#sendflashmessage) | **Post** /message/flash | 
*OneTimePasswordAPI* | [**OtpMessageidGet**](docs/OneTimePasswordAPI.md#otpmessageidget) | **Get** /otp/{messageid} | 
*OneTimePasswordAPI* | [**OtpSendPost**](docs/OneTimePasswordAPI.md#otpsendpost) | **Post** /otp/send | 
*OneTimePasswordAPI* | [**OtpVerifyPost**](docs/OneTimePasswordAPI.md#otpverifypost) | **Post** /otp/verify | 
*UtilsAPI* | [**UtilsErrorsErrorcodeGet**](docs/UtilsAPI.md#utilserrorserrorcodeget) | **Get** /utils/errors/{errorcode} | 
*UtilsAPI* | [**UtilsTestGet**](docs/UtilsAPI.md#utilstestget) | **Get** /utils/test | 


## Documentation For Models

 - [BatchMessage](docs/BatchMessage.md)
 - [BatchMessageResponse](docs/BatchMessageResponse.md)
 - [CancelledMessageResponse](docs/CancelledMessageResponse.md)
 - [CreditsResponse](docs/CreditsResponse.md)
 - [DeletedMessageResponse](docs/DeletedMessageResponse.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [ExtendedErrorModel](docs/ExtendedErrorModel.md)
 - [Message](docs/Message.md)
 - [MessageMetadata](docs/MessageMetadata.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [MessageResponseFailurereason](docs/MessageResponseFailurereason.md)
 - [MetaData](docs/MetaData.md)
 - [OTP](docs/OTP.md)
 - [OTPResponse](docs/OTPResponse.md)
 - [OTPVerify](docs/OTPVerify.md)
 - [OTPVerifyResponse](docs/OTPVerifyResponse.md)
 - [Query](docs/Query.md)
 - [QueryMetadata](docs/QueryMetadata.md)
 - [ScheduledBatchResponse](docs/ScheduledBatchResponse.md)
 - [ScheduledMessage](docs/ScheduledMessage.md)
 - [ScheduledMessageResponse](docs/ScheduledMessageResponse.md)
 - [ScheduledMessagesResponse](docs/ScheduledMessagesResponse.md)
 - [ScheduledMessagesResponseMessage](docs/ScheduledMessagesResponseMessage.md)
 - [SendMessageResponse](docs/SendMessageResponse.md)
 - [TestResponse](docs/TestResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### JWT

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@thesmsworks.co.uk

