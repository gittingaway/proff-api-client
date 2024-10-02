# Listing

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessUnitId** | **string** | Company business unit ID (DK only). | [optional] [default to null]
**CompanyId** | **string** | Company ID | [optional] [default to null]
**Email** | **string** | Company main email. | [optional] [default to null]
**HomePage** | **string** | Company main internet home page. | [optional] [default to null]
**MarketingProtection** | **bool** | Central Business Register (Denmark) marked company as protected against being contacted for advertisement purpose  Applicable only for Denmark. | [optional] [default to null]
**Name** | **string** | Company name. | [optional] [default to null]
**Profession** | **string** | The profession associated with this listing node. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumbers** | [***TelephoneNumbers**](TelephoneNumbers.md) |  | [optional] [default to null]
**TextLine** | [***ListingTextLine**](ListingTextLine.md) |  | [optional] [default to null]
**SubListings** | [**[]Listing**](Listing.md) | The children of this listing tree node. Circular relationship | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

