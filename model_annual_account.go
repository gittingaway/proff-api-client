/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// An annual account for a company or an enterprise.
type AnnualAccount struct {
	// The code for the account being incomplete
	AccIncompleteCode string `json:"accIncompleteCode,omitempty"`
	// The comment description for the account being incomplete
	AccIncompleteDesc string `json:"accIncompleteDesc,omitempty"`
	// Whether or not this is a combined annual account for an enterprise.
	Combined bool `json:"combined,omitempty"`
	// The code for the currency the account amounts are in.
	Currency string `json:"currency,omitempty"`
	// Accounting reported period. Number of months.
	Length string `json:"length,omitempty"`
	// accounting reported period. (year and month YYYY-MM format)
	Period string `json:"period,omitempty"`
	// The year this annual account is for. On standard ECS format.
	Year string `json:"year,omitempty"`
	PeriodStart *DateOnly `json:"periodStart,omitempty"`
	PeriodEnd *DateOnly `json:"periodEnd,omitempty"`
	// An account used in company accounting.
	Accounts []Account `json:"accounts,omitempty"`
}
