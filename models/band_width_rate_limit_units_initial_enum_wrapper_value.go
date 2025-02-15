// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BandWidthRateLimitUnitsInitialEnumWrapperValue band width rate limit units initial enum wrapper value
//
// swagger:model BandWidthRateLimitUnitsInitialEnumWrapperValue
type BandWidthRateLimitUnitsInitialEnumWrapperValue string

const (

	// BandWidthRateLimitUnitsInitialEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	BandWidthRateLimitUnitsInitialEnumWrapperValueUNKNOWN BandWidthRateLimitUnitsInitialEnumWrapperValue = "UNKNOWN"

	// BandWidthRateLimitUnitsInitialEnumWrapperValueByte captures enum value "byte"
	BandWidthRateLimitUnitsInitialEnumWrapperValueByte BandWidthRateLimitUnitsInitialEnumWrapperValue = "byte"

	// BandWidthRateLimitUnitsInitialEnumWrapperValueKilobyte captures enum value "kilobyte"
	BandWidthRateLimitUnitsInitialEnumWrapperValueKilobyte BandWidthRateLimitUnitsInitialEnumWrapperValue = "kilobyte"
)

// for schema
var bandWidthRateLimitUnitsInitialEnumWrapperValueEnum []interface{}

func init() {
	var res []BandWidthRateLimitUnitsInitialEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","byte","kilobyte"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bandWidthRateLimitUnitsInitialEnumWrapperValueEnum = append(bandWidthRateLimitUnitsInitialEnumWrapperValueEnum, v)
	}
}

func (m BandWidthRateLimitUnitsInitialEnumWrapperValue) validateBandWidthRateLimitUnitsInitialEnumWrapperValueEnum(path, location string, value BandWidthRateLimitUnitsInitialEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, bandWidthRateLimitUnitsInitialEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this band width rate limit units initial enum wrapper value
func (m BandWidthRateLimitUnitsInitialEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBandWidthRateLimitUnitsInitialEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this band width rate limit units initial enum wrapper value based on context it is used
func (m BandWidthRateLimitUnitsInitialEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
