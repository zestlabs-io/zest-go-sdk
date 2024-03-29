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

// DistrconfigGetPoolDistributionResponse distrconfig get pool distribution response
//
// swagger:model distrconfigGetPoolDistributionResponse
type DistrconfigGetPoolDistributionResponse struct {

	// db Url
	DbURL string `json:"dbUrl,omitempty"`

	// The type of the pool - User, Filtered(with tags), Global (for the app)
	PoolType *DistrconfigPoolType `json:"poolType,omitempty"`

	// Tag prefix used for determining if user assignments match the record
	TagPrefix string `json:"tagPrefix,omitempty"`
}

// Validate validates this distrconfig get pool distribution response
func (m *DistrconfigGetPoolDistributionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePoolType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigGetPoolDistributionResponse) validatePoolType(formats strfmt.Registry) error {
	if swag.IsZero(m.PoolType) { // not required
		return nil
	}

	if m.PoolType != nil {
		if err := m.PoolType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poolType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poolType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this distrconfig get pool distribution response based on the context it is used
func (m *DistrconfigGetPoolDistributionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePoolType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigGetPoolDistributionResponse) contextValidatePoolType(ctx context.Context, formats strfmt.Registry) error {

	if m.PoolType != nil {
		if err := m.PoolType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("poolType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("poolType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigGetPoolDistributionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigGetPoolDistributionResponse) UnmarshalBinary(b []byte) error {
	var res DistrconfigGetPoolDistributionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
