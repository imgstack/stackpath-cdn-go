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

// CreateSiteScriptReader is a Reader for the CreateSiteScript structure.
type CreateSiteScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSiteScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateSiteScriptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateSiteScriptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateSiteScriptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSiteScriptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSiteScriptOK creates a CreateSiteScriptOK with default headers values
func NewCreateSiteScriptOK() *CreateSiteScriptOK {
	return &CreateSiteScriptOK{}
}

/*CreateSiteScriptOK handles this case with default header values.

CreateSiteScriptOK create site script o k
*/
type CreateSiteScriptOK struct {
	Payload *models.CdnCreateSiteScriptResponse
}

func (o *CreateSiteScriptOK) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] createSiteScriptOK  %+v", 200, o.Payload)
}

func (o *CreateSiteScriptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnCreateSiteScriptResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteScriptUnauthorized creates a CreateSiteScriptUnauthorized with default headers values
func NewCreateSiteScriptUnauthorized() *CreateSiteScriptUnauthorized {
	return &CreateSiteScriptUnauthorized{}
}

/*CreateSiteScriptUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type CreateSiteScriptUnauthorized struct {
	Payload *models.APIStatus
}

func (o *CreateSiteScriptUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] createSiteScriptUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateSiteScriptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteScriptInternalServerError creates a CreateSiteScriptInternalServerError with default headers values
func NewCreateSiteScriptInternalServerError() *CreateSiteScriptInternalServerError {
	return &CreateSiteScriptInternalServerError{}
}

/*CreateSiteScriptInternalServerError handles this case with default header values.

Internal server error.
*/
type CreateSiteScriptInternalServerError struct {
	Payload *models.APIStatus
}

func (o *CreateSiteScriptInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] createSiteScriptInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateSiteScriptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSiteScriptDefault creates a CreateSiteScriptDefault with default headers values
func NewCreateSiteScriptDefault(code int) *CreateSiteScriptDefault {
	return &CreateSiteScriptDefault{
		_statusCode: code,
	}
}

/*CreateSiteScriptDefault handles this case with default header values.

Default error structure.
*/
type CreateSiteScriptDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the create site script default response
func (o *CreateSiteScriptDefault) Code() int {
	return o._statusCode
}

func (o *CreateSiteScriptDefault) Error() string {
	return fmt.Sprintf("[POST /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] CreateSiteScript default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSiteScriptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
