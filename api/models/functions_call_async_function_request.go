// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsCallAsyncFunctionRequest functions call async function request
//
// swagger:model functionsCallAsyncFunctionRequest
type FunctionsCallAsyncFunctionRequest struct {

	// call
	Call *FunctionsCallFunctionRequest `json:"call,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this functions call async function request
func (m *FunctionsCallAsyncFunctionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionsCallAsyncFunctionRequest) validateCall(formats strfmt.Registry) error {

	if swag.IsZero(m.Call) { // not required
		return nil
	}

	if m.Call != nil {
		if err := m.Call.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("call")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsCallAsyncFunctionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsCallAsyncFunctionRequest) UnmarshalBinary(b []byte) error {
	var res FunctionsCallAsyncFunctionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
