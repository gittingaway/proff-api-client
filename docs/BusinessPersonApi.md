# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPersonsBusinessCountryGet**](BusinessPersonApi.md#ApiPersonsBusinessCountryGet) | **Get** /api/persons/business/{country} | List of business person resources.
[**ApiPersonsBusinessCountryIdGet**](BusinessPersonApi.md#ApiPersonsBusinessCountryIdGet) | **Get** /api/persons/business/{country}/{id} | Lookup a specific business persons resource.
[**ApiPersonsBusinessCountryIdNetworkcandidatesGet**](BusinessPersonApi.md#ApiPersonsBusinessCountryIdNetworkcandidatesGet) | **Get** /api/persons/business/{country}/{id}/networkcandidates | Get network resource links for this specific business person resource, against other business persons matching query.
[**ApiPersonsBusinessNetworkCountryGet**](BusinessPersonApi.md#ApiPersonsBusinessNetworkCountryGet) | **Get** /api/persons/business/network/{country} | Presents the shortest business network path between to business persons.The two persons are coupled together by one or more mutual company and one or more mutually related business person. The person network will typically start and end with the persons ids defined in the input parameter set.
[**ApiPersonsBusinessSuggestionsCountryGet**](BusinessPersonApi.md#ApiPersonsBusinessSuggestionsCountryGet) | **Get** /api/persons/business/suggestions/{country} | Get a list of suggested business person names based on the given suggestion term. No limit on how few letters the suggestion resource will respond to is enforced, but it is generally recommended (and it may be enforced in the future) to only querying the resource after at least 3 letters are provide.

# **ApiPersonsBusinessCountryGet**
> BusinessPersonList ApiPersonsBusinessCountryGet(ctx, country, optional)
List of business person resources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
 **optional** | ***BusinessPersonApiApiPersonsBusinessCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BusinessPersonApiApiPersonsBusinessCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**| Free text search query. Case insensitive. Required if &#x27;nationalId&#x27; is not set. | 
 **nationalId** | **optional.String**| Query on a persons national id. Required if &#x27;query&#x27; is not set. | 
 **birthDate** | **optional.String**| Person birth date. Date format yyyy-MM-dd | 
 **sort** | **optional.String**| Order in which to sort the returned list.  Possible values: RELEVANCE, DATE_OF_BIRTH, ALPHABETICAL | 
 **noOfNearbyPages** | **optional.Int32**| Preferred number of pagination uris. | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 
 **filterPageSize** | **optional.Int32**| Overrides page size parameter for embedded filter links in the response. | 
 **sortPageSize** | **optional.Int32**| Overrides page size parameter for embedded sort links in the response. | 

### Return type

[**BusinessPersonList**](BusinessPersonList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPersonsBusinessCountryIdGet**
> BusinessPerson ApiPersonsBusinessCountryIdGet(ctx, country, id)
Lookup a specific business persons resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **id** | **string**| Unique identifier for the entity. | 

### Return type

[**BusinessPerson**](BusinessPerson.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPersonsBusinessCountryIdNetworkcandidatesGet**
> []NetworkCandidate ApiPersonsBusinessCountryIdNetworkcandidatesGet(ctx, country, id, query)
Get network resource links for this specific business person resource, against other business persons matching query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **id** | **string**| The unique identifier for the entity. | 
  **query** | **string**| The free text search query. Case insensitive. | 

### Return type

[**[]NetworkCandidate**](NetworkCandidate.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPersonsBusinessNetworkCountryGet**
> []BusinessNetworkConnectionNode ApiPersonsBusinessNetworkCountryGet(ctx, country, from, to)
Presents the shortest business network path between to business persons.The two persons are coupled together by one or more mutual company and one or more mutually related business person. The person network will typically start and end with the persons ids defined in the input parameter set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
  **from** | **string**| The id defining the start point of the network path. | 
  **to** | **string**| The id defining the end point of the network path. | 

### Return type

[**[]BusinessNetworkConnectionNode**](BusinessNetworkConnectionNode.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiPersonsBusinessSuggestionsCountryGet**
> []Suggestion ApiPersonsBusinessSuggestionsCountryGet(ctx, country, query)
Get a list of suggested business person names based on the given suggestion term. No limit on how few letters the suggestion resource will respond to is enforced, but it is generally recommended (and it may be enforced in the future) to only querying the resource after at least 3 letters are provide.

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

