// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnUpdateScopeRuleConfigurationResponse cdn update scope rule configuration response
// swagger:model cdnUpdateScopeRuleConfigurationResponse
type CdnUpdateScopeRuleConfigurationResponse struct {

	// configuration
	Configuration *CustconfConfiguration `json:"configuration,omitempty"`
}

// Validate validates this cdn update scope rule configuration response
func (m *CdnUpdateScopeRuleConfigurationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnUpdateScopeRuleConfigurationResponse) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnUpdateScopeRuleConfigurationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnUpdateScopeRuleConfigurationResponse) UnmarshalBinary(b []byte) error {
	var res CdnUpdateScopeRuleConfigurationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
