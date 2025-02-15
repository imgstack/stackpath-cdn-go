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

// DisconnectScopeFromOriginReader is a Reader for the DisconnectScopeFromOrigin structure.
type DisconnectScopeFromOriginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisconnectScopeFromOriginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDisconnectScopeFromOriginNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDisconnectScopeFromOriginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDisconnectScopeFromOriginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDisconnectScopeFromOriginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisconnectScopeFromOriginNoContent creates a DisconnectScopeFromOriginNoContent with default headers values
func NewDisconnectScopeFromOriginNoContent() *DisconnectScopeFromOriginNoContent {
	return &DisconnectScopeFromOriginNoContent{}
}

/* DisconnectScopeFromOriginNoContent describes a response with status code 204, with default header values.

No content
*/
type DisconnectScopeFromOriginNoContent struct {
}

func (o *DisconnectScopeFromOriginNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id}][%d] disconnectScopeFromOriginNoContent ", 204)
}

func (o *DisconnectScopeFromOriginNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisconnectScopeFromOriginUnauthorized creates a DisconnectScopeFromOriginUnauthorized with default headers values
func NewDisconnectScopeFromOriginUnauthorized() *DisconnectScopeFromOriginUnauthorized {
	return &DisconnectScopeFromOriginUnauthorized{}
}

/* DisconnectScopeFromOriginUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type DisconnectScopeFromOriginUnauthorized struct {
	Payload *models.APIStatus
}

func (o *DisconnectScopeFromOriginUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id}][%d] disconnectScopeFromOriginUnauthorized  %+v", 401, o.Payload)
}
func (o *DisconnectScopeFromOriginUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *DisconnectScopeFromOriginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisconnectScopeFromOriginInternalServerError creates a DisconnectScopeFromOriginInternalServerError with default headers values
func NewDisconnectScopeFromOriginInternalServerError() *DisconnectScopeFromOriginInternalServerError {
	return &DisconnectScopeFromOriginInternalServerError{}
}

/* DisconnectScopeFromOriginInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DisconnectScopeFromOriginInternalServerError struct {
	Payload *models.APIStatus
}

func (o *DisconnectScopeFromOriginInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id}][%d] disconnectScopeFromOriginInternalServerError  %+v", 500, o.Payload)
}
func (o *DisconnectScopeFromOriginInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *DisconnectScopeFromOriginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisconnectScopeFromOriginDefault creates a DisconnectScopeFromOriginDefault with default headers values
func NewDisconnectScopeFromOriginDefault(code int) *DisconnectScopeFromOriginDefault {
	return &DisconnectScopeFromOriginDefault{
		_statusCode: code,
	}
}

/* DisconnectScopeFromOriginDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type DisconnectScopeFromOriginDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the disconnect scope from origin default response
func (o *DisconnectScopeFromOriginDefault) Code() int {
	return o._statusCode
}

func (o *DisconnectScopeFromOriginDefault) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/origins/{origin_id}][%d] DisconnectScopeFromOrigin default  %+v", o._statusCode, o.Payload)
}
func (o *DisconnectScopeFromOriginDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *DisconnectScopeFromOriginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
