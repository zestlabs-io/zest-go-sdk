// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigUnassignAppFromUsersRequest distrconfig unassign app from users request
//
// swagger:model distrconfigUnassignAppFromUsersRequest
type DistrconfigUnassignAppFromUsersRequest struct {

	// app ID
	AppID string `json:"appID,omitempty"`

	// user i ds
	UserIDs []string `json:"userIDs"`
}

// Validate validates this distrconfig unassign app from users request
func (m *DistrconfigUnassignAppFromUsersRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this distrconfig unassign app from users request based on context it is used
func (m *DistrconfigUnassignAppFromUsersRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigUnassignAppFromUsersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigUnassignAppFromUsersRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigUnassignAppFromUsersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
