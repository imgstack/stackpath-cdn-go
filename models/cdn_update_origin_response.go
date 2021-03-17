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

// CdnUpdateOriginResponse The response from a request to update an origin
//
// swagger:model cdnUpdateOriginResponse
type CdnUpdateOriginResponse struct {

	// origin
	Origin *SchemacdnOrigin `json:"origin,omitempty"`
}

// Validate validates this cdn update origin response
func (m *CdnUpdateOriginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnUpdateOriginResponse) validateOrigin(formats strfmt.Registry) error {
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

// ContextValidate validate this cdn update origin response based on the context it is used
func (m *CdnUpdateOriginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnUpdateOriginResponse) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CdnUpdateOriginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnUpdateOriginResponse) UnmarshalBinary(b []byte) error {
	var res CdnUpdateOriginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
