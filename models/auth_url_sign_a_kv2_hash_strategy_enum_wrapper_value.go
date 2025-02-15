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

// AuthURLSignAKv2HashStrategyEnumWrapperValue auth Url sign a kv2 hash strategy enum wrapper value
//
// swagger:model AuthUrlSignAKv2HashStrategyEnumWrapperValue
type AuthURLSignAKv2HashStrategyEnumWrapperValue string

const (

	// AuthURLSignAKv2HashStrategyEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	AuthURLSignAKv2HashStrategyEnumWrapperValueUNKNOWN AuthURLSignAKv2HashStrategyEnumWrapperValue = "UNKNOWN"

	// AuthURLSignAKv2HashStrategyEnumWrapperValueSha1 captures enum value "sha1"
	AuthURLSignAKv2HashStrategyEnumWrapperValueSha1 AuthURLSignAKv2HashStrategyEnumWrapperValue = "sha1"

	// AuthURLSignAKv2HashStrategyEnumWrapperValueSha256 captures enum value "sha256"
	AuthURLSignAKv2HashStrategyEnumWrapperValueSha256 AuthURLSignAKv2HashStrategyEnumWrapperValue = "sha256"

	// AuthURLSignAKv2HashStrategyEnumWrapperValueMd5 captures enum value "md5"
	AuthURLSignAKv2HashStrategyEnumWrapperValueMd5 AuthURLSignAKv2HashStrategyEnumWrapperValue = "md5"
)

// for schema
var authUrlSignAKv2HashStrategyEnumWrapperValueEnum []interface{}

func init() {
	var res []AuthURLSignAKv2HashStrategyEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","sha1","sha256","md5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authUrlSignAKv2HashStrategyEnumWrapperValueEnum = append(authUrlSignAKv2HashStrategyEnumWrapperValueEnum, v)
	}
}

func (m AuthURLSignAKv2HashStrategyEnumWrapperValue) validateAuthURLSignAKv2HashStrategyEnumWrapperValueEnum(path, location string, value AuthURLSignAKv2HashStrategyEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, authUrlSignAKv2HashStrategyEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this auth Url sign a kv2 hash strategy enum wrapper value
func (m AuthURLSignAKv2HashStrategyEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthURLSignAKv2HashStrategyEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this auth Url sign a kv2 hash strategy enum wrapper value based on context it is used
func (m AuthURLSignAKv2HashStrategyEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
