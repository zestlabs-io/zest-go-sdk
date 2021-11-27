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

// DistrconfigValidatePoolDataRequest distrconfig validate pool data request
//
// swagger:model distrconfigValidatePoolDataRequest
type DistrconfigValidatePoolDataRequest struct {

	// data
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// http://goessner.net/articles/JsonPath/ expression to extract the primary key
	PkExtractExpression string `json:"pkExtractExpression,omitempty"`

	// The type of the pool - User, Filtered(with tags), Global (for the app)
	PoolType *DistrconfigPoolType `json:"poolType,omitempty"`

	// If set - should contain http://json-schema.org/ schema, that will validate every document
	// Format: byte
	Schema strfmt.Base64 `json:"schema,omitempty"`

	// http://goessner.net/articles/JsonPath/ expression to extract tags from the documents. For user type pools this is the user id
	TagExtractExpression string `json:"tagExtractExpression,omitempty"`
}

// Validate validates this distrconfig validate pool data request
func (m *DistrconfigValidatePoolDataRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePoolType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigValidatePoolDataRequest) validatePoolType(formats strfmt.Registry) error {
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

// ContextValidate validate this distrconfig validate pool data request based on the context it is used
func (m *DistrconfigValidatePoolDataRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePoolType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigValidatePoolDataRequest) contextValidatePoolType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DistrconfigValidatePoolDataRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigValidatePoolDataRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigValidatePoolDataRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
