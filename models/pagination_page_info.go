// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationPageInfo Information about a paginated response
//
// This is modeled after the GraphQL Relay spec to support both cursor based pagination and traditional offset based pagination.
//
// swagger:model paginationPageInfo
type PaginationPageInfo struct {

	// The cursor for the last item in the set of data returned
	EndCursor string `json:"endCursor,omitempty"`

	// Whether or not another page of data is available
	HasNextPage bool `json:"hasNextPage"`

	// Whether or not a previous page of data exists
	HasPreviousPage bool `json:"hasPreviousPage"`

	// The cursor for the first item in the set of data returned
	StartCursor string `json:"startCursor,omitempty"`

	// The total number of items in the dataset
	TotalCount string `json:"totalCount,omitempty"`
}

// Validate validates this pagination page info
func (m *PaginationPageInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagination page info based on context it is used
func (m *PaginationPageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationPageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationPageInfo) UnmarshalBinary(b []byte) error {
	var res PaginationPageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
