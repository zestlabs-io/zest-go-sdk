// Code generated by go-swagger; DO NOT EDIT.

package auth_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// AuthServiceCreateAccessKeyReader is a Reader for the AuthServiceCreateAccessKey structure.
type AuthServiceCreateAccessKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceCreateAccessKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceCreateAccessKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceCreateAccessKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceCreateAccessKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceCreateAccessKeyOK creates a AuthServiceCreateAccessKeyOK with default headers values
func NewAuthServiceCreateAccessKeyOK() *AuthServiceCreateAccessKeyOK {
	return &AuthServiceCreateAccessKeyOK{}
}

/*AuthServiceCreateAccessKeyOK handles this case with default header values.

A successful response.
*/
type AuthServiceCreateAccessKeyOK struct {
	Payload *models.V1CreateAccessKeyResponse
}

func (o *AuthServiceCreateAccessKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/accesskey][%d] authServiceCreateAccessKeyOK  %+v", 200, o.Payload)
}

func (o *AuthServiceCreateAccessKeyOK) GetPayload() *models.V1CreateAccessKeyResponse {
	return o.Payload
}

func (o *AuthServiceCreateAccessKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CreateAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreateAccessKeyNotFound creates a AuthServiceCreateAccessKeyNotFound with default headers values
func NewAuthServiceCreateAccessKeyNotFound() *AuthServiceCreateAccessKeyNotFound {
	return &AuthServiceCreateAccessKeyNotFound{}
}

/*AuthServiceCreateAccessKeyNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceCreateAccessKeyNotFound struct {
	Payload string
}

func (o *AuthServiceCreateAccessKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/accesskey][%d] authServiceCreateAccessKeyNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceCreateAccessKeyNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceCreateAccessKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreateAccessKeyDefault creates a AuthServiceCreateAccessKeyDefault with default headers values
func NewAuthServiceCreateAccessKeyDefault(code int) *AuthServiceCreateAccessKeyDefault {
	return &AuthServiceCreateAccessKeyDefault{
		_statusCode: code,
	}
}

/*AuthServiceCreateAccessKeyDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceCreateAccessKeyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service create access key default response
func (o *AuthServiceCreateAccessKeyDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceCreateAccessKeyDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/accesskey][%d] AuthService_CreateAccessKey default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceCreateAccessKeyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceCreateAccessKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
