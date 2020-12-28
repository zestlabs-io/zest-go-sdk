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

// New creates a new pool data service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pool data service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PoolDataServiceBulkCreate(params *PoolDataServiceBulkCreateParams) (*PoolDataServiceBulkCreateOK, error)

	PoolDataServiceBulkDelete(params *PoolDataServiceBulkDeleteParams) (*PoolDataServiceBulkDeleteOK, error)

	PoolDataServiceBulkUpdate(params *PoolDataServiceBulkUpdateParams) (*PoolDataServiceBulkUpdateOK, error)

	PoolDataServiceDeleteAttachment(params *PoolDataServiceDeleteAttachmentParams) (*PoolDataServiceDeleteAttachmentOK, error)

	PoolDataServiceGet(params *PoolDataServiceGetParams) (*PoolDataServiceGetOK, error)

	PoolDataServiceGetAttachment(params *PoolDataServiceGetAttachmentParams, writer io.Writer) (*PoolDataServiceGetAttachmentOK, error)

	PoolDataServiceList(params *PoolDataServiceListParams) (*PoolDataServiceListOK, error)

	PoolDataServiceStoreAttachment(params *PoolDataServiceStoreAttachmentParams) (*PoolDataServiceStoreAttachmentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PoolDataServiceBulkCreate bulks create mobile data records in a pool
*/
func (a *Client) PoolDataServiceBulkCreate(params *PoolDataServiceBulkCreateParams) (*PoolDataServiceBulkCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceBulkCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_BulkCreate",
		Method:             "POST",
		PathPattern:        "/api/data/_r/{poolId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceBulkCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceBulkCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_BulkCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceBulkDelete bulks delete records from mobile data pool
*/
func (a *Client) PoolDataServiceBulkDelete(params *PoolDataServiceBulkDeleteParams) (*PoolDataServiceBulkDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceBulkDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_BulkDelete",
		Method:             "DELETE",
		PathPattern:        "/api/data/_r/{poolId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceBulkDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceBulkDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_BulkDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceBulkUpdate bulks update records in mobile data pool
*/
func (a *Client) PoolDataServiceBulkUpdate(params *PoolDataServiceBulkUpdateParams) (*PoolDataServiceBulkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceBulkUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_BulkUpdate",
		Method:             "PUT",
		PathPattern:        "/api/data/_r/{poolId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceBulkUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceBulkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_BulkUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceDeleteAttachment deletes attachment
*/
func (a *Client) PoolDataServiceDeleteAttachment(params *PoolDataServiceDeleteAttachmentParams) (*PoolDataServiceDeleteAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceDeleteAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_DeleteAttachment",
		Method:             "DELETE",
		PathPattern:        "/api/data/_r/{poolId}/{id}/{attname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceDeleteAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceDeleteAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_DeleteAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceGet gets single mobile data record
*/
func (a *Client) PoolDataServiceGet(params *PoolDataServiceGetParams) (*PoolDataServiceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_Get",
		Method:             "GET",
		PathPattern:        "/api/data/_r/{poolId}/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceGetAttachment gets attachment
*/
func (a *Client) PoolDataServiceGetAttachment(params *PoolDataServiceGetAttachmentParams, writer io.Writer) (*PoolDataServiceGetAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceGetAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_GetAttachment",
		Method:             "GET",
		PathPattern:        "/api/data/_r/{poolId}/{id}/{attname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceGetAttachmentReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceGetAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_GetAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceList lists mobile data records
*/
func (a *Client) PoolDataServiceList(params *PoolDataServiceListParams) (*PoolDataServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_List",
		Method:             "GET",
		PathPattern:        "/api/data/_r/{poolId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PoolDataServiceStoreAttachment stores attachment
*/
func (a *Client) PoolDataServiceStoreAttachment(params *PoolDataServiceStoreAttachmentParams) (*PoolDataServiceStoreAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolDataServiceStoreAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PoolDataService_StoreAttachment",
		Method:             "PUT",
		PathPattern:        "/api/data/_r/{poolId}/{id}/{attname}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PoolDataServiceStoreAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolDataServiceStoreAttachmentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PoolDataService_StoreAttachment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
