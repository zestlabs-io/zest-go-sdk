// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AccessKey v1 access key
//
// swagger:model v1AccessKey
type V1AccessKey struct {

	// access key ID
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// last used at
	LastUsedAt string `json:"lastUsedAt,omitempty"`
}

// Validate validates this v1 access key
func (m *V1AccessKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AccessKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AccessKey) UnmarshalBinary(b []byte) error {
	var res V1AccessKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
