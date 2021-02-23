// Code generated by go-swagger; DO NOT EDIT.

package distr_config_service

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

// NewDistrConfigServiceUpdatePool2Params creates a new DistrConfigServiceUpdatePool2Params object
// with the default values initialized.
func NewDistrConfigServiceUpdatePool2Params() *DistrConfigServiceUpdatePool2Params {
	var ()
	return &DistrConfigServiceUpdatePool2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDistrConfigServiceUpdatePool2ParamsWithTimeout creates a new DistrConfigServiceUpdatePool2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDistrConfigServiceUpdatePool2ParamsWithTimeout(timeout time.Duration) *DistrConfigServiceUpdatePool2Params {
	var ()
	return &DistrConfigServiceUpdatePool2Params{

		timeout: timeout,
	}
}

// NewDistrConfigServiceUpdatePool2ParamsWithContext creates a new DistrConfigServiceUpdatePool2Params object
// with the default values initialized, and the ability to set a context for a request
func NewDistrConfigServiceUpdatePool2ParamsWithContext(ctx context.Context) *DistrConfigServiceUpdatePool2Params {
	var ()
	return &DistrConfigServiceUpdatePool2Params{

		Context: ctx,
	}
}

// NewDistrConfigServiceUpdatePool2ParamsWithHTTPClient creates a new DistrConfigServiceUpdatePool2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDistrConfigServiceUpdatePool2ParamsWithHTTPClient(client *http.Client) *DistrConfigServiceUpdatePool2Params {
	var ()
	return &DistrConfigServiceUpdatePool2Params{
		HTTPClient: client,
	}
}

/*DistrConfigServiceUpdatePool2Params contains all the parameters to send to the API endpoint
for the distr config service update pool2 operation typically these are written to a http.Request
*/
type DistrConfigServiceUpdatePool2Params struct {

	/*Body*/
	Body *models.DistrconfigDataPool
	/*DataPoolID
	  Primary key together with accountId - should be unique by client - used also as name of the pool

	*/
	DataPoolID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) WithTimeout(timeout time.Duration) *DistrConfigServiceUpdatePool2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) WithContext(ctx context.Context) *DistrConfigServiceUpdatePool2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) WithHTTPClient(client *http.Client) *DistrConfigServiceUpdatePool2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) WithBody(body *models.DistrconfigDataPool) *DistrConfigServiceUpdatePool2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) SetBody(body *models.DistrconfigDataPool) {
	o.Body = body
}

// WithDataPoolID adds the dataPoolID to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) WithDataPoolID(dataPoolID string) *DistrConfigServiceUpdatePool2Params {
	o.SetDataPoolID(dataPoolID)
	return o
}

// SetDataPoolID adds the dataPoolId to the distr config service update pool2 params
func (o *DistrConfigServiceUpdatePool2Params) SetDataPoolID(dataPoolID string) {
	o.DataPoolID = dataPoolID
}

// WriteToRequest writes these params to a swagger request
func (o *DistrConfigServiceUpdatePool2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dataPool.id
	if err := r.SetPathParam("dataPool.id", o.DataPoolID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
