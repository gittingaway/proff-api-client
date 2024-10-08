/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// An Atom-style link that describes the relationships between resources.
type PageLink struct {
	// An Atom-style link that describes the relationships between resources.
	PageNumber int32 `json:"pageNumber,omitempty"`
	// The hypertext reference (URI) of the linked resource.
	Href string `json:"href,omitempty"`
	// The type of the relationship between the origin and the target of the link. For example the relationship on a link  to the next page would be \"next\" and the relationship on a link to company details would be \"details\".
	Rel string `json:"rel,omitempty"`
}
