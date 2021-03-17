// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/renderinc/stackpath-cdn-go/models"
)

// GetSitesReader is a Reader for the GetSites structure.
type GetSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSitesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSitesOK creates a GetSitesOK with default headers values
func NewGetSitesOK() *GetSitesOK {
	return &GetSitesOK{}
}

/* GetSitesOK describes a response with status code 200, with default header values.

GetSitesOK get sites o k
*/
type GetSitesOK struct {
	Payload *models.CdnGetSitesResponse
}

func (o *GetSitesOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites][%d] getSitesOK  %+v", 200, o.Payload)
}
func (o *GetSitesOK) GetPayload() *models.CdnGetSitesResponse {
	return o.Payload
}

func (o *GetSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetSitesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesUnauthorized creates a GetSitesUnauthorized with default headers values
func NewGetSitesUnauthorized() *GetSitesUnauthorized {
	return &GetSitesUnauthorized{}
}

/* GetSitesUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetSitesUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetSitesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites][%d] getSitesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSitesUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesInternalServerError creates a GetSitesInternalServerError with default headers values
func NewGetSitesInternalServerError() *GetSitesInternalServerError {
	return &GetSitesInternalServerError{}
}

/* GetSitesInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetSitesInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetSitesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites][%d] getSitesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSitesInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSitesDefault creates a GetSitesDefault with default headers values
func NewGetSitesDefault(code int) *GetSitesDefault {
	return &GetSitesDefault{
		_statusCode: code,
	}
}

/* GetSitesDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetSitesDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get sites default response
func (o *GetSitesDefault) Code() int {
	return o._statusCode
}

func (o *GetSitesDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites][%d] GetSites default  %+v", o._statusCode, o.Payload)
}
func (o *GetSitesDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSitesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
