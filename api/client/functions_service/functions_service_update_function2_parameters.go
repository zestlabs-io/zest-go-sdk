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

// NewFunctionsServiceUpdateFunction2Params creates a new FunctionsServiceUpdateFunction2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFunctionsServiceUpdateFunction2Params() *FunctionsServiceUpdateFunction2Params {
	return &FunctionsServiceUpdateFunction2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewFunctionsServiceUpdateFunction2ParamsWithTimeout creates a new FunctionsServiceUpdateFunction2Params object
// with the ability to set a timeout on a request.
func NewFunctionsServiceUpdateFunction2ParamsWithTimeout(timeout time.Duration) *FunctionsServiceUpdateFunction2Params {
	return &FunctionsServiceUpdateFunction2Params{
		timeout: timeout,
	}
}

// NewFunctionsServiceUpdateFunction2ParamsWithContext creates a new FunctionsServiceUpdateFunction2Params object
// with the ability to set a context for a request.
func NewFunctionsServiceUpdateFunction2ParamsWithContext(ctx context.Context) *FunctionsServiceUpdateFunction2Params {
	return &FunctionsServiceUpdateFunction2Params{
		Context: ctx,
	}
}

// NewFunctionsServiceUpdateFunction2ParamsWithHTTPClient creates a new FunctionsServiceUpdateFunction2Params object
// with the ability to set a custom HTTPClient for a request.
func NewFunctionsServiceUpdateFunction2ParamsWithHTTPClient(client *http.Client) *FunctionsServiceUpdateFunction2Params {
	return &FunctionsServiceUpdateFunction2Params{
		HTTPClient: client,
	}
}

/* FunctionsServiceUpdateFunction2Params contains all the parameters to send to the API endpoint
   for the functions service update function2 operation.

   Typically these are written to a http.Request.
*/
type FunctionsServiceUpdateFunction2Params struct {

	// Body.
	Body *models.FunctionsFunction

	// FunctionID.
	FunctionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the functions service update function2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceUpdateFunction2Params) WithDefaults() *FunctionsServiceUpdateFunction2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the functions service update function2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FunctionsServiceUpdateFunction2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) WithTimeout(timeout time.Duration) *FunctionsServiceUpdateFunction2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) WithContext(ctx context.Context) *FunctionsServiceUpdateFunction2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) WithHTTPClient(client *http.Client) *FunctionsServiceUpdateFunction2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) WithBody(body *models.FunctionsFunction) *FunctionsServiceUpdateFunction2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) SetBody(body *models.FunctionsFunction) {
	o.Body = body
}

// WithFunctionID adds the functionID to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) WithFunctionID(functionID string) *FunctionsServiceUpdateFunction2Params {
	o.SetFunctionID(functionID)
	return o
}

// SetFunctionID adds the functionId to the functions service update function2 params
func (o *FunctionsServiceUpdateFunction2Params) SetFunctionID(functionID string) {
	o.FunctionID = functionID
}

// WriteToRequest writes these params to a swagger request
func (o *FunctionsServiceUpdateFunction2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param function.id
	if err := r.SetPathParam("function.id", o.FunctionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
