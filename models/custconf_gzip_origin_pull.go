// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfGzipOriginPull The compressed origin pull policy allows you to enable the CDN caching servers to request compressed assets from your origin. By enabling this policy, the CDN caching servers send your origin the HTTP Accept-Encoding header with the gzip code (Accept-Encoding: gzip).
//
// swagger:model custconfGzipOriginPull
type CustconfGzipOriginPull struct {

	// This enables support for compressed origin pull
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf gzip origin pull
func (m *CustconfGzipOriginPull) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf gzip origin pull based on context it is used
func (m *CustconfGzipOriginPull) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfGzipOriginPull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfGzipOriginPull) UnmarshalBinary(b []byte) error {
	var res CustconfGzipOriginPull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
