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

// NewAuthServiceDeletePolicyParams creates a new AuthServiceDeletePolicyParams object
// with the default values initialized.
func NewAuthServiceDeletePolicyParams() *AuthServiceDeletePolicyParams {
	var ()
	return &AuthServiceDeletePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceDeletePolicyParamsWithTimeout creates a new AuthServiceDeletePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceDeletePolicyParamsWithTimeout(timeout time.Duration) *AuthServiceDeletePolicyParams {
	var ()
	return &AuthServiceDeletePolicyParams{

		timeout: timeout,
	}
}

// NewAuthServiceDeletePolicyParamsWithContext creates a new AuthServiceDeletePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceDeletePolicyParamsWithContext(ctx context.Context) *AuthServiceDeletePolicyParams {
	var ()
	return &AuthServiceDeletePolicyParams{

		Context: ctx,
	}
}

// NewAuthServiceDeletePolicyParamsWithHTTPClient creates a new AuthServiceDeletePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceDeletePolicyParamsWithHTTPClient(client *http.Client) *AuthServiceDeletePolicyParams {
	var ()
	return &AuthServiceDeletePolicyParams{
		HTTPClient: client,
	}
}

/*AuthServiceDeletePolicyParams contains all the parameters to send to the API endpoint
for the auth service delete policy operation typically these are written to a http.Request
*/
type AuthServiceDeletePolicyParams struct {

	/*PolicyID*/
	PolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) WithTimeout(timeout time.Duration) *AuthServiceDeletePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) WithContext(ctx context.Context) *AuthServiceDeletePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) WithHTTPClient(client *http.Client) *AuthServiceDeletePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyID adds the policyID to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) WithPolicyID(policyID string) *AuthServiceDeletePolicyParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the auth service delete policy params
func (o *AuthServiceDeletePolicyParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceDeletePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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