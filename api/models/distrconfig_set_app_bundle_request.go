// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DistrconfigSetAppBundleRequest distrconfig set app bundle request
//
// swagger:model distrconfigSetAppBundleRequest
type DistrconfigSetAppBundleRequest struct {

	// bundle Url
	BundleURL string `json:"bundleUrl,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this distrconfig set app bundle request
func (m *DistrconfigSetAppBundleRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DistrconfigSetAppBundleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistrconfigSetAppBundleRequest) UnmarshalBinary(b []byte) error {
	var res DistrconfigSetAppBundleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
