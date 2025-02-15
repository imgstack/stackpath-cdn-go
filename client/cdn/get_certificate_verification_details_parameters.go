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
)

// NewGetCertificateVerificationDetailsParams creates a new GetCertificateVerificationDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCertificateVerificationDetailsParams() *GetCertificateVerificationDetailsParams {
	return &GetCertificateVerificationDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCertificateVerificationDetailsParamsWithTimeout creates a new GetCertificateVerificationDetailsParams object
// with the ability to set a timeout on a request.
func NewGetCertificateVerificationDetailsParamsWithTimeout(timeout time.Duration) *GetCertificateVerificationDetailsParams {
	return &GetCertificateVerificationDetailsParams{
		timeout: timeout,
	}
}

// NewGetCertificateVerificationDetailsParamsWithContext creates a new GetCertificateVerificationDetailsParams object
// with the ability to set a context for a request.
func NewGetCertificateVerificationDetailsParamsWithContext(ctx context.Context) *GetCertificateVerificationDetailsParams {
	return &GetCertificateVerificationDetailsParams{
		Context: ctx,
	}
}

// NewGetCertificateVerificationDetailsParamsWithHTTPClient creates a new GetCertificateVerificationDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCertificateVerificationDetailsParamsWithHTTPClient(client *http.Client) *GetCertificateVerificationDetailsParams {
	return &GetCertificateVerificationDetailsParams{
		HTTPClient: client,
	}
}

/* GetCertificateVerificationDetailsParams contains all the parameters to send to the API endpoint
   for the get certificate verification details operation.

   Typically these are written to a http.Request.
*/
type GetCertificateVerificationDetailsParams struct {

	/* CertificateID.

	   The ID of the site to retrieve SSL certificate verification details for
	*/
	CertificateID string

	/* StackID.

	   The ID of the stack containing the site to retrieve SSL certificate verification details for
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get certificate verification details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificateVerificationDetailsParams) WithDefaults() *GetCertificateVerificationDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get certificate verification details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCertificateVerificationDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) WithTimeout(timeout time.Duration) *GetCertificateVerificationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) WithContext(ctx context.Context) *GetCertificateVerificationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) WithHTTPClient(client *http.Client) *GetCertificateVerificationDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateID adds the certificateID to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) WithCertificateID(certificateID string) *GetCertificateVerificationDetailsParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WithStackID adds the stackID to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) WithStackID(stackID string) *GetCertificateVerificationDetailsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get certificate verification details params
func (o *GetCertificateVerificationDetailsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCertificateVerificationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
