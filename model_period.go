/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Performed calls in a period
type Period struct {
	Start *DateOnly `json:"start,omitempty"`
	End *DateOnly `json:"end,omitempty"`
	// Total calls - within the period
	Calls int32 `json:"calls,omitempty"`
}
