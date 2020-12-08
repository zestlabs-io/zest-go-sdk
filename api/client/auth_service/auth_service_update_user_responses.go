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

// AuthServiceUpdateUserReader is a Reader for the AuthServiceUpdateUser structure.
type AuthServiceUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceUpdateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceUpdateUserOK creates a AuthServiceUpdateUserOK with default headers values
func NewAuthServiceUpdateUserOK() *AuthServiceUpdateUserOK {
	return &AuthServiceUpdateUserOK{}
}

/*AuthServiceUpdateUserOK handles this case with default header values.

A successful response.
*/
type AuthServiceUpdateUserOK struct {
	Payload models.V1UpdateUserResponse
}

func (o *AuthServiceUpdateUserOK) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/user/{user.userID}][%d] authServiceUpdateUserOK  %+v", 200, o.Payload)
}

func (o *AuthServiceUpdateUserOK) GetPayload() models.V1UpdateUserResponse {
	return o.Payload
}

func (o *AuthServiceUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceUpdateUserNotFound creates a AuthServiceUpdateUserNotFound with default headers values
func NewAuthServiceUpdateUserNotFound() *AuthServiceUpdateUserNotFound {
	return &AuthServiceUpdateUserNotFound{}
}

/*AuthServiceUpdateUserNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceUpdateUserNotFound struct {
	Payload string
}

func (o *AuthServiceUpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/user/{user.userID}][%d] authServiceUpdateUserNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceUpdateUserNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceUpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceUpdateUserDefault creates a AuthServiceUpdateUserDefault with default headers values
func NewAuthServiceUpdateUserDefault(code int) *AuthServiceUpdateUserDefault {
	return &AuthServiceUpdateUserDefault{
		_statusCode: code,
	}
}

/*AuthServiceUpdateUserDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceUpdateUserDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service update user default response
func (o *AuthServiceUpdateUserDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceUpdateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/user/{user.userID}][%d] AuthService_UpdateUser default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceUpdateUserDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceUpdateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
