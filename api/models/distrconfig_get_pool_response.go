// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigGetPoolResponse Contains data pool task data specified in by ID request
//
// swagger:model distrconfigGetPoolResponse
type DistrconfigGetPoolResponse struct {

	// data pool
	DataPool *DistrconfigDataPool `json:"dataPool,omitempty"`
}

// Validate validates this distrconfig get pool response
func (m *DistrconfigGetPoolResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataPool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigGetPoolResponse) validateDataPool(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPool) { // not required
		return nil
	}

	if m.DataPool != nil {
		if err := m.DataPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this distrconfig get pool response based on the context it is used
func (m *DistrconfigGetPoolResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigGetPoolResponse) contextValidateDataPool(ctx context.Context, formats strfmt.Registry) error {

	if m.DataPool != nil {
		if err := m.DataPool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataPool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigGetPoolResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigGetPoolResponse) UnmarshalBinary(b []byte) error {
	var res DistrconfigGetPoolResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
