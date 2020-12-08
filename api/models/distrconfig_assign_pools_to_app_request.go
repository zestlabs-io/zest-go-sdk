// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigAssignPoolsToAppRequest distrconfig assign pools to app request
//
// swagger:model distrconfigAssignPoolsToAppRequest
type DistrconfigAssignPoolsToAppRequest struct {

	// app ID
	AppID string `json:"appID,omitempty"`

	// pool i ds
	PoolIDs []string `json:"poolIDs"`
}

// Validate validates this distrconfig assign pools to app request
func (m *DistrconfigAssignPoolsToAppRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigAssignPoolsToAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigAssignPoolsToAppRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigAssignPoolsToAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
