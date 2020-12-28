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

// PoolDataServiceDeleteAttachmentReader is a Reader for the PoolDataServiceDeleteAttachment structure.
type PoolDataServiceDeleteAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolDataServiceDeleteAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoolDataServiceDeleteAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPoolDataServiceDeleteAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPoolDataServiceDeleteAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPoolDataServiceDeleteAttachmentOK creates a PoolDataServiceDeleteAttachmentOK with default headers values
func NewPoolDataServiceDeleteAttachmentOK() *PoolDataServiceDeleteAttachmentOK {
	return &PoolDataServiceDeleteAttachmentOK{}
}

/*PoolDataServiceDeleteAttachmentOK handles this case with default header values.

A successful response.
*/
type PoolDataServiceDeleteAttachmentOK struct {
	Payload string
}

func (o *PoolDataServiceDeleteAttachmentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceDeleteAttachmentOK  %+v", 200, o.Payload)
}

func (o *PoolDataServiceDeleteAttachmentOK) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceDeleteAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceDeleteAttachmentBadRequest creates a PoolDataServiceDeleteAttachmentBadRequest with default headers values
func NewPoolDataServiceDeleteAttachmentBadRequest() *PoolDataServiceDeleteAttachmentBadRequest {
	return &PoolDataServiceDeleteAttachmentBadRequest{}
}

/*PoolDataServiceDeleteAttachmentBadRequest handles this case with default header values.

Returned when input parameters are not provided
*/
type PoolDataServiceDeleteAttachmentBadRequest struct {
	Payload string
}

func (o *PoolDataServiceDeleteAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceDeleteAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *PoolDataServiceDeleteAttachmentBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceDeleteAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceDeleteAttachmentNotFound creates a PoolDataServiceDeleteAttachmentNotFound with default headers values
func NewPoolDataServiceDeleteAttachmentNotFound() *PoolDataServiceDeleteAttachmentNotFound {
	return &PoolDataServiceDeleteAttachmentNotFound{}
}

/*PoolDataServiceDeleteAttachmentNotFound handles this case with default header values.

Returned when the document does not exist.
*/
type PoolDataServiceDeleteAttachmentNotFound struct {
	Payload string
}

func (o *PoolDataServiceDeleteAttachmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/data/_r/{poolId}/{id}/{attname}][%d] poolDataServiceDeleteAttachmentNotFound  %+v", 404, o.Payload)
}

func (o *PoolDataServiceDeleteAttachmentNotFound) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceDeleteAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
