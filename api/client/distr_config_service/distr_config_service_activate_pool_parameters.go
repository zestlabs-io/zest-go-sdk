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

	"github.com/zestlabs-io/zest-go-sdk/api/models"
)

// NewDistrConfigServiceActivatePoolParams creates a new DistrConfigServiceActivatePoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDistrConfigServiceActivatePoolParams() *DistrConfigServiceActivatePoolParams {
	return &DistrConfigServiceActivatePoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceActivatePoolParamsWithTimeout creates a new DistrConfigServiceActivatePoolParams object
// with the ability to set a timeout on a request.
func NewDistrConfigServiceActivatePoolParamsWithTimeout(timeout time.Duration) *DistrConfigServiceActivatePoolParams {
	return &DistrConfigServiceActivatePoolParams{
		timeout: timeout,
	}
}

// NewDistrConfigServiceActivatePoolParamsWithContext creates a new DistrConfigServiceActivatePoolParams object
// with the ability to set a context for a request.
func NewDistrConfigServiceActivatePoolParamsWithContext(ctx context.Context) *DistrConfigServiceActivatePoolParams {
	return &DistrConfigServiceActivatePoolParams{
		Context: ctx,
	}
}

// NewDistrConfigServiceActivatePoolParamsWithHTTPClient creates a new DistrConfigServiceActivatePoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewDistrConfigServiceActivatePoolParamsWithHTTPClient(client *http.Client) *DistrConfigServiceActivatePoolParams {
	return &DistrConfigServiceActivatePoolParams{
		HTTPClient: client,
	}
}

/* DistrConfigServiceActivatePoolParams contains all the parameters to send to the API endpoint
   for the distr config service activate pool operation.

   Typically these are written to a http.Request.
*/
type DistrConfigServiceActivatePoolParams struct {

	// Body.
	Body *models.DistrconfigActivatePoolRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the distr config service activate pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistrConfigServiceActivatePoolParams) WithDefaults() *DistrConfigServiceActivatePoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the distr config service activate pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DistrConfigServiceActivatePoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) WithTimeout(timeout time.Duration) *DistrConfigServiceActivatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) WithContext(ctx context.Context) *DistrConfigServiceActivatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) WithHTTPClient(client *http.Client) *DistrConfigServiceActivatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) WithBody(body *models.DistrconfigActivatePoolRequest) *DistrConfigServiceActivatePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the distr config service activate pool params
func (o *DistrConfigServiceActivatePoolParams) SetBody(body *models.DistrconfigActivatePoolRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceActivatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
