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

// AuthServiceDeleteUserReader is a Reader for the AuthServiceDeleteUser structure.
type AuthServiceDeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceDeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceDeleteUserOK creates a AuthServiceDeleteUserOK with default headers values
func NewAuthServiceDeleteUserOK() *AuthServiceDeleteUserOK {
	return &AuthServiceDeleteUserOK{}
}

/* AuthServiceDeleteUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceDeleteUserOK struct {
	Payload models.V1DeleteUserResponse
}

func (o *AuthServiceDeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/user/{userID}][%d] authServiceDeleteUserOK  %+v", 200, o.Payload)
}
func (o *AuthServiceDeleteUserOK) GetPayload() models.V1DeleteUserResponse {
	return o.Payload
}

func (o *AuthServiceDeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceDeleteUserNotFound creates a AuthServiceDeleteUserNotFound with default headers values
func NewAuthServiceDeleteUserNotFound() *AuthServiceDeleteUserNotFound {
	return &AuthServiceDeleteUserNotFound{}
}

/* AuthServiceDeleteUserNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceDeleteUserNotFound struct {
	Payload string
}

func (o *AuthServiceDeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/user/{userID}][%d] authServiceDeleteUserNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceDeleteUserNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceDeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceDeleteUserDefault creates a AuthServiceDeleteUserDefault with default headers values
func NewAuthServiceDeleteUserDefault(code int) *AuthServiceDeleteUserDefault {
	return &AuthServiceDeleteUserDefault{
		_statusCode: code,
	}
}

/* AuthServiceDeleteUserDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceDeleteUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service delete user default response
func (o *AuthServiceDeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceDeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/user/{userID}][%d] AuthService_DeleteUser default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceDeleteUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceDeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
