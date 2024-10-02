# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompaniesOwnerCountryIdGet**](CompanyOwnerApi.md#ApiCompaniesOwnerCountryIdGet) | **Get** /api/companies/owner/{country}/{id} | CompanyOwner

# **ApiCompaniesOwnerCountryIdGet**
> CompanyRealOwner ApiCompaniesOwnerCountryIdGet(ctx, country, id)
CompanyOwner

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| The country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: \&quot;NO\&quot;, \&quot;DK\&quot; | 
  **id** | **int32**|  | 

### Return type

[**CompanyRealOwner**](CompanyRealOwner.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

