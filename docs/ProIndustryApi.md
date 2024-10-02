# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiIndustriesEniroproCountryGet**](ProIndustryApi.md#ApiIndustriesEniroproCountryGet) | **Get** /api/industries/eniropro/{country} | EniroProIndustries resource holds information about industries that are know to EniroPro. EniroProIndustry are  distinct branches of trade, that a company can be associated with.
[**ApiIndustriesEniroproCountryIdFactsGet**](ProIndustryApi.md#ApiIndustriesEniroproCountryIdFactsGet) | **Get** /api/industries/eniropro/{country}/{id}/facts | Get facts for an industry with the specified id
[**ApiIndustriesEniroproCountryIdGet**](ProIndustryApi.md#ApiIndustriesEniroproCountryIdGet) | **Get** /api/industries/eniropro/{country}/{id} | An EniroPro industry.
[**ApiIndustriesEniroproSuggestionsCountryGet**](ProIndustryApi.md#ApiIndustriesEniroproSuggestionsCountryGet) | **Get** /api/industries/eniropro/suggestions/{country} | Get a list of suggested industry names based on the given suggestion term. No limit on how few letters the  suggestion resource will respond to is enforced, but it is generally recommended (and it may be enforced in the  future) to only querying the resource after at least 3 letters are provided.

# **ApiIndustriesEniroproCountryGet**
> []EniroProIndustry ApiIndustriesEniroproCountryGet(ctx, country, optional)
EniroProIndustries resource holds information about industries that are know to EniroPro. EniroProIndustry are  distinct branches of trade, that a company can be associated with.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
 **optional** | ***ProIndustryApiApiIndustriesEniroproCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProIndustryApiApiIndustriesEniroproCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| (Exact) name of the industry | 
 **categoryType** | **optional.String**| Which category type the industry belongs to. Defaults to MAIN  Possible values: TOPCATEGORY, CATEGORY, MAIN, SPECIALITY | 
 **noOfNearbyPages** | **optional.Int32**| Preferred number of pagination uris. | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 
 **filterPageSize** | **optional.Int32**| Overrides page size parameter for embedded filter links in the response. | 
 **sortPageSize** | **optional.Int32**| Overrides page size parameter for embedded sort links in the response. | 

### Return type

[**[]EniroProIndustry**](EniroProIndustry.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiIndustriesEniroproCountryIdFactsGet**
> EniroProIndustryFacts ApiIndustriesEniroproCountryIdFactsGet(ctx, country, id)
Get facts for an industry with the specified id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **id** | **string**| Unique identifier for the entity. | 

### Return type

[**EniroProIndustryFacts**](EniroProIndustryFacts.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiIndustriesEniroproCountryIdGet**
> EniroProIndustry ApiIndustriesEniroproCountryIdGet(ctx, country, id)
An EniroPro industry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **id** | **string**| Unique identifier for the entity. | 

### Return type

[**EniroProIndustry**](EniroProIndustry.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiIndustriesEniroproSuggestionsCountryGet**
> []Suggestion ApiIndustriesEniroproSuggestionsCountryGet(ctx, country, query)
Get a list of suggested industry names based on the given suggestion term. No limit on how few letters the  suggestion resource will respond to is enforced, but it is generally recommended (and it may be enforced in the  future) to only querying the resource after at least 3 letters are provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **query** | **string**| The free text search query. Case insensitive. | 

### Return type

[**[]Suggestion**](Suggestion.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

