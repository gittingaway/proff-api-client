# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiChangesRegisterCountryGet**](ChangesApi.md#ApiChangesRegisterCountryGet) | **Get** /api/changes/register/{country} | Changed companies and business units

# **ApiChangesRegisterCountryGet**
> ChangesResult ApiChangesRegisterCountryGet(ctx, country, optional)
Changed companies and business units

Changes to publicly registered companies are normally updated once a day.<br />                Companies are considered changed when they are removed from the RegisterCompany resource.  Companies are removed from the RegisterCompany resource after having been marked deleted for over 12 months.  New established companies are considered to be a change and will be included.  Deletion of companies is considered to be a change and will be included.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| The country to search within. Must be a two letter uppercase ISO 3166-1-alpha-2 country code.  Possible values: \&quot;NO\&quot; | 
 **optional** | ***ChangesApiApiChangesRegisterCountryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChangesApiApiChangesRegisterCountryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scrollId** | **optional.String**| Specify a scroll Id to advance to the next batch of results.  If specified any other parameter is ignored.                Generally | 
 **types** | [**optional.Interface of []string**](string.md)| Type of changes returned in the search result.  Choose one or more of the following values:    ADDED: gets new documents    REMOVED: gets removed documents    MUTATED: gets changed documents    If MUTATED is specified all fields are considered.  See fields for details on overriding considered fields.                By default all added, removed and changed documents are returned. | 
 **withPatch** | **optional.Bool**| The JSONPatch should be returned with each ChangedDoc element | 
 **since** | **optional.String**| To include ChangedDocs older than 1 day set this date. We keep a history of 30 days of changes.  If you have missed changes older 30 days, you should fetch the original document again.     Format: yyyy-MM-dd | 
 **fields** | [**optional.Interface of []string**](string.md)| Which fields to look for changes in. Only applies for MUTATED documents.    If MUTATED documents are requested, the paths that are searched for changes can be overridden by specifying this parameter one or more times.   E.g. specify fields&#x3D;/name&amp;fields&#x3D;/phoneNumbers if you are only interested in changes to name or phoneNumbers.                The paths refers to fields in the RegisterCompany schema.  They are used as prefixes, so /personRoles includes any change to any element in the array.    If you only specify this parameter, all ADDED, REMOVED and MUTATED (where this parameter matches) documents are returned.    If you want to exclude ADDED and REMOVED documents you have to specify the types parameter (with the MUTATED value). | 

### Return type

[**ChangesResult**](ChangesResult.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

