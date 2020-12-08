// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FunctionsGetFunctionResponse functions get function response
//
// swagger:model functionsGetFunctionResponse
type FunctionsGetFunctionResponse struct {

	// function
	Function *FunctionsFunction `json:"function,omitempty"`
}

// Validate validates this functions get function response
func (m *FunctionsGetFunctionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FunctionsGetFunctionResponse) validateFunction(formats strfmt.Registry) error {

	if swag.IsZero(m.Function) { // not required
		return nil
	}

	if m.Function != nil {
		if err := m.Function.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FunctionsGetFunctionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionsGetFunctionResponse) UnmarshalBinary(b []byte) error {
	var res FunctionsGetFunctionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}