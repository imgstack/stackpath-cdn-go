// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnConnectScopeToOriginRequest A request to associate an origin with a CDN site scope
//
// swagger:model cdnConnectScopeToOriginRequest
type CdnConnectScopeToOriginRequest struct {

	// origin
	Origin *CdnConnectScopeToOriginRequestOrigin `json:"origin,omitempty"`

	// The ID of an existing origin to associate with a scope
	//
	// This is useful for connecting to a shared origin.
	OriginID string `json:"originId,omitempty"`

	// The origin's priority to the scope
	//
	// If a CDN scope is powered by more than one origin, then the one with the lower priority number takes higher precedence.
	Priority int32 `json:"priority,omitempty"`
}

// Validate validates this cdn connect scope to origin request
func (m *CdnConnectScopeToOriginRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnConnectScopeToOriginRequest) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cdn connect scope to origin request based on the context it is used
func (m *CdnConnectScopeToOriginRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnConnectScopeToOriginRequest) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if m.Origin != nil {
		if err := m.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnConnectScopeToOriginRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnConnectScopeToOriginRequest) UnmarshalBinary(b []byte) error {
	var res CdnConnectScopeToOriginRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
