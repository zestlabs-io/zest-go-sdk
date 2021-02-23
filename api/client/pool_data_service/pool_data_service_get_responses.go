// Code generated by go-swagger; DO NOT EDIT.

package pool_data_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// PoolDataServiceGetReader is a Reader for the PoolDataServiceGet structure.
type PoolDataServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolDataServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoolDataServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPoolDataServiceGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPoolDataServiceGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPoolDataServiceGetOK creates a PoolDataServiceGetOK with default headers values
func NewPoolDataServiceGetOK() *PoolDataServiceGetOK {
	return &PoolDataServiceGetOK{}
}

/*PoolDataServiceGetOK handles this case with default header values.

A successful response.
*/
type PoolDataServiceGetOK struct {
	Payload *models.DataGetResponse
}

func (o *PoolDataServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}][%d] poolDataServiceGetOK  %+v", 200, o.Payload)
}

func (o *PoolDataServiceGetOK) GetPayload() *models.DataGetResponse {
	return o.Payload
}

func (o *PoolDataServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceGetBadRequest creates a PoolDataServiceGetBadRequest with default headers values
func NewPoolDataServiceGetBadRequest() *PoolDataServiceGetBadRequest {
	return &PoolDataServiceGetBadRequest{}
}

/*PoolDataServiceGetBadRequest handles this case with default header values.

Returned when input parameters are not provided
*/
type PoolDataServiceGetBadRequest struct {
	Payload string
}

func (o *PoolDataServiceGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}][%d] poolDataServiceGetBadRequest  %+v", 400, o.Payload)
}

func (o *PoolDataServiceGetBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceGetNotFound creates a PoolDataServiceGetNotFound with default headers values
func NewPoolDataServiceGetNotFound() *PoolDataServiceGetNotFound {
	return &PoolDataServiceGetNotFound{}
}

/*PoolDataServiceGetNotFound handles this case with default header values.

Returned when the document does not exist.
*/
type PoolDataServiceGetNotFound struct {
	Payload string
}

func (o *PoolDataServiceGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/data/_r/{poolId}/{id}][%d] poolDataServiceGetNotFound  %+v", 404, o.Payload)
}

func (o *PoolDataServiceGetNotFound) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
