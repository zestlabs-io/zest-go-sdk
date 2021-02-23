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

// NewAuthServiceDeleteAccessKeyParams creates a new AuthServiceDeleteAccessKeyParams object
// with the default values initialized.
func NewAuthServiceDeleteAccessKeyParams() *AuthServiceDeleteAccessKeyParams {
	var ()
	return &AuthServiceDeleteAccessKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceDeleteAccessKeyParamsWithTimeout creates a new AuthServiceDeleteAccessKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAuthServiceDeleteAccessKeyParamsWithTimeout(timeout time.Duration) *AuthServiceDeleteAccessKeyParams {
	var ()
	return &AuthServiceDeleteAccessKeyParams{

		timeout: timeout,
	}
}

// NewAuthServiceDeleteAccessKeyParamsWithContext creates a new AuthServiceDeleteAccessKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewAuthServiceDeleteAccessKeyParamsWithContext(ctx context.Context) *AuthServiceDeleteAccessKeyParams {
	var ()
	return &AuthServiceDeleteAccessKeyParams{

		Context: ctx,
	}
}

// NewAuthServiceDeleteAccessKeyParamsWithHTTPClient creates a new AuthServiceDeleteAccessKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAuthServiceDeleteAccessKeyParamsWithHTTPClient(client *http.Client) *AuthServiceDeleteAccessKeyParams {
	var ()
	return &AuthServiceDeleteAccessKeyParams{
		HTTPClient: client,
	}
}

/*AuthServiceDeleteAccessKeyParams contains all the parameters to send to the API endpoint
for the auth service delete access key operation typically these are written to a http.Request
*/
type AuthServiceDeleteAccessKeyParams struct {

	/*AccessKeyID*/
	AccessKeyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) WithTimeout(timeout time.Duration) *AuthServiceDeleteAccessKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) WithContext(ctx context.Context) *AuthServiceDeleteAccessKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) WithHTTPClient(client *http.Client) *AuthServiceDeleteAccessKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyID adds the accessKeyID to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) WithAccessKeyID(accessKeyID string) *AuthServiceDeleteAccessKeyParams {
	o.SetAccessKeyID(accessKeyID)
	return o
}

// SetAccessKeyID adds the accessKeyId to the auth service delete access key params
func (o *AuthServiceDeleteAccessKeyParams) SetAccessKeyID(accessKeyID string) {
	o.AccessKeyID = accessKeyID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceDeleteAccessKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accessKeyID
	if err := r.SetPathParam("accessKeyID", o.AccessKeyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
