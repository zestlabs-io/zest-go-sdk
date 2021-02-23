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

// NewDistrConfigServiceGetUserParams creates a new DistrConfigServiceGetUserParams object
// with the default values initialized.
func NewDistrConfigServiceGetUserParams() *DistrConfigServiceGetUserParams {
	var ()
	return &DistrConfigServiceGetUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceGetUserParamsWithTimeout creates a new DistrConfigServiceGetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceGetUserParamsWithTimeout(timeout time.Duration) *DistrConfigServiceGetUserParams {
	var ()
	return &DistrConfigServiceGetUserParams{

		timeout: timeout,
	}
}

// NewDistrConfigServiceGetUserParamsWithContext creates a new DistrConfigServiceGetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceGetUserParamsWithContext(ctx context.Context) *DistrConfigServiceGetUserParams {
	var ()
	return &DistrConfigServiceGetUserParams{

		Context: ctx,
	}
}

// NewDistrConfigServiceGetUserParamsWithHTTPClient creates a new DistrConfigServiceGetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceGetUserParamsWithHTTPClient(client *http.Client) *DistrConfigServiceGetUserParams {
	var ()
	return &DistrConfigServiceGetUserParams{
		HTTPClient: client,
	}
}

/*DistrConfigServiceGetUserParams contains all the parameters to send to the API endpoint
for the distr config service get user operation typically these are written to a http.Request
*/
type DistrConfigServiceGetUserParams struct {

	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) WithTimeout(timeout time.Duration) *DistrConfigServiceGetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) WithContext(ctx context.Context) *DistrConfigServiceGetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) WithHTTPClient(client *http.Client) *DistrConfigServiceGetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) WithUserID(userID string) *DistrConfigServiceGetUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the distr config service get user params
func (o *DistrConfigServiceGetUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceGetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userID
	if err := r.SetPathParam("userID", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
