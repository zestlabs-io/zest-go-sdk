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

// NewDistrConfigServiceGetPoolParams creates a new DistrConfigServiceGetPoolParams object
// with the default values initialized.
func NewDistrConfigServiceGetPoolParams() *DistrConfigServiceGetPoolParams {
	var ()
	return &DistrConfigServiceGetPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceGetPoolParamsWithTimeout creates a new DistrConfigServiceGetPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceGetPoolParamsWithTimeout(timeout time.Duration) *DistrConfigServiceGetPoolParams {
	var ()
	return &DistrConfigServiceGetPoolParams{

		timeout: timeout,
	}
}

// NewDistrConfigServiceGetPoolParamsWithContext creates a new DistrConfigServiceGetPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceGetPoolParamsWithContext(ctx context.Context) *DistrConfigServiceGetPoolParams {
	var ()
	return &DistrConfigServiceGetPoolParams{

		Context: ctx,
	}
}

// NewDistrConfigServiceGetPoolParamsWithHTTPClient creates a new DistrConfigServiceGetPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceGetPoolParamsWithHTTPClient(client *http.Client) *DistrConfigServiceGetPoolParams {
	var ()
	return &DistrConfigServiceGetPoolParams{
		HTTPClient: client,
	}
}

/*DistrConfigServiceGetPoolParams contains all the parameters to send to the API endpoint
for the distr config service get pool operation typically these are written to a http.Request
*/
type DistrConfigServiceGetPoolParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) WithTimeout(timeout time.Duration) *DistrConfigServiceGetPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) WithContext(ctx context.Context) *DistrConfigServiceGetPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) WithHTTPClient(client *http.Client) *DistrConfigServiceGetPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) WithID(id string) *DistrConfigServiceGetPoolParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the distr config service get pool params
func (o *DistrConfigServiceGetPoolParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceGetPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
