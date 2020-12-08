// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RemoveUsersFromRoleRequest v1 remove users from role request
//
// swagger:model v1RemoveUsersFromRoleRequest
type V1RemoveUsersFromRoleRequest struct {

	// role ID
	RoleID string `json:"roleID,omitempty"`

	// user i ds
	UserIDs []string `json:"userIDs"`
}

// Validate validates this v1 remove users from role request
func (m *V1RemoveUsersFromRoleRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RemoveUsersFromRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RemoveUsersFromRoleRequest) UnmarshalBinary(b []byte) error {
	var res V1RemoveUsersFromRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
