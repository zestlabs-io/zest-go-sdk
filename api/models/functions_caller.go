// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsCaller functions caller
//
// swagger:model functionsCaller
type FunctionsCaller struct {

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// policies
	Policies string `json:"policies,omitempty"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this functions caller
func (m *FunctionsCaller) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this functions caller based on context it is used
func (m *FunctionsCaller) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsCaller) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsCaller) UnmarshalBinary(b []byte) error {
	var res FunctionsCaller
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
