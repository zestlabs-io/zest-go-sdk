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

// NewAuthServiceCreatePolicyParams creates a new AuthServiceCreatePolicyParams object
// with the default values initialized.
func NewAuthServiceCreatePolicyParams() *AuthServiceCreatePolicyParams {
	var ()
	return &AuthServiceCreatePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceCreatePolicyParamsWithTimeout creates a new AuthServiceCreatePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceCreatePolicyParamsWithTimeout(timeout time.Duration) *AuthServiceCreatePolicyParams {
	var ()
	return &AuthServiceCreatePolicyParams{

		timeout: timeout,
	}
}

// NewAuthServiceCreatePolicyParamsWithContext creates a new AuthServiceCreatePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceCreatePolicyParamsWithContext(ctx context.Context) *AuthServiceCreatePolicyParams {
	var ()
	return &AuthServiceCreatePolicyParams{

		Context: ctx,
	}
}

// NewAuthServiceCreatePolicyParamsWithHTTPClient creates a new AuthServiceCreatePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceCreatePolicyParamsWithHTTPClient(client *http.Client) *AuthServiceCreatePolicyParams {
	var ()
	return &AuthServiceCreatePolicyParams{
		HTTPClient: client,
	}
}

/*AuthServiceCreatePolicyParams contains all the parameters to send to the API endpoint
for the auth service create policy operation typically these are written to a http.Request
*/
type AuthServiceCreatePolicyParams struct {

	/*Body*/
	Body *models.V1CreatePolicyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) WithTimeout(timeout time.Duration) *AuthServiceCreatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) WithContext(ctx context.Context) *AuthServiceCreatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) WithHTTPClient(client *http.Client) *AuthServiceCreatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) WithBody(body *models.V1CreatePolicyRequest) *AuthServiceCreatePolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the auth service create policy params
func (o *AuthServiceCreatePolicyParams) SetBody(body *models.V1CreatePolicyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceCreatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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