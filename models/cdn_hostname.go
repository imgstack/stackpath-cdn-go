// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CdnHostname cdn hostname
// swagger:model cdnHostname
type CdnHostname struct {

	// domain
	Domain string `json:"domain,omitempty"`
}

// Validate validates this cdn hostname
func (m *CdnHostname) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnHostname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnHostname) UnmarshalBinary(b []byte) error {
	var res CdnHostname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
