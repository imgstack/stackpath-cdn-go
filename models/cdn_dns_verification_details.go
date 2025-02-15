// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnDNSVerificationDetails DNS-based domain ownership verification details
//
// swagger:model cdnDnsVerificationDetails
type CdnDNSVerificationDetails struct {

	// A list of DNS records that will validate domain ownership
	DNSRecords []string `json:"dnsRecords,omitempty"`
}

// Validate validates this cdn Dns verification details
func (m *CdnDNSVerificationDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdn Dns verification details based on context it is used
func (m *CdnDNSVerificationDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnDNSVerificationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnDNSVerificationDetails) UnmarshalBinary(b []byte) error {
	var res CdnDNSVerificationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
