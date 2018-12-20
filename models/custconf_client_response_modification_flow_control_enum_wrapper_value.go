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

// CustconfClientResponseModificationFlowControlEnumWrapperValue custconf client response modification flow control enum wrapper value
// swagger:model custconfClientResponseModificationFlowControlEnumWrapperValue
type CustconfClientResponseModificationFlowControlEnumWrapperValue string

const (

	// CustconfClientResponseModificationFlowControlEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	CustconfClientResponseModificationFlowControlEnumWrapperValueUNKNOWN CustconfClientResponseModificationFlowControlEnumWrapperValue = "UNKNOWN"

	// CustconfClientResponseModificationFlowControlEnumWrapperValueNext captures enum value "next"
	CustconfClientResponseModificationFlowControlEnumWrapperValueNext CustconfClientResponseModificationFlowControlEnumWrapperValue = "next"

	// CustconfClientResponseModificationFlowControlEnumWrapperValueBreak captures enum value "break"
	CustconfClientResponseModificationFlowControlEnumWrapperValueBreak CustconfClientResponseModificationFlowControlEnumWrapperValue = "break"
)

// for schema
var custconfClientResponseModificationFlowControlEnumWrapperValueEnum []interface{}

func init() {
	var res []CustconfClientResponseModificationFlowControlEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","next","break"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		custconfClientResponseModificationFlowControlEnumWrapperValueEnum = append(custconfClientResponseModificationFlowControlEnumWrapperValueEnum, v)
	}
}

func (m CustconfClientResponseModificationFlowControlEnumWrapperValue) validateCustconfClientResponseModificationFlowControlEnumWrapperValueEnum(path, location string, value CustconfClientResponseModificationFlowControlEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, custconfClientResponseModificationFlowControlEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this custconf client response modification flow control enum wrapper value
func (m CustconfClientResponseModificationFlowControlEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustconfClientResponseModificationFlowControlEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
