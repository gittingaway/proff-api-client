# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCompaniesLiquidationestablishmentsCountryGet**](LiquidationEstablishmentsApi.md#ApiCompaniesLiquidationestablishmentsCountryGet) | **Get** /api/companies/liquidationestablishments/{country} | Return all establishments and liquidations for the past two years

# **ApiCompaniesLiquidationestablishmentsCountryGet**
> LiquidationAndEstablishments ApiCompaniesLiquidationestablishmentsCountryGet(ctx, country, optional)
Return all establishments and liquidations for the past two years

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: NO, SE, DK, FI | 
 **optional** | ***LiquidationEstablishmentsApiApiCompaniesLiquidationestablishmentsCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LiquidationEstablishmentsApiApiCompaniesLiquidationestablishmentsCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **industryCode** | **optional.String**| Industry code | 
 **location** | **optional.String**| Location | 

### Return type

[**LiquidationAndEstablishments**](LiquidationAndEstablishments.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

