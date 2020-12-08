// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsCallFunctionResponse functions call function response
//
// swagger:model functionsCallFunctionResponse
type FunctionsCallFunctionResponse struct {

	// body
	Body interface{} `json:"body,omitempty"`

	// err
	Err string `json:"err,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this functions call function response
func (m *FunctionsCallFunctionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsCallFunctionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsCallFunctionResponse) UnmarshalBinary(b []byte) error {
	var res FunctionsCallFunctionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
