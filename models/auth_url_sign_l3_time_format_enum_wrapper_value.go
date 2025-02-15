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

// AuthURLSignL3TimeFormatEnumWrapperValue auth Url sign l3 time format enum wrapper value
//
// swagger:model AuthUrlSignL3TimeFormatEnumWrapperValue
type AuthURLSignL3TimeFormatEnumWrapperValue string

const (

	// AuthURLSignL3TimeFormatEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	AuthURLSignL3TimeFormatEnumWrapperValueUNKNOWN AuthURLSignL3TimeFormatEnumWrapperValue = "UNKNOWN"

	// AuthURLSignL3TimeFormatEnumWrapperValueEpoch captures enum value "epoch"
	AuthURLSignL3TimeFormatEnumWrapperValueEpoch AuthURLSignL3TimeFormatEnumWrapperValue = "epoch"

	// AuthURLSignL3TimeFormatEnumWrapperValueDatetime captures enum value "datetime"
	AuthURLSignL3TimeFormatEnumWrapperValueDatetime AuthURLSignL3TimeFormatEnumWrapperValue = "datetime"
)

// for schema
var authUrlSignL3TimeFormatEnumWrapperValueEnum []interface{}

func init() {
	var res []AuthURLSignL3TimeFormatEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","epoch","datetime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authUrlSignL3TimeFormatEnumWrapperValueEnum = append(authUrlSignL3TimeFormatEnumWrapperValueEnum, v)
	}
}

func (m AuthURLSignL3TimeFormatEnumWrapperValue) validateAuthURLSignL3TimeFormatEnumWrapperValueEnum(path, location string, value AuthURLSignL3TimeFormatEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, authUrlSignL3TimeFormatEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this auth Url sign l3 time format enum wrapper value
func (m AuthURLSignL3TimeFormatEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthURLSignL3TimeFormatEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this auth Url sign l3 time format enum wrapper value based on context it is used
func (m AuthURLSignL3TimeFormatEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
