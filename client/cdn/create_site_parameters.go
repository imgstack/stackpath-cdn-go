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

// NewCreateSiteParams creates a new CreateSiteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSiteParams() *CreateSiteParams {
	return &CreateSiteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSiteParamsWithTimeout creates a new CreateSiteParams object
// with the ability to set a timeout on a request.
func NewCreateSiteParamsWithTimeout(timeout time.Duration) *CreateSiteParams {
	return &CreateSiteParams{
		timeout: timeout,
	}
}

// NewCreateSiteParamsWithContext creates a new CreateSiteParams object
// with the ability to set a context for a request.
func NewCreateSiteParamsWithContext(ctx context.Context) *CreateSiteParams {
	return &CreateSiteParams{
		Context: ctx,
	}
}

// NewCreateSiteParamsWithHTTPClient creates a new CreateSiteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSiteParamsWithHTTPClient(client *http.Client) *CreateSiteParams {
	return &CreateSiteParams{
		HTTPClient: client,
	}
}

/* CreateSiteParams contains all the parameters to send to the API endpoint
   for the create site operation.

   Typically these are written to a http.Request.
*/
type CreateSiteParams struct {

	// Body.
	Body *models.CdnCreateSiteRequest

	/* StackID.

	   The ID of the stack to create a new CDN site on
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSiteParams) WithDefaults() *CreateSiteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create site params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSiteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create site params
func (o *CreateSiteParams) WithTimeout(timeout time.Duration) *CreateSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create site params
func (o *CreateSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create site params
func (o *CreateSiteParams) WithContext(ctx context.Context) *CreateSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create site params
func (o *CreateSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create site params
func (o *CreateSiteParams) WithHTTPClient(client *http.Client) *CreateSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create site params
func (o *CreateSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create site params
func (o *CreateSiteParams) WithBody(body *models.CdnCreateSiteRequest) *CreateSiteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create site params
func (o *CreateSiteParams) SetBody(body *models.CdnCreateSiteRequest) {
	o.Body = body
}

// WithStackID adds the stackID to the create site params
func (o *CreateSiteParams) WithStackID(stackID string) *CreateSiteParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the create site params
func (o *CreateSiteParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
