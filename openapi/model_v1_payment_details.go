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
// V1PaymentDetails struct for V1PaymentDetails
type V1PaymentDetails struct {
	Type PaymentDetailsPaymentType `json:"type,omitempty"`
	Info string `json:"info,omitempty"`
}
