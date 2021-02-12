# DeleteDhcpOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSetId** | **string** | The ID of the DHCP options set you want to delete. | 
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewDeleteDhcpOptionsRequest

`func NewDeleteDhcpOptionsRequest(dhcpOptionsSetId string, ) *DeleteDhcpOptionsRequest`

NewDeleteDhcpOptionsRequest instantiates a new DeleteDhcpOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDhcpOptionsRequestWithDefaults

`func NewDeleteDhcpOptionsRequestWithDefaults() *DeleteDhcpOptionsRequest`

NewDeleteDhcpOptionsRequestWithDefaults instantiates a new DeleteDhcpOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpOptionsSetId

`func (o *DeleteDhcpOptionsRequest) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *DeleteDhcpOptionsRequest) GetDhcpOptionsSetIdOk() (*string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptionsSetId

`func (o *DeleteDhcpOptionsRequest) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId sets DhcpOptionsSetId field to given value.


### GetDryRun

`func (o *DeleteDhcpOptionsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteDhcpOptionsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *DeleteDhcpOptionsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *DeleteDhcpOptionsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


