// Code generated by go-swagger; DO NOT EDIT.

package data

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

// NewGetStatusFormatParams creates a new GetStatusFormatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStatusFormatParams() *GetStatusFormatParams {
	return &GetStatusFormatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusFormatParamsWithTimeout creates a new GetStatusFormatParams object
// with the ability to set a timeout on a request.
func NewGetStatusFormatParamsWithTimeout(timeout time.Duration) *GetStatusFormatParams {
	return &GetStatusFormatParams{
		timeout: timeout,
	}
}

// NewGetStatusFormatParamsWithContext creates a new GetStatusFormatParams object
// with the ability to set a context for a request.
func NewGetStatusFormatParamsWithContext(ctx context.Context) *GetStatusFormatParams {
	return &GetStatusFormatParams{
		Context: ctx,
	}
}

// NewGetStatusFormatParamsWithHTTPClient creates a new GetStatusFormatParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStatusFormatParamsWithHTTPClient(client *http.Client) *GetStatusFormatParams {
	return &GetStatusFormatParams{
		HTTPClient: client,
	}
}

/* GetStatusFormatParams contains all the parameters to send to the API endpoint
   for the get status format operation.

   Typically these are written to a http.Request.
*/
type GetStatusFormatParams struct {

	/* Format.

	   format code (file extension)

	   Format: string
	*/
	Format string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get status format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusFormatParams) WithDefaults() *GetStatusFormatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get status format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStatusFormatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get status format params
func (o *GetStatusFormatParams) WithTimeout(timeout time.Duration) *GetStatusFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get status format params
func (o *GetStatusFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get status format params
func (o *GetStatusFormatParams) WithContext(ctx context.Context) *GetStatusFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get status format params
func (o *GetStatusFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get status format params
func (o *GetStatusFormatParams) WithHTTPClient(client *http.Client) *GetStatusFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get status format params
func (o *GetStatusFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get status format params
func (o *GetStatusFormatParams) WithFormat(format string) *GetStatusFormatParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get status format params
func (o *GetStatusFormatParams) SetFormat(format string) {
	o.Format = format
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param format
	if err := r.SetPathParam("format", o.Format); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
