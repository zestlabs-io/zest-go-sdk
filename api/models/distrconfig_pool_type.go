// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DistrconfigPoolType distrconfig pool type
//
// swagger:model distrconfigPoolType
type DistrconfigPoolType string

const (

	// DistrconfigPoolTypeUNKNOWN captures enum value "UNKNOWN"
	DistrconfigPoolTypeUNKNOWN DistrconfigPoolType = "UNKNOWN"

	// DistrconfigPoolTypeGLOBAL captures enum value "GLOBAL"
	DistrconfigPoolTypeGLOBAL DistrconfigPoolType = "GLOBAL"

	// DistrconfigPoolTypeFILTERED captures enum value "FILTERED"
	DistrconfigPoolTypeFILTERED DistrconfigPoolType = "FILTERED"

	// DistrconfigPoolTypeUSER captures enum value "USER"
	DistrconfigPoolTypeUSER DistrconfigPoolType = "USER"
)

// for schema
var distrconfigPoolTypeEnum []interface{}

func init() {
	var res []DistrconfigPoolType
	if err := json.Unmarshal([]byte(`["UNKNOWN","GLOBAL","FILTERED","USER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		distrconfigPoolTypeEnum = append(distrconfigPoolTypeEnum, v)
	}
}

func (m DistrconfigPoolType) validateDistrconfigPoolTypeEnum(path, location string, value DistrconfigPoolType) error {
	if err := validate.EnumCase(path, location, value, distrconfigPoolTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this distrconfig pool type
func (m DistrconfigPoolType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDistrconfigPoolTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this distrconfig pool type based on context it is used
func (m DistrconfigPoolType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
