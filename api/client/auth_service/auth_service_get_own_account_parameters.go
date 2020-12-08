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

// NewAuthServiceGetOwnAccountParams creates a new AuthServiceGetOwnAccountParams object
// with the default values initialized.
func NewAuthServiceGetOwnAccountParams() *AuthServiceGetOwnAccountParams {

	return &AuthServiceGetOwnAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetOwnAccountParamsWithTimeout creates a new AuthServiceGetOwnAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceGetOwnAccountParamsWithTimeout(timeout time.Duration) *AuthServiceGetOwnAccountParams {

	return &AuthServiceGetOwnAccountParams{

		timeout: timeout,
	}
}

// NewAuthServiceGetOwnAccountParamsWithContext creates a new AuthServiceGetOwnAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceGetOwnAccountParamsWithContext(ctx context.Context) *AuthServiceGetOwnAccountParams {

	return &AuthServiceGetOwnAccountParams{

		Context: ctx,
	}
}

// NewAuthServiceGetOwnAccountParamsWithHTTPClient creates a new AuthServiceGetOwnAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceGetOwnAccountParamsWithHTTPClient(client *http.Client) *AuthServiceGetOwnAccountParams {

	return &AuthServiceGetOwnAccountParams{
		HTTPClient: client,
	}
}

/*AuthServiceGetOwnAccountParams contains all the parameters to send to the API endpoint
for the auth service get own account operation typically these are written to a http.Request
*/
type AuthServiceGetOwnAccountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) WithTimeout(timeout time.Duration) *AuthServiceGetOwnAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) WithContext(ctx context.Context) *AuthServiceGetOwnAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) WithHTTPClient(client *http.Client) *AuthServiceGetOwnAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get own account params
func (o *AuthServiceGetOwnAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetOwnAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
