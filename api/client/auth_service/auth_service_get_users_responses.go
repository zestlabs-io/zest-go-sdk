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

// AuthServiceGetUsersReader is a Reader for the AuthServiceGetUsers structure.
type AuthServiceGetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceGetUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetUsersOK creates a AuthServiceGetUsersOK with default headers values
func NewAuthServiceGetUsersOK() *AuthServiceGetUsersOK {
	return &AuthServiceGetUsersOK{}
}

/*AuthServiceGetUsersOK handles this case with default header values.

A successful response.
*/
type AuthServiceGetUsersOK struct {
	Payload *models.V1GetUsersResponse
}

func (o *AuthServiceGetUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/users][%d] authServiceGetUsersOK  %+v", 200, o.Payload)
}

func (o *AuthServiceGetUsersOK) GetPayload() *models.V1GetUsersResponse {
	return o.Payload
}

func (o *AuthServiceGetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetUsersNotFound creates a AuthServiceGetUsersNotFound with default headers values
func NewAuthServiceGetUsersNotFound() *AuthServiceGetUsersNotFound {
	return &AuthServiceGetUsersNotFound{}
}

/*AuthServiceGetUsersNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetUsersNotFound struct {
	Payload string
}

func (o *AuthServiceGetUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/users][%d] authServiceGetUsersNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceGetUsersNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetUsersDefault creates a AuthServiceGetUsersDefault with default headers values
func NewAuthServiceGetUsersDefault(code int) *AuthServiceGetUsersDefault {
	return &AuthServiceGetUsersDefault{
		_statusCode: code,
	}
}

/*AuthServiceGetUsersDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceGetUsersDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get users default response
func (o *AuthServiceGetUsersDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetUsersDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/users][%d] AuthService_GetUsers default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceGetUsersDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
