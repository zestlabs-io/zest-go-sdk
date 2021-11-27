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

// LogEntrySeverity log entry severity
//
// swagger:model LogEntrySeverity
type LogEntrySeverity string

func NewLogEntrySeverity(value LogEntrySeverity) *LogEntrySeverity {
	v := value
	return &v
}

const (

	// LogEntrySeverityTRACE captures enum value "TRACE"
	LogEntrySeverityTRACE LogEntrySeverity = "TRACE"

	// LogEntrySeverityINFO captures enum value "INFO"
	LogEntrySeverityINFO LogEntrySeverity = "INFO"

	// LogEntrySeverityWARNING captures enum value "WARNING"
	LogEntrySeverityWARNING LogEntrySeverity = "WARNING"

	// LogEntrySeverityERROR captures enum value "ERROR"
	LogEntrySeverityERROR LogEntrySeverity = "ERROR"

	// LogEntrySeverityAUDIT captures enum value "AUDIT"
	LogEntrySeverityAUDIT LogEntrySeverity = "AUDIT"

	// LogEntrySeverityALL captures enum value "ALL"
	LogEntrySeverityALL LogEntrySeverity = "ALL"
)

// for schema
var logEntrySeverityEnum []interface{}

func init() {
	var res []LogEntrySeverity
	if err := json.Unmarshal([]byte(`["TRACE","INFO","WARNING","ERROR","AUDIT","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logEntrySeverityEnum = append(logEntrySeverityEnum, v)
	}
}

func (m LogEntrySeverity) validateLogEntrySeverityEnum(path, location string, value LogEntrySeverity) error {
	if err := validate.EnumCase(path, location, value, logEntrySeverityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this log entry severity
func (m LogEntrySeverity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLogEntrySeverityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this log entry severity based on context it is used
func (m LogEntrySeverity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
