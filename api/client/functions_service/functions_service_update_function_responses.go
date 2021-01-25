// Code generated by go-swagger; DO NOT EDIT.

package functions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// FunctionsServiceUpdateFunctionReader is a Reader for the FunctionsServiceUpdateFunction structure.
type FunctionsServiceUpdateFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FunctionsServiceUpdateFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFunctionsServiceUpdateFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFunctionsServiceUpdateFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFunctionsServiceUpdateFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFunctionsServiceUpdateFunctionOK creates a FunctionsServiceUpdateFunctionOK with default headers values
func NewFunctionsServiceUpdateFunctionOK() *FunctionsServiceUpdateFunctionOK {
	return &FunctionsServiceUpdateFunctionOK{}
}

/* FunctionsServiceUpdateFunctionOK describes a response with status code 200, with default header values.

A successful response.
*/
type FunctionsServiceUpdateFunctionOK struct {
	Payload models.FunctionsUpdateFunctionResponse
}

func (o *FunctionsServiceUpdateFunctionOK) Error() string {
	return fmt.Sprintf("[PUT /api/func/v1/function][%d] functionsServiceUpdateFunctionOK  %+v", 200, o.Payload)
}
func (o *FunctionsServiceUpdateFunctionOK) GetPayload() models.FunctionsUpdateFunctionResponse {
	return o.Payload
}

func (o *FunctionsServiceUpdateFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceUpdateFunctionNotFound creates a FunctionsServiceUpdateFunctionNotFound with default headers values
func NewFunctionsServiceUpdateFunctionNotFound() *FunctionsServiceUpdateFunctionNotFound {
	return &FunctionsServiceUpdateFunctionNotFound{}
}

/* FunctionsServiceUpdateFunctionNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type FunctionsServiceUpdateFunctionNotFound struct {
	Payload string
}

func (o *FunctionsServiceUpdateFunctionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/func/v1/function][%d] functionsServiceUpdateFunctionNotFound  %+v", 404, o.Payload)
}
func (o *FunctionsServiceUpdateFunctionNotFound) GetPayload() string {
	return o.Payload
}

func (o *FunctionsServiceUpdateFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceUpdateFunctionDefault creates a FunctionsServiceUpdateFunctionDefault with default headers values
func NewFunctionsServiceUpdateFunctionDefault(code int) *FunctionsServiceUpdateFunctionDefault {
	return &FunctionsServiceUpdateFunctionDefault{
		_statusCode: code,
	}
}

/* FunctionsServiceUpdateFunctionDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type FunctionsServiceUpdateFunctionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the functions service update function default response
func (o *FunctionsServiceUpdateFunctionDefault) Code() int {
	return o._statusCode
}

func (o *FunctionsServiceUpdateFunctionDefault) Error() string {
	return fmt.Sprintf("[PUT /api/func/v1/function][%d] FunctionsService_UpdateFunction default  %+v", o._statusCode, o.Payload)
}
func (o *FunctionsServiceUpdateFunctionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *FunctionsServiceUpdateFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
