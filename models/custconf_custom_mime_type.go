// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfCustomMimeType The custom mime type policy allows you to map file extensions to specific mime types for the CDN caching servers to use when delivering assets. The mime types you map using this policy may also be limited to specific response codes to address scenarios in which the mime type changes based on the response code.
//
// swagger:model custconfCustomMimeType
type CustconfCustomMimeType struct {

	// String of values delimited by a ',' character. A comma separated list of status codes that apply to this policy
	Code string `json:"code,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// String of values delimited by a ',' character. This is a comma separated list of file extension and mime type pairs that describe the mime type mapping for the CDN caching servers. The file extension and mime type values should be delimited by a colon (:). For example, to map files ending with jpg to the image/jpeg mime type and files ending with png to image/png, set this value to: jpg:image/jpeg,png:image/png.
	ExtensionMap string `json:"extensionMap,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`
}

// Validate validates this custconf custom mime type
func (m *CustconfCustomMimeType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf custom mime type based on context it is used
func (m *CustconfCustomMimeType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfCustomMimeType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfCustomMimeType) UnmarshalBinary(b []byte) error {
	var res CustconfCustomMimeType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
