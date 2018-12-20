// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfBandWidthRateLimitUnits The Bandwidth Throttling Units policy allows you to override the default
// units used by the CDN when processing the bandwidth throttling policies.
// swagger:model custconfBandWidthRateLimitUnits
type CustconfBandWidthRateLimitUnits struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// initial
	Initial BandWidthRateLimitUnitsInitialEnumWrapperValue `json:"initial,omitempty"`

	// sustained
	Sustained BandWidthRateLimitUnitsSustainedEnumWrapperValue `json:"sustained,omitempty"`
}

// Validate validates this custconf band width rate limit units
func (m *CustconfBandWidthRateLimitUnits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInitial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSustained(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfBandWidthRateLimitUnits) validateInitial(formats strfmt.Registry) error {

	if swag.IsZero(m.Initial) { // not required
		return nil
	}

	if err := m.Initial.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("initial")
		}
		return err
	}

	return nil
}

func (m *CustconfBandWidthRateLimitUnits) validateSustained(formats strfmt.Registry) error {

	if swag.IsZero(m.Sustained) { // not required
		return nil
	}

	if err := m.Sustained.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sustained")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfBandWidthRateLimitUnits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfBandWidthRateLimitUnits) UnmarshalBinary(b []byte) error {
	var res CustconfBandWidthRateLimitUnits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
