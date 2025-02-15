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

// GetScopeRuleConfigurationReader is a Reader for the GetScopeRuleConfiguration structure.
type GetScopeRuleConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScopeRuleConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScopeRuleConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScopeRuleConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScopeRuleConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetScopeRuleConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScopeRuleConfigurationOK creates a GetScopeRuleConfigurationOK with default headers values
func NewGetScopeRuleConfigurationOK() *GetScopeRuleConfigurationOK {
	return &GetScopeRuleConfigurationOK{}
}

/* GetScopeRuleConfigurationOK describes a response with status code 200, with default header values.

GetScopeRuleConfigurationOK get scope rule configuration o k
*/
type GetScopeRuleConfigurationOK struct {
	Payload *models.CdnGetScopeRuleConfigurationResponse
}

func (o *GetScopeRuleConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] getScopeRuleConfigurationOK  %+v", 200, o.Payload)
}
func (o *GetScopeRuleConfigurationOK) GetPayload() *models.CdnGetScopeRuleConfigurationResponse {
	return o.Payload
}

func (o *GetScopeRuleConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetScopeRuleConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleConfigurationUnauthorized creates a GetScopeRuleConfigurationUnauthorized with default headers values
func NewGetScopeRuleConfigurationUnauthorized() *GetScopeRuleConfigurationUnauthorized {
	return &GetScopeRuleConfigurationUnauthorized{}
}

/* GetScopeRuleConfigurationUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetScopeRuleConfigurationUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetScopeRuleConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] getScopeRuleConfigurationUnauthorized  %+v", 401, o.Payload)
}
func (o *GetScopeRuleConfigurationUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleConfigurationInternalServerError creates a GetScopeRuleConfigurationInternalServerError with default headers values
func NewGetScopeRuleConfigurationInternalServerError() *GetScopeRuleConfigurationInternalServerError {
	return &GetScopeRuleConfigurationInternalServerError{}
}

/* GetScopeRuleConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetScopeRuleConfigurationInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetScopeRuleConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] getScopeRuleConfigurationInternalServerError  %+v", 500, o.Payload)
}
func (o *GetScopeRuleConfigurationInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScopeRuleConfigurationDefault creates a GetScopeRuleConfigurationDefault with default headers values
func NewGetScopeRuleConfigurationDefault(code int) *GetScopeRuleConfigurationDefault {
	return &GetScopeRuleConfigurationDefault{
		_statusCode: code,
	}
}

/* GetScopeRuleConfigurationDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetScopeRuleConfigurationDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get scope rule configuration default response
func (o *GetScopeRuleConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *GetScopeRuleConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] GetScopeRuleConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *GetScopeRuleConfigurationDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetScopeRuleConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
