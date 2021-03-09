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

// NewGetCompactJSONParams creates a new GetCompactJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCompactJSONParams() *GetCompactJSONParams {
	return &GetCompactJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCompactJSONParamsWithTimeout creates a new GetCompactJSONParams object
// with the ability to set a timeout on a request.
func NewGetCompactJSONParamsWithTimeout(timeout time.Duration) *GetCompactJSONParams {
	return &GetCompactJSONParams{
		timeout: timeout,
	}
}

// NewGetCompactJSONParamsWithContext creates a new GetCompactJSONParams object
// with the ability to set a context for a request.
func NewGetCompactJSONParamsWithContext(ctx context.Context) *GetCompactJSONParams {
	return &GetCompactJSONParams{
		Context: ctx,
	}
}

// NewGetCompactJSONParamsWithHTTPClient creates a new GetCompactJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCompactJSONParamsWithHTTPClient(client *http.Client) *GetCompactJSONParams {
	return &GetCompactJSONParams{
		HTTPClient: client,
	}
}

/* GetCompactJSONParams contains all the parameters to send to the API endpoint
   for the get compact JSON operation.

   Typically these are written to a http.Request.
*/
type GetCompactJSONParams struct {

	/* Altitude.

	   Whole meters above sea level

	   Format: integer
	*/
	Altitude *int64

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

// WithDefaults hydrates default values in the get compact JSON params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompactJSONParams) WithDefaults() *GetCompactJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compact JSON params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCompactJSONParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compact JSON params
func (o *GetCompactJSONParams) WithTimeout(timeout time.Duration) *GetCompactJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compact JSON params
func (o *GetCompactJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compact JSON params
func (o *GetCompactJSONParams) WithContext(ctx context.Context) *GetCompactJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compact JSON params
func (o *GetCompactJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compact JSON params
func (o *GetCompactJSONParams) WithHTTPClient(client *http.Client) *GetCompactJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compact JSON params
func (o *GetCompactJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAltitude adds the altitude to the get compact JSON params
func (o *GetCompactJSONParams) WithAltitude(altitude *int64) *GetCompactJSONParams {
	o.SetAltitude(altitude)
	return o
}

// SetAltitude adds the altitude to the get compact JSON params
func (o *GetCompactJSONParams) SetAltitude(altitude *int64) {
	o.Altitude = altitude
}

// WithLat adds the lat to the get compact JSON params
func (o *GetCompactJSONParams) WithLat(lat float32) *GetCompactJSONParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the get compact JSON params
func (o *GetCompactJSONParams) SetLat(lat float32) {
	o.Lat = lat
}

// WithLon adds the lon to the get compact JSON params
func (o *GetCompactJSONParams) WithLon(lon float32) *GetCompactJSONParams {
	o.SetLon(lon)
	return o
}

// SetLon adds the lon to the get compact JSON params
func (o *GetCompactJSONParams) SetLon(lon float32) {
	o.Lon = lon
}

// WriteToRequest writes these params to a swagger request
func (o *GetCompactJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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