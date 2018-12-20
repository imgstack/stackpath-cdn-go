// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfResponseFilter custconf response filter
// swagger:model custconfResponseFilter
type CustconfResponseFilter struct {

	// filter name
	FilterName string `json:"filterName,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	PopFilter string `json:"popFilter,omitempty"`

	// String of values deliminated by a ',' character.
	RegionFilter string `json:"regionFilter,omitempty"`

	// String of values deliminated by a ',' character. This is a pattern match
	// expression for each status code this policy applies to.  For example, 2*,
	// 3* applies this policy to all 200 and 300 level HTTP responses from your
	// origin.
	StatusCodeFilter string `json:"statusCodeFilter,omitempty"`
}

// Validate validates this custconf response filter
func (m *CustconfResponseFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfResponseFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfResponseFilter) UnmarshalBinary(b []byte) error {
	var res CustconfResponseFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
