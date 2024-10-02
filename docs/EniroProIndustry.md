# EniroProIndustry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryType** | **string** |    The category type this industry belongs to.  Possible values:  &lt;value&gt;TOPCATEGORY&lt;/value&gt;,  &lt;value&gt;CATEGORY&lt;/value&gt;,  &lt;value&gt;MAIN&lt;/value&gt;,  &lt;value&gt;SPECIALITY&lt;/value&gt; | [optional] [default to null]
**Description** | **string** | Industry description | [optional] [default to null]
**IndustryCode** | **string** | Industry Code | [optional] [default to null]
**Name** | **string** | The official name of the industry. E.g. \&quot;Hairdresser\&quot;. | [optional] [default to null]
**NumberOfCompanies** | **int32** | Number of companies in this industry. | [optional] [default to null]
**Details** | [***Link**](Link.md) |  | [optional] [default to null]
**Companies** | [***Link**](Link.md) |  | [optional] [default to null]
**Facts** | [***Link**](Link.md) |  | [optional] [default to null]
**Children** | [**[]Link**](Link.md) | Industries. Circular relationship. | [optional] [default to null]
**Parents** | [**[]Link**](Link.md) | Parent industries. Circular relationship. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

