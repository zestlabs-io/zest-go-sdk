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

// NewAuthServiceDeleteUserParams creates a new AuthServiceDeleteUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthServiceDeleteUserParams() *AuthServiceDeleteUserParams {
	return &AuthServiceDeleteUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthServiceDeleteUserParamsWithTimeout creates a new AuthServiceDeleteUserParams object
// with the ability to set a timeout on a request.
func NewAuthServiceDeleteUserParamsWithTimeout(timeout time.Duration) *AuthServiceDeleteUserParams {
	return &AuthServiceDeleteUserParams{
		timeout: timeout,
	}
}

// NewAuthServiceDeleteUserParamsWithContext creates a new AuthServiceDeleteUserParams object
// with the ability to set a context for a request.
func NewAuthServiceDeleteUserParamsWithContext(ctx context.Context) *AuthServiceDeleteUserParams {
	return &AuthServiceDeleteUserParams{
		Context: ctx,
	}
}

// NewAuthServiceDeleteUserParamsWithHTTPClient creates a new AuthServiceDeleteUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthServiceDeleteUserParamsWithHTTPClient(client *http.Client) *AuthServiceDeleteUserParams {
	return &AuthServiceDeleteUserParams{
		HTTPClient: client,
	}
}

/* AuthServiceDeleteUserParams contains all the parameters to send to the API endpoint
   for the auth service delete user operation.

   Typically these are written to a http.Request.
*/
type AuthServiceDeleteUserParams struct {

	// UserID.
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth service delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceDeleteUserParams) WithDefaults() *AuthServiceDeleteUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth service delete user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthServiceDeleteUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth service delete user params
func (o *AuthServiceDeleteUserParams) WithTimeout(timeout time.Duration) *AuthServiceDeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth service delete user params
func (o *AuthServiceDeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth service delete user params
func (o *AuthServiceDeleteUserParams) WithContext(ctx context.Context) *AuthServiceDeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth service delete user params
func (o *AuthServiceDeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth service delete user params
func (o *AuthServiceDeleteUserParams) WithHTTPClient(client *http.Client) *AuthServiceDeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth service delete user params
func (o *AuthServiceDeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the auth service delete user params
func (o *AuthServiceDeleteUserParams) WithUserID(userID string) *AuthServiceDeleteUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the auth service delete user params
func (o *AuthServiceDeleteUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AuthServiceDeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userID
	if err := r.SetPathParam("userID", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
