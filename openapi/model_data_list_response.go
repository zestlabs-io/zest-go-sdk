/*
 * Zest API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: contact@zestlabs.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DataListResponse struct for DataListResponse
type DataListResponse struct {
	// The number of rows in the database
	TotalRows float32 `json:"total_rows,omitempty"`
	// Current offset
	Offset float32 `json:"offset,omitempty"`
	Rows []DataDocument `json:"rows,omitempty"`
}
