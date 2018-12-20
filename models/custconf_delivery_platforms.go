// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfDeliveryPlatforms custconf delivery platforms
// swagger:model custconfDeliveryPlatforms
type CustconfDeliveryPlatforms struct {

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	Platforms string `json:"platforms,omitempty"`
}

// Validate validates this custconf delivery platforms
func (m *CustconfDeliveryPlatforms) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfDeliveryPlatforms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfDeliveryPlatforms) UnmarshalBinary(b []byte) error {
	var res CustconfDeliveryPlatforms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
