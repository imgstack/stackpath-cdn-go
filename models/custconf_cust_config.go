// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfCustConfig custconf cust config
// swagger:model custconfCustConfig
type CustconfCustConfig struct {

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// micro seconds
	MicroSeconds string `json:"microSeconds,omitempty"`
}

// Validate validates this custconf cust config
func (m *CustconfCustConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfCustConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfCustConfig) UnmarshalBinary(b []byte) error {
	var res CustconfCustConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
