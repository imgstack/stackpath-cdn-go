// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnScanOriginRequest A request to scan a URL from the StackPath edge network
//
// swagger:model cdnScanOriginRequest
type CdnScanOriginRequest struct {

	// The URL to scan
	Domain string `json:"domain,omitempty"`
}

// Validate validates this cdn scan origin request
func (m *CdnScanOriginRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdn scan origin request based on context it is used
func (m *CdnScanOriginRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnScanOriginRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnScanOriginRequest) UnmarshalBinary(b []byte) error {
	var res CdnScanOriginRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
