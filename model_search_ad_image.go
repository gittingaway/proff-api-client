/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// A search ad image associated with an company.
type SearchAdImage struct {
	// Search ad description.
	Description string `json:"description,omitempty"`
	// Search ad image file name.
	ImageFileName string `json:"imageFileName,omitempty"`
	// Search ad keyword.
	Keyword string `json:"keyword,omitempty"`
}
