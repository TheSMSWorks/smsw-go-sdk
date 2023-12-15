# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the messages you would like returned (either &#39;SENT&#39;, &#39;DELIVERED&#39;, &#39;EXPIRED&#39;, &#39;UNDELIVERABLE&#39;, &#39;REJECTED&#39; or &#39;INCOMING&#39;) | [optional] 
**Credits** | Pointer to **float32** | The number of credits used on the message. Floating point number. | [optional] 
**Destination** | Pointer to **string** | The phone number of the recipient. Start UK numbers with 44 and drop the leading 0. | [optional] 
**Sender** | Pointer to **string** | The sender of the message (this can be the configured sender name for an outbound message or the senders phone number for an inbound message). | [optional] 
**Keyword** | Pointer to **string** | The keyword used in the inbound message | [optional] 
**From** | Pointer to **string** | The date-time from which you would like matching messages | [optional] 
**To** | Pointer to **string** | The date-time to which you would like matching messages | [optional] 
**Limit** | Pointer to **float32** | The maximum number of messages that you would like returned in this call. The default is 1000. | [optional] 
**Skip** | Pointer to **float32** | The number of results you would like to ignore before returning messages. In combination with the &#39;limit&#39; parameter his can be used to page results, so that you can deal with a limited number in your logic at each time. | [optional] 
**Unread** | Pointer to **bool** | In queries for incoming messages (&#39;status&#39; is &#39;INCOMING&#39;), specify whether you explicitly want unread messages (true) or read messages (false). Omit this parameter in other circumstances. | [optional] 
**Metadata** | Pointer to [**QueryMetadata**](QueryMetadata.md) |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Query) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Query) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Query) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Query) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCredits

`func (o *Query) GetCredits() float32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Query) GetCreditsOk() (*float32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Query) SetCredits(v float32)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *Query) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDestination

`func (o *Query) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Query) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Query) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Query) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetSender

`func (o *Query) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Query) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Query) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Query) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetKeyword

`func (o *Query) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *Query) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *Query) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *Query) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetFrom

`func (o *Query) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Query) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Query) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Query) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Query) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Query) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Query) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Query) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetLimit

`func (o *Query) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Query) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Query) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Query) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSkip

`func (o *Query) GetSkip() float32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *Query) GetSkipOk() (*float32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *Query) SetSkip(v float32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *Query) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetUnread

`func (o *Query) GetUnread() bool`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *Query) GetUnreadOk() (*bool, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *Query) SetUnread(v bool)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *Query) HasUnread() bool`

HasUnread returns a boolean if a field has been set.

### GetMetadata

`func (o *Query) GetMetadata() QueryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Query) GetMetadataOk() (*QueryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Query) SetMetadata(v QueryMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Query) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


