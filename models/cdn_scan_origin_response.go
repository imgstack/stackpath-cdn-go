// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CdnScanOriginResponse cdn scan origin response
// swagger:model cdnScanOriginResponse
type CdnScanOriginResponse struct {

	// domain in use
	DomainInUse bool `json:"domainInUse,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// ssl details
	SslDetails *ScanOriginResponseOriginScanSSLDetails `json:"sslDetails,omitempty"`
}

// Validate validates this cdn scan origin response
func (m *CdnScanOriginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSslDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnScanOriginResponse) validateSslDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.SslDetails) { // not required
		return nil
	}

	if m.SslDetails != nil {
		if err := m.SslDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sslDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnScanOriginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnScanOriginResponse) UnmarshalBinary(b []byte) error {
	var res CdnScanOriginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
