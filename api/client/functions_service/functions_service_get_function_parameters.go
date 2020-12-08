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

// NewFunctionsServiceGetFunctionParams creates a new FunctionsServiceGetFunctionParams object
// with the default values initialized.
func NewFunctionsServiceGetFunctionParams() *FunctionsServiceGetFunctionParams {
	var ()
	return &FunctionsServiceGetFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsServiceGetFunctionParamsWithTimeout creates a new FunctionsServiceGetFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFunctionsServiceGetFunctionParamsWithTimeout(timeout time.Duration) *FunctionsServiceGetFunctionParams {
	var ()
	return &FunctionsServiceGetFunctionParams{

		timeout: timeout,
	}
}

// NewFunctionsServiceGetFunctionParamsWithContext creates a new FunctionsServiceGetFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFunctionsServiceGetFunctionParamsWithContext(ctx context.Context) *FunctionsServiceGetFunctionParams {
	var ()
	return &FunctionsServiceGetFunctionParams{

		Context: ctx,
	}
}

// NewFunctionsServiceGetFunctionParamsWithHTTPClient creates a new FunctionsServiceGetFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFunctionsServiceGetFunctionParamsWithHTTPClient(client *http.Client) *FunctionsServiceGetFunctionParams {
	var ()
	return &FunctionsServiceGetFunctionParams{
		HTTPClient: client,
	}
}

/*FunctionsServiceGetFunctionParams contains all the parameters to send to the API endpoint
for the functions service get function operation typically these are written to a http.Request
*/
type FunctionsServiceGetFunctionParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) WithTimeout(timeout time.Duration) *FunctionsServiceGetFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) WithContext(ctx context.Context) *FunctionsServiceGetFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) WithHTTPClient(client *http.Client) *FunctionsServiceGetFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) WithID(id string) *FunctionsServiceGetFunctionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the functions service get function params
func (o *FunctionsServiceGetFunctionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsServiceGetFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
