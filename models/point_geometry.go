// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PointGeometry GeoJSON point type
//
// swagger:model PointGeometry
type PointGeometry struct {

	// [longitude, latitude, altitude]. All numbers in decimal.
	// Example: [60.5,11.59,1001]
	// Required: true
	// Min Items: 2
	Coordinates []float64 `json:"coordinates"`

	// type
	// Required: true
	// Enum: [Point]
	Type *string `json:"type"`
}

// Validate validates this point geometry
func (m *PointGeometry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PointGeometry) validateCoordinates(formats strfmt.Registry) error {

	if err := validate.Required("coordinates", "body", m.Coordinates); err != nil {
		return err
	}

	iCoordinatesSize := int64(len(m.Coordinates))

	if err := validate.MinItems("coordinates", "body", iCoordinatesSize, 2); err != nil {
		return err
	}

	return nil
}

var pointGeometryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Point"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pointGeometryTypeTypePropEnum = append(pointGeometryTypeTypePropEnum, v)
	}
}

const (

	// PointGeometryTypePoint captures enum value "Point"
	PointGeometryTypePoint string = "Point"
)

// prop value enum
func (m *PointGeometry) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pointGeometryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PointGeometry) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this point geometry based on context it is used
func (m *PointGeometry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PointGeometry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PointGeometry) UnmarshalBinary(b []byte) error {
	var res PointGeometry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
