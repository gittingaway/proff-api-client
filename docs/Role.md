# Role

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int64** | The organisation number of the associated company. | [optional] [default to null]
**CompanyName** | **string** | The complete and formal name of the company that this role belongs to. | [optional] [default to null]
**RoleName** | **string** | The name of the role. | [optional] [default to null]
**Company** | [***Link**](Link.md) |  | [optional] [default to null]
**CompanyAccounts** | [**[]AnnualAccount**](AnnualAccount.md) | Annual accounts for the company. Note that only a subset of accounting posts are available, for a complete  set the company must be retrieved. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

