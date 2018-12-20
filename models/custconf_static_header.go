// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfStaticHeader The static header injection policy allows you to insert headers into the CDN
// request and response processor.
// swagger:model custconfStaticHeader
type CustconfStaticHeader struct {

	// This is the static HTTP header you want inserted into the CDN request.
	ClientRequest string `json:"clientRequest,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is the static HTTP header you want inserted into the CDN response.
	HTTP string `json:"http,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// This is the HTTP header you want inserted into the origin pull request.
	OriginPull string `json:"originPull,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`
}

// Validate validates this custconf static header
func (m *CustconfStaticHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfStaticHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfStaticHeader) UnmarshalBinary(b []byte) error {
	var res CustconfStaticHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
