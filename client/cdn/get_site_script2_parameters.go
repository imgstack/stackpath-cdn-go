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

// NewGetSiteScript2Params creates a new GetSiteScript2Params object
// with the default values initialized.
func NewGetSiteScript2Params() *GetSiteScript2Params {
	var ()
	return &GetSiteScript2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteScript2ParamsWithTimeout creates a new GetSiteScript2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteScript2ParamsWithTimeout(timeout time.Duration) *GetSiteScript2Params {
	var ()
	return &GetSiteScript2Params{

		timeout: timeout,
	}
}

// NewGetSiteScript2ParamsWithContext creates a new GetSiteScript2Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteScript2ParamsWithContext(ctx context.Context) *GetSiteScript2Params {
	var ()
	return &GetSiteScript2Params{

		Context: ctx,
	}
}

// NewGetSiteScript2ParamsWithHTTPClient creates a new GetSiteScript2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteScript2ParamsWithHTTPClient(client *http.Client) *GetSiteScript2Params {
	var ()
	return &GetSiteScript2Params{
		HTTPClient: client,
	}
}

/*GetSiteScript2Params contains all the parameters to send to the API endpoint
for the get site script2 operation typically these are written to a http.Request
*/
type GetSiteScript2Params struct {

	/*ScriptID*/
	ScriptID string
	/*ScriptVersion*/
	ScriptVersion string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site script2 params
func (o *GetSiteScript2Params) WithTimeout(timeout time.Duration) *GetSiteScript2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site script2 params
func (o *GetSiteScript2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site script2 params
func (o *GetSiteScript2Params) WithContext(ctx context.Context) *GetSiteScript2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site script2 params
func (o *GetSiteScript2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site script2 params
func (o *GetSiteScript2Params) WithHTTPClient(client *http.Client) *GetSiteScript2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site script2 params
func (o *GetSiteScript2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScriptID adds the scriptID to the get site script2 params
func (o *GetSiteScript2Params) WithScriptID(scriptID string) *GetSiteScript2Params {
	o.SetScriptID(scriptID)
	return o
}

// SetScriptID adds the scriptId to the get site script2 params
func (o *GetSiteScript2Params) SetScriptID(scriptID string) {
	o.ScriptID = scriptID
}

// WithScriptVersion adds the scriptVersion to the get site script2 params
func (o *GetSiteScript2Params) WithScriptVersion(scriptVersion string) *GetSiteScript2Params {
	o.SetScriptVersion(scriptVersion)
	return o
}

// SetScriptVersion adds the scriptVersion to the get site script2 params
func (o *GetSiteScript2Params) SetScriptVersion(scriptVersion string) {
	o.ScriptVersion = scriptVersion
}

// WithSiteID adds the siteID to the get site script2 params
func (o *GetSiteScript2Params) WithSiteID(siteID string) *GetSiteScript2Params {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get site script2 params
func (o *GetSiteScript2Params) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the get site script2 params
func (o *GetSiteScript2Params) WithStackID(stackID string) *GetSiteScript2Params {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get site script2 params
func (o *GetSiteScript2Params) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteScript2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param script_id
	if err := r.SetPathParam("script_id", o.ScriptID); err != nil {
		return err
	}

	// path param script_version
	if err := r.SetPathParam("script_version", o.ScriptVersion); err != nil {
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
