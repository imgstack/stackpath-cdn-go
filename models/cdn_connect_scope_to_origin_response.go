// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnConnectScopeToOriginResponse cdn connect scope to origin response
// swagger:model cdnConnectScopeToOriginResponse
type CdnConnectScopeToOriginResponse struct {

	// scope origin
	ScopeOrigin *CdnScopeOrigin `json:"scopeOrigin,omitempty"`
}

// Validate validates this cdn connect scope to origin response
func (m *CdnConnectScopeToOriginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScopeOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnConnectScopeToOriginResponse) validateScopeOrigin(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeOrigin) { // not required
		return nil
	}

	if m.ScopeOrigin != nil {
		if err := m.ScopeOrigin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scopeOrigin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnConnectScopeToOriginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnConnectScopeToOriginResponse) UnmarshalBinary(b []byte) error {
	var res CdnConnectScopeToOriginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
