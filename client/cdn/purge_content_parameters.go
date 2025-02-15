// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/renderinc/stackpath-cdn-go/models"
)

// NewPurgeContentParams creates a new PurgeContentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPurgeContentParams() *PurgeContentParams {
	return &PurgeContentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPurgeContentParamsWithTimeout creates a new PurgeContentParams object
// with the ability to set a timeout on a request.
func NewPurgeContentParamsWithTimeout(timeout time.Duration) *PurgeContentParams {
	return &PurgeContentParams{
		timeout: timeout,
	}
}

// NewPurgeContentParamsWithContext creates a new PurgeContentParams object
// with the ability to set a context for a request.
func NewPurgeContentParamsWithContext(ctx context.Context) *PurgeContentParams {
	return &PurgeContentParams{
		Context: ctx,
	}
}

// NewPurgeContentParamsWithHTTPClient creates a new PurgeContentParams object
// with the ability to set a custom HTTPClient for a request.
func NewPurgeContentParamsWithHTTPClient(client *http.Client) *PurgeContentParams {
	return &PurgeContentParams{
		HTTPClient: client,
	}
}

/* PurgeContentParams contains all the parameters to send to the API endpoint
   for the purge content operation.

   Typically these are written to a http.Request.
*/
type PurgeContentParams struct {

	// Body.
	Body *models.CdnPurgeContentRequest

	/* StackID.

	   The ID of the stack containing the sites to purge items from
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the purge content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PurgeContentParams) WithDefaults() *PurgeContentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the purge content params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PurgeContentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the purge content params
func (o *PurgeContentParams) WithTimeout(timeout time.Duration) *PurgeContentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purge content params
func (o *PurgeContentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purge content params
func (o *PurgeContentParams) WithContext(ctx context.Context) *PurgeContentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purge content params
func (o *PurgeContentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purge content params
func (o *PurgeContentParams) WithHTTPClient(client *http.Client) *PurgeContentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purge content params
func (o *PurgeContentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the purge content params
func (o *PurgeContentParams) WithBody(body *models.CdnPurgeContentRequest) *PurgeContentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the purge content params
func (o *PurgeContentParams) SetBody(body *models.CdnPurgeContentRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the purge content params
func (o *PurgeContentParams) WithStackID(stackID string) *PurgeContentParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the purge content params
func (o *PurgeContentParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *PurgeContentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
