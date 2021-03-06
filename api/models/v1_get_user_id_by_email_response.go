// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GetUserIDByEmailResponse v1 get user ID by email response
//
// swagger:model v1GetUserIDByEmailResponse
type V1GetUserIDByEmailResponse struct {

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this v1 get user ID by email response
func (m *V1GetUserIDByEmailResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GetUserIDByEmailResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetUserIDByEmailResponse) UnmarshalBinary(b []byte) error {
	var res V1GetUserIDByEmailResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
