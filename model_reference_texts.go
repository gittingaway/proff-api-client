/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Texts describing the references between an entity and other entities.
type ReferenceTexts struct {
	// A text containing a reference from the current entity to another one.
	From string `json:"from,omitempty"`
	// A text containing a reference to an entity related to the current one.
	See string `json:"see,omitempty"`
	// A text containing a reference to the current entity from another one.
	To string `json:"to,omitempty"`
}
