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
// V1RemovePoliciesFromUserRequest struct for V1RemovePoliciesFromUserRequest
type V1RemovePoliciesFromUserRequest struct {
	UserID string `json:"userID,omitempty"`
	PolicyIDs []string `json:"policyIDs,omitempty"`
}
