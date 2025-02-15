// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCHelpLink stackpath rpc help link
//
// swagger:model stackpath.rpc.Help.Link
type StackpathRPCHelpLink struct {

	// description
	Description string `json:"description,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this stackpath rpc help link
func (m *StackpathRPCHelpLink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stackpath rpc help link based on context it is used
func (m *StackpathRPCHelpLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCHelpLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCHelpLink) UnmarshalBinary(b []byte) error {
	var res StackpathRPCHelpLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
