/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// A shareholder in a company. Can be a person or a company.
type Shareholder struct {
	// The organisation number of this shareholder.
	CompanyId string `json:"companyId,omitempty"`
	// The first name of this shareholder.
	FirstName string `json:"firstName,omitempty"`
	// The last name of this shareholder.
	LastName string `json:"lastName,omitempty"`
	// The name of this shareholder.
	Name string `json:"name,omitempty"`
	// The number of shares this shareholder owns in the company.
	NumberOfShares float64 `json:"numberOfShares,omitempty"`
	// The percentage of the total number of company shares that is owned by this shareholder. Formatted for view.
	Share string `json:"share,omitempty"`
	Details *Link `json:"details,omitempty"`
}
