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

// UpdateSiteCertificateHostsReader is a Reader for the UpdateSiteCertificateHosts structure.
type UpdateSiteCertificateHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteCertificateHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateSiteCertificateHostsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSiteCertificateHostsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSiteCertificateHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSiteCertificateHostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSiteCertificateHostsNoContent creates a UpdateSiteCertificateHostsNoContent with default headers values
func NewUpdateSiteCertificateHostsNoContent() *UpdateSiteCertificateHostsNoContent {
	return &UpdateSiteCertificateHostsNoContent{}
}

/* UpdateSiteCertificateHostsNoContent describes a response with status code 204, with default header values.

No content
*/
type UpdateSiteCertificateHostsNoContent struct {
}

func (o *UpdateSiteCertificateHostsNoContent) Error() string {
	return fmt.Sprintf("[PUT /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts][%d] updateSiteCertificateHostsNoContent ", 204)
}

func (o *UpdateSiteCertificateHostsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSiteCertificateHostsUnauthorized creates a UpdateSiteCertificateHostsUnauthorized with default headers values
func NewUpdateSiteCertificateHostsUnauthorized() *UpdateSiteCertificateHostsUnauthorized {
	return &UpdateSiteCertificateHostsUnauthorized{}
}

/* UpdateSiteCertificateHostsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateSiteCertificateHostsUnauthorized struct {
	Payload *models.APIStatus
}

func (o *UpdateSiteCertificateHostsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts][%d] updateSiteCertificateHostsUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateSiteCertificateHostsUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *UpdateSiteCertificateHostsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSiteCertificateHostsInternalServerError creates a UpdateSiteCertificateHostsInternalServerError with default headers values
func NewUpdateSiteCertificateHostsInternalServerError() *UpdateSiteCertificateHostsInternalServerError {
	return &UpdateSiteCertificateHostsInternalServerError{}
}

/* UpdateSiteCertificateHostsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type UpdateSiteCertificateHostsInternalServerError struct {
	Payload *models.APIStatus
}

func (o *UpdateSiteCertificateHostsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts][%d] updateSiteCertificateHostsInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateSiteCertificateHostsInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *UpdateSiteCertificateHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSiteCertificateHostsDefault creates a UpdateSiteCertificateHostsDefault with default headers values
func NewUpdateSiteCertificateHostsDefault(code int) *UpdateSiteCertificateHostsDefault {
	return &UpdateSiteCertificateHostsDefault{
		_statusCode: code,
	}
}

/* UpdateSiteCertificateHostsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type UpdateSiteCertificateHostsDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the update site certificate hosts default response
func (o *UpdateSiteCertificateHostsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteCertificateHostsDefault) Error() string {
	return fmt.Sprintf("[PUT /cdn/v1/stacks/{stack_id}/sites/{site_id}/certificates/{certificate_id}/hosts][%d] UpdateSiteCertificateHosts default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateSiteCertificateHostsDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *UpdateSiteCertificateHostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
