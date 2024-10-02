/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// A candidate for matching against a business person in a network search.
type NetworkCandidate struct {
	NetworkLink *Link `json:"networkLink,omitempty"`
	Person *NetworkCandidatePerson `json:"person,omitempty"`
}
