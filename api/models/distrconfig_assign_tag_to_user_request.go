// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigAssignTagToUserRequest distrconfig assign tag to user request
//
// swagger:model distrconfigAssignTagToUserRequest
type DistrconfigAssignTagToUserRequest struct {

	// pool ID
	PoolID string `json:"poolID,omitempty"`

	// tag value
	TagValue string `json:"tagValue,omitempty"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this distrconfig assign tag to user request
func (m *DistrconfigAssignTagToUserRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigAssignTagToUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigAssignTagToUserRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigAssignTagToUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}