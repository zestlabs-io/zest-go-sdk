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

// NewAppscapeServiceGetMetricsMetaParams creates a new AppscapeServiceGetMetricsMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppscapeServiceGetMetricsMetaParams() *AppscapeServiceGetMetricsMetaParams {
	return &AppscapeServiceGetMetricsMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppscapeServiceGetMetricsMetaParamsWithTimeout creates a new AppscapeServiceGetMetricsMetaParams object
// with the ability to set a timeout on a request.
func NewAppscapeServiceGetMetricsMetaParamsWithTimeout(timeout time.Duration) *AppscapeServiceGetMetricsMetaParams {
	return &AppscapeServiceGetMetricsMetaParams{
		timeout: timeout,
	}
}

// NewAppscapeServiceGetMetricsMetaParamsWithContext creates a new AppscapeServiceGetMetricsMetaParams object
// with the ability to set a context for a request.
func NewAppscapeServiceGetMetricsMetaParamsWithContext(ctx context.Context) *AppscapeServiceGetMetricsMetaParams {
	return &AppscapeServiceGetMetricsMetaParams{
		Context: ctx,
	}
}

// NewAppscapeServiceGetMetricsMetaParamsWithHTTPClient creates a new AppscapeServiceGetMetricsMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppscapeServiceGetMetricsMetaParamsWithHTTPClient(client *http.Client) *AppscapeServiceGetMetricsMetaParams {
	return &AppscapeServiceGetMetricsMetaParams{
		HTTPClient: client,
	}
}

/* AppscapeServiceGetMetricsMetaParams contains all the parameters to send to the API endpoint
   for the appscape service get metrics meta operation.

   Typically these are written to a http.Request.
*/
type AppscapeServiceGetMetricsMetaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the appscape service get metrics meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppscapeServiceGetMetricsMetaParams) WithDefaults() *AppscapeServiceGetMetricsMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the appscape service get metrics meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppscapeServiceGetMetricsMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) WithTimeout(timeout time.Duration) *AppscapeServiceGetMetricsMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) WithContext(ctx context.Context) *AppscapeServiceGetMetricsMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) WithHTTPClient(client *http.Client) *AppscapeServiceGetMetricsMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the appscape service get metrics meta params
func (o *AppscapeServiceGetMetricsMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AppscapeServiceGetMetricsMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
