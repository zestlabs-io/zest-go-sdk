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

// DataListResponse data list response
//
// swagger:model dataListResponse
type DataListResponse struct {

	// Current offset
	Offset float64 `json:"offset,omitempty"`

	// rows
	Rows []*DataDocument `json:"rows"`

	// The number of rows in the database
	TotalRows float64 `json:"total_rows,omitempty"`
}

// Validate validates this data list response
func (m *DataListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataListResponse) validateRows(formats strfmt.Registry) error {

	if swag.IsZero(m.Rows) { // not required
		return nil
	}

	for i := 0; i < len(m.Rows); i++ {
		if swag.IsZero(m.Rows[i]) { // not required
			continue
		}

		if m.Rows[i] != nil {
			if err := m.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataListResponse) UnmarshalBinary(b []byte) error {
	var res DataListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
