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

// AuthServiceSetPasswordReader is a Reader for the AuthServiceSetPassword structure.
type AuthServiceSetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceSetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceSetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceSetPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceSetPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceSetPasswordOK creates a AuthServiceSetPasswordOK with default headers values
func NewAuthServiceSetPasswordOK() *AuthServiceSetPasswordOK {
	return &AuthServiceSetPasswordOK{}
}

/* AuthServiceSetPasswordOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceSetPasswordOK struct {
	Payload models.V1SetPasswordResponse
}

func (o *AuthServiceSetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/set-password][%d] authServiceSetPasswordOK  %+v", 200, o.Payload)
}
func (o *AuthServiceSetPasswordOK) GetPayload() models.V1SetPasswordResponse {
	return o.Payload
}

func (o *AuthServiceSetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceSetPasswordNotFound creates a AuthServiceSetPasswordNotFound with default headers values
func NewAuthServiceSetPasswordNotFound() *AuthServiceSetPasswordNotFound {
	return &AuthServiceSetPasswordNotFound{}
}

/* AuthServiceSetPasswordNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceSetPasswordNotFound struct {
	Payload string
}

func (o *AuthServiceSetPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/set-password][%d] authServiceSetPasswordNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceSetPasswordNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceSetPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceSetPasswordDefault creates a AuthServiceSetPasswordDefault with default headers values
func NewAuthServiceSetPasswordDefault(code int) *AuthServiceSetPasswordDefault {
	return &AuthServiceSetPasswordDefault{
		_statusCode: code,
	}
}

/* AuthServiceSetPasswordDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceSetPasswordDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service set password default response
func (o *AuthServiceSetPasswordDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceSetPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/set-password][%d] AuthService_SetPassword default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceSetPasswordDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceSetPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
