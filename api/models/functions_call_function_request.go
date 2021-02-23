// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsCallFunctionRequest functions call function request
//
// swagger:model functionsCallFunctionRequest
type FunctionsCallFunctionRequest struct {

	// id
	ID string `json:"id,omitempty"`

	// payload
	Payload interface{} `json:"payload,omitempty"`

	// payload JSON
	// Format: byte
	PayloadJSON strfmt.Base64 `json:"payloadJSON,omitempty"`
}

// Validate validates this functions call function request
func (m *FunctionsCallFunctionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsCallFunctionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsCallFunctionRequest) UnmarshalBinary(b []byte) error {
	var res FunctionsCallFunctionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
