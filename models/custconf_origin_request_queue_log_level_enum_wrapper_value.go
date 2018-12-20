// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CustconfOriginRequestQueueLogLevelEnumWrapperValue custconf origin request queue log level enum wrapper value
// swagger:model custconfOriginRequestQueueLogLevelEnumWrapperValue
type CustconfOriginRequestQueueLogLevelEnumWrapperValue string

const (

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueUNKNOWN CustconfOriginRequestQueueLogLevelEnumWrapperValue = "UNKNOWN"

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueDebug captures enum value "debug"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueDebug CustconfOriginRequestQueueLogLevelEnumWrapperValue = "debug"

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueInfo captures enum value "info"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueInfo CustconfOriginRequestQueueLogLevelEnumWrapperValue = "info"

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueWarning captures enum value "warning"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueWarning CustconfOriginRequestQueueLogLevelEnumWrapperValue = "warning"

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueError captures enum value "error"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueError CustconfOriginRequestQueueLogLevelEnumWrapperValue = "error"

	// CustconfOriginRequestQueueLogLevelEnumWrapperValueCrit captures enum value "crit"
	CustconfOriginRequestQueueLogLevelEnumWrapperValueCrit CustconfOriginRequestQueueLogLevelEnumWrapperValue = "crit"
)

// for schema
var custconfOriginRequestQueueLogLevelEnumWrapperValueEnum []interface{}

func init() {
	var res []CustconfOriginRequestQueueLogLevelEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","debug","info","warning","error","crit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		custconfOriginRequestQueueLogLevelEnumWrapperValueEnum = append(custconfOriginRequestQueueLogLevelEnumWrapperValueEnum, v)
	}
}

func (m CustconfOriginRequestQueueLogLevelEnumWrapperValue) validateCustconfOriginRequestQueueLogLevelEnumWrapperValueEnum(path, location string, value CustconfOriginRequestQueueLogLevelEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, custconfOriginRequestQueueLogLevelEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this custconf origin request queue log level enum wrapper value
func (m CustconfOriginRequestQueueLogLevelEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustconfOriginRequestQueueLogLevelEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
