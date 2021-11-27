// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsCreateFunctionRequest functions create function request
//
// swagger:model functionsCreateFunctionRequest
type FunctionsCreateFunctionRequest struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// func type
	FuncType string `json:"funcType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this functions create function request
func (m *FunctionsCreateFunctionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this functions create function request based on context it is used
func (m *FunctionsCreateFunctionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsCreateFunctionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsCreateFunctionRequest) UnmarshalBinary(b []byte) error {
	var res FunctionsCreateFunctionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
