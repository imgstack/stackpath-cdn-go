// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfBandWidthLimit The pattern based bandwidth throttling policy allows you to limit the transfer rate of assets to end users based on a set of rules matching the request's HTTP User-Agent and/or the path. Each rule must be expressed in the following format: <User-Agent Pattern>:<path pattern 1, path pattern 2>. For example, the pattern:  "Mozilla*:*.mp3,*dir*.exe|*IE*:*.jpg,*.zip|*ios 6*:* " will match all MP3 files and EXE files containing the substring "dir" that are requested by a User-Agent containing the substring "Mozilla," all JPG and ZIP files requested by a User-Agent containing the substring "IE," and all requests made by User-Agents containing the substring "ios 6."
//
// swagger:model custconfBandWidthLimit
type CustconfBandWidthLimit struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a '|' character. These are pattern match rules to use for applying rate limiting on requests.
	Rule string `json:"rule,omitempty"`

	// These are the initial bytes (ri) and the sustained rate (rs) query string parameters to use for this rule. Example: ri=100,rs=1000
	Values string `json:"values,omitempty"`
}

// Validate validates this custconf band width limit
func (m *CustconfBandWidthLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf band width limit based on context it is used
func (m *CustconfBandWidthLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfBandWidthLimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfBandWidthLimit) UnmarshalBinary(b []byte) error {
	var res CustconfBandWidthLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
