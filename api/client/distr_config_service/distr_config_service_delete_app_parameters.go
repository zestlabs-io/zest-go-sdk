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

// NewDistrConfigServiceDeleteAppParams creates a new DistrConfigServiceDeleteAppParams object
// with the default values initialized.
func NewDistrConfigServiceDeleteAppParams() *DistrConfigServiceDeleteAppParams {
	var ()
	return &DistrConfigServiceDeleteAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceDeleteAppParamsWithTimeout creates a new DistrConfigServiceDeleteAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceDeleteAppParamsWithTimeout(timeout time.Duration) *DistrConfigServiceDeleteAppParams {
	var ()
	return &DistrConfigServiceDeleteAppParams{

		timeout: timeout,
	}
}

// NewDistrConfigServiceDeleteAppParamsWithContext creates a new DistrConfigServiceDeleteAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceDeleteAppParamsWithContext(ctx context.Context) *DistrConfigServiceDeleteAppParams {
	var ()
	return &DistrConfigServiceDeleteAppParams{

		Context: ctx,
	}
}

// NewDistrConfigServiceDeleteAppParamsWithHTTPClient creates a new DistrConfigServiceDeleteAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceDeleteAppParamsWithHTTPClient(client *http.Client) *DistrConfigServiceDeleteAppParams {
	var ()
	return &DistrConfigServiceDeleteAppParams{
		HTTPClient: client,
	}
}

/*DistrConfigServiceDeleteAppParams contains all the parameters to send to the API endpoint
for the distr config service delete app operation typically these are written to a http.Request
*/
type DistrConfigServiceDeleteAppParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) WithTimeout(timeout time.Duration) *DistrConfigServiceDeleteAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) WithContext(ctx context.Context) *DistrConfigServiceDeleteAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) WithHTTPClient(client *http.Client) *DistrConfigServiceDeleteAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) WithID(id string) *DistrConfigServiceDeleteAppParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the distr config service delete app params
func (o *DistrConfigServiceDeleteAppParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceDeleteAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
