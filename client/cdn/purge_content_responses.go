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

// PurgeContentReader is a Reader for the PurgeContent structure.
type PurgeContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurgeContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPurgeContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPurgeContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPurgeContentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPurgeContentOK creates a PurgeContentOK with default headers values
func NewPurgeContentOK() *PurgeContentOK {
	return &PurgeContentOK{}
}

/* PurgeContentOK describes a response with status code 200, with default header values.

PurgeContentOK purge content o k
*/
type PurgeContentOK struct {
	Payload *models.CdnPurgeContentResponse
}

func (o *PurgeContentOK) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/purge][%d] purgeContentOK  %+v", 200, o.Payload)
}
func (o *PurgeContentOK) GetPayload() *models.CdnPurgeContentResponse {
	return o.Payload
}

func (o *PurgeContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnPurgeContentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeContentUnauthorized creates a PurgeContentUnauthorized with default headers values
func NewPurgeContentUnauthorized() *PurgeContentUnauthorized {
	return &PurgeContentUnauthorized{}
}

/* PurgeContentUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type PurgeContentUnauthorized struct {
	Payload *models.APIStatus
}

func (o *PurgeContentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/purge][%d] purgeContentUnauthorized  %+v", 401, o.Payload)
}
func (o *PurgeContentUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *PurgeContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeContentInternalServerError creates a PurgeContentInternalServerError with default headers values
func NewPurgeContentInternalServerError() *PurgeContentInternalServerError {
	return &PurgeContentInternalServerError{}
}

/* PurgeContentInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type PurgeContentInternalServerError struct {
	Payload *models.APIStatus
}

func (o *PurgeContentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/purge][%d] purgeContentInternalServerError  %+v", 500, o.Payload)
}
func (o *PurgeContentInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *PurgeContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeContentDefault creates a PurgeContentDefault with default headers values
func NewPurgeContentDefault(code int) *PurgeContentDefault {
	return &PurgeContentDefault{
		_statusCode: code,
	}
}

/* PurgeContentDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type PurgeContentDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the purge content default response
func (o *PurgeContentDefault) Code() int {
	return o._statusCode
}

func (o *PurgeContentDefault) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/purge][%d] PurgeContent default  %+v", o._statusCode, o.Payload)
}
func (o *PurgeContentDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *PurgeContentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
