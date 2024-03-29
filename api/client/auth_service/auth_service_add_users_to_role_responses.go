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

// AuthServiceAddUsersToRoleReader is a Reader for the AuthServiceAddUsersToRole structure.
type AuthServiceAddUsersToRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceAddUsersToRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceAddUsersToRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceAddUsersToRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceAddUsersToRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceAddUsersToRoleOK creates a AuthServiceAddUsersToRoleOK with default headers values
func NewAuthServiceAddUsersToRoleOK() *AuthServiceAddUsersToRoleOK {
	return &AuthServiceAddUsersToRoleOK{}
}

/* AuthServiceAddUsersToRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceAddUsersToRoleOK struct {
	Payload *models.V1AddUsersToRoleResponse
}

func (o *AuthServiceAddUsersToRoleOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/add][%d] authServiceAddUsersToRoleOK  %+v", 200, o.Payload)
}
func (o *AuthServiceAddUsersToRoleOK) GetPayload() *models.V1AddUsersToRoleResponse {
	return o.Payload
}

func (o *AuthServiceAddUsersToRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AddUsersToRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceAddUsersToRoleNotFound creates a AuthServiceAddUsersToRoleNotFound with default headers values
func NewAuthServiceAddUsersToRoleNotFound() *AuthServiceAddUsersToRoleNotFound {
	return &AuthServiceAddUsersToRoleNotFound{}
}

/* AuthServiceAddUsersToRoleNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceAddUsersToRoleNotFound struct {
	Payload string
}

func (o *AuthServiceAddUsersToRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/add][%d] authServiceAddUsersToRoleNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceAddUsersToRoleNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceAddUsersToRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceAddUsersToRoleDefault creates a AuthServiceAddUsersToRoleDefault with default headers values
func NewAuthServiceAddUsersToRoleDefault(code int) *AuthServiceAddUsersToRoleDefault {
	return &AuthServiceAddUsersToRoleDefault{
		_statusCode: code,
	}
}

/* AuthServiceAddUsersToRoleDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceAddUsersToRoleDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service add users to role default response
func (o *AuthServiceAddUsersToRoleDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceAddUsersToRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/add][%d] AuthService_AddUsersToRole default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceAddUsersToRoleDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceAddUsersToRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
