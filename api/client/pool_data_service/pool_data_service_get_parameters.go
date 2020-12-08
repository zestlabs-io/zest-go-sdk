// Code generated by go-swagger; DO NOT EDIT.

package pool_data_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPoolDataServiceGetParams creates a new PoolDataServiceGetParams object
// with the default values initialized.
func NewPoolDataServiceGetParams() *PoolDataServiceGetParams {
	var ()
	return &PoolDataServiceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPoolDataServiceGetParamsWithTimeout creates a new PoolDataServiceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoolDataServiceGetParamsWithTimeout(timeout time.Duration) *PoolDataServiceGetParams {
	var ()
	return &PoolDataServiceGetParams{

		timeout: timeout,
	}
}

// NewPoolDataServiceGetParamsWithContext creates a new PoolDataServiceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoolDataServiceGetParamsWithContext(ctx context.Context) *PoolDataServiceGetParams {
	var ()
	return &PoolDataServiceGetParams{

		Context: ctx,
	}
}

// NewPoolDataServiceGetParamsWithHTTPClient creates a new PoolDataServiceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPoolDataServiceGetParamsWithHTTPClient(client *http.Client) *PoolDataServiceGetParams {
	var ()
	return &PoolDataServiceGetParams{
		HTTPClient: client,
	}
}

/*PoolDataServiceGetParams contains all the parameters to send to the API endpoint
for the pool data service get operation typically these are written to a http.Request
*/
type PoolDataServiceGetParams struct {

	/*ID
	  The ID (Primary Key) of the record

	*/
	ID string
	/*PoolID
	  Pool ID (e.g. orders)

	*/
	PoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pool data service get params
func (o *PoolDataServiceGetParams) WithTimeout(timeout time.Duration) *PoolDataServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool data service get params
func (o *PoolDataServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pool data service get params
func (o *PoolDataServiceGetParams) WithContext(ctx context.Context) *PoolDataServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool data service get params
func (o *PoolDataServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pool data service get params
func (o *PoolDataServiceGetParams) WithHTTPClient(client *http.Client) *PoolDataServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pool data service get params
func (o *PoolDataServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pool data service get params
func (o *PoolDataServiceGetParams) WithID(id string) *PoolDataServiceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pool data service get params
func (o *PoolDataServiceGetParams) SetID(id string) {
	o.ID = id
}

// WithPoolID adds the poolID to the pool data service get params
func (o *PoolDataServiceGetParams) WithPoolID(poolID string) *PoolDataServiceGetParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the pool data service get params
func (o *PoolDataServiceGetParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WriteToRequest writes these params to a swagger request
func (o *PoolDataServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param poolId
	if err := r.SetPathParam("poolId", o.PoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}