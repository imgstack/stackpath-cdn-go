// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnGetSiteDNSTargetsResponse The response from a request to retrieve a CDN site's DNS CNAME targets
//
// swagger:model cdnGetSiteDnsTargetsResponse
type CdnGetSiteDNSTargetsResponse struct {

	// The requested DNS CNAME targets
	//
	// A site's hostname should point to these CNAME targets in order for traffic to be sent through StackPath's edge nodes.
	Addresses []string `json:"addresses,omitempty"`
}

// Validate validates this cdn get site Dns targets response
func (m *CdnGetSiteDNSTargetsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdn get site Dns targets response based on context it is used
func (m *CdnGetSiteDNSTargetsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnGetSiteDNSTargetsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnGetSiteDNSTargetsResponse) UnmarshalBinary(b []byte) error {
	var res CdnGetSiteDNSTargetsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
