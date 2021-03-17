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

// GetScopeRuleReader is a Reader for the GetScopeRule structure.
type GetScopeRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScopeRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScopeRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScopeRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScopeRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetScopeRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScopeRuleOK creates a GetScopeRuleOK with default headers values
func NewGetScopeRuleOK() *GetScopeRuleOK {
	return &GetScopeRuleOK{}
}

/* GetScopeRuleOK describes a response with status code 200, with default header values.

GetScopeRuleOK get scope rule o k
*/
type GetScopeRuleOK struct {
	Payload *models.CdnGetScopeRuleResponse
}

func (o *GetScopeRuleOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}][%d] getScopeRuleOK  %+v", 200, o.Payload)
}
func (o *GetScopeRuleOK) GetPayload() *models.CdnGetScopeRuleResponse {
	return o.Payload
}

func (o *GetScopeRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetScopeRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleUnauthorized creates a GetScopeRuleUnauthorized with default headers values
func NewGetScopeRuleUnauthorized() *GetScopeRuleUnauthorized {
	return &GetScopeRuleUnauthorized{}
}

/* GetScopeRuleUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetScopeRuleUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetScopeRuleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}][%d] getScopeRuleUnauthorized  %+v", 401, o.Payload)
}
func (o *GetScopeRuleUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleInternalServerError creates a GetScopeRuleInternalServerError with default headers values
func NewGetScopeRuleInternalServerError() *GetScopeRuleInternalServerError {
	return &GetScopeRuleInternalServerError{}
}

/* GetScopeRuleInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetScopeRuleInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetScopeRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}][%d] getScopeRuleInternalServerError  %+v", 500, o.Payload)
}
func (o *GetScopeRuleInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleDefault creates a GetScopeRuleDefault with default headers values
func NewGetScopeRuleDefault(code int) *GetScopeRuleDefault {
	return &GetScopeRuleDefault{
		_statusCode: code,
	}
}

/* GetScopeRuleDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetScopeRuleDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get scope rule default response
func (o *GetScopeRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetScopeRuleDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}][%d] GetScopeRule default  %+v", o._statusCode, o.Payload)
}
func (o *GetScopeRuleDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
