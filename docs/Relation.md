# Relation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShareType** | **string** | Share type. E.g.: \&quot;A-aksjer\&quot; | [optional] [default to null]
**TotalShares** | **int64** | Total number of shares in current company. | [optional] [default to null]
**NumberOfShares** | **int64** | Number of shares owned by this shareholder in current company. | [optional] [default to null]
**SharePercentage** | **float64** | Share percentage owned by this shareholder in current company. | [optional] [default to null]
**Direct** | **bool** | Indicates if ownerships are restricted to direct ownerships. [true | false] | [optional] [default to null]
**SourceDate** | **float64** |  | [optional] [default to null]
**UpdateDate** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Id** | **string** | Shareholder-ID. Used when fetching more information for this shareholder. | [optional] [default to null]
**OrganisationNumber** | **string** | Organisation number. Only available for company shareholder. | [optional] [default to null]
**BirthYear** | **string** | Four-digit birth year. Only available for person shareholder. | [optional] [default to null]
**BirthDate** | [**time.Time**](time.Time.md) | Only available for person shareholder. | [optional] [default to null]
**Name** | **string** | Company- or person name. | [optional] [default to null]
**Address** | **string** | Shareholder address. E.g.: \&quot;Drammensveien 1\&quot; | [optional] [default to null]
**ZipCode** | **string** | Shareholder postal code. E.g.: \&quot;1234\&quot; | [optional] [default to null]
**Location** | **string** | Shareholder city. E.g. \&quot;OSLO\&quot;: | [optional] [default to null]
**Country** | **string** | Shareholder ISO2 country code. E.g.: \&quot;NO\&quot; | [optional] [default to null]
**EntityType** | **string** | Describes type of shareholder. [\&quot;Person\&quot; | \&quot;Company\&quot; | \&quot;Unknown\&quot;] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

