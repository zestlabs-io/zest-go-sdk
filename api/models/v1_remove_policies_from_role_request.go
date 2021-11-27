// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1RemovePoliciesFromRoleRequest v1 remove policies from role request
//
// swagger:model v1RemovePoliciesFromRoleRequest
type V1RemovePoliciesFromRoleRequest struct {

	// policy i ds
	PolicyIDs []string `json:"policyIDs"`

	// role ID
	RoleID string `json:"roleID,omitempty"`
}

// Validate validates this v1 remove policies from role request
func (m *V1RemovePoliciesFromRoleRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 remove policies from role request based on context it is used
func (m *V1RemovePoliciesFromRoleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1RemovePoliciesFromRoleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RemovePoliciesFromRoleRequest) UnmarshalBinary(b []byte) error {
	var res V1RemovePoliciesFromRoleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
