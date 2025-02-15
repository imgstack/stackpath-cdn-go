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

// NewUpdateCertificateParams creates a new UpdateCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCertificateParams() *UpdateCertificateParams {
	return &UpdateCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCertificateParamsWithTimeout creates a new UpdateCertificateParams object
// with the ability to set a timeout on a request.
func NewUpdateCertificateParamsWithTimeout(timeout time.Duration) *UpdateCertificateParams {
	return &UpdateCertificateParams{
		timeout: timeout,
	}
}

// NewUpdateCertificateParamsWithContext creates a new UpdateCertificateParams object
// with the ability to set a context for a request.
func NewUpdateCertificateParamsWithContext(ctx context.Context) *UpdateCertificateParams {
	return &UpdateCertificateParams{
		Context: ctx,
	}
}

// NewUpdateCertificateParamsWithHTTPClient creates a new UpdateCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCertificateParamsWithHTTPClient(client *http.Client) *UpdateCertificateParams {
	return &UpdateCertificateParams{
		HTTPClient: client,
	}
}

/* UpdateCertificateParams contains all the parameters to send to the API endpoint
   for the update certificate operation.

   Typically these are written to a http.Request.
*/
type UpdateCertificateParams struct {

	// Body.
	Body *models.CdnUpdateCertificateRequest

	/* CertificateID.

	   The ID of SSL certificate to update
	*/
	CertificateID string

	/* StackID.

	   The ID of the stack containing the SSL certificate to update
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCertificateParams) WithDefaults() *UpdateCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update certificate params
func (o *UpdateCertificateParams) WithTimeout(timeout time.Duration) *UpdateCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update certificate params
func (o *UpdateCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update certificate params
func (o *UpdateCertificateParams) WithContext(ctx context.Context) *UpdateCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update certificate params
func (o *UpdateCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update certificate params
func (o *UpdateCertificateParams) WithHTTPClient(client *http.Client) *UpdateCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update certificate params
func (o *UpdateCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update certificate params
func (o *UpdateCertificateParams) WithBody(body *models.CdnUpdateCertificateRequest) *UpdateCertificateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update certificate params
func (o *UpdateCertificateParams) SetBody(body *models.CdnUpdateCertificateRequest) {
	o.Body = body
}

// WithCertificateID adds the certificateID to the update certificate params
func (o *UpdateCertificateParams) WithCertificateID(certificateID string) *UpdateCertificateParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the update certificate params
func (o *UpdateCertificateParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WithStackID adds the stackID to the update certificate params
func (o *UpdateCertificateParams) WithStackID(stackID string) *UpdateCertificateParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the update certificate params
func (o *UpdateCertificateParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param certificate_id
	if err := r.SetPathParam("certificate_id", o.CertificateID); err != nil {
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
