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

// NewDistrConfigServiceAssignAppToUsersParams creates a new DistrConfigServiceAssignAppToUsersParams object
// with the default values initialized.
func NewDistrConfigServiceAssignAppToUsersParams() *DistrConfigServiceAssignAppToUsersParams {
	var ()
	return &DistrConfigServiceAssignAppToUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceAssignAppToUsersParamsWithTimeout creates a new DistrConfigServiceAssignAppToUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceAssignAppToUsersParamsWithTimeout(timeout time.Duration) *DistrConfigServiceAssignAppToUsersParams {
	var ()
	return &DistrConfigServiceAssignAppToUsersParams{

		timeout: timeout,
	}
}

// NewDistrConfigServiceAssignAppToUsersParamsWithContext creates a new DistrConfigServiceAssignAppToUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceAssignAppToUsersParamsWithContext(ctx context.Context) *DistrConfigServiceAssignAppToUsersParams {
	var ()
	return &DistrConfigServiceAssignAppToUsersParams{

		Context: ctx,
	}
}

// NewDistrConfigServiceAssignAppToUsersParamsWithHTTPClient creates a new DistrConfigServiceAssignAppToUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceAssignAppToUsersParamsWithHTTPClient(client *http.Client) *DistrConfigServiceAssignAppToUsersParams {
	var ()
	return &DistrConfigServiceAssignAppToUsersParams{
		HTTPClient: client,
	}
}

/*DistrConfigServiceAssignAppToUsersParams contains all the parameters to send to the API endpoint
for the distr config service assign app to users operation typically these are written to a http.Request
*/
type DistrConfigServiceAssignAppToUsersParams struct {

	/*Body*/
	Body *models.DistrconfigAssignAppToUsersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) WithTimeout(timeout time.Duration) *DistrConfigServiceAssignAppToUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) WithContext(ctx context.Context) *DistrConfigServiceAssignAppToUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) WithHTTPClient(client *http.Client) *DistrConfigServiceAssignAppToUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) WithBody(body *models.DistrconfigAssignAppToUsersRequest) *DistrConfigServiceAssignAppToUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the distr config service assign app to users params
func (o *DistrConfigServiceAssignAppToUsersParams) SetBody(body *models.DistrconfigAssignAppToUsersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceAssignAppToUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
