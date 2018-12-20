// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnCreateScopeHostnameResponse cdn create scope hostname response
// swagger:model cdnCreateScopeHostnameResponse
type CdnCreateScopeHostnameResponse struct {

	// hostname
	Hostname *CdnHostname `json:"hostname,omitempty"`
}

// Validate validates this cdn create scope hostname response
func (m *CdnCreateScopeHostnameResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnCreateScopeHostnameResponse) validateHostname(formats strfmt.Registry) error {

	if swag.IsZero(m.Hostname) { // not required
		return nil
	}

	if m.Hostname != nil {
		if err := m.Hostname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostname")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnCreateScopeHostnameResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnCreateScopeHostnameResponse) UnmarshalBinary(b []byte) error {
	var res CdnCreateScopeHostnameResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
