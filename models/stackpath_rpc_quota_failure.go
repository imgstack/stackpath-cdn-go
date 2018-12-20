// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackpathRPCQuotaFailure stackpath rpc quota failure
// swagger:model stackpath.rpc.QuotaFailure
type StackpathRPCQuotaFailure struct {

	// violations
	Violations []*StackpathRPCQuotaFailureViolation `json:"violations"`
}

// AtType gets the at type of this subtype
func (m *StackpathRPCQuotaFailure) AtType() string {
	return "stackpath.rpc.QuotaFailure"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCQuotaFailure) SetAtType(val string) {

}

// Violations gets the violations of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCQuotaFailure) UnmarshalJSON(raw []byte) error {
	var data struct {

		// violations
		Violations []*StackpathRPCQuotaFailureViolation `json:"violations"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StackpathRPCQuotaFailure

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.Violations = data.Violations

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCQuotaFailure) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// violations
		Violations []*StackpathRPCQuotaFailureViolation `json:"violations"`
	}{

		Violations: m.Violations,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type"`
	}{

		AtType: m.AtType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this stackpath rpc quota failure
func (m *StackpathRPCQuotaFailure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateViolations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackpathRPCQuotaFailure) validateViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.Violations) { // not required
		return nil
	}

	for i := 0; i < len(m.Violations); i++ {
		if swag.IsZero(m.Violations[i]) { // not required
			continue
		}

		if m.Violations[i] != nil {
			if err := m.Violations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("violations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCQuotaFailure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCQuotaFailure) UnmarshalBinary(b []byte) error {
	var res StackpathRPCQuotaFailure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
