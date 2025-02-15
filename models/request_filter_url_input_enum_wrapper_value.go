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

// RequestFilterURLInputEnumWrapperValue request filter Url input enum wrapper value
//
// swagger:model RequestFilterUrlInputEnumWrapperValue
type RequestFilterURLInputEnumWrapperValue string

const (

	// RequestFilterURLInputEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	RequestFilterURLInputEnumWrapperValueUNKNOWN RequestFilterURLInputEnumWrapperValue = "UNKNOWN"

	// RequestFilterURLInputEnumWrapperValueFULLURL captures enum value "FULL_URL"
	RequestFilterURLInputEnumWrapperValueFULLURL RequestFilterURLInputEnumWrapperValue = "FULL_URL"

	// RequestFilterURLInputEnumWrapperValueFULLFILEPATH captures enum value "FULL_FILE_PATH"
	RequestFilterURLInputEnumWrapperValueFULLFILEPATH RequestFilterURLInputEnumWrapperValue = "FULL_FILE_PATH"

	// RequestFilterURLInputEnumWrapperValueFULLFILEPATHLEGACY captures enum value "FULL_FILE_PATH_LEGACY"
	RequestFilterURLInputEnumWrapperValueFULLFILEPATHLEGACY RequestFilterURLInputEnumWrapperValue = "FULL_FILE_PATH_LEGACY"
)

// for schema
var requestFilterUrlInputEnumWrapperValueEnum []interface{}

func init() {
	var res []RequestFilterURLInputEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","FULL_URL","FULL_FILE_PATH","FULL_FILE_PATH_LEGACY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestFilterUrlInputEnumWrapperValueEnum = append(requestFilterUrlInputEnumWrapperValueEnum, v)
	}
}

func (m RequestFilterURLInputEnumWrapperValue) validateRequestFilterURLInputEnumWrapperValueEnum(path, location string, value RequestFilterURLInputEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, requestFilterUrlInputEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this request filter Url input enum wrapper value
func (m RequestFilterURLInputEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRequestFilterURLInputEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this request filter Url input enum wrapper value based on context it is used
func (m RequestFilterURLInputEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
