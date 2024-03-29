// Code generated by go-swagger; DO NOT EDIT.

package auth_service

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

// NewAuthServiceCheckUsernameExistsParams creates a new AuthServiceCheckUsernameExistsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceCheckUsernameExistsParams() *AuthServiceCheckUsernameExistsParams {
	return &AuthServiceCheckUsernameExistsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceCheckUsernameExistsParamsWithTimeout creates a new AuthServiceCheckUsernameExistsParams object
// with the ability to set a timeout on a request.
func NewAuthServiceCheckUsernameExistsParamsWithTimeout(timeout time.Duration) *AuthServiceCheckUsernameExistsParams {
	return &AuthServiceCheckUsernameExistsParams{
		timeout: timeout,
	}
}

// NewAuthServiceCheckUsernameExistsParamsWithContext creates a new AuthServiceCheckUsernameExistsParams object
// with the ability to set a context for a request.
func NewAuthServiceCheckUsernameExistsParamsWithContext(ctx context.Context) *AuthServiceCheckUsernameExistsParams {
	return &AuthServiceCheckUsernameExistsParams{
		Context: ctx,
	}
}

// NewAuthServiceCheckUsernameExistsParamsWithHTTPClient creates a new AuthServiceCheckUsernameExistsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceCheckUsernameExistsParamsWithHTTPClient(client *http.Client) *AuthServiceCheckUsernameExistsParams {
	return &AuthServiceCheckUsernameExistsParams{
		HTTPClient: client,
	}
}

/* AuthServiceCheckUsernameExistsParams contains all the parameters to send to the API endpoint
   for the auth service check username exists operation.

   Typically these are written to a http.Request.
*/
type AuthServiceCheckUsernameExistsParams struct {

	// Body.
	Body *models.V1CheckUsernameExistsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service check username exists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceCheckUsernameExistsParams) WithDefaults() *AuthServiceCheckUsernameExistsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service check username exists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceCheckUsernameExistsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) WithTimeout(timeout time.Duration) *AuthServiceCheckUsernameExistsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) WithContext(ctx context.Context) *AuthServiceCheckUsernameExistsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) WithHTTPClient(client *http.Client) *AuthServiceCheckUsernameExistsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) WithBody(body *models.V1CheckUsernameExistsRequest) *AuthServiceCheckUsernameExistsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auth service check username exists params
func (o *AuthServiceCheckUsernameExistsParams) SetBody(body *models.V1CheckUsernameExistsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceCheckUsernameExistsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
