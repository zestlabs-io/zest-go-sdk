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

// AuthServiceGetUserAccessKeysReader is a Reader for the AuthServiceGetUserAccessKeys structure.
type AuthServiceGetUserAccessKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetUserAccessKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetUserAccessKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceGetUserAccessKeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetUserAccessKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetUserAccessKeysOK creates a AuthServiceGetUserAccessKeysOK with default headers values
func NewAuthServiceGetUserAccessKeysOK() *AuthServiceGetUserAccessKeysOK {
	return &AuthServiceGetUserAccessKeysOK{}
}

/* AuthServiceGetUserAccessKeysOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceGetUserAccessKeysOK struct {
	Payload *models.V1GetUserAccessKeysResponse
}

func (o *AuthServiceGetUserAccessKeysOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/user/accesskeys/{userID}][%d] authServiceGetUserAccessKeysOK  %+v", 200, o.Payload)
}
func (o *AuthServiceGetUserAccessKeysOK) GetPayload() *models.V1GetUserAccessKeysResponse {
	return o.Payload
}

func (o *AuthServiceGetUserAccessKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetUserAccessKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetUserAccessKeysNotFound creates a AuthServiceGetUserAccessKeysNotFound with default headers values
func NewAuthServiceGetUserAccessKeysNotFound() *AuthServiceGetUserAccessKeysNotFound {
	return &AuthServiceGetUserAccessKeysNotFound{}
}

/* AuthServiceGetUserAccessKeysNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetUserAccessKeysNotFound struct {
	Payload string
}

func (o *AuthServiceGetUserAccessKeysNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/user/accesskeys/{userID}][%d] authServiceGetUserAccessKeysNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceGetUserAccessKeysNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetUserAccessKeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetUserAccessKeysDefault creates a AuthServiceGetUserAccessKeysDefault with default headers values
func NewAuthServiceGetUserAccessKeysDefault(code int) *AuthServiceGetUserAccessKeysDefault {
	return &AuthServiceGetUserAccessKeysDefault{
		_statusCode: code,
	}
}

/* AuthServiceGetUserAccessKeysDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceGetUserAccessKeysDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get user access keys default response
func (o *AuthServiceGetUserAccessKeysDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetUserAccessKeysDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/user/accesskeys/{userID}][%d] AuthService_GetUserAccessKeys default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceGetUserAccessKeysDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetUserAccessKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
