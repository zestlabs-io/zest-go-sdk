// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1DataPoint A DataPoint represents a single data point in a metrics series and consists
// of a timestamp and the value for the metric at the specific time.
//
// swagger:model v1DataPoint
type V1DataPoint struct {

	// t
	T string `json:"t,omitempty"`

	// v
	V float64 `json:"v,omitempty"`
}

// Validate validates this v1 data point
func (m *V1DataPoint) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1DataPoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1DataPoint) UnmarshalBinary(b []byte) error {
	var res V1DataPoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}