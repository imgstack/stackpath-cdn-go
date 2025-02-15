// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfAuthURLSignAliCloudA custconf auth Url sign ali cloud a
//
// swagger:model custconfAuthUrlSignAliCloudA
type CustconfAuthURLSignAliCloudA struct {

	// enabled
	Enabled bool `json:"enabled"`

	// expiration extension
	ExpirationExtension int64 `json:"expirationExtension,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// include params before token
	IncludeParamsBeforeToken bool `json:"includeParamsBeforeToken"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// pass phrase
	PassPhrase string `json:"passPhrase,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// token field
	TokenField string `json:"tokenField,omitempty"`
}

// Validate validates this custconf auth Url sign ali cloud a
func (m *CustconfAuthURLSignAliCloudA) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf auth Url sign ali cloud a based on context it is used
func (m *CustconfAuthURLSignAliCloudA) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLSignAliCloudA) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLSignAliCloudA) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLSignAliCloudA
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
