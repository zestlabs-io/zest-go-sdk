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

// AuthServiceGetRolesReader is a Reader for the AuthServiceGetRoles structure.
type AuthServiceGetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceGetRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetRolesOK creates a AuthServiceGetRolesOK with default headers values
func NewAuthServiceGetRolesOK() *AuthServiceGetRolesOK {
	return &AuthServiceGetRolesOK{}
}

/*AuthServiceGetRolesOK handles this case with default header values.

A successful response.
*/
type AuthServiceGetRolesOK struct {
	Payload *models.V1GetRolesResponse
}

func (o *AuthServiceGetRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/roles][%d] authServiceGetRolesOK  %+v", 200, o.Payload)
}

func (o *AuthServiceGetRolesOK) GetPayload() *models.V1GetRolesResponse {
	return o.Payload
}

func (o *AuthServiceGetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetRolesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetRolesNotFound creates a AuthServiceGetRolesNotFound with default headers values
func NewAuthServiceGetRolesNotFound() *AuthServiceGetRolesNotFound {
	return &AuthServiceGetRolesNotFound{}
}

/*AuthServiceGetRolesNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetRolesNotFound struct {
	Payload string
}

func (o *AuthServiceGetRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/roles][%d] authServiceGetRolesNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceGetRolesNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetRolesDefault creates a AuthServiceGetRolesDefault with default headers values
func NewAuthServiceGetRolesDefault(code int) *AuthServiceGetRolesDefault {
	return &AuthServiceGetRolesDefault{
		_statusCode: code,
	}
}

/*AuthServiceGetRolesDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceGetRolesDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get roles default response
func (o *AuthServiceGetRolesDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetRolesDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/roles][%d] AuthService_GetRoles default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceGetRolesDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
