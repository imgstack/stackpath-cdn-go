// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfFailSafeOriginPull custconf fail safe origin pull
// swagger:model custconfFailSafeOriginPull
type CustconfFailSafeOriginPull struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// String of values deliminated by a ',' character.
	StatusCodeMatch string `json:"statusCodeMatch,omitempty"`
}

// Validate validates this custconf fail safe origin pull
func (m *CustconfFailSafeOriginPull) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfFailSafeOriginPull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfFailSafeOriginPull) UnmarshalBinary(b []byte) error {
	var res CustconfFailSafeOriginPull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
