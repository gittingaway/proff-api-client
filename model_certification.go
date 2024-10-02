/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Company certification.
type Certification struct {
	// The certification code associated with this certification.
	Code string `json:"code,omitempty"`
	// The name, location or id of a logo associated with this certification.
	Logo string `json:"logo,omitempty"`
	// A description of the certificate.
	Text string `json:"text,omitempty"`
	// A url to more information about this certification.
	Url string `json:"url,omitempty"`
}
