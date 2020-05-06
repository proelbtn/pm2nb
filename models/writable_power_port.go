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

// WritablePowerPort writable power port
//
// swagger:model WritablePowerPort
type WritablePowerPort struct {

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Maximum draw
	//
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// tags
	Tags []string `json:"tags"`

	// Type
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-20p nema-l6-30p nema-l6-50p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o]
	Type string `json:"type,omitempty"`
}

// Validate validates this writable power port
func (m *WritablePowerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *WritablePowerPort) validateAllocatedDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", int64(*m.AllocatedDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateCable(formats strfmt.Registry) error {

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

var writablePowerPortTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerPortTypeConnectionStatusPropEnum = append(writablePowerPortTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritablePowerPort) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writablePowerPortTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerPort) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateMaximumDraw(formats strfmt.Registry) error {

	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", int64(*m.MaximumDraw), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", int64(*m.MaximumDraw), 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateTags(formats strfmt.Registry) error {

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

var writablePowerPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-20p","nema-l6-30p","nema-l6-50p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerPortTypeTypePropEnum = append(writablePowerPortTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerPortTypeIec60320C6 captures enum value "iec-60320-c6"
	WritablePowerPortTypeIec60320C6 string = "iec-60320-c6"

	// WritablePowerPortTypeIec60320C8 captures enum value "iec-60320-c8"
	WritablePowerPortTypeIec60320C8 string = "iec-60320-c8"

	// WritablePowerPortTypeIec60320C14 captures enum value "iec-60320-c14"
	WritablePowerPortTypeIec60320C14 string = "iec-60320-c14"

	// WritablePowerPortTypeIec60320C16 captures enum value "iec-60320-c16"
	WritablePowerPortTypeIec60320C16 string = "iec-60320-c16"

	// WritablePowerPortTypeIec60320C20 captures enum value "iec-60320-c20"
	WritablePowerPortTypeIec60320C20 string = "iec-60320-c20"

	// WritablePowerPortTypeIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerPortTypeIec60309pne4h string = "iec-60309-p-n-e-4h"

	// WritablePowerPortTypeIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerPortTypeIec60309pne6h string = "iec-60309-p-n-e-6h"

	// WritablePowerPortTypeIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerPortTypeIec60309pne9h string = "iec-60309-p-n-e-9h"

	// WritablePowerPortTypeIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerPortTypeIec603092pe4h string = "iec-60309-2p-e-4h"

	// WritablePowerPortTypeIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerPortTypeIec603092pe6h string = "iec-60309-2p-e-6h"

	// WritablePowerPortTypeIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerPortTypeIec603092pe9h string = "iec-60309-2p-e-9h"

	// WritablePowerPortTypeIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerPortTypeIec603093pe4h string = "iec-60309-3p-e-4h"

	// WritablePowerPortTypeIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerPortTypeIec603093pe6h string = "iec-60309-3p-e-6h"

	// WritablePowerPortTypeIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerPortTypeIec603093pe9h string = "iec-60309-3p-e-9h"

	// WritablePowerPortTypeIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerPortTypeIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerPortTypeIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerPortTypeIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerPortTypeIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerPortTypeIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerPortTypeNema515p captures enum value "nema-5-15p"
	WritablePowerPortTypeNema515p string = "nema-5-15p"

	// WritablePowerPortTypeNema520p captures enum value "nema-5-20p"
	WritablePowerPortTypeNema520p string = "nema-5-20p"

	// WritablePowerPortTypeNema530p captures enum value "nema-5-30p"
	WritablePowerPortTypeNema530p string = "nema-5-30p"

	// WritablePowerPortTypeNema550p captures enum value "nema-5-50p"
	WritablePowerPortTypeNema550p string = "nema-5-50p"

	// WritablePowerPortTypeNema615p captures enum value "nema-6-15p"
	WritablePowerPortTypeNema615p string = "nema-6-15p"

	// WritablePowerPortTypeNema620p captures enum value "nema-6-20p"
	WritablePowerPortTypeNema620p string = "nema-6-20p"

	// WritablePowerPortTypeNema630p captures enum value "nema-6-30p"
	WritablePowerPortTypeNema630p string = "nema-6-30p"

	// WritablePowerPortTypeNema650p captures enum value "nema-6-50p"
	WritablePowerPortTypeNema650p string = "nema-6-50p"

	// WritablePowerPortTypeNemaL515p captures enum value "nema-l5-15p"
	WritablePowerPortTypeNemaL515p string = "nema-l5-15p"

	// WritablePowerPortTypeNemaL520p captures enum value "nema-l5-20p"
	WritablePowerPortTypeNemaL520p string = "nema-l5-20p"

	// WritablePowerPortTypeNemaL530p captures enum value "nema-l5-30p"
	WritablePowerPortTypeNemaL530p string = "nema-l5-30p"

	// WritablePowerPortTypeNemaL550p captures enum value "nema-l5-50p"
	WritablePowerPortTypeNemaL550p string = "nema-l5-50p"

	// WritablePowerPortTypeNemaL620p captures enum value "nema-l6-20p"
	WritablePowerPortTypeNemaL620p string = "nema-l6-20p"

	// WritablePowerPortTypeNemaL630p captures enum value "nema-l6-30p"
	WritablePowerPortTypeNemaL630p string = "nema-l6-30p"

	// WritablePowerPortTypeNemaL650p captures enum value "nema-l6-50p"
	WritablePowerPortTypeNemaL650p string = "nema-l6-50p"

	// WritablePowerPortTypeCs6361c captures enum value "cs6361c"
	WritablePowerPortTypeCs6361c string = "cs6361c"

	// WritablePowerPortTypeCs6365c captures enum value "cs6365c"
	WritablePowerPortTypeCs6365c string = "cs6365c"

	// WritablePowerPortTypeCs8165c captures enum value "cs8165c"
	WritablePowerPortTypeCs8165c string = "cs8165c"

	// WritablePowerPortTypeCs8265c captures enum value "cs8265c"
	WritablePowerPortTypeCs8265c string = "cs8265c"

	// WritablePowerPortTypeCs8365c captures enum value "cs8365c"
	WritablePowerPortTypeCs8365c string = "cs8365c"

	// WritablePowerPortTypeCs8465c captures enum value "cs8465c"
	WritablePowerPortTypeCs8465c string = "cs8465c"

	// WritablePowerPortTypeItae captures enum value "ita-e"
	WritablePowerPortTypeItae string = "ita-e"

	// WritablePowerPortTypeItaf captures enum value "ita-f"
	WritablePowerPortTypeItaf string = "ita-f"

	// WritablePowerPortTypeItaEf captures enum value "ita-ef"
	WritablePowerPortTypeItaEf string = "ita-ef"

	// WritablePowerPortTypeItag captures enum value "ita-g"
	WritablePowerPortTypeItag string = "ita-g"

	// WritablePowerPortTypeItah captures enum value "ita-h"
	WritablePowerPortTypeItah string = "ita-h"

	// WritablePowerPortTypeItai captures enum value "ita-i"
	WritablePowerPortTypeItai string = "ita-i"

	// WritablePowerPortTypeItaj captures enum value "ita-j"
	WritablePowerPortTypeItaj string = "ita-j"

	// WritablePowerPortTypeItak captures enum value "ita-k"
	WritablePowerPortTypeItak string = "ita-k"

	// WritablePowerPortTypeItal captures enum value "ita-l"
	WritablePowerPortTypeItal string = "ita-l"

	// WritablePowerPortTypeItam captures enum value "ita-m"
	WritablePowerPortTypeItam string = "ita-m"

	// WritablePowerPortTypeItan captures enum value "ita-n"
	WritablePowerPortTypeItan string = "ita-n"

	// WritablePowerPortTypeItao captures enum value "ita-o"
	WritablePowerPortTypeItao string = "ita-o"
)

// prop value enum
func (m *WritablePowerPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePowerPortTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerPort) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerPort) UnmarshalBinary(b []byte) error {
	var res WritablePowerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
