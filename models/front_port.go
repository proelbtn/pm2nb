// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrontPort front port
//
// swagger:model FrontPort
type FrontPort struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// rear port
	// Required: true
	RearPort *FrontPortRearPort `json:"rear_port"`

	// Rear port position
	// Maximum: 64
	// Minimum: 1
	RearPortPosition int64 `json:"rear_port_position,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// type
	// Required: true
	Type *FrontPortType `json:"type"`
}

// Validate validates this front port
func (m *FrontPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRearPortPosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
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

func (m *FrontPort) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
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

func (m *FrontPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateRearPort(formats strfmt.Registry) error {

	if err := validate.Required("rear_port", "body", m.RearPort); err != nil {
		return err
	}

	if m.RearPort != nil {
		if err := m.RearPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rear_port")
			}
			return err
		}
	}

	return nil
}

func (m *FrontPort) validateRearPortPosition(formats strfmt.Registry) error {

	if swag.IsZero(m.RearPortPosition) { // not required
		return nil
	}

	if err := validate.MinimumInt("rear_port_position", "body", int64(m.RearPortPosition), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rear_port_position", "body", int64(m.RearPortPosition), 64, false); err != nil {
		return err
	}

	return nil
}

func (m *FrontPort) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *FrontPort) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrontPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPort) UnmarshalBinary(b []byte) error {
	var res FrontPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FrontPortType Type
//
// swagger:model FrontPortType
type FrontPortType struct {

	// label
	// Required: true
	// Enum: [8P8C 110 Punch BNC MRJ21 FC LC LC/APC LSH LSH/APC MPO MTRJ SC SC/APC ST]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 110-punch bnc mrj21 fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st]
	Value *string `json:"value"`
}

// Validate validates this front port type
func (m *FrontPortType) Validate(formats strfmt.Registry) error {
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

var frontPortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","110 Punch","BNC","MRJ21","FC","LC","LC/APC","LSH","LSH/APC","MPO","MTRJ","SC","SC/APC","ST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTypeTypeLabelPropEnum = append(frontPortTypeTypeLabelPropEnum, v)
	}
}

const (

	// FrontPortTypeLabelNr8P8C captures enum value "8P8C"
	FrontPortTypeLabelNr8P8C string = "8P8C"

	// FrontPortTypeLabelNr110Punch captures enum value "110 Punch"
	FrontPortTypeLabelNr110Punch string = "110 Punch"

	// FrontPortTypeLabelBNC captures enum value "BNC"
	FrontPortTypeLabelBNC string = "BNC"

	// FrontPortTypeLabelMRJ21 captures enum value "MRJ21"
	FrontPortTypeLabelMRJ21 string = "MRJ21"

	// FrontPortTypeLabelFC captures enum value "FC"
	FrontPortTypeLabelFC string = "FC"

	// FrontPortTypeLabelLC captures enum value "LC"
	FrontPortTypeLabelLC string = "LC"

	// FrontPortTypeLabelLCAPC captures enum value "LC/APC"
	FrontPortTypeLabelLCAPC string = "LC/APC"

	// FrontPortTypeLabelLSH captures enum value "LSH"
	FrontPortTypeLabelLSH string = "LSH"

	// FrontPortTypeLabelLSHAPC captures enum value "LSH/APC"
	FrontPortTypeLabelLSHAPC string = "LSH/APC"

	// FrontPortTypeLabelMPO captures enum value "MPO"
	FrontPortTypeLabelMPO string = "MPO"

	// FrontPortTypeLabelMTRJ captures enum value "MTRJ"
	FrontPortTypeLabelMTRJ string = "MTRJ"

	// FrontPortTypeLabelSC captures enum value "SC"
	FrontPortTypeLabelSC string = "SC"

	// FrontPortTypeLabelSCAPC captures enum value "SC/APC"
	FrontPortTypeLabelSCAPC string = "SC/APC"

	// FrontPortTypeLabelST captures enum value "ST"
	FrontPortTypeLabelST string = "ST"
)

// prop value enum
func (m *FrontPortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontPortTypeTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var frontPortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","110-punch","bnc","mrj21","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontPortTypeTypeValuePropEnum = append(frontPortTypeTypeValuePropEnum, v)
	}
}

const (

	// FrontPortTypeValueNr8p8c captures enum value "8p8c"
	FrontPortTypeValueNr8p8c string = "8p8c"

	// FrontPortTypeValueNr110Punch captures enum value "110-punch"
	FrontPortTypeValueNr110Punch string = "110-punch"

	// FrontPortTypeValueBnc captures enum value "bnc"
	FrontPortTypeValueBnc string = "bnc"

	// FrontPortTypeValueMrj21 captures enum value "mrj21"
	FrontPortTypeValueMrj21 string = "mrj21"

	// FrontPortTypeValueFc captures enum value "fc"
	FrontPortTypeValueFc string = "fc"

	// FrontPortTypeValueLc captures enum value "lc"
	FrontPortTypeValueLc string = "lc"

	// FrontPortTypeValueLcApc captures enum value "lc-apc"
	FrontPortTypeValueLcApc string = "lc-apc"

	// FrontPortTypeValueLsh captures enum value "lsh"
	FrontPortTypeValueLsh string = "lsh"

	// FrontPortTypeValueLshApc captures enum value "lsh-apc"
	FrontPortTypeValueLshApc string = "lsh-apc"

	// FrontPortTypeValueMpo captures enum value "mpo"
	FrontPortTypeValueMpo string = "mpo"

	// FrontPortTypeValueMtrj captures enum value "mtrj"
	FrontPortTypeValueMtrj string = "mtrj"

	// FrontPortTypeValueSc captures enum value "sc"
	FrontPortTypeValueSc string = "sc"

	// FrontPortTypeValueScApc captures enum value "sc-apc"
	FrontPortTypeValueScApc string = "sc-apc"

	// FrontPortTypeValueSt captures enum value "st"
	FrontPortTypeValueSt string = "st"
)

// prop value enum
func (m *FrontPortType) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, frontPortTypeTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FrontPortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrontPortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontPortType) UnmarshalBinary(b []byte) error {
	var res FrontPortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
