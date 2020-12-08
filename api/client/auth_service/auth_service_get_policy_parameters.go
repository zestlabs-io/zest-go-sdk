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

// NewAuthServiceGetPolicyParams creates a new AuthServiceGetPolicyParams object
// with the default values initialized.
func NewAuthServiceGetPolicyParams() *AuthServiceGetPolicyParams {
	var ()
	return &AuthServiceGetPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceGetPolicyParamsWithTimeout creates a new AuthServiceGetPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceGetPolicyParamsWithTimeout(timeout time.Duration) *AuthServiceGetPolicyParams {
	var ()
	return &AuthServiceGetPolicyParams{

		timeout: timeout,
	}
}

// NewAuthServiceGetPolicyParamsWithContext creates a new AuthServiceGetPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceGetPolicyParamsWithContext(ctx context.Context) *AuthServiceGetPolicyParams {
	var ()
	return &AuthServiceGetPolicyParams{

		Context: ctx,
	}
}

// NewAuthServiceGetPolicyParamsWithHTTPClient creates a new AuthServiceGetPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceGetPolicyParamsWithHTTPClient(client *http.Client) *AuthServiceGetPolicyParams {
	var ()
	return &AuthServiceGetPolicyParams{
		HTTPClient: client,
	}
}

/*AuthServiceGetPolicyParams contains all the parameters to send to the API endpoint
for the auth service get policy operation typically these are written to a http.Request
*/
type AuthServiceGetPolicyParams struct {

	/*PolicyID*/
	PolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service get policy params
func (o *AuthServiceGetPolicyParams) WithTimeout(timeout time.Duration) *AuthServiceGetPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service get policy params
func (o *AuthServiceGetPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service get policy params
func (o *AuthServiceGetPolicyParams) WithContext(ctx context.Context) *AuthServiceGetPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service get policy params
func (o *AuthServiceGetPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service get policy params
func (o *AuthServiceGetPolicyParams) WithHTTPClient(client *http.Client) *AuthServiceGetPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service get policy params
func (o *AuthServiceGetPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyID adds the policyID to the auth service get policy params
func (o *AuthServiceGetPolicyParams) WithPolicyID(policyID string) *AuthServiceGetPolicyParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the auth service get policy params
func (o *AuthServiceGetPolicyParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceGetPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param policyID
	if err := r.SetPathParam("policyID", o.PolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}