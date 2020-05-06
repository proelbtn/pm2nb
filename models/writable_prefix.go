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

// WritablePrefix writable prefix
//
// swagger:model WritablePrefix
type WritablePrefix struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Family
	// Read Only: true
	Family string `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is a pool
	//
	// All IP addresses within this prefix are considered usable
	IsPool bool `json:"is_pool,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Prefix
	//
	// IPv4 or IPv6 network with mask
	// Required: true
	Prefix *string `json:"prefix"`

	// Role
	//
	// The primary function of this prefix
	Role *int64 `json:"role,omitempty"`

	// Site
	Site *int64 `json:"site,omitempty"`

	// Status
	//
	// Operational status of this prefix
	// Enum: [container active reserved deprecated]
	Status string `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// VLAN
	Vlan *int64 `json:"vlan,omitempty"`

	// VRF
	Vrf *int64 `json:"vrf,omitempty"`
}

// Validate validates this writable prefix
func (m *WritablePrefix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePrefix) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePrefix) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePrefix) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePrefix) validatePrefix(formats strfmt.Registry) error {

	if err := validate.Required("prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

var writablePrefixTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["container","active","reserved","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePrefixTypeStatusPropEnum = append(writablePrefixTypeStatusPropEnum, v)
	}
}

const (

	// WritablePrefixStatusContainer captures enum value "container"
	WritablePrefixStatusContainer string = "container"

	// WritablePrefixStatusActive captures enum value "active"
	WritablePrefixStatusActive string = "active"

	// WritablePrefixStatusReserved captures enum value "reserved"
	WritablePrefixStatusReserved string = "reserved"

	// WritablePrefixStatusDeprecated captures enum value "deprecated"
	WritablePrefixStatusDeprecated string = "deprecated"
)

// prop value enum
func (m *WritablePrefix) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writablePrefixTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritablePrefix) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritablePrefix) validateTags(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *WritablePrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePrefix) UnmarshalBinary(b []byte) error {
	var res WritablePrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
