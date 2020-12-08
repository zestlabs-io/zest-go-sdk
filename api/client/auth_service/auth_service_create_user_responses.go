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

// AuthServiceCreateUserReader is a Reader for the AuthServiceCreateUser structure.
type AuthServiceCreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceCreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceCreateUserOK creates a AuthServiceCreateUserOK with default headers values
func NewAuthServiceCreateUserOK() *AuthServiceCreateUserOK {
	return &AuthServiceCreateUserOK{}
}

/*AuthServiceCreateUserOK handles this case with default header values.

A successful response.
*/
type AuthServiceCreateUserOK struct {
	Payload *models.V1CreateUserResponse
}

func (o *AuthServiceCreateUserOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/user][%d] authServiceCreateUserOK  %+v", 200, o.Payload)
}

func (o *AuthServiceCreateUserOK) GetPayload() *models.V1CreateUserResponse {
	return o.Payload
}

func (o *AuthServiceCreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CreateUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreateUserNotFound creates a AuthServiceCreateUserNotFound with default headers values
func NewAuthServiceCreateUserNotFound() *AuthServiceCreateUserNotFound {
	return &AuthServiceCreateUserNotFound{}
}

/*AuthServiceCreateUserNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceCreateUserNotFound struct {
	Payload string
}

func (o *AuthServiceCreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/user][%d] authServiceCreateUserNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceCreateUserNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceCreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreateUserDefault creates a AuthServiceCreateUserDefault with default headers values
func NewAuthServiceCreateUserDefault(code int) *AuthServiceCreateUserDefault {
	return &AuthServiceCreateUserDefault{
		_statusCode: code,
	}
}

/*AuthServiceCreateUserDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceCreateUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service create user default response
func (o *AuthServiceCreateUserDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceCreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/user][%d] AuthService_CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceCreateUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceCreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}