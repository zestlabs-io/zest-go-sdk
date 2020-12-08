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

// FunctionsServiceGetFunctionReader is a Reader for the FunctionsServiceGetFunction structure.
type FunctionsServiceGetFunctionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FunctionsServiceGetFunctionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFunctionsServiceGetFunctionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFunctionsServiceGetFunctionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFunctionsServiceGetFunctionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFunctionsServiceGetFunctionOK creates a FunctionsServiceGetFunctionOK with default headers values
func NewFunctionsServiceGetFunctionOK() *FunctionsServiceGetFunctionOK {
	return &FunctionsServiceGetFunctionOK{}
}

/*FunctionsServiceGetFunctionOK handles this case with default header values.

A successful response.
*/
type FunctionsServiceGetFunctionOK struct {
	Payload *models.FunctionsGetFunctionResponse
}

func (o *FunctionsServiceGetFunctionOK) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/function/{id}][%d] functionsServiceGetFunctionOK  %+v", 200, o.Payload)
}

func (o *FunctionsServiceGetFunctionOK) GetPayload() *models.FunctionsGetFunctionResponse {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionsGetFunctionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceGetFunctionNotFound creates a FunctionsServiceGetFunctionNotFound with default headers values
func NewFunctionsServiceGetFunctionNotFound() *FunctionsServiceGetFunctionNotFound {
	return &FunctionsServiceGetFunctionNotFound{}
}

/*FunctionsServiceGetFunctionNotFound handles this case with default header values.

Returned when the resource does not exist.
*/
type FunctionsServiceGetFunctionNotFound struct {
	Payload string
}

func (o *FunctionsServiceGetFunctionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/function/{id}][%d] functionsServiceGetFunctionNotFound  %+v", 404, o.Payload)
}

func (o *FunctionsServiceGetFunctionNotFound) GetPayload() string {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceGetFunctionDefault creates a FunctionsServiceGetFunctionDefault with default headers values
func NewFunctionsServiceGetFunctionDefault(code int) *FunctionsServiceGetFunctionDefault {
	return &FunctionsServiceGetFunctionDefault{
		_statusCode: code,
	}
}

/*FunctionsServiceGetFunctionDefault handles this case with default header values.

An unexpected error response
*/
type FunctionsServiceGetFunctionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the functions service get function default response
func (o *FunctionsServiceGetFunctionDefault) Code() int {
	return o._statusCode
}

func (o *FunctionsServiceGetFunctionDefault) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/function/{id}][%d] FunctionsService_GetFunction default  %+v", o._statusCode, o.Payload)
}

func (o *FunctionsServiceGetFunctionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}