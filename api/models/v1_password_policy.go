// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1PasswordPolicy v1 password policy
//
// swagger:model v1PasswordPolicy
type V1PasswordPolicy struct {

	// min length
	MinLength string `json:"minLength,omitempty"`

	// use lower letters
	UseLowerLetters bool `json:"useLowerLetters,omitempty"`

	// use numbers
	UseNumbers bool `json:"useNumbers,omitempty"`

	// use special charecters
	UseSpecialCharecters bool `json:"useSpecialCharecters,omitempty"`

	// use upper letters
	UseUpperLetters bool `json:"useUpperLetters,omitempty"`
}

// Validate validates this v1 password policy
func (m *V1PasswordPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1PasswordPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PasswordPolicy) UnmarshalBinary(b []byte) error {
	var res V1PasswordPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
