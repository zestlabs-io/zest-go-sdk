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
// DistrconfigGetPoolDistributionResponse struct for DistrconfigGetPoolDistributionResponse
type DistrconfigGetPoolDistributionResponse struct {
	DbUrl string `json:"dbUrl,omitempty"`
	PoolType DistrconfigPoolType `json:"poolType,omitempty"`
	TagPrefix string `json:"tagPrefix,omitempty"`
}