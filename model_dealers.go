/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// The dealers of a company's products.
type Dealers struct {
	// Get a list of names of dealers of a company's products.
	DealerNames []string `json:"dealerNames,omitempty"`
	// A url to information about dealers of a company's products.
	InformationUrl string `json:"informationUrl,omitempty"`
}
