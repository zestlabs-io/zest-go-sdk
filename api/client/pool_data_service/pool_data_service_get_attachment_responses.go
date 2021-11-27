// Code generated by go-swagger; DO NOT EDIT.

package pool_data_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PoolDataServiceGetAttachmentReader is a Reader for the PoolDataServiceGetAttachment structure.
type PoolDataServiceGetAttachmentReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *PoolDataServiceGetAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoolDataServiceGetAttachmentOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPoolDataServiceGetAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPoolDataServiceGetAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPoolDataServiceGetAttachmentOK creates a PoolDataServiceGetAttachmentOK with default headers values
func NewPoolDataServiceGetAttachmentOK(writer io.Writer) *PoolDataServiceGetAttachmentOK {
	return &PoolDataServiceGetAttachmentOK{

		Payload: writer,
	}
}

/* PoolDataServiceGetAttachmentOK describes a response with status code 200, with default header values.

A successful response.
*/
type PoolDataServiceGetAttachmentOK struct {
	Payload io.Writer
}

func (o *PoolDataServiceGetAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceGetAttachmentOK  %+v", 200, o.Payload)
}
func (o *PoolDataServiceGetAttachmentOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *PoolDataServiceGetAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceGetAttachmentBadRequest creates a PoolDataServiceGetAttachmentBadRequest with default headers values
func NewPoolDataServiceGetAttachmentBadRequest() *PoolDataServiceGetAttachmentBadRequest {
	return &PoolDataServiceGetAttachmentBadRequest{}
}

/* PoolDataServiceGetAttachmentBadRequest describes a response with status code 400, with default header values.

Returned when input parameters are not provided
*/
type PoolDataServiceGetAttachmentBadRequest struct {
	Payload string
}

func (o *PoolDataServiceGetAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceGetAttachmentBadRequest  %+v", 400, o.Payload)
}
func (o *PoolDataServiceGetAttachmentBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceGetAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceGetAttachmentNotFound creates a PoolDataServiceGetAttachmentNotFound with default headers values
func NewPoolDataServiceGetAttachmentNotFound() *PoolDataServiceGetAttachmentNotFound {
	return &PoolDataServiceGetAttachmentNotFound{}
}

/* PoolDataServiceGetAttachmentNotFound describes a response with status code 404, with default header values.

Returned when the document does not exist.
*/
type PoolDataServiceGetAttachmentNotFound struct {
	Payload string
}

func (o *PoolDataServiceGetAttachmentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceGetAttachmentNotFound  %+v", 404, o.Payload)
}
func (o *PoolDataServiceGetAttachmentNotFound) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceGetAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
