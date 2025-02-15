// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PurgeContentRequestItem An individual item to purge from the CDN
//
// swagger:model PurgeContentRequestItem
type PurgeContentRequestItem struct {

	// A list of HTTP response headers from the origin that should exist for its content to be purged
	Headers []string `json:"headers,omitempty"`

	// Whether or not to mark the asset as expired and re-validate instead of deleting
	InvalidateOnly bool `json:"invalidateOnly"`

	// Whether or not to purge dynamic versions of assets
	//
	// This is ignored if recursive is true.
	PurgeAllDynamic bool `json:"purgeAllDynamic"`

	// purge selector
	PurgeSelector []*PurgeContentRequestPurgeSelector `json:"purgeSelector,omitempty"`

	// Whether or not to recursively delete content from the CDN
	Recursive bool `json:"recursive"`

	// The URL at which to delete content
	URL string `json:"url,omitempty"`
}

// Validate validates this purge content request item
func (m *PurgeContentRequestItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePurgeSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurgeContentRequestItem) validatePurgeSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.PurgeSelector) { // not required
		return nil
	}

	for i := 0; i < len(m.PurgeSelector); i++ {
		if swag.IsZero(m.PurgeSelector[i]) { // not required
			continue
		}

		if m.PurgeSelector[i] != nil {
			if err := m.PurgeSelector[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("purgeSelector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this purge content request item based on the context it is used
func (m *PurgeContentRequestItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePurgeSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PurgeContentRequestItem) contextValidatePurgeSelector(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PurgeSelector); i++ {

		if m.PurgeSelector[i] != nil {
			if err := m.PurgeSelector[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("purgeSelector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PurgeContentRequestItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurgeContentRequestItem) UnmarshalBinary(b []byte) error {
	var res PurgeContentRequestItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
