# EniroProIndustryFacts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryType** | **string** | The category type this industry belongs to.  Possible values:  &lt;value&gt;TOPCATEGORY&lt;/value&gt;  ,  &lt;value&gt;CATEGORY&lt;/value&gt;  ,  &lt;value&gt;MAIN&lt;/value&gt;  ,  &lt;value&gt;SPECIALITY&lt;/value&gt; | [optional] [default to null]
**Country** | **string** | Country of origin for for this industry facts. A two letter uppercase ISO 3166-1-alpha-2 country code. | [optional] [default to null]
**Description** | **string** | Industry description | [optional] [default to null]
**IndustryCode** | **string** | Industry Code | [optional] [default to null]
**Name** | **string** | The official name of the industry. E.g. \&quot;Hairdresser\&quot;. | [optional] [default to null]
**NumberOfCompanies** | **int32** | Number of companies in this industry. | [optional] [default to null]
**Children** | [**[]EniroProIndustry**](EniroProIndustry.md) | An industry, any department or branch of art, occupation or business. | [optional] [default to null]
**Link** | [***Link**](Link.md) |  | [optional] [default to null]
**FinancialInfo** | [**[]FinancialInfo**](FinancialInfo.md) | Annual financial info for an attribute of an industry. | [optional] [default to null]
**Parents** | [**[]EniroProIndustry**](EniroProIndustry.md) | Parent industries. Circular relationship. | [optional] [default to null]
**TopEmployeesCompanies** | [**[]TopEmployeesCompany**](TopEmployeesCompany.md) | An top employer in an industry. A company that is among the ones with the most employees in the industry. | [optional] [default to null]
**TopTurnoverCompany** | [**[]TopTurnoverCompany**](TopTurnoverCompany.md) | An top employer in an industry. A company that is among the ones with the most employees in the industry. | [optional] [default to null]
**Companies** | [***Link**](Link.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

