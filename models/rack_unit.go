// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RackUnit rack unit
//
// swagger:model RackUnit
type RackUnit struct {

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// face
	Face *RackUnitFace `json:"face,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Read Only: true
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this rack unit
func (m *RackUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackUnit) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *RackUnit) validateFace(formats strfmt.Registry) error {

	if swag.IsZero(m.Face) { // not required
		return nil
	}

	if m.Face != nil {
		if err := m.Face.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("face")
			}
			return err
		}
	}

	return nil
}

func (m *RackUnit) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", string(m.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackUnit) UnmarshalBinary(b []byte) error {
	var res RackUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackUnitFace Face
//
// swagger:model RackUnitFace
type RackUnitFace struct {

	// label
	// Required: true
	// Enum: [Front Rear]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [front rear]
	Value *string `json:"value"`
}

// Validate validates this rack unit face
func (m *RackUnitFace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rackUnitFaceTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Front","Rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackUnitFaceTypeLabelPropEnum = append(rackUnitFaceTypeLabelPropEnum, v)
	}
}

const (

	// RackUnitFaceLabelFront captures enum value "Front"
	RackUnitFaceLabelFront string = "Front"

	// RackUnitFaceLabelRear captures enum value "Rear"
	RackUnitFaceLabelRear string = "Rear"
)

// prop value enum
func (m *RackUnitFace) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rackUnitFaceTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RackUnitFace) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("face"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rackUnitFaceTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["front","rear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rackUnitFaceTypeValuePropEnum = append(rackUnitFaceTypeValuePropEnum, v)
	}
}

const (

	// RackUnitFaceValueFront captures enum value "front"
	RackUnitFaceValueFront string = "front"

	// RackUnitFaceValueRear captures enum value "rear"
	RackUnitFaceValueRear string = "rear"
)

// prop value enum
func (m *RackUnitFace) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rackUnitFaceTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RackUnitFace) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("face"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("face"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackUnitFace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackUnitFace) UnmarshalBinary(b []byte) error {
	var res RackUnitFace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}