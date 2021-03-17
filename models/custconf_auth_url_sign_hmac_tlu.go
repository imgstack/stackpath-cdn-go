// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfAuthURLSignHmacTlu custconf auth Url sign hmac tlu
//
// swagger:model custconfAuthUrlSignHmacTlu
type CustconfAuthURLSignHmacTlu struct {

	// algorithm Id map
	AlgorithmIDMap map[string]*CustconfAuthURLSignHmacTluAlgorithmIDMapEnumWrapperValue `json:"algorithmIdMap,omitempty"`

	// algorithm Id parameter name
	AlgorithmIDParameterName string `json:"algorithmIdParameterName,omitempty"`

	// digest parameter name
	DigestParameterName string `json:"digestParameterName,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// expire parameter name
	ExpireParameterName string `json:"expireParameterName,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// key Id parameter name
	KeyIDParameterName string `json:"keyIdParameterName,omitempty"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// symmetric key Id map
	SymmetricKeyIDMap map[string]string `json:"symmetricKeyIdMap,omitempty"`
}

// Validate validates this custconf auth Url sign hmac tlu
func (m *CustconfAuthURLSignHmacTlu) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithmIDMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthURLSignHmacTlu) validateAlgorithmIDMap(formats strfmt.Registry) error {
	if swag.IsZero(m.AlgorithmIDMap) { // not required
		return nil
	}

	for k := range m.AlgorithmIDMap {

		if swag.IsZero(m.AlgorithmIDMap[k]) { // not required
			continue
		}
		if val, ok := m.AlgorithmIDMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custconf auth Url sign hmac tlu based on the context it is used
func (m *CustconfAuthURLSignHmacTlu) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlgorithmIDMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthURLSignHmacTlu) contextValidateAlgorithmIDMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.AlgorithmIDMap {

		if val, ok := m.AlgorithmIDMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLSignHmacTlu) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLSignHmacTlu) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLSignHmacTlu
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
