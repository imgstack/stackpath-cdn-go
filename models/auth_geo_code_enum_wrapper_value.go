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

// AuthGeoCodeEnumWrapperValue auth geo code enum wrapper value
//
// swagger:model AuthGeoCodeEnumWrapperValue
type AuthGeoCodeEnumWrapperValue string

const (

	// AuthGeoCodeEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	AuthGeoCodeEnumWrapperValueUNKNOWN AuthGeoCodeEnumWrapperValue = "UNKNOWN"

	// AuthGeoCodeEnumWrapperValueCountryCode captures enum value "countryCode"
	AuthGeoCodeEnumWrapperValueCountryCode AuthGeoCodeEnumWrapperValue = "countryCode"

	// AuthGeoCodeEnumWrapperValueRegion captures enum value "region"
	AuthGeoCodeEnumWrapperValueRegion AuthGeoCodeEnumWrapperValue = "region"

	// AuthGeoCodeEnumWrapperValueSubdivisionCodes captures enum value "subdivisionCodes"
	AuthGeoCodeEnumWrapperValueSubdivisionCodes AuthGeoCodeEnumWrapperValue = "subdivisionCodes"

	// AuthGeoCodeEnumWrapperValueCity captures enum value "city"
	AuthGeoCodeEnumWrapperValueCity AuthGeoCodeEnumWrapperValue = "city"

	// AuthGeoCodeEnumWrapperValuePostalCode captures enum value "postalCode"
	AuthGeoCodeEnumWrapperValuePostalCode AuthGeoCodeEnumWrapperValue = "postalCode"

	// AuthGeoCodeEnumWrapperValueContinentCode captures enum value "continentCode"
	AuthGeoCodeEnumWrapperValueContinentCode AuthGeoCodeEnumWrapperValue = "continentCode"

	// AuthGeoCodeEnumWrapperValueTimeZone captures enum value "timeZone"
	AuthGeoCodeEnumWrapperValueTimeZone AuthGeoCodeEnumWrapperValue = "timeZone"

	// AuthGeoCodeEnumWrapperValueDmaCode captures enum value "dmaCode"
	AuthGeoCodeEnumWrapperValueDmaCode AuthGeoCodeEnumWrapperValue = "dmaCode"

	// AuthGeoCodeEnumWrapperValueAreaCode captures enum value "areaCode"
	AuthGeoCodeEnumWrapperValueAreaCode AuthGeoCodeEnumWrapperValue = "areaCode"
)

// for schema
var authGeoCodeEnumWrapperValueEnum []interface{}

func init() {
	var res []AuthGeoCodeEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","countryCode","region","subdivisionCodes","city","postalCode","continentCode","timeZone","dmaCode","areaCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authGeoCodeEnumWrapperValueEnum = append(authGeoCodeEnumWrapperValueEnum, v)
	}
}

func (m AuthGeoCodeEnumWrapperValue) validateAuthGeoCodeEnumWrapperValueEnum(path, location string, value AuthGeoCodeEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, authGeoCodeEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this auth geo code enum wrapper value
func (m AuthGeoCodeEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthGeoCodeEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this auth geo code enum wrapper value based on context it is used
func (m AuthGeoCodeEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
