# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiShareholdersEniroproCountryGet**](ShareholdersApi.md#ApiShareholdersEniroproCountryGet) | **Get** /api/shareholders/eniropro/{country} | Query for both company- and person shareholders.
[**ApiShareholdersEniroproCountryOwnersCompanyIdGet**](ShareholdersApi.md#ApiShareholdersEniroproCountryOwnersCompanyIdGet) | **Get** /api/shareholders/eniropro/{country}/owners/{companyId} | Get owners for a specific company.
[**ApiShareholdersEniroproCountryOwnershipsIdGet**](ShareholdersApi.md#ApiShareholdersEniroproCountryOwnershipsIdGet) | **Get** /api/shareholders/eniropro/{country}/ownerships/{id} | Get shares owned by a specific person or company.  Owner id can be found in the id field of entity
[**ApiShareholdersEniroproCountryRoleRoleIdGet**](ShareholdersApi.md#ApiShareholdersEniroproCountryRoleRoleIdGet) | **Get** /api/shareholders/eniropro/{country}/role/{roleId} | Lookup shareholders by Role-ID

# **ApiShareholdersEniroproCountryGet**
> ShareholdersResult ApiShareholdersEniroproCountryGet(ctx, country, name, optional)
Query for both company- and person shareholders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code. Possible values: \&quot;NO\&quot; | 
  **name** | **string**| Company name or person name | 
 **optional** | ***ShareholdersApiApiShareholdersEniroproCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareholdersApiApiShareholdersEniroproCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 

### Return type

[**ShareholdersResult**](ShareholdersResult.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiShareholdersEniroproCountryOwnersCompanyIdGet**
> LookupResult ApiShareholdersEniroproCountryOwnersCompanyIdGet(ctx, country, companyId, optional)
Get owners for a specific company.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code. Possible values: \&quot;NO\&quot; | 
  **companyId** | **string**| Company organisation number | 
 **optional** | ***ShareholdersApiApiShareholdersEniroproCountryOwnersCompanyIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareholdersApiApiShareholdersEniroproCountryOwnersCompanyIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direct** | **optional.Bool**| Restrict to direct ownerships | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 

### Return type

[**LookupResult**](LookupResult.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiShareholdersEniroproCountryOwnershipsIdGet**
> LookupResult ApiShareholdersEniroproCountryOwnershipsIdGet(ctx, country, id, optional)
Get shares owned by a specific person or company.  Owner id can be found in the id field of entity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: \&quot;NO\&quot; | 
  **id** | **string**| Organisation number or ownerId | 
 **optional** | ***ShareholdersApiApiShareholdersEniroproCountryOwnershipsIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ShareholdersApiApiShareholdersEniroproCountryOwnershipsIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direct** | **optional.Bool**| Restrict to direct ownerships | 
 **pageSize** | **optional.Int32**| Preferred number of resources returned for a particular query - defaults to 10. | 
 **pageNumber** | **optional.Int32**| Requested page number - defaults to 1. | 

### Return type

[**LookupResult**](LookupResult.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiShareholdersEniroproCountryRoleRoleIdGet**
> ShareholdersResult ApiShareholdersEniroproCountryRoleRoleIdGet(ctx, country, roleId)
Lookup shareholders by Role-ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: \&quot;NO\&quot; | 
  **roleId** | **string**| Role-ID for person | 

### Return type

[**ShareholdersResult**](ShareholdersResult.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

