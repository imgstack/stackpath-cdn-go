// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfStreamChunkedEncodingResponse custconf stream chunked encoding response
// swagger:model custconfStreamChunkedEncodingResponse
type CustconfStreamChunkedEncodingResponse struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf stream chunked encoding response
func (m *CustconfStreamChunkedEncodingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfStreamChunkedEncodingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfStreamChunkedEncodingResponse) UnmarshalBinary(b []byte) error {
	var res CustconfStreamChunkedEncodingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
