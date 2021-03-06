// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStartUpdateParams creates a new StartUpdateParams object
// with the default values initialized.
func NewStartUpdateParams() *StartUpdateParams {
	var ()
	return &StartUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartUpdateParamsWithTimeout creates a new StartUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartUpdateParamsWithTimeout(timeout time.Duration) *StartUpdateParams {
	var ()
	return &StartUpdateParams{

		timeout: timeout,
	}
}

// NewStartUpdateParamsWithContext creates a new StartUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartUpdateParamsWithContext(ctx context.Context) *StartUpdateParams {
	var ()
	return &StartUpdateParams{

		Context: ctx,
	}
}

// NewStartUpdateParamsWithHTTPClient creates a new StartUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartUpdateParamsWithHTTPClient(client *http.Client) *StartUpdateParams {
	var ()
	return &StartUpdateParams{
		HTTPClient: client,
	}
}

/*StartUpdateParams contains all the parameters to send to the API endpoint
for the start update operation typically these are written to a http.Request
*/
type StartUpdateParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start update params
func (o *StartUpdateParams) WithTimeout(timeout time.Duration) *StartUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start update params
func (o *StartUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start update params
func (o *StartUpdateParams) WithContext(ctx context.Context) *StartUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start update params
func (o *StartUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start update params
func (o *StartUpdateParams) WithHTTPClient(client *http.Client) *StartUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start update params
func (o *StartUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start update params
func (o *StartUpdateParams) WithBody(body interface{}) *StartUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start update params
func (o *StartUpdateParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
