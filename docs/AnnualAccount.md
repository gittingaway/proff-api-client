# AnnualAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccIncompleteCode** | **string** | The code for the account being incomplete | [optional] [default to null]
**AccIncompleteDesc** | **string** | The comment description for the account being incomplete | [optional] [default to null]
**Combined** | **bool** | Whether or not this is a combined annual account for an enterprise. | [optional] [default to null]
**Currency** | **string** | The code for the currency the account amounts are in. | [optional] [default to null]
**Length** | **string** | Accounting reported period. Number of months. | [optional] [default to null]
**Period** | **string** | accounting reported period. (year and month YYYY-MM format) | [optional] [default to null]
**Year** | **string** | The year this annual account is for. On standard ECS format. | [optional] [default to null]
**PeriodStart** | [***DateOnly**](DateOnly.md) |  | [optional] [default to null]
**PeriodEnd** | [***DateOnly**](DateOnly.md) |  | [optional] [default to null]
**Accounts** | [**[]Account**](Account.md) | An account used in company accounting. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

