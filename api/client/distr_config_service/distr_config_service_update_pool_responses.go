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

// DistrConfigServiceUpdatePoolReader is a Reader for the DistrConfigServiceUpdatePool structure.
type DistrConfigServiceUpdatePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistrConfigServiceUpdatePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistrConfigServiceUpdatePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDistrConfigServiceUpdatePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDistrConfigServiceUpdatePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDistrConfigServiceUpdatePoolOK creates a DistrConfigServiceUpdatePoolOK with default headers values
func NewDistrConfigServiceUpdatePoolOK() *DistrConfigServiceUpdatePoolOK {
	return &DistrConfigServiceUpdatePoolOK{}
}

/* DistrConfigServiceUpdatePoolOK describes a response with status code 200, with default header values.

A successful response.
*/
type DistrConfigServiceUpdatePoolOK struct {
	Payload models.DistrconfigUpdatePoolResponse
}

func (o *DistrConfigServiceUpdatePoolOK) Error() string {
	return fmt.Sprintf("[PUT /api/distribution/v1/pool][%d] distrConfigServiceUpdatePoolOK  %+v", 200, o.Payload)
}
func (o *DistrConfigServiceUpdatePoolOK) GetPayload() models.DistrconfigUpdatePoolResponse {
	return o.Payload
}

func (o *DistrConfigServiceUpdatePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceUpdatePoolNotFound creates a DistrConfigServiceUpdatePoolNotFound with default headers values
func NewDistrConfigServiceUpdatePoolNotFound() *DistrConfigServiceUpdatePoolNotFound {
	return &DistrConfigServiceUpdatePoolNotFound{}
}

/* DistrConfigServiceUpdatePoolNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DistrConfigServiceUpdatePoolNotFound struct {
	Payload string
}

func (o *DistrConfigServiceUpdatePoolNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/distribution/v1/pool][%d] distrConfigServiceUpdatePoolNotFound  %+v", 404, o.Payload)
}
func (o *DistrConfigServiceUpdatePoolNotFound) GetPayload() string {
	return o.Payload
}

func (o *DistrConfigServiceUpdatePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceUpdatePoolDefault creates a DistrConfigServiceUpdatePoolDefault with default headers values
func NewDistrConfigServiceUpdatePoolDefault(code int) *DistrConfigServiceUpdatePoolDefault {
	return &DistrConfigServiceUpdatePoolDefault{
		_statusCode: code,
	}
}

/* DistrConfigServiceUpdatePoolDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type DistrConfigServiceUpdatePoolDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the distr config service update pool default response
func (o *DistrConfigServiceUpdatePoolDefault) Code() int {
	return o._statusCode
}

func (o *DistrConfigServiceUpdatePoolDefault) Error() string {
	return fmt.Sprintf("[PUT /api/distribution/v1/pool][%d] DistrConfigService_UpdatePool default  %+v", o._statusCode, o.Payload)
}
func (o *DistrConfigServiceUpdatePoolDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DistrConfigServiceUpdatePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
