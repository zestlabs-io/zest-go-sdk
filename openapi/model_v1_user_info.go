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
// V1UserInfo struct for V1UserInfo
type V1UserInfo struct {
	AccountID string `json:"accountID,omitempty"`
	UserID string `json:"userID,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Policies []V1Policy `json:"policies,omitempty"`
}
