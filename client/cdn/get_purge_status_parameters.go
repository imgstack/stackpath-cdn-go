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

// NewGetPurgeStatusParams creates a new GetPurgeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPurgeStatusParams() *GetPurgeStatusParams {
	return &GetPurgeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPurgeStatusParamsWithTimeout creates a new GetPurgeStatusParams object
// with the ability to set a timeout on a request.
func NewGetPurgeStatusParamsWithTimeout(timeout time.Duration) *GetPurgeStatusParams {
	return &GetPurgeStatusParams{
		timeout: timeout,
	}
}

// NewGetPurgeStatusParamsWithContext creates a new GetPurgeStatusParams object
// with the ability to set a context for a request.
func NewGetPurgeStatusParamsWithContext(ctx context.Context) *GetPurgeStatusParams {
	return &GetPurgeStatusParams{
		Context: ctx,
	}
}

// NewGetPurgeStatusParamsWithHTTPClient creates a new GetPurgeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPurgeStatusParamsWithHTTPClient(client *http.Client) *GetPurgeStatusParams {
	return &GetPurgeStatusParams{
		HTTPClient: client,
	}
}

/* GetPurgeStatusParams contains all the parameters to send to the API endpoint
   for the get purge status operation.

   Typically these are written to a http.Request.
*/
type GetPurgeStatusParams struct {

	/* PurgeID.

	   The ID of the purge request to check the status of
	*/
	PurgeID string

	/* StackID.

	   The ID of the stack the purge request was made against
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get purge status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurgeStatusParams) WithDefaults() *GetPurgeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get purge status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPurgeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get purge status params
func (o *GetPurgeStatusParams) WithTimeout(timeout time.Duration) *GetPurgeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get purge status params
func (o *GetPurgeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get purge status params
func (o *GetPurgeStatusParams) WithContext(ctx context.Context) *GetPurgeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get purge status params
func (o *GetPurgeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get purge status params
func (o *GetPurgeStatusParams) WithHTTPClient(client *http.Client) *GetPurgeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get purge status params
func (o *GetPurgeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPurgeID adds the purgeID to the get purge status params
func (o *GetPurgeStatusParams) WithPurgeID(purgeID string) *GetPurgeStatusParams {
	o.SetPurgeID(purgeID)
	return o
}

// SetPurgeID adds the purgeId to the get purge status params
func (o *GetPurgeStatusParams) SetPurgeID(purgeID string) {
	o.PurgeID = purgeID
}

// WithStackID adds the stackID to the get purge status params
func (o *GetPurgeStatusParams) WithStackID(stackID string) *GetPurgeStatusParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get purge status params
func (o *GetPurgeStatusParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPurgeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param purge_id
	if err := r.SetPathParam("purge_id", o.PurgeID); err != nil {
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
