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

// CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue custconf auth Url sign hmac tlu algorithm Id map enum wrapper value
// swagger:model custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValue
type CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue string

const (

	// CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueUNKNOWN CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue = "UNKNOWN"

	// CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueHmacsha1 captures enum value "hmacsha1"
	CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueHmacsha1 CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue = "hmacsha1"

	// CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueHmacsha256 captures enum value "hmacsha256"
	CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueHmacsha256 CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue = "hmacsha256"
)

// for schema
var custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValueEnum []interface{}

func init() {
	var res []CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","hmacsha1","hmacsha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValueEnum = append(custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValueEnum, v)
	}
}

func (m CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue) validateCustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueEnum(path, location string, value CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, custconfAuthUrlSignHmacTluAlgorithmIdMapEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this custconf auth Url sign hmac tlu algorithm Id map enum wrapper value
func (m CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
