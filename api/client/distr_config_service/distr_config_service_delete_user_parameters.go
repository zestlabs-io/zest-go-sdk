// Code generated by go-swagger; DO NOT EDIT.

package distr_config_service

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

// NewDistrConfigServiceDeleteUserParams creates a new DistrConfigServiceDeleteUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDistrConfigServiceDeleteUserParams() *DistrConfigServiceDeleteUserParams {
	return &DistrConfigServiceDeleteUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceDeleteUserParamsWithTimeout creates a new DistrConfigServiceDeleteUserParams object
// with the ability to set a timeout on a request.
func NewDistrConfigServiceDeleteUserParamsWithTimeout(timeout time.Duration) *DistrConfigServiceDeleteUserParams {
	return &DistrConfigServiceDeleteUserParams{
		timeout: timeout,
	}
}

// NewDistrConfigServiceDeleteUserParamsWithContext creates a new DistrConfigServiceDeleteUserParams object
// with the ability to set a context for a request.
func NewDistrConfigServiceDeleteUserParamsWithContext(ctx context.Context) *DistrConfigServiceDeleteUserParams {
	return &DistrConfigServiceDeleteUserParams{
		Context: ctx,
	}
}

// NewDistrConfigServiceDeleteUserParamsWithHTTPClient creates a new DistrConfigServiceDeleteUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewDistrConfigServiceDeleteUserParamsWithHTTPClient(client *http.Client) *DistrConfigServiceDeleteUserParams {
	return &DistrConfigServiceDeleteUserParams{
		HTTPClient: client,
	}
}

/* DistrConfigServiceDeleteUserParams contains all the parameters to send to the API endpoint
   for the distr config service delete user operation.

   Typically these are written to a http.Request.
*/
type DistrConfigServiceDeleteUserParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the distr config service delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistrConfigServiceDeleteUserParams) WithDefaults() *DistrConfigServiceDeleteUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the distr config service delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistrConfigServiceDeleteUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) WithTimeout(timeout time.Duration) *DistrConfigServiceDeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) WithContext(ctx context.Context) *DistrConfigServiceDeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) WithHTTPClient(client *http.Client) *DistrConfigServiceDeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) WithID(id string) *DistrConfigServiceDeleteUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the distr config service delete user params
func (o *DistrConfigServiceDeleteUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceDeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
