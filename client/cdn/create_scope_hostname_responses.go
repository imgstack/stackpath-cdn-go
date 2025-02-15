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

// CreateScopeHostnameReader is a Reader for the CreateScopeHostname structure.
type CreateScopeHostnameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScopeHostnameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateScopeHostnameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateScopeHostnameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateScopeHostnameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateScopeHostnameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateScopeHostnameOK creates a CreateScopeHostnameOK with default headers values
func NewCreateScopeHostnameOK() *CreateScopeHostnameOK {
	return &CreateScopeHostnameOK{}
}

/* CreateScopeHostnameOK describes a response with status code 200, with default header values.

CreateScopeHostnameOK create scope hostname o k
*/
type CreateScopeHostnameOK struct {
	Payload *models.CdnCreateScopeHostnameResponse
}

func (o *CreateScopeHostnameOK) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames][%d] createScopeHostnameOK  %+v", 200, o.Payload)
}
func (o *CreateScopeHostnameOK) GetPayload() *models.CdnCreateScopeHostnameResponse {
	return o.Payload
}

func (o *CreateScopeHostnameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnCreateScopeHostnameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScopeHostnameUnauthorized creates a CreateScopeHostnameUnauthorized with default headers values
func NewCreateScopeHostnameUnauthorized() *CreateScopeHostnameUnauthorized {
	return &CreateScopeHostnameUnauthorized{}
}

/* CreateScopeHostnameUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type CreateScopeHostnameUnauthorized struct {
	Payload *models.APIStatus
}

func (o *CreateScopeHostnameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames][%d] createScopeHostnameUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateScopeHostnameUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *CreateScopeHostnameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScopeHostnameInternalServerError creates a CreateScopeHostnameInternalServerError with default headers values
func NewCreateScopeHostnameInternalServerError() *CreateScopeHostnameInternalServerError {
	return &CreateScopeHostnameInternalServerError{}
}

/* CreateScopeHostnameInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type CreateScopeHostnameInternalServerError struct {
	Payload *models.APIStatus
}

func (o *CreateScopeHostnameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames][%d] createScopeHostnameInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateScopeHostnameInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *CreateScopeHostnameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScopeHostnameDefault creates a CreateScopeHostnameDefault with default headers values
func NewCreateScopeHostnameDefault(code int) *CreateScopeHostnameDefault {
	return &CreateScopeHostnameDefault{
		_statusCode: code,
	}
}

/* CreateScopeHostnameDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type CreateScopeHostnameDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the create scope hostname default response
func (o *CreateScopeHostnameDefault) Code() int {
	return o._statusCode
}

func (o *CreateScopeHostnameDefault) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/hostnames][%d] CreateScopeHostname default  %+v", o._statusCode, o.Payload)
}
func (o *CreateScopeHostnameDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *CreateScopeHostnameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
