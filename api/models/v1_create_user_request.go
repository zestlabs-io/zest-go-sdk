// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CreateUserRequest v1 create user request
//
// swagger:model v1CreateUserRequest
type V1CreateUserRequest struct {

	// email
	Email string `json:"email,omitempty"`

	// federation ID
	FederationID string `json:"federationID,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this v1 create user request
func (m *V1CreateUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CreateUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CreateUserRequest) UnmarshalBinary(b []byte) error {
	var res V1CreateUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
