// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ResetPasswordRequest v1 reset password request
//
// swagger:model v1ResetPasswordRequest
type V1ResetPasswordRequest struct {

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this v1 reset password request
func (m *V1ResetPasswordRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ResetPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ResetPasswordRequest) UnmarshalBinary(b []byte) error {
	var res V1ResetPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}