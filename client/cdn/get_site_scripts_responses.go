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

// GetSiteScriptsReader is a Reader for the GetSiteScripts structure.
type GetSiteScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSiteScriptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSiteScriptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSiteScriptsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteScriptsOK creates a GetSiteScriptsOK with default headers values
func NewGetSiteScriptsOK() *GetSiteScriptsOK {
	return &GetSiteScriptsOK{}
}

/* GetSiteScriptsOK describes a response with status code 200, with default header values.

GetSiteScriptsOK get site scripts o k
*/
type GetSiteScriptsOK struct {
	Payload *models.CdnGetSiteScriptsResponse
}

func (o *GetSiteScriptsOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] getSiteScriptsOK  %+v", 200, o.Payload)
}
func (o *GetSiteScriptsOK) GetPayload() *models.CdnGetSiteScriptsResponse {
	return o.Payload
}

func (o *GetSiteScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetSiteScriptsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScriptsUnauthorized creates a GetSiteScriptsUnauthorized with default headers values
func NewGetSiteScriptsUnauthorized() *GetSiteScriptsUnauthorized {
	return &GetSiteScriptsUnauthorized{}
}

/* GetSiteScriptsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetSiteScriptsUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetSiteScriptsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] getSiteScriptsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSiteScriptsUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScriptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScriptsInternalServerError creates a GetSiteScriptsInternalServerError with default headers values
func NewGetSiteScriptsInternalServerError() *GetSiteScriptsInternalServerError {
	return &GetSiteScriptsInternalServerError{}
}

/* GetSiteScriptsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetSiteScriptsInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetSiteScriptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] getSiteScriptsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSiteScriptsInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScriptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScriptsDefault creates a GetSiteScriptsDefault with default headers values
func NewGetSiteScriptsDefault(code int) *GetSiteScriptsDefault {
	return &GetSiteScriptsDefault{
		_statusCode: code,
	}
}

/* GetSiteScriptsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetSiteScriptsDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get site scripts default response
func (o *GetSiteScriptsDefault) Code() int {
	return o._statusCode
}

func (o *GetSiteScriptsDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts][%d] GetSiteScripts default  %+v", o._statusCode, o.Payload)
}
func (o *GetSiteScriptsDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScriptsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
