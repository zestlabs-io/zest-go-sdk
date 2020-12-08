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

// AuthServiceUpdateAccountReader is a Reader for the AuthServiceUpdateAccount structure.
type AuthServiceUpdateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceUpdateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceUpdateAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceUpdateAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceUpdateAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceUpdateAccountOK creates a AuthServiceUpdateAccountOK with default headers values
func NewAuthServiceUpdateAccountOK() *AuthServiceUpdateAccountOK {
	return &AuthServiceUpdateAccountOK{}
}

/*AuthServiceUpdateAccountOK handles this case with default header values.

A successful response.
*/
type AuthServiceUpdateAccountOK struct {
	Payload models.V1UpdateAccountResponse
}

func (o *AuthServiceUpdateAccountOK) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/account/{account.accountID}][%d] authServiceUpdateAccountOK  %+v", 200, o.Payload)
}

func (o *AuthServiceUpdateAccountOK) GetPayload() models.V1UpdateAccountResponse {
	return o.Payload
}

func (o *AuthServiceUpdateAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceUpdateAccountNotFound creates a AuthServiceUpdateAccountNotFound with default headers values
func NewAuthServiceUpdateAccountNotFound() *AuthServiceUpdateAccountNotFound {
	return &AuthServiceUpdateAccountNotFound{}
}

/*AuthServiceUpdateAccountNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceUpdateAccountNotFound struct {
	Payload string
}

func (o *AuthServiceUpdateAccountNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/account/{account.accountID}][%d] authServiceUpdateAccountNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceUpdateAccountNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceUpdateAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceUpdateAccountDefault creates a AuthServiceUpdateAccountDefault with default headers values
func NewAuthServiceUpdateAccountDefault(code int) *AuthServiceUpdateAccountDefault {
	return &AuthServiceUpdateAccountDefault{
		_statusCode: code,
	}
}

/*AuthServiceUpdateAccountDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceUpdateAccountDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service update account default response
func (o *AuthServiceUpdateAccountDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceUpdateAccountDefault) Error() string {
	return fmt.Sprintf("[PUT /api/auth/v1/account/{account.accountID}][%d] AuthService_UpdateAccount default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceUpdateAccountDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceUpdateAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}