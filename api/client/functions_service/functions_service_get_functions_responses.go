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

// FunctionsServiceGetFunctionsReader is a Reader for the FunctionsServiceGetFunctions structure.
type FunctionsServiceGetFunctionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FunctionsServiceGetFunctionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFunctionsServiceGetFunctionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFunctionsServiceGetFunctionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFunctionsServiceGetFunctionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFunctionsServiceGetFunctionsOK creates a FunctionsServiceGetFunctionsOK with default headers values
func NewFunctionsServiceGetFunctionsOK() *FunctionsServiceGetFunctionsOK {
	return &FunctionsServiceGetFunctionsOK{}
}

/* FunctionsServiceGetFunctionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type FunctionsServiceGetFunctionsOK struct {
	Payload *models.FunctionsGetFunctionsResponse
}

func (o *FunctionsServiceGetFunctionsOK) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/functions][%d] functionsServiceGetFunctionsOK  %+v", 200, o.Payload)
}
func (o *FunctionsServiceGetFunctionsOK) GetPayload() *models.FunctionsGetFunctionsResponse {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FunctionsGetFunctionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceGetFunctionsNotFound creates a FunctionsServiceGetFunctionsNotFound with default headers values
func NewFunctionsServiceGetFunctionsNotFound() *FunctionsServiceGetFunctionsNotFound {
	return &FunctionsServiceGetFunctionsNotFound{}
}

/* FunctionsServiceGetFunctionsNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type FunctionsServiceGetFunctionsNotFound struct {
	Payload string
}

func (o *FunctionsServiceGetFunctionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/functions][%d] functionsServiceGetFunctionsNotFound  %+v", 404, o.Payload)
}
func (o *FunctionsServiceGetFunctionsNotFound) GetPayload() string {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionsServiceGetFunctionsDefault creates a FunctionsServiceGetFunctionsDefault with default headers values
func NewFunctionsServiceGetFunctionsDefault(code int) *FunctionsServiceGetFunctionsDefault {
	return &FunctionsServiceGetFunctionsDefault{
		_statusCode: code,
	}
}

/* FunctionsServiceGetFunctionsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type FunctionsServiceGetFunctionsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the functions service get functions default response
func (o *FunctionsServiceGetFunctionsDefault) Code() int {
	return o._statusCode
}

func (o *FunctionsServiceGetFunctionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/func/v1/functions][%d] FunctionsService_GetFunctions default  %+v", o._statusCode, o.Payload)
}
func (o *FunctionsServiceGetFunctionsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *FunctionsServiceGetFunctionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
