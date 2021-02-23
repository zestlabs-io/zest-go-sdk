// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserInfo v1 user info
//
// swagger:model v1UserInfo
type V1UserInfo struct {

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// logout Url
	LogoutURL string `json:"logoutUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// policies
	Policies []*V1Policy `json:"policies"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this v1 user info
func (m *V1UserInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UserInfo) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserInfo) UnmarshalBinary(b []byte) error {
	var res V1UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
