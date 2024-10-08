/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Company listing data as it appears in the paper catalogue. The listings are a composite tree structure of arbitrary  width and depth.  The email, homepage and telephone number fields are mutually exclusive, so only one of them will appear in a single  listing node.
type Listing struct {
	// Company business unit ID (DK only).
	BusinessUnitId string `json:"businessUnitId,omitempty"`
	// Company ID
	CompanyId string `json:"companyId,omitempty"`
	// Company main email.
	Email string `json:"email,omitempty"`
	// Company main internet home page.
	HomePage string `json:"homePage,omitempty"`
	// Central Business Register (Denmark) marked company as protected against being contacted for advertisement purpose  Applicable only for Denmark.
	MarketingProtection bool `json:"marketingProtection,omitempty"`
	// Company name.
	Name string `json:"name,omitempty"`
	// The profession associated with this listing node.
	Profession string `json:"profession,omitempty"`
	Address *Address `json:"address,omitempty"`
	PhoneNumbers *TelephoneNumbers `json:"phoneNumbers,omitempty"`
	TextLine *ListingTextLine `json:"textLine,omitempty"`
	// The children of this listing tree node. Circular relationship
	SubListings []Listing `json:"subListings,omitempty"`
}
