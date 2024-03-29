// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AddUsersToRoleResponse v1 add users to role response
//
// swagger:model v1AddUsersToRoleResponse
type V1AddUsersToRoleResponse struct {

	// failed user i ds
	FailedUserIDs []string `json:"failedUserIDs"`
}

// Validate validates this v1 add users to role response
func (m *V1AddUsersToRoleResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 add users to role response based on context it is used
func (m *V1AddUsersToRoleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AddUsersToRoleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AddUsersToRoleResponse) UnmarshalBinary(b []byte) error {
	var res V1AddUsersToRoleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
