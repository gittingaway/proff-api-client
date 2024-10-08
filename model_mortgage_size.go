/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

type MortgageSize struct {
	// The count of the aggregated mortgage type
	Count int64 `json:"count,omitempty"`
	// The sum of the aggregated mortgage type
	Sum int64 `json:"sum,omitempty"`
}
