// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfResponseHeader The custom HTTP response headers policy allows you to specify a list of HTTP
// headers you want the CDN caching servers to include in the response to
// clients.
// swagger:model custconfResponseHeader
type CustconfResponseHeader struct {

	// This gives the ability to disable the ETag header on the response.
	EnableETag bool `json:"enableETag,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// A pipe delimited list of rules that instructs the CDN caching servers to
	// include a content-disposition header. The rules included in this setting
	// must be entered in the following format: Content-Disposition:<User
	// Agent>:<file extension 1>, <file extension 2>.  For example, to send the
	// Content-Disposition header for all Mozilla browsers on the delivery of mp3,
	// exe, tar, zip, gz and rar files, you would enter the following in the
	// field: Content-Disposition:Mozilla*:mp3,exe,tar,zip,gz,rar
	HTTP string `json:"http,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf response header
func (m *CustconfResponseHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfResponseHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfResponseHeader) UnmarshalBinary(b []byte) error {
	var res CustconfResponseHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
