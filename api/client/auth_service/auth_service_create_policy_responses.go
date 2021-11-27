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

// AuthServiceCreatePolicyReader is a Reader for the AuthServiceCreatePolicy structure.
type AuthServiceCreatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceCreatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceCreatePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceCreatePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceCreatePolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceCreatePolicyOK creates a AuthServiceCreatePolicyOK with default headers values
func NewAuthServiceCreatePolicyOK() *AuthServiceCreatePolicyOK {
	return &AuthServiceCreatePolicyOK{}
}

/* AuthServiceCreatePolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceCreatePolicyOK struct {
	Payload *models.V1CreatePolicyResponse
}

func (o *AuthServiceCreatePolicyOK) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/policy][%d] authServiceCreatePolicyOK  %+v", 200, o.Payload)
}
func (o *AuthServiceCreatePolicyOK) GetPayload() *models.V1CreatePolicyResponse {
	return o.Payload
}

func (o *AuthServiceCreatePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CreatePolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreatePolicyNotFound creates a AuthServiceCreatePolicyNotFound with default headers values
func NewAuthServiceCreatePolicyNotFound() *AuthServiceCreatePolicyNotFound {
	return &AuthServiceCreatePolicyNotFound{}
}

/* AuthServiceCreatePolicyNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceCreatePolicyNotFound struct {
	Payload string
}

func (o *AuthServiceCreatePolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/policy][%d] authServiceCreatePolicyNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceCreatePolicyNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceCreatePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceCreatePolicyDefault creates a AuthServiceCreatePolicyDefault with default headers values
func NewAuthServiceCreatePolicyDefault(code int) *AuthServiceCreatePolicyDefault {
	return &AuthServiceCreatePolicyDefault{
		_statusCode: code,
	}
}

/* AuthServiceCreatePolicyDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceCreatePolicyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service create policy default response
func (o *AuthServiceCreatePolicyDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceCreatePolicyDefault) Error() string {
	return fmt.Sprintf("[POST /api/auth/v1/policy][%d] AuthService_CreatePolicy default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceCreatePolicyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceCreatePolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
