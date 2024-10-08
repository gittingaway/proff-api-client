/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// RegisterBusinessCategoryDTO
type RegisterBusinessCategory struct {
	// TBusiness category type (e.g. \"NACE\")
	BusinessCategoryType string `json:"businessCategoryType,omitempty"`
	// CountryId
	CountryId string `json:"countryId,omitempty"`
	// CountryId
	GroupId string `json:"groupId,omitempty"`
	// Group Text
	GroupText string `json:"groupText,omitempty"`
	// Id
	Id string `json:"id,omitempty"`
	// Text
	Text string `json:"text,omitempty"`
}
