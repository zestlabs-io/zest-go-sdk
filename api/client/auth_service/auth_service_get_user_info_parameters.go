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

// NewAuthServiceGetUserInfoParams creates a new AuthServiceGetUserInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceGetUserInfoParams() *AuthServiceGetUserInfoParams {
	return &AuthServiceGetUserInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetUserInfoParamsWithTimeout creates a new AuthServiceGetUserInfoParams object
// with the ability to set a timeout on a request.
func NewAuthServiceGetUserInfoParamsWithTimeout(timeout time.Duration) *AuthServiceGetUserInfoParams {
	return &AuthServiceGetUserInfoParams{
		timeout: timeout,
	}
}

// NewAuthServiceGetUserInfoParamsWithContext creates a new AuthServiceGetUserInfoParams object
// with the ability to set a context for a request.
func NewAuthServiceGetUserInfoParamsWithContext(ctx context.Context) *AuthServiceGetUserInfoParams {
	return &AuthServiceGetUserInfoParams{
		Context: ctx,
	}
}

// NewAuthServiceGetUserInfoParamsWithHTTPClient creates a new AuthServiceGetUserInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceGetUserInfoParamsWithHTTPClient(client *http.Client) *AuthServiceGetUserInfoParams {
	return &AuthServiceGetUserInfoParams{
		HTTPClient: client,
	}
}

/* AuthServiceGetUserInfoParams contains all the parameters to send to the API endpoint
   for the auth service get user info operation.

   Typically these are written to a http.Request.
*/
type AuthServiceGetUserInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service get user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetUserInfoParams) WithDefaults() *AuthServiceGetUserInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service get user info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceGetUserInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) WithTimeout(timeout time.Duration) *AuthServiceGetUserInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) WithContext(ctx context.Context) *AuthServiceGetUserInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) WithHTTPClient(client *http.Client) *AuthServiceGetUserInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get user info params
func (o *AuthServiceGetUserInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetUserInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
