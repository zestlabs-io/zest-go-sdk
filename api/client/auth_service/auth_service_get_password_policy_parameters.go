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

// NewAuthServiceGetPasswordPolicyParams creates a new AuthServiceGetPasswordPolicyParams object
// with the default values initialized.
func NewAuthServiceGetPasswordPolicyParams() *AuthServiceGetPasswordPolicyParams {

	return &AuthServiceGetPasswordPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetPasswordPolicyParamsWithTimeout creates a new AuthServiceGetPasswordPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceGetPasswordPolicyParamsWithTimeout(timeout time.Duration) *AuthServiceGetPasswordPolicyParams {

	return &AuthServiceGetPasswordPolicyParams{

		timeout: timeout,
	}
}

// NewAuthServiceGetPasswordPolicyParamsWithContext creates a new AuthServiceGetPasswordPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceGetPasswordPolicyParamsWithContext(ctx context.Context) *AuthServiceGetPasswordPolicyParams {

	return &AuthServiceGetPasswordPolicyParams{

		Context: ctx,
	}
}

// NewAuthServiceGetPasswordPolicyParamsWithHTTPClient creates a new AuthServiceGetPasswordPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceGetPasswordPolicyParamsWithHTTPClient(client *http.Client) *AuthServiceGetPasswordPolicyParams {

	return &AuthServiceGetPasswordPolicyParams{
		HTTPClient: client,
	}
}

/*AuthServiceGetPasswordPolicyParams contains all the parameters to send to the API endpoint
for the auth service get password policy operation typically these are written to a http.Request
*/
type AuthServiceGetPasswordPolicyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) WithTimeout(timeout time.Duration) *AuthServiceGetPasswordPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) WithContext(ctx context.Context) *AuthServiceGetPasswordPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) WithHTTPClient(client *http.Client) *AuthServiceGetPasswordPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get password policy params
func (o *AuthServiceGetPasswordPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetPasswordPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
