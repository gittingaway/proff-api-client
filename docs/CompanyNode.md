# CompanyNode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganisationNumber** | **string** | Company organisation number. | [optional] [default to null]
**Name** | **string** | Company name. | [optional] [default to null]
**CountryCode** | **string** | Company country. Uppercase ISO 3166-1-alpha-2 country code. | [optional] [default to null]
**OwnedPercentage** | **float64** | Percentage of the company shares owned by it&#x27;s parent. | [optional] [default to null]
**InactiveDate** | **string** | If a company is inactive this field has the inactivation date. ISO 8601 date format. | [optional] [default to null]
**Sub** | [**[]CompanyNode**](CompanyNode.md) | Companies that this company has more than 50% ownership in, e.g. subsidiaries. | [optional] [default to null]
**CompanyAccounts** | [**[]AnnualAccount**](AnnualAccount.md) | Annual account information for the company. Note that the structure is just a summary with a selection of important accounting figures for the most recent years. | [optional] [default to null]
**Link** | [***Link**](Link.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

