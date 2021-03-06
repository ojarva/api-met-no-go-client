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
	"github.com/go-openapi/swag"
)

// NewGetCompactFormatParams creates a new GetCompactFormatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompactFormatParams() *GetCompactFormatParams {
	return &GetCompactFormatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompactFormatParamsWithTimeout creates a new GetCompactFormatParams object
// with the ability to set a timeout on a request.
func NewGetCompactFormatParamsWithTimeout(timeout time.Duration) *GetCompactFormatParams {
	return &GetCompactFormatParams{
		timeout: timeout,
	}
}

// NewGetCompactFormatParamsWithContext creates a new GetCompactFormatParams object
// with the ability to set a context for a request.
func NewGetCompactFormatParamsWithContext(ctx context.Context) *GetCompactFormatParams {
	return &GetCompactFormatParams{
		Context: ctx,
	}
}

// NewGetCompactFormatParamsWithHTTPClient creates a new GetCompactFormatParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompactFormatParamsWithHTTPClient(client *http.Client) *GetCompactFormatParams {
	return &GetCompactFormatParams{
		HTTPClient: client,
	}
}

/* GetCompactFormatParams contains all the parameters to send to the API endpoint
   for the get compact format operation.

   Typically these are written to a http.Request.
*/
type GetCompactFormatParams struct {

	/* Altitude.

	   Whole meters above sea level

	   Format: integer
	*/
	Altitude *int64

	/* Format.

	   format code (file extension)

	   Format: string
	*/
	Format string

	/* Lat.

	   Latitude

	   Format: float
	*/
	Lat float32

	/* Lon.

	   Longitude

	   Format: float
	*/
	Lon float32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compact format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompactFormatParams) WithDefaults() *GetCompactFormatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compact format params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompactFormatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compact format params
func (o *GetCompactFormatParams) WithTimeout(timeout time.Duration) *GetCompactFormatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compact format params
func (o *GetCompactFormatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compact format params
func (o *GetCompactFormatParams) WithContext(ctx context.Context) *GetCompactFormatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compact format params
func (o *GetCompactFormatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compact format params
func (o *GetCompactFormatParams) WithHTTPClient(client *http.Client) *GetCompactFormatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compact format params
func (o *GetCompactFormatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAltitude adds the altitude to the get compact format params
func (o *GetCompactFormatParams) WithAltitude(altitude *int64) *GetCompactFormatParams {
	o.SetAltitude(altitude)
	return o
}

// SetAltitude adds the altitude to the get compact format params
func (o *GetCompactFormatParams) SetAltitude(altitude *int64) {
	o.Altitude = altitude
}

// WithFormat adds the format to the get compact format params
func (o *GetCompactFormatParams) WithFormat(format string) *GetCompactFormatParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get compact format params
func (o *GetCompactFormatParams) SetFormat(format string) {
	o.Format = format
}

// WithLat adds the lat to the get compact format params
func (o *GetCompactFormatParams) WithLat(lat float32) *GetCompactFormatParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the get compact format params
func (o *GetCompactFormatParams) SetLat(lat float32) {
	o.Lat = lat
}

// WithLon adds the lon to the get compact format params
func (o *GetCompactFormatParams) WithLon(lon float32) *GetCompactFormatParams {
	o.SetLon(lon)
	return o
}

// SetLon adds the lon to the get compact format params
func (o *GetCompactFormatParams) SetLon(lon float32) {
	o.Lon = lon
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompactFormatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Altitude != nil {

		// query param altitude
		var qrAltitude int64

		if o.Altitude != nil {
			qrAltitude = *o.Altitude
		}
		qAltitude := swag.FormatInt64(qrAltitude)
		if qAltitude != "" {

			if err := r.SetQueryParam("altitude", qAltitude); err != nil {
				return err
			}
		}
	}

	// path param format
	if err := r.SetPathParam("format", o.Format); err != nil {
		return err
	}

	// query param lat
	qrLat := o.Lat
	qLat := swag.FormatFloat32(qrLat)
	if qLat != "" {

		if err := r.SetQueryParam("lat", qLat); err != nil {
			return err
		}
	}

	// query param lon
	qrLon := o.Lon
	qLon := swag.FormatFloat32(qrLon)
	if qLon != "" {

		if err := r.SetQueryParam("lon", qLon); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
