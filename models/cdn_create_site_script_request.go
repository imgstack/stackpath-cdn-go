// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CdnCreateSiteScriptRequest A request to create a new EdgeEngine script
//
// swagger:model cdnCreateSiteScriptRequest
type CdnCreateSiteScriptRequest struct {

	// The contents of the new EdgeEngine script
	// Format: byte
	Code strfmt.Base64 `json:"code,omitempty"`

	// The name of the new EdgeEngine script
	Name string `json:"name,omitempty"`

	// The HTTP request paths that are handled by the new EdgeEngine script
	Paths []string `json:"paths,omitempty"`
}

// Validate validates this cdn create site script request
func (m *CdnCreateSiteScriptRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cdn create site script request based on context it is used
func (m *CdnCreateSiteScriptRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CdnCreateSiteScriptRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnCreateSiteScriptRequest) UnmarshalBinary(b []byte) error {
	var res CdnCreateSiteScriptRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
