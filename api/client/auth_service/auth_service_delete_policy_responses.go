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

// AuthServiceDeletePolicyReader is a Reader for the AuthServiceDeletePolicy structure.
type AuthServiceDeletePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceDeletePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceDeletePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceDeletePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceDeletePolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceDeletePolicyOK creates a AuthServiceDeletePolicyOK with default headers values
func NewAuthServiceDeletePolicyOK() *AuthServiceDeletePolicyOK {
	return &AuthServiceDeletePolicyOK{}
}

/* AuthServiceDeletePolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type AuthServiceDeletePolicyOK struct {
	Payload models.V1DeletePolicyResponse
}

func (o *AuthServiceDeletePolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/policy/{policyID}][%d] authServiceDeletePolicyOK  %+v", 200, o.Payload)
}
func (o *AuthServiceDeletePolicyOK) GetPayload() models.V1DeletePolicyResponse {
	return o.Payload
}

func (o *AuthServiceDeletePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceDeletePolicyNotFound creates a AuthServiceDeletePolicyNotFound with default headers values
func NewAuthServiceDeletePolicyNotFound() *AuthServiceDeletePolicyNotFound {
	return &AuthServiceDeletePolicyNotFound{}
}

/* AuthServiceDeletePolicyNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type AuthServiceDeletePolicyNotFound struct {
	Payload string
}

func (o *AuthServiceDeletePolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/policy/{policyID}][%d] authServiceDeletePolicyNotFound  %+v", 404, o.Payload)
}
func (o *AuthServiceDeletePolicyNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceDeletePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceDeletePolicyDefault creates a AuthServiceDeletePolicyDefault with default headers values
func NewAuthServiceDeletePolicyDefault(code int) *AuthServiceDeletePolicyDefault {
	return &AuthServiceDeletePolicyDefault{
		_statusCode: code,
	}
}

/* AuthServiceDeletePolicyDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AuthServiceDeletePolicyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service delete policy default response
func (o *AuthServiceDeletePolicyDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceDeletePolicyDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/auth/v1/policy/{policyID}][%d] AuthService_DeletePolicy default  %+v", o._statusCode, o.Payload)
}
func (o *AuthServiceDeletePolicyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceDeletePolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
