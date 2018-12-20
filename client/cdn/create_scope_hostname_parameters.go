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

// NewCreateScopeHostnameParams creates a new CreateScopeHostnameParams object
// with the default values initialized.
func NewCreateScopeHostnameParams() *CreateScopeHostnameParams {
	var ()
	return &CreateScopeHostnameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateScopeHostnameParamsWithTimeout creates a new CreateScopeHostnameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateScopeHostnameParamsWithTimeout(timeout time.Duration) *CreateScopeHostnameParams {
	var ()
	return &CreateScopeHostnameParams{

		timeout: timeout,
	}
}

// NewCreateScopeHostnameParamsWithContext creates a new CreateScopeHostnameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateScopeHostnameParamsWithContext(ctx context.Context) *CreateScopeHostnameParams {
	var ()
	return &CreateScopeHostnameParams{

		Context: ctx,
	}
}

// NewCreateScopeHostnameParamsWithHTTPClient creates a new CreateScopeHostnameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateScopeHostnameParamsWithHTTPClient(client *http.Client) *CreateScopeHostnameParams {
	var ()
	return &CreateScopeHostnameParams{
		HTTPClient: client,
	}
}

/*CreateScopeHostnameParams contains all the parameters to send to the API endpoint
for the create scope hostname operation typically these are written to a http.Request
*/
type CreateScopeHostnameParams struct {

	/*Body*/
	Body *models.CdnCreateScopeHostnameRequest
	/*ScopeID*/
	ScopeID string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create scope hostname params
func (o *CreateScopeHostnameParams) WithTimeout(timeout time.Duration) *CreateScopeHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create scope hostname params
func (o *CreateScopeHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create scope hostname params
func (o *CreateScopeHostnameParams) WithContext(ctx context.Context) *CreateScopeHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create scope hostname params
func (o *CreateScopeHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create scope hostname params
func (o *CreateScopeHostnameParams) WithHTTPClient(client *http.Client) *CreateScopeHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create scope hostname params
func (o *CreateScopeHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create scope hostname params
func (o *CreateScopeHostnameParams) WithBody(body *models.CdnCreateScopeHostnameRequest) *CreateScopeHostnameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create scope hostname params
func (o *CreateScopeHostnameParams) SetBody(body *models.CdnCreateScopeHostnameRequest) {
	o.Body = body
}

// WithScopeID adds the scopeID to the create scope hostname params
func (o *CreateScopeHostnameParams) WithScopeID(scopeID string) *CreateScopeHostnameParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the create scope hostname params
func (o *CreateScopeHostnameParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the create scope hostname params
func (o *CreateScopeHostnameParams) WithSiteID(siteID string) *CreateScopeHostnameParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create scope hostname params
func (o *CreateScopeHostnameParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the create scope hostname params
func (o *CreateScopeHostnameParams) WithStackID(stackID string) *CreateScopeHostnameParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the create scope hostname params
func (o *CreateScopeHostnameParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateScopeHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param scope_id
	if err := r.SetPathParam("scope_id", o.ScopeID); err != nil {
		return err
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
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
