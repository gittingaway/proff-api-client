/*
 * Proff API
 *
 * You need a valid token to access Proff API. Contact sales for more information. Token header: Authorization : Token [YOUR_TOKEN]
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package proff

type ShortAnalysisSection struct {
	Finansieringsgrad1 float64 `json:"finansieringsgrad1,omitempty"`
	KapitalensOmloepshastighet float64 `json:"kapitalensOmloepshastighet,omitempty"`
	ResultatAvDriften float64 `json:"resultatAvDriften,omitempty"`
	OmsetningPerLoennskrone float64 `json:"omsetningPerLoennskrone,omitempty"`
	Bruttofortjeneste float64 `json:"bruttofortjeneste,omitempty"`
	OmsetningPerAnsatt float64 `json:"omsetningPerAnsatt,omitempty"`
}
