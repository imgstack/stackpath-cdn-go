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

// GetPopsReader is a Reader for the GetPops structure.
type GetPopsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPopsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPopsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPopsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPopsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPopsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPopsOK creates a GetPopsOK with default headers values
func NewGetPopsOK() *GetPopsOK {
	return &GetPopsOK{}
}

/* GetPopsOK describes a response with status code 200, with default header values.

GetPopsOK get pops o k
*/
type GetPopsOK struct {
	Payload *models.CdnGetPopsResponse
}

func (o *GetPopsOK) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/pops][%d] getPopsOK  %+v", 200, o.Payload)
}
func (o *GetPopsOK) GetPayload() *models.CdnGetPopsResponse {
	return o.Payload
}

func (o *GetPopsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnGetPopsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPopsUnauthorized creates a GetPopsUnauthorized with default headers values
func NewGetPopsUnauthorized() *GetPopsUnauthorized {
	return &GetPopsUnauthorized{}
}

/* GetPopsUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type GetPopsUnauthorized struct {
	Payload *models.APIStatus
}

func (o *GetPopsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/pops][%d] getPopsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPopsUnauthorized) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetPopsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPopsInternalServerError creates a GetPopsInternalServerError with default headers values
func NewGetPopsInternalServerError() *GetPopsInternalServerError {
	return &GetPopsInternalServerError{}
}

/* GetPopsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetPopsInternalServerError struct {
	Payload *models.APIStatus
}

func (o *GetPopsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/pops][%d] getPopsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPopsInternalServerError) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetPopsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPopsDefault creates a GetPopsDefault with default headers values
func NewGetPopsDefault(code int) *GetPopsDefault {
	return &GetPopsDefault{
		_statusCode: code,
	}
}

/* GetPopsDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type GetPopsDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the get pops default response
func (o *GetPopsDefault) Code() int {
	return o._statusCode
}

func (o *GetPopsDefault) Error() string {
	return fmt.Sprintf("[GET /cdn/v1/pops][%d] GetPops default  %+v", o._statusCode, o.Payload)
}
func (o *GetPopsDefault) GetPayload() *models.APIStatus {
	return o.Payload
}

func (o *GetPopsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
