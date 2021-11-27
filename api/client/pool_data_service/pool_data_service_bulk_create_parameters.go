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

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// NewPoolDataServiceBulkCreateParams creates a new PoolDataServiceBulkCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPoolDataServiceBulkCreateParams() *PoolDataServiceBulkCreateParams {
	return &PoolDataServiceBulkCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPoolDataServiceBulkCreateParamsWithTimeout creates a new PoolDataServiceBulkCreateParams object
// with the ability to set a timeout on a request.
func NewPoolDataServiceBulkCreateParamsWithTimeout(timeout time.Duration) *PoolDataServiceBulkCreateParams {
	return &PoolDataServiceBulkCreateParams{
		timeout: timeout,
	}
}

// NewPoolDataServiceBulkCreateParamsWithContext creates a new PoolDataServiceBulkCreateParams object
// with the ability to set a context for a request.
func NewPoolDataServiceBulkCreateParamsWithContext(ctx context.Context) *PoolDataServiceBulkCreateParams {
	return &PoolDataServiceBulkCreateParams{
		Context: ctx,
	}
}

// NewPoolDataServiceBulkCreateParamsWithHTTPClient creates a new PoolDataServiceBulkCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPoolDataServiceBulkCreateParamsWithHTTPClient(client *http.Client) *PoolDataServiceBulkCreateParams {
	return &PoolDataServiceBulkCreateParams{
		HTTPClient: client,
	}
}

/* PoolDataServiceBulkCreateParams contains all the parameters to send to the API endpoint
   for the pool data service bulk create operation.

   Typically these are written to a http.Request.
*/
type PoolDataServiceBulkCreateParams struct {

	// Body.
	Body models.DataBulkCreateRequest

	/* PoolID.

	   Pool ID (e.g. orders)
	*/
	PoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pool data service bulk create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolDataServiceBulkCreateParams) WithDefaults() *PoolDataServiceBulkCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pool data service bulk create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolDataServiceBulkCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) WithTimeout(timeout time.Duration) *PoolDataServiceBulkCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) WithContext(ctx context.Context) *PoolDataServiceBulkCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) WithHTTPClient(client *http.Client) *PoolDataServiceBulkCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) WithBody(body models.DataBulkCreateRequest) *PoolDataServiceBulkCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) SetBody(body models.DataBulkCreateRequest) {
	o.Body = body
}

// WithPoolID adds the poolID to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) WithPoolID(poolID string) *PoolDataServiceBulkCreateParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the pool data service bulk create params
func (o *PoolDataServiceBulkCreateParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WriteToRequest writes these params to a swagger request
func (o *PoolDataServiceBulkCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
