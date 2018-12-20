// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfOriginPullShieldOverride Suspending origin shield allows a Highwinds administrator to disable origin
// shielding for a customer without overriding the customer's configured origin
// shield value.
// swagger:model custconfOriginPullShieldOverride
type CustconfOriginPullShieldOverride struct {

	// Suspends origin shielding functionality.
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf origin pull shield override
func (m *CustconfOriginPullShieldOverride) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfOriginPullShieldOverride) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfOriginPullShieldOverride) UnmarshalBinary(b []byte) error {
	var res CustconfOriginPullShieldOverride
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
