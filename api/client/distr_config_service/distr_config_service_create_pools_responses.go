// Code generated by go-swagger; DO NOT EDIT.

package distr_config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// DistrConfigServiceCreatePoolsReader is a Reader for the DistrConfigServiceCreatePools structure.
type DistrConfigServiceCreatePoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistrConfigServiceCreatePoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistrConfigServiceCreatePoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDistrConfigServiceCreatePoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDistrConfigServiceCreatePoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDistrConfigServiceCreatePoolsOK creates a DistrConfigServiceCreatePoolsOK with default headers values
func NewDistrConfigServiceCreatePoolsOK() *DistrConfigServiceCreatePoolsOK {
	return &DistrConfigServiceCreatePoolsOK{}
}

/* DistrConfigServiceCreatePoolsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DistrConfigServiceCreatePoolsOK struct {
	Payload models.DistrconfigCreatePoolsResponse
}

func (o *DistrConfigServiceCreatePoolsOK) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pools][%d] distrConfigServiceCreatePoolsOK  %+v", 200, o.Payload)
}
func (o *DistrConfigServiceCreatePoolsOK) GetPayload() models.DistrconfigCreatePoolsResponse {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceCreatePoolsNotFound creates a DistrConfigServiceCreatePoolsNotFound with default headers values
func NewDistrConfigServiceCreatePoolsNotFound() *DistrConfigServiceCreatePoolsNotFound {
	return &DistrConfigServiceCreatePoolsNotFound{}
}

/* DistrConfigServiceCreatePoolsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DistrConfigServiceCreatePoolsNotFound struct {
	Payload string
}

func (o *DistrConfigServiceCreatePoolsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pools][%d] distrConfigServiceCreatePoolsNotFound  %+v", 404, o.Payload)
}
func (o *DistrConfigServiceCreatePoolsNotFound) GetPayload() string {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceCreatePoolsDefault creates a DistrConfigServiceCreatePoolsDefault with default headers values
func NewDistrConfigServiceCreatePoolsDefault(code int) *DistrConfigServiceCreatePoolsDefault {
	return &DistrConfigServiceCreatePoolsDefault{
		_statusCode: code,
	}
}

/* DistrConfigServiceCreatePoolsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type DistrConfigServiceCreatePoolsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the distr config service create pools default response
func (o *DistrConfigServiceCreatePoolsDefault) Code() int {
	return o._statusCode
}

func (o *DistrConfigServiceCreatePoolsDefault) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pools][%d] DistrConfigService_CreatePools default  %+v", o._statusCode, o.Payload)
}
func (o *DistrConfigServiceCreatePoolsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
