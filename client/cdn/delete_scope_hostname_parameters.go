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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteScopeHostnameParams creates a new DeleteScopeHostnameParams object
// with the default values initialized.
func NewDeleteScopeHostnameParams() *DeleteScopeHostnameParams {
	var ()
	return &DeleteScopeHostnameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScopeHostnameParamsWithTimeout creates a new DeleteScopeHostnameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScopeHostnameParamsWithTimeout(timeout time.Duration) *DeleteScopeHostnameParams {
	var ()
	return &DeleteScopeHostnameParams{

		timeout: timeout,
	}
}

// NewDeleteScopeHostnameParamsWithContext creates a new DeleteScopeHostnameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScopeHostnameParamsWithContext(ctx context.Context) *DeleteScopeHostnameParams {
	var ()
	return &DeleteScopeHostnameParams{

		Context: ctx,
	}
}

// NewDeleteScopeHostnameParamsWithHTTPClient creates a new DeleteScopeHostnameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScopeHostnameParamsWithHTTPClient(client *http.Client) *DeleteScopeHostnameParams {
	var ()
	return &DeleteScopeHostnameParams{
		HTTPClient: client,
	}
}

/*DeleteScopeHostnameParams contains all the parameters to send to the API endpoint
for the delete scope hostname operation typically these are written to a http.Request
*/
type DeleteScopeHostnameParams struct {

	/*DisableTransparentMode*/
	DisableTransparentMode *bool
	/*Domain*/
	Domain string
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

// WithTimeout adds the timeout to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithTimeout(timeout time.Duration) *DeleteScopeHostnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithContext(ctx context.Context) *DeleteScopeHostnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithHTTPClient(client *http.Client) *DeleteScopeHostnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisableTransparentMode adds the disableTransparentMode to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithDisableTransparentMode(disableTransparentMode *bool) *DeleteScopeHostnameParams {
	o.SetDisableTransparentMode(disableTransparentMode)
	return o
}

// SetDisableTransparentMode adds the disableTransparentMode to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetDisableTransparentMode(disableTransparentMode *bool) {
	o.DisableTransparentMode = disableTransparentMode
}

// WithDomain adds the domain to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithDomain(domain string) *DeleteScopeHostnameParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithScopeID adds the scopeID to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithScopeID(scopeID string) *DeleteScopeHostnameParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithSiteID(siteID string) *DeleteScopeHostnameParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the delete scope hostname params
func (o *DeleteScopeHostnameParams) WithStackID(stackID string) *DeleteScopeHostnameParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete scope hostname params
func (o *DeleteScopeHostnameParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScopeHostnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisableTransparentMode != nil {

		// query param disable_transparent_mode
		var qrDisableTransparentMode bool
		if o.DisableTransparentMode != nil {
			qrDisableTransparentMode = *o.DisableTransparentMode
		}
		qDisableTransparentMode := swag.FormatBool(qrDisableTransparentMode)
		if qDisableTransparentMode != "" {
			if err := r.SetQueryParam("disable_transparent_mode", qDisableTransparentMode); err != nil {
				return err
			}
		}

	}

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
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
