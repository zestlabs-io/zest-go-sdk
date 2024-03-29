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

// DistrConfigServiceUnassignAppFromUsersReader is a Reader for the DistrConfigServiceUnassignAppFromUsers structure.
type DistrConfigServiceUnassignAppFromUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistrConfigServiceUnassignAppFromUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistrConfigServiceUnassignAppFromUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDistrConfigServiceUnassignAppFromUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDistrConfigServiceUnassignAppFromUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDistrConfigServiceUnassignAppFromUsersOK creates a DistrConfigServiceUnassignAppFromUsersOK with default headers values
func NewDistrConfigServiceUnassignAppFromUsersOK() *DistrConfigServiceUnassignAppFromUsersOK {
	return &DistrConfigServiceUnassignAppFromUsersOK{}
}

/* DistrConfigServiceUnassignAppFromUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type DistrConfigServiceUnassignAppFromUsersOK struct {
	Payload *models.DistrconfigUnassignAppFromUsersResponse
}

func (o *DistrConfigServiceUnassignAppFromUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/app/users/unassign][%d] distrConfigServiceUnassignAppFromUsersOK  %+v", 200, o.Payload)
}
func (o *DistrConfigServiceUnassignAppFromUsersOK) GetPayload() *models.DistrconfigUnassignAppFromUsersResponse {
	return o.Payload
}

func (o *DistrConfigServiceUnassignAppFromUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrconfigUnassignAppFromUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceUnassignAppFromUsersNotFound creates a DistrConfigServiceUnassignAppFromUsersNotFound with default headers values
func NewDistrConfigServiceUnassignAppFromUsersNotFound() *DistrConfigServiceUnassignAppFromUsersNotFound {
	return &DistrConfigServiceUnassignAppFromUsersNotFound{}
}

/* DistrConfigServiceUnassignAppFromUsersNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type DistrConfigServiceUnassignAppFromUsersNotFound struct {
	Payload string
}

func (o *DistrConfigServiceUnassignAppFromUsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/app/users/unassign][%d] distrConfigServiceUnassignAppFromUsersNotFound  %+v", 404, o.Payload)
}
func (o *DistrConfigServiceUnassignAppFromUsersNotFound) GetPayload() string {
	return o.Payload
}

func (o *DistrConfigServiceUnassignAppFromUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceUnassignAppFromUsersDefault creates a DistrConfigServiceUnassignAppFromUsersDefault with default headers values
func NewDistrConfigServiceUnassignAppFromUsersDefault(code int) *DistrConfigServiceUnassignAppFromUsersDefault {
	return &DistrConfigServiceUnassignAppFromUsersDefault{
		_statusCode: code,
	}
}

/* DistrConfigServiceUnassignAppFromUsersDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type DistrConfigServiceUnassignAppFromUsersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the distr config service unassign app from users default response
func (o *DistrConfigServiceUnassignAppFromUsersDefault) Code() int {
	return o._statusCode
}

func (o *DistrConfigServiceUnassignAppFromUsersDefault) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/app/users/unassign][%d] DistrConfigService_UnassignAppFromUsers default  %+v", o._statusCode, o.Payload)
}
func (o *DistrConfigServiceUnassignAppFromUsersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DistrConfigServiceUnassignAppFromUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
