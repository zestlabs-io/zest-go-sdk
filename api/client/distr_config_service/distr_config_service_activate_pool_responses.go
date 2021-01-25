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

// DistrConfigServiceActivatePoolReader is a Reader for the DistrConfigServiceActivatePool structure.
type DistrConfigServiceActivatePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistrConfigServiceActivatePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistrConfigServiceActivatePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDistrConfigServiceActivatePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDistrConfigServiceActivatePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDistrConfigServiceActivatePoolOK creates a DistrConfigServiceActivatePoolOK with default headers values
func NewDistrConfigServiceActivatePoolOK() *DistrConfigServiceActivatePoolOK {
	return &DistrConfigServiceActivatePoolOK{}
}

/* DistrConfigServiceActivatePoolOK describes a response with status code 200, with default header values.

A successful response.
*/
type DistrConfigServiceActivatePoolOK struct {
	Payload models.DistrconfigActivatePoolResponse
}

func (o *DistrConfigServiceActivatePoolOK) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool/activate][%d] distrConfigServiceActivatePoolOK  %+v", 200, o.Payload)
}
func (o *DistrConfigServiceActivatePoolOK) GetPayload() models.DistrconfigActivatePoolResponse {
	return o.Payload
}

func (o *DistrConfigServiceActivatePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceActivatePoolNotFound creates a DistrConfigServiceActivatePoolNotFound with default headers values
func NewDistrConfigServiceActivatePoolNotFound() *DistrConfigServiceActivatePoolNotFound {
	return &DistrConfigServiceActivatePoolNotFound{}
}

/* DistrConfigServiceActivatePoolNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DistrConfigServiceActivatePoolNotFound struct {
	Payload string
}

func (o *DistrConfigServiceActivatePoolNotFound) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool/activate][%d] distrConfigServiceActivatePoolNotFound  %+v", 404, o.Payload)
}
func (o *DistrConfigServiceActivatePoolNotFound) GetPayload() string {
	return o.Payload
}

func (o *DistrConfigServiceActivatePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceActivatePoolDefault creates a DistrConfigServiceActivatePoolDefault with default headers values
func NewDistrConfigServiceActivatePoolDefault(code int) *DistrConfigServiceActivatePoolDefault {
	return &DistrConfigServiceActivatePoolDefault{
		_statusCode: code,
	}
}

/* DistrConfigServiceActivatePoolDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type DistrConfigServiceActivatePoolDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the distr config service activate pool default response
func (o *DistrConfigServiceActivatePoolDefault) Code() int {
	return o._statusCode
}

func (o *DistrConfigServiceActivatePoolDefault) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool/activate][%d] DistrConfigService_ActivatePool default  %+v", o._statusCode, o.Payload)
}
func (o *DistrConfigServiceActivatePoolDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DistrConfigServiceActivatePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
