// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RoleUpdate v1 role update
//
// swagger:model v1RoleUpdate
type V1RoleUpdate struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// role ID
	RoleID string `json:"roleID,omitempty"`
}

// Validate validates this v1 role update
func (m *V1RoleUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 role update based on context it is used
func (m *V1RoleUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RoleUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RoleUpdate) UnmarshalBinary(b []byte) error {
	var res V1RoleUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
