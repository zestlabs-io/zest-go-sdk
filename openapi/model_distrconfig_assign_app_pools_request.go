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
// DistrconfigAssignAppPoolsRequest struct for DistrconfigAssignAppPoolsRequest
type DistrconfigAssignAppPoolsRequest struct {
	AppId string `json:"appId,omitempty"`
	PoolIds []string `json:"poolIds,omitempty"`
}
