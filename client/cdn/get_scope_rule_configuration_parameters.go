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

// NewGetScopeRuleConfigurationParams creates a new GetScopeRuleConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScopeRuleConfigurationParams() *GetScopeRuleConfigurationParams {
	return &GetScopeRuleConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScopeRuleConfigurationParamsWithTimeout creates a new GetScopeRuleConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetScopeRuleConfigurationParamsWithTimeout(timeout time.Duration) *GetScopeRuleConfigurationParams {
	return &GetScopeRuleConfigurationParams{
		timeout: timeout,
	}
}

// NewGetScopeRuleConfigurationParamsWithContext creates a new GetScopeRuleConfigurationParams object
// with the ability to set a context for a request.
func NewGetScopeRuleConfigurationParamsWithContext(ctx context.Context) *GetScopeRuleConfigurationParams {
	return &GetScopeRuleConfigurationParams{
		Context: ctx,
	}
}

// NewGetScopeRuleConfigurationParamsWithHTTPClient creates a new GetScopeRuleConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScopeRuleConfigurationParamsWithHTTPClient(client *http.Client) *GetScopeRuleConfigurationParams {
	return &GetScopeRuleConfigurationParams{
		HTTPClient: client,
	}
}

/* GetScopeRuleConfigurationParams contains all the parameters to send to the API endpoint
   for the get scope rule configuration operation.

   Typically these are written to a http.Request.
*/
type GetScopeRuleConfigurationParams struct {

	/* RuleID.

	   The ID of the EdgeRule to retrieve configuration for
	*/
	RuleID string

	/* ScopeID.

	   The ID of the CDN site scope to retrieve an EdgeRule's configuration from
	*/
	ScopeID string

	/* SiteID.

	   The ID of the site to retrieve an EdgeRule's configuration from
	*/
	SiteID string

	/* StackID.

	   The ID of the stack containing the site to retrieve an EdgeRule's configuration from
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scope rule configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScopeRuleConfigurationParams) WithDefaults() *GetScopeRuleConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scope rule configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScopeRuleConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithTimeout(timeout time.Duration) *GetScopeRuleConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithContext(ctx context.Context) *GetScopeRuleConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithHTTPClient(client *http.Client) *GetScopeRuleConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithRuleID(ruleID string) *GetScopeRuleConfigurationParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithScopeID adds the scopeID to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithScopeID(scopeID string) *GetScopeRuleConfigurationParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithSiteID(siteID string) *GetScopeRuleConfigurationParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) WithStackID(stackID string) *GetScopeRuleConfigurationParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get scope rule configuration params
func (o *GetScopeRuleConfigurationParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScopeRuleConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rule_id
	if err := r.SetPathParam("rule_id", o.RuleID); err != nil {
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
