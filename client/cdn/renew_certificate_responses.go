// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// RenewCertificateReader is a Reader for the RenewCertificate structure.
type RenewCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenewCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRenewCertificateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRenewCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRenewCertificateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRenewCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRenewCertificateNoContent creates a RenewCertificateNoContent with default headers values
func NewRenewCertificateNoContent() *RenewCertificateNoContent {
	return &RenewCertificateNoContent{}
}

/*RenewCertificateNoContent handles this case with default header values.

No content
*/
type RenewCertificateNoContent struct {
}

func (o *RenewCertificateNoContent) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew][%d] renewCertificateNoContent ", 204)
}

func (o *RenewCertificateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenewCertificateUnauthorized creates a RenewCertificateUnauthorized with default headers values
func NewRenewCertificateUnauthorized() *RenewCertificateUnauthorized {
	return &RenewCertificateUnauthorized{}
}

/*RenewCertificateUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type RenewCertificateUnauthorized struct {
	Payload *models.APIStatus
}

func (o *RenewCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew][%d] renewCertificateUnauthorized  %+v", 401, o.Payload)
}

func (o *RenewCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenewCertificateInternalServerError creates a RenewCertificateInternalServerError with default headers values
func NewRenewCertificateInternalServerError() *RenewCertificateInternalServerError {
	return &RenewCertificateInternalServerError{}
}

/*RenewCertificateInternalServerError handles this case with default header values.

Internal server error.
*/
type RenewCertificateInternalServerError struct {
	Payload *models.APIStatus
}

func (o *RenewCertificateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew][%d] renewCertificateInternalServerError  %+v", 500, o.Payload)
}

func (o *RenewCertificateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenewCertificateDefault creates a RenewCertificateDefault with default headers values
func NewRenewCertificateDefault(code int) *RenewCertificateDefault {
	return &RenewCertificateDefault{
		_statusCode: code,
	}
}

/*RenewCertificateDefault handles this case with default header values.

Default error structure.
*/
type RenewCertificateDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the renew certificate default response
func (o *RenewCertificateDefault) Code() int {
	return o._statusCode
}

func (o *RenewCertificateDefault) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/certificates/{certificate_id}/renew][%d] RenewCertificate default  %+v", o._statusCode, o.Payload)
}

func (o *RenewCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
