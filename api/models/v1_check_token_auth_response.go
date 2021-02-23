// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1CheckTokenAuthResponse v1 check token auth response
//
// swagger:model v1CheckTokenAuthResponse
type V1CheckTokenAuthResponse struct {

	// error
	Error string `json:"error,omitempty"`

	// token
	// Format: byte
	Token strfmt.Base64 `json:"token,omitempty"`
}

// Validate validates this v1 check token auth response
func (m *V1CheckTokenAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1CheckTokenAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CheckTokenAuthResponse) UnmarshalBinary(b []byte) error {
	var res V1CheckTokenAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
