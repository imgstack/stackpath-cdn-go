// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfOriginRequestModification custconf origin request modification
// swagger:model custconfOriginRequestModification
type CustconfOriginRequestModification struct {

	// String of values delimited by a '|' character. This is the static HTTP header you want inserted into the CDN request. Start value with "append:", "replace:" or "create:" which defines if Header will be added, replaced or create if not exists. Default is append.
	AddHeaders string `json:"addHeaders,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// flow control
	FlowControl string `json:"flowControl,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// header pattern
	HeaderPattern string `json:"headerPattern,omitempty"`

	// header rewrite
	HeaderRewrite string `json:"headerRewrite,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// url pattern
	URLPattern string `json:"urlPattern,omitempty"`

	// url rewrite
	URLRewrite string `json:"urlRewrite,omitempty"`
}

// Validate validates this custconf origin request modification
func (m *CustconfOriginRequestModification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfOriginRequestModification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfOriginRequestModification) UnmarshalBinary(b []byte) error {
	var res CustconfOriginRequestModification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
