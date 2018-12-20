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
)

// NewGetPopsParams creates a new GetPopsParams object
// with the default values initialized.
func NewGetPopsParams() *GetPopsParams {

	return &GetPopsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPopsParamsWithTimeout creates a new GetPopsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPopsParamsWithTimeout(timeout time.Duration) *GetPopsParams {

	return &GetPopsParams{

		timeout: timeout,
	}
}

// NewGetPopsParamsWithContext creates a new GetPopsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPopsParamsWithContext(ctx context.Context) *GetPopsParams {

	return &GetPopsParams{

		Context: ctx,
	}
}

// NewGetPopsParamsWithHTTPClient creates a new GetPopsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPopsParamsWithHTTPClient(client *http.Client) *GetPopsParams {

	return &GetPopsParams{
		HTTPClient: client,
	}
}

/*GetPopsParams contains all the parameters to send to the API endpoint
for the get pops operation typically these are written to a http.Request
*/
type GetPopsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pops params
func (o *GetPopsParams) WithTimeout(timeout time.Duration) *GetPopsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pops params
func (o *GetPopsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pops params
func (o *GetPopsParams) WithContext(ctx context.Context) *GetPopsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pops params
func (o *GetPopsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pops params
func (o *GetPopsParams) WithHTTPClient(client *http.Client) *GetPopsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pops params
func (o *GetPopsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPopsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
