# EniroProCompany

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertType** | **string** | Type of advert this company listing is. | [optional] [default to null]
**BusinessUnitId** | **string** | Company business unit ID (DK only). | [optional] [default to null]
**CompanyId** | **string** | Company Id | [optional] [default to null]
**CustomerId** | **string** | The company customer id. | [optional] [default to null]
**Description** | **string** | A promotional description of the company. | [optional] [default to null]
**Email** | **string** | Company main email. | [optional] [default to null]
**HomePage** | **string** | Company main internet home page. | [optional] [default to null]
**LineId** | **string** | The unique Norwegian id for this company. | [optional] [default to null]
**ListingId** | **string** | The id of the listing this company belongs to. | [optional] [default to null]
**ListingType** | **string** | The Type of this company listing. | [optional] [default to null]
**MarketingProtection** | **bool** | Central Business Register (Denmark) marked company as protected against being contacted for advertisement purpose  Applicable only for Denmark. | [optional] [default to null]
**Name** | **string** | Company name. | [optional] [default to null]
**OrganisationNumber** | **string** | Company organisation number. | [optional] [default to null]
**SOrganisationNumber** | **string** | Company secret organisation number.  Companies may have a secret/private organisation number;  e.g. Swedish ENK. | [optional] [default to null]
**Rank** | **float64** | The rank of this company in the context of a search result. | [optional] [default to null]
**SalesRank** | **int64** | Sales rank reflects the amount of purchases for this company. High rank means high amount. | [optional] [default to null]
**CurrentIndustry** | [***EniroProIndustry**](EniroProIndustry.md) |  | [optional] [default to null]
**Link** | [***Link**](Link.md) |  | [optional] [default to null]
**Industries** | [**[]EniroProCompanyIndustry**](EniroProCompanyIndustry.md) | An industry to which companies are related, with additional customer specific data | [optional] [default to null]
**Location** | [***Location**](Location.md) |  | [optional] [default to null]
**MainImage** | [***Image**](Image.md) |  | [optional] [default to null]
**PhoneNumbers** | [***TelephoneNumbers**](TelephoneNumbers.md) |  | [optional] [default to null]
**PostalAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Products** | [**[]Product**](Product.md) | A paid product for this company. | [optional] [default to null]
**ReferenceTexts** | [***ReferenceTexts**](ReferenceTexts.md) |  | [optional] [default to null]
**SearchAdImages** | [**[]SearchAdImage**](SearchAdImage.md) | A search ad image associated with an company. | [optional] [default to null]
**VisitorAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**RegisterListing** | [***RegisterCompany**](RegisterCompany.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

