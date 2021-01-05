# UndeleteMasterKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKey** | Pointer to [**MasterKey**](MasterKey.md) |  | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### NewUndeleteMasterKeyResponse

`func NewUndeleteMasterKeyResponse() *UndeleteMasterKeyResponse`

NewUndeleteMasterKeyResponse instantiates a new UndeleteMasterKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUndeleteMasterKeyResponseWithDefaults

`func NewUndeleteMasterKeyResponseWithDefaults() *UndeleteMasterKeyResponse`

NewUndeleteMasterKeyResponseWithDefaults instantiates a new UndeleteMasterKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKey

`func (o *UndeleteMasterKeyResponse) GetMasterKey() MasterKey`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *UndeleteMasterKeyResponse) GetMasterKeyOk() (*MasterKey, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *UndeleteMasterKeyResponse) SetMasterKey(v MasterKey)`

SetMasterKey sets MasterKey field to given value.

### HasMasterKey

`func (o *UndeleteMasterKeyResponse) HasMasterKey() bool`

HasMasterKey returns a boolean if a field has been set.

### GetResponseContext

`func (o *UndeleteMasterKeyResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *UndeleteMasterKeyResponse) GetResponseContextOk() (*ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContext

`func (o *UndeleteMasterKeyResponse) SetResponseContext(v ResponseContext)`

SetResponseContext sets ResponseContext field to given value.

### HasResponseContext

`func (o *UndeleteMasterKeyResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


