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

// AuthServiceCheckTokenAuthReader is a Reader for the AuthServiceCheckTokenAuth structure.
type AuthServiceCheckTokenAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceCheckTokenAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceCheckTokenAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceCheckTokenAuthNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceCheckTokenAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceCheckTokenAuthOK creates a AuthServiceCheckTokenAuthOK with default headers values
func NewAuthServiceCheckTokenAuthOK() *AuthServiceCheckTokenAuthOK {
	return &AuthServiceCheckTokenAuthOK{}
}

/*AuthServiceCheckTokenAuthOK handles this case with default header values.

A successful response.
*/
type AuthServiceCheckTokenAuthOK struct {
	Payload *models.V1CheckTokenAuthResponse
}

func (o *AuthServiceCheckTokenAuthOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/check-token][%d] authServiceCheckTokenAuthOK  %+v", 200, o.Payload)
}

func (o *AuthServiceCheckTokenAuthOK) GetPayload() *models.V1CheckTokenAuthResponse {
	return o.Payload
}

func (o *AuthServiceCheckTokenAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CheckTokenAuthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCheckTokenAuthNotFound creates a AuthServiceCheckTokenAuthNotFound with default headers values
func NewAuthServiceCheckTokenAuthNotFound() *AuthServiceCheckTokenAuthNotFound {
	return &AuthServiceCheckTokenAuthNotFound{}
}

/*AuthServiceCheckTokenAuthNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceCheckTokenAuthNotFound struct {
	Payload string
}

func (o *AuthServiceCheckTokenAuthNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/check-token][%d] authServiceCheckTokenAuthNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceCheckTokenAuthNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceCheckTokenAuthNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCheckTokenAuthDefault creates a AuthServiceCheckTokenAuthDefault with default headers values
func NewAuthServiceCheckTokenAuthDefault(code int) *AuthServiceCheckTokenAuthDefault {
	return &AuthServiceCheckTokenAuthDefault{
		_statusCode: code,
	}
}

/*AuthServiceCheckTokenAuthDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceCheckTokenAuthDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service check token auth default response
func (o *AuthServiceCheckTokenAuthDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceCheckTokenAuthDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/check-token][%d] AuthService_CheckTokenAuth default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceCheckTokenAuthDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceCheckTokenAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
