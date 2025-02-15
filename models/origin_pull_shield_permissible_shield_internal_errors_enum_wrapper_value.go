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

// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue origin pull shield permissible shield internal errors enum wrapper value
//
// swagger:model OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue
type OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue string

const (

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueUNKNOWN OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "UNKNOWN"

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueNONE captures enum value "NONE"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueNONE OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "NONE"

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueCONNECTIONONLY captures enum value "CONNECTION_ONLY"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueCONNECTIONONLY OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "CONNECTION_ONLY"

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueWRITEONLY captures enum value "WRITE_ONLY"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueWRITEONLY OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "WRITE_ONLY"

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueWRITEREAD captures enum value "WRITE_READ"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueWRITEREAD OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "WRITE_READ"

	// OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueALL captures enum value "ALL"
	OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueALL OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue = "ALL"
)

// for schema
var originPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum []interface{}

func init() {
	var res []OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","NONE","CONNECTION_ONLY","WRITE_ONLY","WRITE_READ","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		originPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum = append(originPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum, v)
	}
}

func (m OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue) validateOriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum(path, location string, value OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, originPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this origin pull shield permissible shield internal errors enum wrapper value
func (m OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this origin pull shield permissible shield internal errors enum wrapper value based on context it is used
func (m OriginPullShieldPermissibleShieldInternalErrorsEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
