// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1FederationConfig v1 federation config
//
// swagger:model v1FederationConfig
type V1FederationConfig struct {

	// fed ID
	FedID string `json:"fedID,omitempty"`

	// fed type
	FedType V1FederationType `json:"fedType,omitempty"`

	// oidc client ID
	OidcClientID string `json:"oidcClientID,omitempty"`

	// oidc client secret
	OidcClientSecret string `json:"oidcClientSecret,omitempty"`

	// OIDC Config
	OidcEndpoint string `json:"oidcEndpoint,omitempty"`

	// oidc scopes
	OidcScopes string `json:"oidcScopes,omitempty"`

	// oidc use secret
	OidcUseSecret bool `json:"oidcUseSecret,omitempty"`
}

// Validate validates this v1 federation config
func (m *V1FederationConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFedType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1FederationConfig) validateFedType(formats strfmt.Registry) error {

	if swag.IsZero(m.FedType) { // not required
		return nil
	}

	if err := m.FedType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fedType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1FederationConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FederationConfig) UnmarshalBinary(b []byte) error {
	var res V1FederationConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
