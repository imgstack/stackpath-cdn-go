// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfAwsSignedS3PostV4 custconf aws signed s3 post v4
// swagger:model custconfAwsSignedS3PostV4
type CustconfAwsSignedS3PostV4 struct {

	// access key Id
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// authentication type
	AuthenticationType CustconfAwsSignedS3PostV4AuthenticationTypeEnumWrapperValue `json:"authenticationType,omitempty"`

	// aws region
	AwsRegion string `json:"awsRegion,omitempty"`

	// aws service
	AwsService string `json:"awsService,omitempty"`

	// bucket identifier
	BucketIdentifier string `json:"bucketIdentifier,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// expire time seconds
	ExpireTimeSeconds int64 `json:"expireTimeSeconds,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// secret access key
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// String of values delimited by a ',' character.
	SignedHeaders string `json:"signedHeaders,omitempty"`
}

// Validate validates this custconf aws signed s3 post v4
func (m *CustconfAwsSignedS3PostV4) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAwsSignedS3PostV4) validateAuthenticationType(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	if err := m.AuthenticationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authenticationType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAwsSignedS3PostV4) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAwsSignedS3PostV4) UnmarshalBinary(b []byte) error {
	var res CustconfAwsSignedS3PostV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
