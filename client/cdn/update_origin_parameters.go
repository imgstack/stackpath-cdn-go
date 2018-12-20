// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// NewUpdateOriginParams creates a new UpdateOriginParams object
// with the default values initialized.
func NewUpdateOriginParams() *UpdateOriginParams {
	var ()
	return &UpdateOriginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOriginParamsWithTimeout creates a new UpdateOriginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOriginParamsWithTimeout(timeout time.Duration) *UpdateOriginParams {
	var ()
	return &UpdateOriginParams{

		timeout: timeout,
	}
}

// NewUpdateOriginParamsWithContext creates a new UpdateOriginParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOriginParamsWithContext(ctx context.Context) *UpdateOriginParams {
	var ()
	return &UpdateOriginParams{

		Context: ctx,
	}
}

// NewUpdateOriginParamsWithHTTPClient creates a new UpdateOriginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOriginParamsWithHTTPClient(client *http.Client) *UpdateOriginParams {
	var ()
	return &UpdateOriginParams{
		HTTPClient: client,
	}
}

/*UpdateOriginParams contains all the parameters to send to the API endpoint
for the update origin operation typically these are written to a http.Request
*/
type UpdateOriginParams struct {

	/*Body*/
	Body *models.CdnUpdateOriginRequest
	/*OriginID*/
	OriginID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update origin params
func (o *UpdateOriginParams) WithTimeout(timeout time.Duration) *UpdateOriginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update origin params
func (o *UpdateOriginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update origin params
func (o *UpdateOriginParams) WithContext(ctx context.Context) *UpdateOriginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update origin params
func (o *UpdateOriginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update origin params
func (o *UpdateOriginParams) WithHTTPClient(client *http.Client) *UpdateOriginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update origin params
func (o *UpdateOriginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update origin params
func (o *UpdateOriginParams) WithBody(body *models.CdnUpdateOriginRequest) *UpdateOriginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update origin params
func (o *UpdateOriginParams) SetBody(body *models.CdnUpdateOriginRequest) {
	o.Body = body
}

// WithOriginID adds the originID to the update origin params
func (o *UpdateOriginParams) WithOriginID(originID string) *UpdateOriginParams {
	o.SetOriginID(originID)
	return o
}

// SetOriginID adds the originId to the update origin params
func (o *UpdateOriginParams) SetOriginID(originID string) {
	o.OriginID = originID
}

// WithStackID adds the stackID to the update origin params
func (o *UpdateOriginParams) WithStackID(stackID string) *UpdateOriginParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update origin params
func (o *UpdateOriginParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOriginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param origin_id
	if err := r.SetPathParam("origin_id", o.OriginID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
