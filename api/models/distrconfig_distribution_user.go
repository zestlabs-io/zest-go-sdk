// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigDistributionUser distrconfig distribution user
//
// swagger:model distrconfigDistributionUser
type DistrconfigDistributionUser struct {

	// active app ID
	ActiveAppID string `json:"activeAppID,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// tag assignments
	TagAssignments []*DistrconfigUserTagAssignment `json:"tagAssignments"`
}

// Validate validates this distrconfig distribution user
func (m *DistrconfigDistributionUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigDistributionUser) validateTagAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.TagAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.TagAssignments); i++ {
		if swag.IsZero(m.TagAssignments[i]) { // not required
			continue
		}

		if m.TagAssignments[i] != nil {
			if err := m.TagAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this distrconfig distribution user based on the context it is used
func (m *DistrconfigDistributionUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTagAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistrconfigDistributionUser) contextValidateTagAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagAssignments); i++ {

		if m.TagAssignments[i] != nil {
			if err := m.TagAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigDistributionUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigDistributionUser) UnmarshalBinary(b []byte) error {
	var res DistrconfigDistributionUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
