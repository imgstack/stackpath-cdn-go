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

// GetSiteScript2Reader is a Reader for the GetSiteScript2 structure.
type GetSiteScript2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteScript2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSiteScript2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSiteScript2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSiteScript2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSiteScript2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSiteScript2OK creates a GetSiteScript2OK with default headers values
func NewGetSiteScript2OK() *GetSiteScript2OK {
	return &GetSiteScript2OK{}
}

/* GetSiteScript2OK describes a response with status code 200, with default header values.

GetSiteScript2OK get site script2 o k
*/
type GetSiteScript2OK struct {
	Payload *models.CdnGetSiteScriptResponse
}

func (o *GetSiteScript2OK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version}][%d] getSiteScript2OK  %+v", 200, o.Payload)
}
func (o *GetSiteScript2OK) GetPayload() *models.CdnGetSiteScriptResponse {
	return o.Payload
}

func (o *GetSiteScript2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetSiteScriptResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScript2Unauthorized creates a GetSiteScript2Unauthorized with default headers values
func NewGetSiteScript2Unauthorized() *GetSiteScript2Unauthorized {
	return &GetSiteScript2Unauthorized{}
}

/* GetSiteScript2Unauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetSiteScript2Unauthorized struct {
	Payload *models.APIStatus
}

func (o *GetSiteScript2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version}][%d] getSiteScript2Unauthorized  %+v", 401, o.Payload)
}
func (o *GetSiteScript2Unauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScript2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScript2InternalServerError creates a GetSiteScript2InternalServerError with default headers values
func NewGetSiteScript2InternalServerError() *GetSiteScript2InternalServerError {
	return &GetSiteScript2InternalServerError{}
}

/* GetSiteScript2InternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetSiteScript2InternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetSiteScript2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version}][%d] getSiteScript2InternalServerError  %+v", 500, o.Payload)
}
func (o *GetSiteScript2InternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScript2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteScript2Default creates a GetSiteScript2Default with default headers values
func NewGetSiteScript2Default(code int) *GetSiteScript2Default {
	return &GetSiteScript2Default{
		_statusCode: code,
	}
}

/* GetSiteScript2Default describes a response with status code -1, with default header values.

Default error structure.
*/
type GetSiteScript2Default struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get site script2 default response
func (o *GetSiteScript2Default) Code() int {
	return o._statusCode
}

func (o *GetSiteScript2Default) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}/{script_version}][%d] GetSiteScript2 default  %+v", o._statusCode, o.Payload)
}
func (o *GetSiteScript2Default) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetSiteScript2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
