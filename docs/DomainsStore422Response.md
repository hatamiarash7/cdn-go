# DomainsStore422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] [default to false]
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**DomainsStore422ResponseErrors**](DomainsStore422ResponseErrors.md) |  | [optional] 

## Methods

### NewDomainsStore422Response

`func NewDomainsStore422Response() *DomainsStore422Response`

NewDomainsStore422Response instantiates a new DomainsStore422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainsStore422ResponseWithDefaults

`func NewDomainsStore422ResponseWithDefaults() *DomainsStore422Response`

NewDomainsStore422ResponseWithDefaults instantiates a new DomainsStore422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DomainsStore422Response) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainsStore422Response) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainsStore422Response) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DomainsStore422Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DomainsStore422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainsStore422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainsStore422Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DomainsStore422Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *DomainsStore422Response) GetErrors() DomainsStore422ResponseErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DomainsStore422Response) GetErrorsOk() (*DomainsStore422ResponseErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DomainsStore422Response) SetErrors(v DomainsStore422ResponseErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DomainsStore422Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


