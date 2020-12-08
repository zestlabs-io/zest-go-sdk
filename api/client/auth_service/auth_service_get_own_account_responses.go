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

// AuthServiceGetOwnAccountReader is a Reader for the AuthServiceGetOwnAccount structure.
type AuthServiceGetOwnAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetOwnAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetOwnAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceGetOwnAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetOwnAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetOwnAccountOK creates a AuthServiceGetOwnAccountOK with default headers values
func NewAuthServiceGetOwnAccountOK() *AuthServiceGetOwnAccountOK {
	return &AuthServiceGetOwnAccountOK{}
}

/*AuthServiceGetOwnAccountOK handles this case with default header values.

A successful response.
*/
type AuthServiceGetOwnAccountOK struct {
	Payload *models.V1GetAccountResponse
}

func (o *AuthServiceGetOwnAccountOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/account][%d] authServiceGetOwnAccountOK  %+v", 200, o.Payload)
}

func (o *AuthServiceGetOwnAccountOK) GetPayload() *models.V1GetAccountResponse {
	return o.Payload
}

func (o *AuthServiceGetOwnAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetOwnAccountNotFound creates a AuthServiceGetOwnAccountNotFound with default headers values
func NewAuthServiceGetOwnAccountNotFound() *AuthServiceGetOwnAccountNotFound {
	return &AuthServiceGetOwnAccountNotFound{}
}

/*AuthServiceGetOwnAccountNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetOwnAccountNotFound struct {
	Payload string
}

func (o *AuthServiceGetOwnAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/account][%d] authServiceGetOwnAccountNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceGetOwnAccountNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetOwnAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetOwnAccountDefault creates a AuthServiceGetOwnAccountDefault with default headers values
func NewAuthServiceGetOwnAccountDefault(code int) *AuthServiceGetOwnAccountDefault {
	return &AuthServiceGetOwnAccountDefault{
		_statusCode: code,
	}
}

/*AuthServiceGetOwnAccountDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceGetOwnAccountDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get own account default response
func (o *AuthServiceGetOwnAccountDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetOwnAccountDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/account][%d] AuthService_GetOwnAccount default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceGetOwnAccountDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetOwnAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}