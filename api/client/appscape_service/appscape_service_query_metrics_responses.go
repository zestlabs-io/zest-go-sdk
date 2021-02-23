// Code generated by go-swagger; DO NOT EDIT.

package appscape_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// AppscapeServiceQueryMetricsReader is a Reader for the AppscapeServiceQueryMetrics structure.
type AppscapeServiceQueryMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppscapeServiceQueryMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppscapeServiceQueryMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppscapeServiceQueryMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppscapeServiceQueryMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppscapeServiceQueryMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppscapeServiceQueryMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppscapeServiceQueryMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppscapeServiceQueryMetricsOK creates a AppscapeServiceQueryMetricsOK with default headers values
func NewAppscapeServiceQueryMetricsOK() *AppscapeServiceQueryMetricsOK {
	return &AppscapeServiceQueryMetricsOK{}
}

/*AppscapeServiceQueryMetricsOK handles this case with default header values.

A successful response.
*/
type AppscapeServiceQueryMetricsOK struct {
	Payload *models.V1QueryMetricsResponse
}

func (o *AppscapeServiceQueryMetricsOK) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] appscapeServiceQueryMetricsOK  %+v", 200, o.Payload)
}

func (o *AppscapeServiceQueryMetricsOK) GetPayload() *models.V1QueryMetricsResponse {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1QueryMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceQueryMetricsBadRequest creates a AppscapeServiceQueryMetricsBadRequest with default headers values
func NewAppscapeServiceQueryMetricsBadRequest() *AppscapeServiceQueryMetricsBadRequest {
	return &AppscapeServiceQueryMetricsBadRequest{}
}

/*AppscapeServiceQueryMetricsBadRequest handles this case with default header values.

Returned when the caller provided incorrect request parameters.
*/
type AppscapeServiceQueryMetricsBadRequest struct {
	Payload string
}

func (o *AppscapeServiceQueryMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] appscapeServiceQueryMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *AppscapeServiceQueryMetricsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceQueryMetricsForbidden creates a AppscapeServiceQueryMetricsForbidden with default headers values
func NewAppscapeServiceQueryMetricsForbidden() *AppscapeServiceQueryMetricsForbidden {
	return &AppscapeServiceQueryMetricsForbidden{}
}

/*AppscapeServiceQueryMetricsForbidden handles this case with default header values.

Returned when the caller is not authorised to perform this call.
*/
type AppscapeServiceQueryMetricsForbidden struct {
	Payload string
}

func (o *AppscapeServiceQueryMetricsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] appscapeServiceQueryMetricsForbidden  %+v", 403, o.Payload)
}

func (o *AppscapeServiceQueryMetricsForbidden) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceQueryMetricsNotFound creates a AppscapeServiceQueryMetricsNotFound with default headers values
func NewAppscapeServiceQueryMetricsNotFound() *AppscapeServiceQueryMetricsNotFound {
	return &AppscapeServiceQueryMetricsNotFound{}
}

/*AppscapeServiceQueryMetricsNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AppscapeServiceQueryMetricsNotFound struct {
	Payload string
}

func (o *AppscapeServiceQueryMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] appscapeServiceQueryMetricsNotFound  %+v", 404, o.Payload)
}

func (o *AppscapeServiceQueryMetricsNotFound) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceQueryMetricsInternalServerError creates a AppscapeServiceQueryMetricsInternalServerError with default headers values
func NewAppscapeServiceQueryMetricsInternalServerError() *AppscapeServiceQueryMetricsInternalServerError {
	return &AppscapeServiceQueryMetricsInternalServerError{}
}

/*AppscapeServiceQueryMetricsInternalServerError handles this case with default header values.

Returned when an error occurred while processing the process.
*/
type AppscapeServiceQueryMetricsInternalServerError struct {
	Payload string
}

func (o *AppscapeServiceQueryMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] appscapeServiceQueryMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *AppscapeServiceQueryMetricsInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceQueryMetricsDefault creates a AppscapeServiceQueryMetricsDefault with default headers values
func NewAppscapeServiceQueryMetricsDefault(code int) *AppscapeServiceQueryMetricsDefault {
	return &AppscapeServiceQueryMetricsDefault{
		_statusCode: code,
	}
}

/*AppscapeServiceQueryMetricsDefault handles this case with default header values.

An unexpected error response
*/
type AppscapeServiceQueryMetricsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the appscape service query metrics default response
func (o *AppscapeServiceQueryMetricsDefault) Code() int {
	return o._statusCode
}

func (o *AppscapeServiceQueryMetricsDefault) Error() string {
	return fmt.Sprintf("[POST /api/appscape/v1/metrics/query][%d] AppscapeService_QueryMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *AppscapeServiceQueryMetricsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AppscapeServiceQueryMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}