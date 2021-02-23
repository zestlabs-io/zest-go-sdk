// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1MetricAggregationOperator v1 metric aggregation operator
//
// swagger:model v1MetricAggregationOperator
type V1MetricAggregationOperator string

const (

	// V1MetricAggregationOperatorSUM captures enum value "SUM"
	V1MetricAggregationOperatorSUM V1MetricAggregationOperator = "SUM"

	// V1MetricAggregationOperatorAVG captures enum value "AVG"
	V1MetricAggregationOperatorAVG V1MetricAggregationOperator = "AVG"

	// V1MetricAggregationOperatorMIN captures enum value "MIN"
	V1MetricAggregationOperatorMIN V1MetricAggregationOperator = "MIN"

	// V1MetricAggregationOperatorMAX captures enum value "MAX"
	V1MetricAggregationOperatorMAX V1MetricAggregationOperator = "MAX"
)

// for schema
var v1MetricAggregationOperatorEnum []interface{}

func init() {
	var res []V1MetricAggregationOperator
	if err := json.Unmarshal([]byte(`["SUM","AVG","MIN","MAX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1MetricAggregationOperatorEnum = append(v1MetricAggregationOperatorEnum, v)
	}
}

func (m V1MetricAggregationOperator) validateV1MetricAggregationOperatorEnum(path, location string, value V1MetricAggregationOperator) error {
	if err := validate.EnumCase(path, location, value, v1MetricAggregationOperatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 metric aggregation operator
func (m V1MetricAggregationOperator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1MetricAggregationOperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
