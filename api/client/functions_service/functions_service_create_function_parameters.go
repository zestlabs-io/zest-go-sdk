// Code generated by go-swagger; DO NOT EDIT.

package functions_service

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

// NewFunctionsServiceCreateFunctionParams creates a new FunctionsServiceCreateFunctionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFunctionsServiceCreateFunctionParams() *FunctionsServiceCreateFunctionParams {
	return &FunctionsServiceCreateFunctionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsServiceCreateFunctionParamsWithTimeout creates a new FunctionsServiceCreateFunctionParams object
// with the ability to set a timeout on a request.
func NewFunctionsServiceCreateFunctionParamsWithTimeout(timeout time.Duration) *FunctionsServiceCreateFunctionParams {
	return &FunctionsServiceCreateFunctionParams{
		timeout: timeout,
	}
}

// NewFunctionsServiceCreateFunctionParamsWithContext creates a new FunctionsServiceCreateFunctionParams object
// with the ability to set a context for a request.
func NewFunctionsServiceCreateFunctionParamsWithContext(ctx context.Context) *FunctionsServiceCreateFunctionParams {
	return &FunctionsServiceCreateFunctionParams{
		Context: ctx,
	}
}

// NewFunctionsServiceCreateFunctionParamsWithHTTPClient creates a new FunctionsServiceCreateFunctionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFunctionsServiceCreateFunctionParamsWithHTTPClient(client *http.Client) *FunctionsServiceCreateFunctionParams {
	return &FunctionsServiceCreateFunctionParams{
		HTTPClient: client,
	}
}

/* FunctionsServiceCreateFunctionParams contains all the parameters to send to the API endpoint
   for the functions service create function operation.

   Typically these are written to a http.Request.
*/
type FunctionsServiceCreateFunctionParams struct {

	// Body.
	Body *models.FunctionsCreateFunctionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the functions service create function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceCreateFunctionParams) WithDefaults() *FunctionsServiceCreateFunctionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the functions service create function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceCreateFunctionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) WithTimeout(timeout time.Duration) *FunctionsServiceCreateFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) WithContext(ctx context.Context) *FunctionsServiceCreateFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) WithHTTPClient(client *http.Client) *FunctionsServiceCreateFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) WithBody(body *models.FunctionsCreateFunctionRequest) *FunctionsServiceCreateFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the functions service create function params
func (o *FunctionsServiceCreateFunctionParams) SetBody(body *models.FunctionsCreateFunctionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsServiceCreateFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
