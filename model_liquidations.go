/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Model containing list of items
type Liquidations struct {
	// Liquidations grouped by yyyy-MM
	Items []DateValueItem `json:"items,omitempty"`
}
