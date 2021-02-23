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

// PoolDataServiceBulkDeleteReader is a Reader for the PoolDataServiceBulkDelete structure.
type PoolDataServiceBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoolDataServiceBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoolDataServiceBulkDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPoolDataServiceBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPoolDataServiceBulkDeleteOK creates a PoolDataServiceBulkDeleteOK with default headers values
func NewPoolDataServiceBulkDeleteOK() *PoolDataServiceBulkDeleteOK {
	return &PoolDataServiceBulkDeleteOK{}
}

/*PoolDataServiceBulkDeleteOK handles this case with default header values.

A successful response.
*/
type PoolDataServiceBulkDeleteOK struct {
	Payload *models.DataBulkDeleteResponse
}

func (o *PoolDataServiceBulkDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/data/_r/{poolId}][%d] poolDataServiceBulkDeleteOK  %+v", 200, o.Payload)
}

func (o *PoolDataServiceBulkDeleteOK) GetPayload() *models.DataBulkDeleteResponse {
	return o.Payload
}

func (o *PoolDataServiceBulkDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataBulkDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPoolDataServiceBulkDeleteBadRequest creates a PoolDataServiceBulkDeleteBadRequest with default headers values
func NewPoolDataServiceBulkDeleteBadRequest() *PoolDataServiceBulkDeleteBadRequest {
	return &PoolDataServiceBulkDeleteBadRequest{}
}

/*PoolDataServiceBulkDeleteBadRequest handles this case with default header values.

Returned when input parameters are not provided
*/
type PoolDataServiceBulkDeleteBadRequest struct {
	Payload string
}

func (o *PoolDataServiceBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/data/_r/{poolId}][%d] poolDataServiceBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PoolDataServiceBulkDeleteBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PoolDataServiceBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
