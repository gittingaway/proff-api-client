# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompaniesRegisterBusinesscategoriesCountryGet**](RegisterCompanyApi.md#ApiCompaniesRegisterBusinesscategoriesCountryGet) | **Get** /api/companies/register/businesscategories/{country} | List of register business category entities
[**ApiCompaniesRegisterCompanytypesCountryGet**](RegisterCompanyApi.md#ApiCompaniesRegisterCompanytypesCountryGet) | **Get** /api/companies/register/companytypes/{country} | Get all official company types.
[**ApiCompaniesRegisterCountryGet**](RegisterCompanyApi.md#ApiCompaniesRegisterCountryGet) | **Get** /api/companies/register/{country} | List of publicly registered companies.
[**ApiCompaniesRegisterCountryIdCorporateStructureGet**](RegisterCompanyApi.md#ApiCompaniesRegisterCountryIdCorporateStructureGet) | **Get** /api/companies/register/{country}/{id}/corporateStructure | Corporate structure look-up for a specific register company resource.
[**ApiCompaniesRegisterCountryIdGet**](RegisterCompanyApi.md#ApiCompaniesRegisterCountryIdGet) | **Get** /api/companies/register/{country}/{id} | Register company resource look-up.

# **ApiCompaniesRegisterBusinesscategoriesCountryGet**
> []RegisterBusinessCategory ApiCompaniesRegisterBusinesscategoriesCountryGet(ctx, country)
List of register business category entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO,SE,DK,FI | 

### Return type

[**[]RegisterBusinessCategory**](RegisterBusinessCategory.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesRegisterCompanytypesCountryGet**
> []RegisterCompanyType ApiCompaniesRegisterCompanytypesCountryGet(ctx, country)
Get all official company types.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO,SE,DK,FI | 

### Return type

[**[]RegisterCompanyType**](RegisterCompanyType.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesRegisterCountryGet**
> RegisterCompanyList ApiCompaniesRegisterCountryGet(ctx, country, optional)
List of publicly registered companies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**|  | 
 **optional** | ***RegisterCompanyApiApiCompaniesRegisterCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterCompanyApiApiCompaniesRegisterCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **optional.String**| How the list should be sorted. If not specified, a default sorting is used.  Table of default sorting.  query parameters  Default sorting  profitFrom, profitTo profitDesc  revenueFrom, revenueTo revenueDesc  establishedYearFrom, establishedYearTo establishedYearAsc  numEmployeesFrom, numEmployeesTo, numEmployeesDesc  companyType, profitDesc  location, profitDesc  query relevance                 When mixing query parameters with different defaults, the default sorting is always the first one  in the table that matches a query parameter.    Possible values:  numEmployeesAsc, numEmployeesDesc, revenueAsc, revenueDesc, establishedYearAsc, establishedYearDesc,  companyNameDesc, naceIndustryName, companyTypeNameDesc, profitAsc, profitDesc, relevance, postNumber | 
 **industryCode** | **optional.String**| Get industry code | 
 **location** | **optional.String**| Filter by a location.  Examples (description in parenthesis is not a part of the value):  \&quot;Viken\&quot; (county), \&quot;Drammen\&quot; (municipality), \&quot;Østlandet\&quot; (country part) | 
 **query** | **optional.String**| Free text query. Case insensitive. Required if none of the other parameters are used. If  this is the only parameter set it is considered to be a \&quot;simple search\&quot; which affects which sorting you can use. | 
 **numEmployeesFrom** | **optional.String**| Number of employees, from | 
 **numEmployeesTo** | **optional.String**| Highest number of employees companies in the list should have. | 
 **revenueFrom** | **optional.String**| Lowest revenue companies in the list should have. | 
 **revenueTo** | **optional.String**| Highest revenue companies in the list should have. | 
 **profitFrom** | **optional.String**| Lowest profit companies in the list should have. | 
 **profitTo** | **optional.String**| Highest profit companies in the list should have. | 
 **establishedYearFrom** | **optional.String**| Earliest year the companies in the list should have been established. | 
 **establishedYearTo** | **optional.String**| Most recent year the companies in the list should have been established. | 
 **companyType** | [**optional.Interface of []string**](string.md)| Array of company types.  Some examples of Norwegian codes: \&quot;AS\&quot;, \&quot;ANS\&quot;, \&quot;EP\&quot; and \&quot;NUF\&quot;. | 
 **industry** | **optional.String**| Get industries | 
 **foundationDateFrom** | **optional.String**| Foundation date from - dd.MM.yyyy format | 
 **foundationDateTo** | **optional.String**| Foundation date to - dd.MM.yyyy format | 
 **liquidationDateFrom** | **optional.String**| Liquidation from - dd.MM.yyyy format | 
 **liquidationDateTo** | **optional.String**| Liquidation date to - dd.MM.yyyy format | 
 **filter** | **optional.String**| Filter. E.g.: status:AKTIVT | 
 **noOfNearbyPages** | **optional.Int32**| Preferred number of pagination uris. | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 
 **filterPageSize** | **optional.Int32**| Overrides page size parameter for embedded filter links in the response. | 
 **sortPageSize** | **optional.Int32**| Overrides page size parameter for embedded sort links in the response. | 

### Return type

[**RegisterCompanyList**](RegisterCompanyList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesRegisterCountryIdCorporateStructureGet**
> CorporateStructure ApiCompaniesRegisterCountryIdCorporateStructureGet(ctx, country, id)
Corporate structure look-up for a specific register company resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK | 
  **id** | **string**| Unique identifier for the entity. | 

### Return type

[**CorporateStructure**](CorporateStructure.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiCompaniesRegisterCountryIdGet**
> RegisterCompanyDetails ApiCompaniesRegisterCountryIdGet(ctx, country, id)
Register company resource look-up.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO,SE,DK,FI | 
  **id** | **string**| Unique identifier for the entity. | 

### Return type

[**RegisterCompanyDetails**](RegisterCompanyDetails.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

