// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigUnassignPoolsFromAppResponse distrconfig unassign pools from app response
//
// swagger:model distrconfigUnassignPoolsFromAppResponse
type DistrconfigUnassignPoolsFromAppResponse struct {

	// failed pool i ds
	FailedPoolIDs []string `json:"failedPoolIDs"`
}

// Validate validates this distrconfig unassign pools from app response
func (m *DistrconfigUnassignPoolsFromAppResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this distrconfig unassign pools from app response based on context it is used
func (m *DistrconfigUnassignPoolsFromAppResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigUnassignPoolsFromAppResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigUnassignPoolsFromAppResponse) UnmarshalBinary(b []byte) error {
	var res DistrconfigUnassignPoolsFromAppResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
