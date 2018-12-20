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

// CustconfOriginResponseModificationFlowControlEnumWrapperValue custconf origin response modification flow control enum wrapper value
// swagger:model custconfOriginResponseModificationFlowControlEnumWrapperValue
type CustconfOriginResponseModificationFlowControlEnumWrapperValue string

const (

	// CustconfOriginResponseModificationFlowControlEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	CustconfOriginResponseModificationFlowControlEnumWrapperValueUNKNOWN CustconfOriginResponseModificationFlowControlEnumWrapperValue = "UNKNOWN"

	// CustconfOriginResponseModificationFlowControlEnumWrapperValueNext captures enum value "next"
	CustconfOriginResponseModificationFlowControlEnumWrapperValueNext CustconfOriginResponseModificationFlowControlEnumWrapperValue = "next"

	// CustconfOriginResponseModificationFlowControlEnumWrapperValueBreak captures enum value "break"
	CustconfOriginResponseModificationFlowControlEnumWrapperValueBreak CustconfOriginResponseModificationFlowControlEnumWrapperValue = "break"
)

// for schema
var custconfOriginResponseModificationFlowControlEnumWrapperValueEnum []interface{}

func init() {
	var res []CustconfOriginResponseModificationFlowControlEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","next","break"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		custconfOriginResponseModificationFlowControlEnumWrapperValueEnum = append(custconfOriginResponseModificationFlowControlEnumWrapperValueEnum, v)
	}
}

func (m CustconfOriginResponseModificationFlowControlEnumWrapperValue) validateCustconfOriginResponseModificationFlowControlEnumWrapperValueEnum(path, location string, value CustconfOriginResponseModificationFlowControlEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, custconfOriginResponseModificationFlowControlEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this custconf origin response modification flow control enum wrapper value
func (m CustconfOriginResponseModificationFlowControlEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustconfOriginResponseModificationFlowControlEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
