# DayOpeningHours

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllDay** | **bool** | true status flag if the company is open the whole day this particular day, and false otherwise. | [optional] [default to null]
**Closed** | **bool** | true status flag if the company is closed the whole day this particular day, and false otherwise. | [optional] [default to null]
**ClosingHour** | **string** | The hour the company closes a particular day. Must be combined with ClosingMinutes to get a complete time. | [optional] [default to null]
**ClosingMinutes** | **string** | The minutes the company closes a particular day. Must be combined with ClosingHour to get a complete time. | [optional] [default to null]
**OpenHours** | **bool** | true status flag if the company has opening hours this particular day. Use a combination of OpeningHour +  OpeningMinutes and ClosingHour + ClosingMinutes to calculate the exact interval of the opening hours. | [optional] [default to null]
**OpeningHour** | **string** | The hour the company opens a particular day. Must be combined with OpeningMinutes to get a complete time. | [optional] [default to null]
**OpeningMinutes** | **string** | The minutes the company opens a particular day. Must be combined with OpeningHour to get a complete time. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

