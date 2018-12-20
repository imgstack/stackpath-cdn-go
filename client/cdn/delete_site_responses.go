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

// DeleteSiteReader is a Reader for the DeleteSite structure.
type DeleteSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSiteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteSiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteSiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSiteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteNoContent creates a DeleteSiteNoContent with default headers values
func NewDeleteSiteNoContent() *DeleteSiteNoContent {
	return &DeleteSiteNoContent{}
}

/*DeleteSiteNoContent handles this case with default header values.

No content
*/
type DeleteSiteNoContent struct {
}

func (o *DeleteSiteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}][%d] deleteSiteNoContent ", 204)
}

func (o *DeleteSiteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteUnauthorized creates a DeleteSiteUnauthorized with default headers values
func NewDeleteSiteUnauthorized() *DeleteSiteUnauthorized {
	return &DeleteSiteUnauthorized{}
}

/*DeleteSiteUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteSiteUnauthorized struct {
	Payload *models.APIStatus
}

func (o *DeleteSiteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}][%d] deleteSiteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSiteInternalServerError creates a DeleteSiteInternalServerError with default headers values
func NewDeleteSiteInternalServerError() *DeleteSiteInternalServerError {
	return &DeleteSiteInternalServerError{}
}

/*DeleteSiteInternalServerError handles this case with default header values.

Internal server error.
*/
type DeleteSiteInternalServerError struct {
	Payload *models.APIStatus
}

func (o *DeleteSiteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}][%d] deleteSiteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSiteDefault creates a DeleteSiteDefault with default headers values
func NewDeleteSiteDefault(code int) *DeleteSiteDefault {
	return &DeleteSiteDefault{
		_statusCode: code,
	}
}

/*DeleteSiteDefault handles this case with default header values.

Default error structure.
*/
type DeleteSiteDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the delete site default response
func (o *DeleteSiteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteDefault) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}][%d] DeleteSite default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSiteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
