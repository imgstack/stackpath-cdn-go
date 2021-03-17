// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfVHost A hostname policy allows you to specify an alternate domain name that you want to use to access content from your CDN container.
//
// swagger:model custconfVHost
type CustconfVHost struct {

	// This is the hostname you want to enable in this policy. Note: You must configure your container's CNAME record with your DNS provider to enable this hostname to deliver content.
	Domain string `json:"domain,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf v host
func (m *CustconfVHost) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf v host based on context it is used
func (m *CustconfVHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfVHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfVHost) UnmarshalBinary(b []byte) error {
	var res CustconfVHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
