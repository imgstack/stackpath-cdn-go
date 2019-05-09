// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfReceiptLogsConfig Delivery receipt log settings
// swagger:model custconfReceiptLogsConfig
type CustconfReceiptLogsConfig struct {

	// enabled
	Enabled bool `json:"enabled"`

	// String of values delimited by a ',' character.
	ExtraLogFields string `json:"extraLogFields,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf receipt logs config
func (m *CustconfReceiptLogsConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfReceiptLogsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfReceiptLogsConfig) UnmarshalBinary(b []byte) error {
	var res CustconfReceiptLogsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
