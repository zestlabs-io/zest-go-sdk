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

// AppscapeServiceGetLogLabelsReader is a Reader for the AppscapeServiceGetLogLabels structure.
type AppscapeServiceGetLogLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppscapeServiceGetLogLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppscapeServiceGetLogLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAppscapeServiceGetLogLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAppscapeServiceGetLogLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAppscapeServiceGetLogLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAppscapeServiceGetLogLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAppscapeServiceGetLogLabelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppscapeServiceGetLogLabelsOK creates a AppscapeServiceGetLogLabelsOK with default headers values
func NewAppscapeServiceGetLogLabelsOK() *AppscapeServiceGetLogLabelsOK {
	return &AppscapeServiceGetLogLabelsOK{}
}

/* AppscapeServiceGetLogLabelsOK describes a response with status code 200, with default header values.

A successful response.
*/
type AppscapeServiceGetLogLabelsOK struct {
	Payload *models.V1GetLogLabelsResponse
}

func (o *AppscapeServiceGetLogLabelsOK) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] appscapeServiceGetLogLabelsOK  %+v", 200, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsOK) GetPayload() *models.V1GetLogLabelsResponse {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetLogLabelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceGetLogLabelsBadRequest creates a AppscapeServiceGetLogLabelsBadRequest with default headers values
func NewAppscapeServiceGetLogLabelsBadRequest() *AppscapeServiceGetLogLabelsBadRequest {
	return &AppscapeServiceGetLogLabelsBadRequest{}
}

/* AppscapeServiceGetLogLabelsBadRequest describes a response with status code 400, with default header values.

Returned when the caller provided incorrect request parameters.
*/
type AppscapeServiceGetLogLabelsBadRequest struct {
	Payload string
}

func (o *AppscapeServiceGetLogLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] appscapeServiceGetLogLabelsBadRequest  %+v", 400, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceGetLogLabelsForbidden creates a AppscapeServiceGetLogLabelsForbidden with default headers values
func NewAppscapeServiceGetLogLabelsForbidden() *AppscapeServiceGetLogLabelsForbidden {
	return &AppscapeServiceGetLogLabelsForbidden{}
}

/* AppscapeServiceGetLogLabelsForbidden describes a response with status code 403, with default header values.

Returned when the caller is not authorised to perform this call.
*/
type AppscapeServiceGetLogLabelsForbidden struct {
	Payload string
}

func (o *AppscapeServiceGetLogLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] appscapeServiceGetLogLabelsForbidden  %+v", 403, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsForbidden) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceGetLogLabelsNotFound creates a AppscapeServiceGetLogLabelsNotFound with default headers values
func NewAppscapeServiceGetLogLabelsNotFound() *AppscapeServiceGetLogLabelsNotFound {
	return &AppscapeServiceGetLogLabelsNotFound{}
}

/* AppscapeServiceGetLogLabelsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AppscapeServiceGetLogLabelsNotFound struct {
	Payload string
}

func (o *AppscapeServiceGetLogLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] appscapeServiceGetLogLabelsNotFound  %+v", 404, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsNotFound) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceGetLogLabelsInternalServerError creates a AppscapeServiceGetLogLabelsInternalServerError with default headers values
func NewAppscapeServiceGetLogLabelsInternalServerError() *AppscapeServiceGetLogLabelsInternalServerError {
	return &AppscapeServiceGetLogLabelsInternalServerError{}
}

/* AppscapeServiceGetLogLabelsInternalServerError describes a response with status code 500, with default header values.

Returned when an error occurred while processing the process.
*/
type AppscapeServiceGetLogLabelsInternalServerError struct {
	Payload string
}

func (o *AppscapeServiceGetLogLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] appscapeServiceGetLogLabelsInternalServerError  %+v", 500, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppscapeServiceGetLogLabelsDefault creates a AppscapeServiceGetLogLabelsDefault with default headers values
func NewAppscapeServiceGetLogLabelsDefault(code int) *AppscapeServiceGetLogLabelsDefault {
	return &AppscapeServiceGetLogLabelsDefault{
		_statusCode: code,
	}
}

/* AppscapeServiceGetLogLabelsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AppscapeServiceGetLogLabelsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the appscape service get log labels default response
func (o *AppscapeServiceGetLogLabelsDefault) Code() int {
	return o._statusCode
}

func (o *AppscapeServiceGetLogLabelsDefault) Error() string {
	return fmt.Sprintf("[GET /api/appscape/v1/logs/labels][%d] AppscapeService_GetLogLabels default  %+v", o._statusCode, o.Payload)
}
func (o *AppscapeServiceGetLogLabelsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AppscapeServiceGetLogLabelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
