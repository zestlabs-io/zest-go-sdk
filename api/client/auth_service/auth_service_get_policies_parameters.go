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
)

// NewAuthServiceGetPoliciesParams creates a new AuthServiceGetPoliciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceGetPoliciesParams() *AuthServiceGetPoliciesParams {
	return &AuthServiceGetPoliciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetPoliciesParamsWithTimeout creates a new AuthServiceGetPoliciesParams object
// with the ability to set a timeout on a request.
func NewAuthServiceGetPoliciesParamsWithTimeout(timeout time.Duration) *AuthServiceGetPoliciesParams {
	return &AuthServiceGetPoliciesParams{
		timeout: timeout,
	}
}

// NewAuthServiceGetPoliciesParamsWithContext creates a new AuthServiceGetPoliciesParams object
// with the ability to set a context for a request.
func NewAuthServiceGetPoliciesParamsWithContext(ctx context.Context) *AuthServiceGetPoliciesParams {
	return &AuthServiceGetPoliciesParams{
		Context: ctx,
	}
}

// NewAuthServiceGetPoliciesParamsWithHTTPClient creates a new AuthServiceGetPoliciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceGetPoliciesParamsWithHTTPClient(client *http.Client) *AuthServiceGetPoliciesParams {
	return &AuthServiceGetPoliciesParams{
		HTTPClient: client,
	}
}

/* AuthServiceGetPoliciesParams contains all the parameters to send to the API endpoint
   for the auth service get policies operation.

   Typically these are written to a http.Request.
*/
type AuthServiceGetPoliciesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service get policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetPoliciesParams) WithDefaults() *AuthServiceGetPoliciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service get policies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetPoliciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) WithTimeout(timeout time.Duration) *AuthServiceGetPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) WithContext(ctx context.Context) *AuthServiceGetPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) WithHTTPClient(client *http.Client) *AuthServiceGetPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get policies params
func (o *AuthServiceGetPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
