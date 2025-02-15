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

// CustconfClientResponseQueueLogLevelEnumWrapperValue custconf client response queue log level enum wrapper value
//
// swagger:model custconfClientResponseQueueLogLevelEnumWrapperValue
type CustconfClientResponseQueueLogLevelEnumWrapperValue string

const (

	// CustconfClientResponseQueueLogLevelEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	CustconfClientResponseQueueLogLevelEnumWrapperValueUNKNOWN CustconfClientResponseQueueLogLevelEnumWrapperValue = "UNKNOWN"

	// CustconfClientResponseQueueLogLevelEnumWrapperValueDebug captures enum value "debug"
	CustconfClientResponseQueueLogLevelEnumWrapperValueDebug CustconfClientResponseQueueLogLevelEnumWrapperValue = "debug"

	// CustconfClientResponseQueueLogLevelEnumWrapperValueInfo captures enum value "info"
	CustconfClientResponseQueueLogLevelEnumWrapperValueInfo CustconfClientResponseQueueLogLevelEnumWrapperValue = "info"

	// CustconfClientResponseQueueLogLevelEnumWrapperValueWarning captures enum value "warning"
	CustconfClientResponseQueueLogLevelEnumWrapperValueWarning CustconfClientResponseQueueLogLevelEnumWrapperValue = "warning"

	// CustconfClientResponseQueueLogLevelEnumWrapperValueError captures enum value "error"
	CustconfClientResponseQueueLogLevelEnumWrapperValueError CustconfClientResponseQueueLogLevelEnumWrapperValue = "error"

	// CustconfClientResponseQueueLogLevelEnumWrapperValueCrit captures enum value "crit"
	CustconfClientResponseQueueLogLevelEnumWrapperValueCrit CustconfClientResponseQueueLogLevelEnumWrapperValue = "crit"
)

// for schema
var custconfClientResponseQueueLogLevelEnumWrapperValueEnum []interface{}

func init() {
	var res []CustconfClientResponseQueueLogLevelEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","debug","info","warning","error","crit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		custconfClientResponseQueueLogLevelEnumWrapperValueEnum = append(custconfClientResponseQueueLogLevelEnumWrapperValueEnum, v)
	}
}

func (m CustconfClientResponseQueueLogLevelEnumWrapperValue) validateCustconfClientResponseQueueLogLevelEnumWrapperValueEnum(path, location string, value CustconfClientResponseQueueLogLevelEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, custconfClientResponseQueueLogLevelEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this custconf client response queue log level enum wrapper value
func (m CustconfClientResponseQueueLogLevelEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustconfClientResponseQueueLogLevelEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this custconf client response queue log level enum wrapper value based on context it is used
func (m CustconfClientResponseQueueLogLevelEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
