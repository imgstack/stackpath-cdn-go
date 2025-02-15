// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StackpathRPCBadRequestFieldViolation stackpath rpc bad request field violation
//
// swagger:model stackpath.rpc.BadRequest.FieldViolation
type StackpathRPCBadRequestFieldViolation struct {

	// description
	Description string `json:"description,omitempty"`

	// field
	Field string `json:"field,omitempty"`
}

// Validate validates this stackpath rpc bad request field violation
func (m *StackpathRPCBadRequestFieldViolation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stackpath rpc bad request field violation based on context it is used
func (m *StackpathRPCBadRequestFieldViolation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCBadRequestFieldViolation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCBadRequestFieldViolation) UnmarshalBinary(b []byte) error {
	var res StackpathRPCBadRequestFieldViolation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
