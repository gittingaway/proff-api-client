# BusinessPersonList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessPersons** | [**[]BusinessPerson**](BusinessPerson.md) | Professional with business roles in companies, and with connections to other professionals in the enterprise world. Roles can typically be: Chairman, committee member, deputy committee member or shareholder (Daglig Leder, Styreformann, Styremedlem, Varamedlem, Aksjon�r). A business connection is a connection this business person has to other businesses persons. | [optional] [default to null]
**Sorting** | [***Sorting**](Sorting.md) |  | [optional] [default to null]
**CorrectedQuery** | **string** | A possible rewrite of the query string made by the search engine. The search result typically reflects this  rewritten query | [optional] [default to null]
**NumberOfHits** | **int32** | Total number of hits available. | [optional] [default to null]
**FilterGroups** | [**[]FilterGroup**](FilterGroup.md) | A set of related filters. The display name is a common description for all the filters in the group. I.e. if the display name in the filter group is: &#x27;Countrypart&#x27; then you can expect all filters to represent a particular country part like: Nord-Norge (North of Norway) or Vestlandet. | [optional] [default to null]
**Pagination** | [***Pagination**](Pagination.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

