// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfReceiptLogs Enable/disable delivery receipt access logs
// swagger:model custconfReceiptLogs
type CustconfReceiptLogs struct {

	// Determines whether receipt access logging be enabled for this customer.
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf receipt logs
func (m *CustconfReceiptLogs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfReceiptLogs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfReceiptLogs) UnmarshalBinary(b []byte) error {
	var res CustconfReceiptLogs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
