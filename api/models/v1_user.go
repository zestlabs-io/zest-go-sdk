// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1User v1 user
//
// swagger:model v1User
type V1User struct {

	// email
	Email string `json:"email,omitempty"`

	// first name
	FirstName string `json:"firstName,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// last login
	LastLogin string `json:"lastLogin,omitempty"`

	// last name
	LastName string `json:"lastName,omitempty"`

	// phone number
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// policy i ds
	PolicyIDs []string `json:"policyIDs"`

	// role i ds
	RoleIDs []string `json:"roleIDs"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this v1 user
func (m *V1User) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1User) UnmarshalBinary(b []byte) error {
	var res V1User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
