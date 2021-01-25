// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateUserResponse v1 create user response
//
// swagger:model v1CreateUserResponse
type V1CreateUserResponse struct {

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this v1 create user response
func (m *V1CreateUserResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 create user response based on context it is used
func (m *V1CreateUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateUserResponse) UnmarshalBinary(b []byte) error {
	var res V1CreateUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
