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

// NewAuthServiceGetRolesParams creates a new AuthServiceGetRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceGetRolesParams() *AuthServiceGetRolesParams {
	return &AuthServiceGetRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetRolesParamsWithTimeout creates a new AuthServiceGetRolesParams object
// with the ability to set a timeout on a request.
func NewAuthServiceGetRolesParamsWithTimeout(timeout time.Duration) *AuthServiceGetRolesParams {
	return &AuthServiceGetRolesParams{
		timeout: timeout,
	}
}

// NewAuthServiceGetRolesParamsWithContext creates a new AuthServiceGetRolesParams object
// with the ability to set a context for a request.
func NewAuthServiceGetRolesParamsWithContext(ctx context.Context) *AuthServiceGetRolesParams {
	return &AuthServiceGetRolesParams{
		Context: ctx,
	}
}

// NewAuthServiceGetRolesParamsWithHTTPClient creates a new AuthServiceGetRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceGetRolesParamsWithHTTPClient(client *http.Client) *AuthServiceGetRolesParams {
	return &AuthServiceGetRolesParams{
		HTTPClient: client,
	}
}

/* AuthServiceGetRolesParams contains all the parameters to send to the API endpoint
   for the auth service get roles operation.

   Typically these are written to a http.Request.
*/
type AuthServiceGetRolesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetRolesParams) WithDefaults() *AuthServiceGetRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service get roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service get roles params
func (o *AuthServiceGetRolesParams) WithTimeout(timeout time.Duration) *AuthServiceGetRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get roles params
func (o *AuthServiceGetRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get roles params
func (o *AuthServiceGetRolesParams) WithContext(ctx context.Context) *AuthServiceGetRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get roles params
func (o *AuthServiceGetRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get roles params
func (o *AuthServiceGetRolesParams) WithHTTPClient(client *http.Client) *AuthServiceGetRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get roles params
func (o *AuthServiceGetRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
