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
// DistrconfigUnassignTagFromUserRequest struct for DistrconfigUnassignTagFromUserRequest
type DistrconfigUnassignTagFromUserRequest struct {
	PoolID string `json:"poolID,omitempty"`
	UserID string `json:"userID,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}
