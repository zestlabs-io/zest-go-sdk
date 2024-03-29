// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreatePolicyResponse v1 create policy response
//
// swagger:model v1CreatePolicyResponse
type V1CreatePolicyResponse struct {

	// policy ID
	PolicyID string `json:"policyID,omitempty"`
}

// Validate validates this v1 create policy response
func (m *V1CreatePolicyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 create policy response based on context it is used
func (m *V1CreatePolicyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreatePolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreatePolicyResponse) UnmarshalBinary(b []byte) error {
	var res V1CreatePolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
