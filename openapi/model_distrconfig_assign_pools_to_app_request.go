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
// DistrconfigAssignPoolsToAppRequest struct for DistrconfigAssignPoolsToAppRequest
type DistrconfigAssignPoolsToAppRequest struct {
	AppID string `json:"appID,omitempty"`
	PoolIDs []string `json:"poolIDs,omitempty"`
}
