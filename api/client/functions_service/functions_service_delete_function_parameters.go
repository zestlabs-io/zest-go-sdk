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
)

// NewFunctionsServiceDeleteFunctionParams creates a new FunctionsServiceDeleteFunctionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFunctionsServiceDeleteFunctionParams() *FunctionsServiceDeleteFunctionParams {
	return &FunctionsServiceDeleteFunctionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsServiceDeleteFunctionParamsWithTimeout creates a new FunctionsServiceDeleteFunctionParams object
// with the ability to set a timeout on a request.
func NewFunctionsServiceDeleteFunctionParamsWithTimeout(timeout time.Duration) *FunctionsServiceDeleteFunctionParams {
	return &FunctionsServiceDeleteFunctionParams{
		timeout: timeout,
	}
}

// NewFunctionsServiceDeleteFunctionParamsWithContext creates a new FunctionsServiceDeleteFunctionParams object
// with the ability to set a context for a request.
func NewFunctionsServiceDeleteFunctionParamsWithContext(ctx context.Context) *FunctionsServiceDeleteFunctionParams {
	return &FunctionsServiceDeleteFunctionParams{
		Context: ctx,
	}
}

// NewFunctionsServiceDeleteFunctionParamsWithHTTPClient creates a new FunctionsServiceDeleteFunctionParams object
// with the ability to set a custom HTTPClient for a request.
func NewFunctionsServiceDeleteFunctionParamsWithHTTPClient(client *http.Client) *FunctionsServiceDeleteFunctionParams {
	return &FunctionsServiceDeleteFunctionParams{
		HTTPClient: client,
	}
}

/* FunctionsServiceDeleteFunctionParams contains all the parameters to send to the API endpoint
   for the functions service delete function operation.

   Typically these are written to a http.Request.
*/
type FunctionsServiceDeleteFunctionParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the functions service delete function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceDeleteFunctionParams) WithDefaults() *FunctionsServiceDeleteFunctionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the functions service delete function params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceDeleteFunctionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) WithTimeout(timeout time.Duration) *FunctionsServiceDeleteFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) WithContext(ctx context.Context) *FunctionsServiceDeleteFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) WithHTTPClient(client *http.Client) *FunctionsServiceDeleteFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) WithID(id string) *FunctionsServiceDeleteFunctionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the functions service delete function params
func (o *FunctionsServiceDeleteFunctionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsServiceDeleteFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
