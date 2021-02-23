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

// DistrConfigServiceCreatePoolReader is a Reader for the DistrConfigServiceCreatePool structure.
type DistrConfigServiceCreatePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DistrConfigServiceCreatePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDistrConfigServiceCreatePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDistrConfigServiceCreatePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDistrConfigServiceCreatePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDistrConfigServiceCreatePoolOK creates a DistrConfigServiceCreatePoolOK with default headers values
func NewDistrConfigServiceCreatePoolOK() *DistrConfigServiceCreatePoolOK {
	return &DistrConfigServiceCreatePoolOK{}
}

/*DistrConfigServiceCreatePoolOK handles this case with default header values.

A successful response.
*/
type DistrConfigServiceCreatePoolOK struct {
	Payload models.DistrconfigCreatePoolResponse
}

func (o *DistrConfigServiceCreatePoolOK) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool][%d] distrConfigServiceCreatePoolOK  %+v", 200, o.Payload)
}

func (o *DistrConfigServiceCreatePoolOK) GetPayload() models.DistrconfigCreatePoolResponse {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceCreatePoolNotFound creates a DistrConfigServiceCreatePoolNotFound with default headers values
func NewDistrConfigServiceCreatePoolNotFound() *DistrConfigServiceCreatePoolNotFound {
	return &DistrConfigServiceCreatePoolNotFound{}
}

/*DistrConfigServiceCreatePoolNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type DistrConfigServiceCreatePoolNotFound struct {
	Payload string
}

func (o *DistrConfigServiceCreatePoolNotFound) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool][%d] distrConfigServiceCreatePoolNotFound  %+v", 404, o.Payload)
}

func (o *DistrConfigServiceCreatePoolNotFound) GetPayload() string {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDistrConfigServiceCreatePoolDefault creates a DistrConfigServiceCreatePoolDefault with default headers values
func NewDistrConfigServiceCreatePoolDefault(code int) *DistrConfigServiceCreatePoolDefault {
	return &DistrConfigServiceCreatePoolDefault{
		_statusCode: code,
	}
}

/*DistrConfigServiceCreatePoolDefault handles this case with default header values.

An unexpected error response
*/
type DistrConfigServiceCreatePoolDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the distr config service create pool default response
func (o *DistrConfigServiceCreatePoolDefault) Code() int {
	return o._statusCode
}

func (o *DistrConfigServiceCreatePoolDefault) Error() string {
	return fmt.Sprintf("[POST /api/distribution/v1/pool][%d] DistrConfigService_CreatePool default  %+v", o._statusCode, o.Payload)
}

func (o *DistrConfigServiceCreatePoolDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *DistrConfigServiceCreatePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
