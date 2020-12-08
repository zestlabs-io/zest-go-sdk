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

// AuthServiceRemoveUsersFromRoleReader is a Reader for the AuthServiceRemoveUsersFromRole structure.
type AuthServiceRemoveUsersFromRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceRemoveUsersFromRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceRemoveUsersFromRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceRemoveUsersFromRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceRemoveUsersFromRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceRemoveUsersFromRoleOK creates a AuthServiceRemoveUsersFromRoleOK with default headers values
func NewAuthServiceRemoveUsersFromRoleOK() *AuthServiceRemoveUsersFromRoleOK {
	return &AuthServiceRemoveUsersFromRoleOK{}
}

/*AuthServiceRemoveUsersFromRoleOK handles this case with default header values.

A successful response.
*/
type AuthServiceRemoveUsersFromRoleOK struct {
	Payload *models.V1RemoveUsersFromRoleResponse
}

func (o *AuthServiceRemoveUsersFromRoleOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/remove][%d] authServiceRemoveUsersFromRoleOK  %+v", 200, o.Payload)
}

func (o *AuthServiceRemoveUsersFromRoleOK) GetPayload() *models.V1RemoveUsersFromRoleResponse {
	return o.Payload
}

func (o *AuthServiceRemoveUsersFromRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RemoveUsersFromRoleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceRemoveUsersFromRoleNotFound creates a AuthServiceRemoveUsersFromRoleNotFound with default headers values
func NewAuthServiceRemoveUsersFromRoleNotFound() *AuthServiceRemoveUsersFromRoleNotFound {
	return &AuthServiceRemoveUsersFromRoleNotFound{}
}

/*AuthServiceRemoveUsersFromRoleNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceRemoveUsersFromRoleNotFound struct {
	Payload string
}

func (o *AuthServiceRemoveUsersFromRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/remove][%d] authServiceRemoveUsersFromRoleNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceRemoveUsersFromRoleNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceRemoveUsersFromRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceRemoveUsersFromRoleDefault creates a AuthServiceRemoveUsersFromRoleDefault with default headers values
func NewAuthServiceRemoveUsersFromRoleDefault(code int) *AuthServiceRemoveUsersFromRoleDefault {
	return &AuthServiceRemoveUsersFromRoleDefault{
		_statusCode: code,
	}
}

/*AuthServiceRemoveUsersFromRoleDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceRemoveUsersFromRoleDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service remove users from role default response
func (o *AuthServiceRemoveUsersFromRoleDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceRemoveUsersFromRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/role/users/remove][%d] AuthService_RemoveUsersFromRole default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceRemoveUsersFromRoleDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceRemoveUsersFromRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}