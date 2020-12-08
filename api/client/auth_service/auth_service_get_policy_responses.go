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

// AuthServiceGetPolicyReader is a Reader for the AuthServiceGetPolicy structure.
type AuthServiceGetPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthServiceGetPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthServiceGetPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAuthServiceGetPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAuthServiceGetPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthServiceGetPolicyOK creates a AuthServiceGetPolicyOK with default headers values
func NewAuthServiceGetPolicyOK() *AuthServiceGetPolicyOK {
	return &AuthServiceGetPolicyOK{}
}

/*AuthServiceGetPolicyOK handles this case with default header values.

A successful response.
*/
type AuthServiceGetPolicyOK struct {
	Payload *models.V1GetPolicyResponse
}

func (o *AuthServiceGetPolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/policy/{policyID}][%d] authServiceGetPolicyOK  %+v", 200, o.Payload)
}

func (o *AuthServiceGetPolicyOK) GetPayload() *models.V1GetPolicyResponse {
	return o.Payload
}

func (o *AuthServiceGetPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetPolicyNotFound creates a AuthServiceGetPolicyNotFound with default headers values
func NewAuthServiceGetPolicyNotFound() *AuthServiceGetPolicyNotFound {
	return &AuthServiceGetPolicyNotFound{}
}

/*AuthServiceGetPolicyNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type AuthServiceGetPolicyNotFound struct {
	Payload string
}

func (o *AuthServiceGetPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/policy/{policyID}][%d] authServiceGetPolicyNotFound  %+v", 404, o.Payload)
}

func (o *AuthServiceGetPolicyNotFound) GetPayload() string {
	return o.Payload
}

func (o *AuthServiceGetPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthServiceGetPolicyDefault creates a AuthServiceGetPolicyDefault with default headers values
func NewAuthServiceGetPolicyDefault(code int) *AuthServiceGetPolicyDefault {
	return &AuthServiceGetPolicyDefault{
		_statusCode: code,
	}
}

/*AuthServiceGetPolicyDefault handles this case with default header values.

An unexpected error response
*/
type AuthServiceGetPolicyDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the auth service get policy default response
func (o *AuthServiceGetPolicyDefault) Code() int {
	return o._statusCode
}

func (o *AuthServiceGetPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /api/auth/v1/policy/{policyID}][%d] AuthService_GetPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *AuthServiceGetPolicyDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *AuthServiceGetPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}