// Code generated by go-swagger; DO NOT EDIT.

package appscape_service

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

// NewAppscapeServiceGetLogLabelsParams creates a new AppscapeServiceGetLogLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppscapeServiceGetLogLabelsParams() *AppscapeServiceGetLogLabelsParams {
	return &AppscapeServiceGetLogLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppscapeServiceGetLogLabelsParamsWithTimeout creates a new AppscapeServiceGetLogLabelsParams object
// with the ability to set a timeout on a request.
func NewAppscapeServiceGetLogLabelsParamsWithTimeout(timeout time.Duration) *AppscapeServiceGetLogLabelsParams {
	return &AppscapeServiceGetLogLabelsParams{
		timeout: timeout,
	}
}

// NewAppscapeServiceGetLogLabelsParamsWithContext creates a new AppscapeServiceGetLogLabelsParams object
// with the ability to set a context for a request.
func NewAppscapeServiceGetLogLabelsParamsWithContext(ctx context.Context) *AppscapeServiceGetLogLabelsParams {
	return &AppscapeServiceGetLogLabelsParams{
		Context: ctx,
	}
}

// NewAppscapeServiceGetLogLabelsParamsWithHTTPClient creates a new AppscapeServiceGetLogLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppscapeServiceGetLogLabelsParamsWithHTTPClient(client *http.Client) *AppscapeServiceGetLogLabelsParams {
	return &AppscapeServiceGetLogLabelsParams{
		HTTPClient: client,
	}
}

/* AppscapeServiceGetLogLabelsParams contains all the parameters to send to the API endpoint
   for the appscape service get log labels operation.

   Typically these are written to a http.Request.
*/
type AppscapeServiceGetLogLabelsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the appscape service get log labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppscapeServiceGetLogLabelsParams) WithDefaults() *AppscapeServiceGetLogLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the appscape service get log labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppscapeServiceGetLogLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) WithTimeout(timeout time.Duration) *AppscapeServiceGetLogLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) WithContext(ctx context.Context) *AppscapeServiceGetLogLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) WithHTTPClient(client *http.Client) *AppscapeServiceGetLogLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the appscape service get log labels params
func (o *AppscapeServiceGetLogLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AppscapeServiceGetLogLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
