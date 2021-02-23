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

// AuthServiceGetClientsReader is a Reader for the AuthServiceGetClients structure.
type AuthServiceGetClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAuthServiceGetClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthServiceGetClientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthServiceGetClientsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetClientsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetClientsOK creates a AuthServiceGetClientsOK with default headers values
func NewAuthServiceGetClientsOK() *AuthServiceGetClientsOK {
	return &AuthServiceGetClientsOK{}
}

/*AuthServiceGetClientsOK handles this case with default header values.

Returned when clients are successfuly fetched.
*/
type AuthServiceGetClientsOK struct {
	Payload *models.V1GetClientsResponse
}

func (o *AuthServiceGetClientsOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/clients][%d] authServiceGetClientsOK  %+v", 200, o.Payload)
}

func (o *AuthServiceGetClientsOK) GetPayload() *models.V1GetClientsResponse {
	return o.Payload
}

func (o *AuthServiceGetClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetClientsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetClientsForbidden creates a AuthServiceGetClientsForbidden with default headers values
func NewAuthServiceGetClientsForbidden() *AuthServiceGetClientsForbidden {
	return &AuthServiceGetClientsForbidden{}
}

/*AuthServiceGetClientsForbidden handles this case with default header values.

Returned when the caller is not allowed to perform this call.
*/
type AuthServiceGetClientsForbidden struct {
	Payload interface{}
}

func (o *AuthServiceGetClientsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/clients][%d] authServiceGetClientsForbidden  %+v", 403, o.Payload)
}

func (o *AuthServiceGetClientsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthServiceGetClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetClientsNotFound creates a AuthServiceGetClientsNotFound with default headers values
func NewAuthServiceGetClientsNotFound() *AuthServiceGetClientsNotFound {
	return &AuthServiceGetClientsNotFound{}
}

/*AuthServiceGetClientsNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetClientsNotFound struct {
	Payload string
}

func (o *AuthServiceGetClientsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/clients][%d] authServiceGetClientsNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceGetClientsNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetClientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetClientsInternalServerError creates a AuthServiceGetClientsInternalServerError with default headers values
func NewAuthServiceGetClientsInternalServerError() *AuthServiceGetClientsInternalServerError {
	return &AuthServiceGetClientsInternalServerError{}
}

/*AuthServiceGetClientsInternalServerError handles this case with default header values.

Returned whenever an internall error occurs.
*/
type AuthServiceGetClientsInternalServerError struct {
	Payload interface{}
}

func (o *AuthServiceGetClientsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/clients][%d] authServiceGetClientsInternalServerError  %+v", 500, o.Payload)
}

func (o *AuthServiceGetClientsInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *AuthServiceGetClientsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetClientsDefault creates a AuthServiceGetClientsDefault with default headers values
func NewAuthServiceGetClientsDefault(code int) *AuthServiceGetClientsDefault {
	return &AuthServiceGetClientsDefault{
		_statusCode: code,
	}
}

/*AuthServiceGetClientsDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceGetClientsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get clients default response
func (o *AuthServiceGetClientsDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetClientsDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/clients][%d] AuthService_GetClients default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceGetClientsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetClientsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
