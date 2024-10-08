/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Date value item class
type DateValueItem struct {
	// Date
	Date string `json:"date,omitempty"`
	// Value
	Value string `json:"value,omitempty"`
}
