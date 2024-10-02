# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompaniesEniroproCountryGet**](ProCompanyApi.md#ApiCompaniesEniroproCountryGet) | **Get** /api/companies/eniropro/{country} | A list of Eniro Pro company resources.
[**ApiCompaniesEniroproCountryIdGet**](ProCompanyApi.md#ApiCompaniesEniroproCountryIdGet) | **Get** /api/companies/eniropro/{country}/{id} | Lookup a specific Eniro Pro company resource.
[**ApiCompaniesEniroproLocationsCountryGet**](ProCompanyApi.md#ApiCompaniesEniroproLocationsCountryGet) | **Get** /api/companies/eniropro/locations/{country} | Root location of a location tree with country as root element.

# **ApiCompaniesEniroproCountryGet**
> EniroProCompanyList ApiCompaniesEniroproCountryGet(ctx, country, optional)
A list of Eniro Pro company resources.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.              Possible values:              &lt;value&gt;NO&lt;/value&gt;,              &lt;value&gt;DK&lt;/value&gt;,              &lt;value&gt;SE&lt;/value&gt; | 
 **optional** | ***ProCompanyApiApiCompaniesEniroproCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProCompanyApiApiCompaniesEniroproCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Free text search on company name. Case insensitive. Mutually exclusive with the &#x27;industry&#x27; parameter. | 
 **industryCode** | **optional.String**| Search on industry code. | 
 **industry** | **optional.String**| Free text search on company industry. Case insensitive. Mutually exclusive with the &#x27;name&#x27; parameter. | 
 **sort** | **optional.String**| How the list should be sorted. Default is companyRank. | 
 **mapBoundingBox** | **optional.String**| Comma-separated list containing 2 pairs of coordinates (long, lat) for a map bounding box. First pair is lower-left corner, second pair is upper-right corner. Sequence: Western longitude, southern latitude, eastern longitude, northern latitude E.g.: 10.7417397,59.9152922,10.8417397,59.9952922 | 
 **filter** | **optional.String**| Custom filter | 
 **expand** | **optional.Bool**| Should the registerListing be expanded in the returned documents | 
 **noOfNearbyPages** | **optional.Int32**| Preferred number of pagination uris. | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 
 **filterPageSize** | **optional.Int32**| Overrides page size parameter for embedded filter links in the response. | 
 **sortPageSize** | **optional.Int32**| Overrides page size parameter for embedded sort links in the response. | 

### Return type

[**EniroProCompanyList**](EniroProCompanyList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesEniroproCountryIdGet**
> EniroProCompanyDetails ApiCompaniesEniroproCountryIdGet(ctx, country, id)
Lookup a specific Eniro Pro company resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.               Possible values:                                  &lt;value&gt;NO&lt;/value&gt;,                                  &lt;value&gt;DK&lt;/value&gt;,                                  &lt;value&gt;SE&lt;/value&gt; | 
  **id** | **string**| The unique identifier for the entity. | 

### Return type

[**EniroProCompanyDetails**](EniroProCompanyDetails.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesEniroproLocationsCountryGet**
> EniroProLocation ApiCompaniesEniroproLocationsCountryGet(ctx, country)
Root location of a location tree with country as root element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.               Possible values:                                  &lt;value&gt;NO&lt;/value&gt;,                                  &lt;value&gt;DK&lt;/value&gt;,                                  &lt;value&gt;SE&lt;/value&gt; | 

### Return type

[**EniroProLocation**](EniroProLocation.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

