// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CustconfAuthURLSign URL Signing policies allow you to restrict access to your content by configuring a "shared secret" with StackPath. This "shared secret" is used to apply an MD5 hashing algorithm on the URL to validate the signature supplied on the request. Since the "shared secret" is only known by the publisher and StackPath, URL signatures cannot be generated by unauthorized users.
//
// swagger:model custconfAuthUrlSign
type CustconfAuthURLSign struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is the name of the query string parameter that contains the Epoch time after which the URL is considered invalid.
	ExpiresField string `json:"expiresField,omitempty"`

	// String of values delimited by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking
	ID string `json:"id,omitempty"`

	// Select this option if you want StackPath to exclude query string parameters specified after the passphrase in the validation process.
	IgnoreFieldsAfterToken bool `json:"ignoreFieldsAfterToken"`

	// This is a query string parameter that contains an IPv4 address to which the published URL will be restricted.
	IPAddressField string `json:"ipAddressField,omitempty"`

	// String of values delimited by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// The shared secret used during the signing process. This value should only be known by StackPath and systems authorized to sign your content.
	PassPhrase string `json:"passPhrase,omitempty"`

	// This is the name of the query string parameter that contains the value of the secret. This query string parameter is only used during the generation and validation of a URL signature and should not be present in the published URL.
	PassPhraseField string `json:"passPhraseField,omitempty"`

	// String of values delimited by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// This is the name of the query string parameter that will be used by the publisher to specify the signature for the URL.
	TokenField string `json:"tokenField,omitempty"`

	// This is a query string parameter used to limit the number of characters in the path that should be considered when validating the URL signature. This feature allows you to reuse the same signature on all assets stored in the same cache path. Setting this value to 0 will strip off the filename in the URL (leaving the trailing slash) when calculating the checksum.
	URILengthField string `json:"uriLengthField,omitempty"`

	// This is a query string parameter used to restrict the published URL to a specific user agent. Publishers can use this feature during the signing process to ensure that only a specific user agent can access the published content. You do not need to specify the user agent on the published URL. StackPath will use the HTTP User-Agent header value during signature validation.
	UserAgentField string `json:"userAgentField,omitempty"`
}

// Validate validates this custconf auth Url sign
func (m *CustconfAuthURLSign) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this custconf auth Url sign based on context it is used
func (m *CustconfAuthURLSign) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLSign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLSign) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLSign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
