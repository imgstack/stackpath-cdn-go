// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfAuthHTTPBasic HTTP basic authentication policies allow you to control access to your content by requiring the end user to enter a valid username and password before gaining access to content.
//
// swagger:model custconfAuthHttpBasic
type CustconfAuthHTTPBasic struct {

	// This is a URL to a resource on the authentication server responsible for authentication of users.
	BindingPoint string `json:"bindingPoint,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// This is the authentication realm given back to the user on requests which do not contain the authentication credentials. Browsers typically display this value to the user when the login credentials are requested.
	Realm string `json:"realm,omitempty"`

	// This is the number of seconds the authentication session will be cached by the browsers. After this time, browsers will be asked to present the user credentials again for re-authentication.
	TTL int64 `json:"ttl,omitempty"`
}

// Validate validates this custconf auth Http basic
func (m *CustconfAuthHTTPBasic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf auth Http basic based on context it is used
func (m *CustconfAuthHTTPBasic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthHTTPBasic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthHTTPBasic) UnmarshalBinary(b []byte) error {
	var res CustconfAuthHTTPBasic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
