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

// NewDistrConfigServiceCreatePoolParams creates a new DistrConfigServiceCreatePoolParams object
// with the default values initialized.
func NewDistrConfigServiceCreatePoolParams() *DistrConfigServiceCreatePoolParams {
	var ()
	return &DistrConfigServiceCreatePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceCreatePoolParamsWithTimeout creates a new DistrConfigServiceCreatePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceCreatePoolParamsWithTimeout(timeout time.Duration) *DistrConfigServiceCreatePoolParams {
	var ()
	return &DistrConfigServiceCreatePoolParams{

		timeout: timeout,
	}
}

// NewDistrConfigServiceCreatePoolParamsWithContext creates a new DistrConfigServiceCreatePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceCreatePoolParamsWithContext(ctx context.Context) *DistrConfigServiceCreatePoolParams {
	var ()
	return &DistrConfigServiceCreatePoolParams{

		Context: ctx,
	}
}

// NewDistrConfigServiceCreatePoolParamsWithHTTPClient creates a new DistrConfigServiceCreatePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceCreatePoolParamsWithHTTPClient(client *http.Client) *DistrConfigServiceCreatePoolParams {
	var ()
	return &DistrConfigServiceCreatePoolParams{
		HTTPClient: client,
	}
}

/*DistrConfigServiceCreatePoolParams contains all the parameters to send to the API endpoint
for the distr config service create pool operation typically these are written to a http.Request
*/
type DistrConfigServiceCreatePoolParams struct {

	/*Body*/
	Body *models.DistrconfigDataPool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) WithTimeout(timeout time.Duration) *DistrConfigServiceCreatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) WithContext(ctx context.Context) *DistrConfigServiceCreatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) WithHTTPClient(client *http.Client) *DistrConfigServiceCreatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) WithBody(body *models.DistrconfigDataPool) *DistrConfigServiceCreatePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the distr config service create pool params
func (o *DistrConfigServiceCreatePoolParams) SetBody(body *models.DistrconfigDataPool) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceCreatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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