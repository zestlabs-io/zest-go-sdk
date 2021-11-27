// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CheckUsernameExistsResponse v1 check username exists response
//
// swagger:model v1CheckUsernameExistsResponse
type V1CheckUsernameExistsResponse struct {

	// exists
	Exists bool `json:"exists,omitempty"`
}

// Validate validates this v1 check username exists response
func (m *V1CheckUsernameExistsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 check username exists response based on context it is used
func (m *V1CheckUsernameExistsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CheckUsernameExistsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CheckUsernameExistsResponse) UnmarshalBinary(b []byte) error {
	var res V1CheckUsernameExistsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
