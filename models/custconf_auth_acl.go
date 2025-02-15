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

// CustconfAuthACL IP address restrictions allow you to configure your CDN container to grant or deny a specific IP addresses or range of IP addresses from accessing content cached in a directory in your CDN container.
//
// swagger:model custconfAuthACL
type CustconfAuthACL struct {

	// access code
	AccessCode *AuthACLAccessCodeEnumWrapperValue `json:"accessCode,omitempty"`

	// client IP src
	ClientIPSrc *AuthACLClientIPSrcEnumWrapperValue `json:"clientIPSrc,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This allows you to specify the name of a HTTP request header which will contain the client IP address to use for this policy.
	Header string `json:"header,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character. This is a list of IP addresses considered for this policy.
	IPList string `json:"ipList,omitempty"`

	// protocol
	Protocol *CustconfAuthACLProtocolEnumWrapperValue `json:"protocol,omitempty"`
}

// Validate validates this custconf auth ACL
func (m *CustconfAuthACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIPSrc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthACL) validateAccessCode(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessCode) { // not required
		return nil
	}

	if m.AccessCode != nil {
		if err := m.AccessCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCode")
			}
			return err
		}
	}

	return nil
}

func (m *CustconfAuthACL) validateClientIPSrc(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientIPSrc) { // not required
		return nil
	}

	if m.ClientIPSrc != nil {
		if err := m.ClientIPSrc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientIPSrc")
			}
			return err
		}
	}

	return nil
}

func (m *CustconfAuthACL) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this custconf auth ACL based on the context it is used
func (m *CustconfAuthACL) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientIPSrc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthACL) contextValidateAccessCode(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessCode != nil {
		if err := m.AccessCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessCode")
			}
			return err
		}
	}

	return nil
}

func (m *CustconfAuthACL) contextValidateClientIPSrc(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientIPSrc != nil {
		if err := m.ClientIPSrc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientIPSrc")
			}
			return err
		}
	}

	return nil
}

func (m *CustconfAuthACL) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {
		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthACL) UnmarshalBinary(b []byte) error {
	var res CustconfAuthACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
