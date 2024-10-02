/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

// Coordinate for a company.
type Coordinate struct {
	// This location longitude- or x-coordinate value - Depending on the coordinate system representation.
	XCoordinate float64 `json:"xCoordinate,omitempty"`
	// This location latitude- or y-coordinate value - Depending on the coordinate system representation.
	YCoordinate float64 `json:"yCoordinate,omitempty"`
	// Specifies which coordinate system the provided coordinates use.
	CoordinateSystem string `json:"coordinateSystem,omitempty"`
}
