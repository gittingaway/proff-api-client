/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Performed calls
type CallStats struct {
	// Calls performed last minute
	LastMinute int32 `json:"lastMinute,omitempty"`
	// Calls performed last hour
	LastHour int32 `json:"lastHour,omitempty"`
	// Calls performed last day
	LastDay int32 `json:"lastDay,omitempty"`
	// Calls performed last month
	LastMonth int32 `json:"lastMonth,omitempty"`
	// Calls performed last year
	LastYear int32 `json:"lastYear,omitempty"`
	Period *Period `json:"period,omitempty"`
}
