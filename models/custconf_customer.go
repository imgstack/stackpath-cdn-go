// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfCustomer custconf customer
//
// swagger:model custconfCustomer
type CustconfCustomer struct {

	// String of values delimited by a ',' character.
	AccessLogFields string `json:"accessLogFields,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character.
	OpLogFields string `json:"opLogFields,omitempty"`

	// String of values delimited by a ',' character.
	ReceiptLogFields string `json:"receiptLogFields,omitempty"`
}

// Validate validates this custconf customer
func (m *CustconfCustomer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf customer based on context it is used
func (m *CustconfCustomer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfCustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfCustomer) UnmarshalBinary(b []byte) error {
	var res CustconfCustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
