// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfRedirectExceptions The redirect response codes policy allows you to specify the HTTP redirect status code the CDN caching server should use when the CDN issues a redirect. Using this policy, you can assign different redirect codes to user agents requesting content.
//
// swagger:model custconfRedirectExceptions
type CustconfRedirectExceptions struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character. This is a comma separated list of user agents and redirect code pairs. The user agent and redirect code values are separated by a colon (:), and you may use wildcards in the user agent field. For example, to map assign a 307 status code to all Chrome browsers, you would specify: *chrome*:307.
	RedirectAgentCode string `json:"redirectAgentCode,omitempty"`
}

// Validate validates this custconf redirect exceptions
func (m *CustconfRedirectExceptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf redirect exceptions based on context it is used
func (m *CustconfRedirectExceptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfRedirectExceptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfRedirectExceptions) UnmarshalBinary(b []byte) error {
	var res CustconfRedirectExceptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
