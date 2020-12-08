// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigAssignAppToUsersRequest distrconfig assign app to users request
//
// swagger:model distrconfigAssignAppToUsersRequest
type DistrconfigAssignAppToUsersRequest struct {

	// app ID
	AppID string `json:"appID,omitempty"`

	// user i ds
	UserIDs []string `json:"userIDs"`
}

// Validate validates this distrconfig assign app to users request
func (m *DistrconfigAssignAppToUsersRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigAssignAppToUsersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigAssignAppToUsersRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigAssignAppToUsersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}