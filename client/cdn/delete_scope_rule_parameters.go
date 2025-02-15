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

// NewDeleteScopeRuleParams creates a new DeleteScopeRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteScopeRuleParams() *DeleteScopeRuleParams {
	return &DeleteScopeRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScopeRuleParamsWithTimeout creates a new DeleteScopeRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteScopeRuleParamsWithTimeout(timeout time.Duration) *DeleteScopeRuleParams {
	return &DeleteScopeRuleParams{
		timeout: timeout,
	}
}

// NewDeleteScopeRuleParamsWithContext creates a new DeleteScopeRuleParams object
// with the ability to set a context for a request.
func NewDeleteScopeRuleParamsWithContext(ctx context.Context) *DeleteScopeRuleParams {
	return &DeleteScopeRuleParams{
		Context: ctx,
	}
}

// NewDeleteScopeRuleParamsWithHTTPClient creates a new DeleteScopeRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteScopeRuleParamsWithHTTPClient(client *http.Client) *DeleteScopeRuleParams {
	return &DeleteScopeRuleParams{
		HTTPClient: client,
	}
}

/* DeleteScopeRuleParams contains all the parameters to send to the API endpoint
   for the delete scope rule operation.

   Typically these are written to a http.Request.
*/
type DeleteScopeRuleParams struct {

	/* RuleID.

	   The ID of the EdgeRule to delete
	*/
	RuleID string

	/* ScopeID.

	   The ID of the CDN site scope to delete an EdgeRule from
	*/
	ScopeID string

	/* SiteID.

	   The ID of the site to delete an EdgeRule from
	*/
	SiteID string

	/* StackID.

	   The ID of the stack containing the site to delete an EdgeRule from
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete scope rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScopeRuleParams) WithDefaults() *DeleteScopeRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete scope rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScopeRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete scope rule params
func (o *DeleteScopeRuleParams) WithTimeout(timeout time.Duration) *DeleteScopeRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scope rule params
func (o *DeleteScopeRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scope rule params
func (o *DeleteScopeRuleParams) WithContext(ctx context.Context) *DeleteScopeRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scope rule params
func (o *DeleteScopeRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scope rule params
func (o *DeleteScopeRuleParams) WithHTTPClient(client *http.Client) *DeleteScopeRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scope rule params
func (o *DeleteScopeRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleID adds the ruleID to the delete scope rule params
func (o *DeleteScopeRuleParams) WithRuleID(ruleID string) *DeleteScopeRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the delete scope rule params
func (o *DeleteScopeRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithScopeID adds the scopeID to the delete scope rule params
func (o *DeleteScopeRuleParams) WithScopeID(scopeID string) *DeleteScopeRuleParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the delete scope rule params
func (o *DeleteScopeRuleParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the delete scope rule params
func (o *DeleteScopeRuleParams) WithSiteID(siteID string) *DeleteScopeRuleParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete scope rule params
func (o *DeleteScopeRuleParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the delete scope rule params
func (o *DeleteScopeRuleParams) WithStackID(stackID string) *DeleteScopeRuleParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete scope rule params
func (o *DeleteScopeRuleParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScopeRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
