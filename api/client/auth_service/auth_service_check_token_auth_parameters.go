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

// NewAuthServiceCheckTokenAuthParams creates a new AuthServiceCheckTokenAuthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceCheckTokenAuthParams() *AuthServiceCheckTokenAuthParams {
	return &AuthServiceCheckTokenAuthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceCheckTokenAuthParamsWithTimeout creates a new AuthServiceCheckTokenAuthParams object
// with the ability to set a timeout on a request.
func NewAuthServiceCheckTokenAuthParamsWithTimeout(timeout time.Duration) *AuthServiceCheckTokenAuthParams {
	return &AuthServiceCheckTokenAuthParams{
		timeout: timeout,
	}
}

// NewAuthServiceCheckTokenAuthParamsWithContext creates a new AuthServiceCheckTokenAuthParams object
// with the ability to set a context for a request.
func NewAuthServiceCheckTokenAuthParamsWithContext(ctx context.Context) *AuthServiceCheckTokenAuthParams {
	return &AuthServiceCheckTokenAuthParams{
		Context: ctx,
	}
}

// NewAuthServiceCheckTokenAuthParamsWithHTTPClient creates a new AuthServiceCheckTokenAuthParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceCheckTokenAuthParamsWithHTTPClient(client *http.Client) *AuthServiceCheckTokenAuthParams {
	return &AuthServiceCheckTokenAuthParams{
		HTTPClient: client,
	}
}

/* AuthServiceCheckTokenAuthParams contains all the parameters to send to the API endpoint
   for the auth service check token auth operation.

   Typically these are written to a http.Request.
*/
type AuthServiceCheckTokenAuthParams struct {

	// Body.
	Body *models.V1CheckTokenAuthRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service check token auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceCheckTokenAuthParams) WithDefaults() *AuthServiceCheckTokenAuthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service check token auth params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceCheckTokenAuthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) WithTimeout(timeout time.Duration) *AuthServiceCheckTokenAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) WithContext(ctx context.Context) *AuthServiceCheckTokenAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) WithHTTPClient(client *http.Client) *AuthServiceCheckTokenAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) WithBody(body *models.V1CheckTokenAuthRequest) *AuthServiceCheckTokenAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auth service check token auth params
func (o *AuthServiceCheckTokenAuthParams) SetBody(body *models.V1CheckTokenAuthRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceCheckTokenAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
