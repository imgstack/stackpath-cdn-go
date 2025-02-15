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

// OriginPullPolicyExpirePolicyEnumWrapperValue origin pull policy expire policy enum wrapper value
//
// swagger:model OriginPullPolicyExpirePolicyEnumWrapperValue
type OriginPullPolicyExpirePolicyEnumWrapperValue string

const (

	// OriginPullPolicyExpirePolicyEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	OriginPullPolicyExpirePolicyEnumWrapperValueUNKNOWN OriginPullPolicyExpirePolicyEnumWrapperValue = "UNKNOWN"

	// OriginPullPolicyExpirePolicyEnumWrapperValueCACHECONTROL captures enum value "CACHE_CONTROL"
	OriginPullPolicyExpirePolicyEnumWrapperValueCACHECONTROL OriginPullPolicyExpirePolicyEnumWrapperValue = "CACHE_CONTROL"

	// OriginPullPolicyExpirePolicyEnumWrapperValueINGEST captures enum value "INGEST"
	OriginPullPolicyExpirePolicyEnumWrapperValueINGEST OriginPullPolicyExpirePolicyEnumWrapperValue = "INGEST"

	// OriginPullPolicyExpirePolicyEnumWrapperValueLASTMODIFY captures enum value "LAST_MODIFY"
	OriginPullPolicyExpirePolicyEnumWrapperValueLASTMODIFY OriginPullPolicyExpirePolicyEnumWrapperValue = "LAST_MODIFY"

	// OriginPullPolicyExpirePolicyEnumWrapperValueNEVEREXPIRE captures enum value "NEVER_EXPIRE"
	OriginPullPolicyExpirePolicyEnumWrapperValueNEVEREXPIRE OriginPullPolicyExpirePolicyEnumWrapperValue = "NEVER_EXPIRE"

	// OriginPullPolicyExpirePolicyEnumWrapperValueDONOTCACHE captures enum value "DO_NOT_CACHE"
	OriginPullPolicyExpirePolicyEnumWrapperValueDONOTCACHE OriginPullPolicyExpirePolicyEnumWrapperValue = "DO_NOT_CACHE"
)

// for schema
var originPullPolicyExpirePolicyEnumWrapperValueEnum []interface{}

func init() {
	var res []OriginPullPolicyExpirePolicyEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","CACHE_CONTROL","INGEST","LAST_MODIFY","NEVER_EXPIRE","DO_NOT_CACHE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		originPullPolicyExpirePolicyEnumWrapperValueEnum = append(originPullPolicyExpirePolicyEnumWrapperValueEnum, v)
	}
}

func (m OriginPullPolicyExpirePolicyEnumWrapperValue) validateOriginPullPolicyExpirePolicyEnumWrapperValueEnum(path, location string, value OriginPullPolicyExpirePolicyEnumWrapperValue) error {
	if err := validate.EnumCase(path, location, value, originPullPolicyExpirePolicyEnumWrapperValueEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this origin pull policy expire policy enum wrapper value
func (m OriginPullPolicyExpirePolicyEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOriginPullPolicyExpirePolicyEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this origin pull policy expire policy enum wrapper value based on context it is used
func (m OriginPullPolicyExpirePolicyEnumWrapperValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
