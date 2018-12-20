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

// NewGetSiteScriptsParams creates a new GetSiteScriptsParams object
// with the default values initialized.
func NewGetSiteScriptsParams() *GetSiteScriptsParams {
	var ()
	return &GetSiteScriptsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteScriptsParamsWithTimeout creates a new GetSiteScriptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteScriptsParamsWithTimeout(timeout time.Duration) *GetSiteScriptsParams {
	var ()
	return &GetSiteScriptsParams{

		timeout: timeout,
	}
}

// NewGetSiteScriptsParamsWithContext creates a new GetSiteScriptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteScriptsParamsWithContext(ctx context.Context) *GetSiteScriptsParams {
	var ()
	return &GetSiteScriptsParams{

		Context: ctx,
	}
}

// NewGetSiteScriptsParamsWithHTTPClient creates a new GetSiteScriptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteScriptsParamsWithHTTPClient(client *http.Client) *GetSiteScriptsParams {
	var ()
	return &GetSiteScriptsParams{
		HTTPClient: client,
	}
}

/*GetSiteScriptsParams contains all the parameters to send to the API endpoint
for the get site scripts operation typically these are written to a http.Request
*/
type GetSiteScriptsParams struct {

	/*PageRequestAfter
	  after is the cursor value after which data will be returned.

	*/
	PageRequestAfter *string
	/*PageRequestFilter
	  filter will accept sql style constraints.

	*/
	PageRequestFilter *string
	/*PageRequestFirst
	  first is the number of items desired.

	*/
	PageRequestFirst *string
	/*PageRequestSortBy
	  sort_by will sort the response by the given field.

	*/
	PageRequestSortBy *string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site scripts params
func (o *GetSiteScriptsParams) WithTimeout(timeout time.Duration) *GetSiteScriptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site scripts params
func (o *GetSiteScriptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site scripts params
func (o *GetSiteScriptsParams) WithContext(ctx context.Context) *GetSiteScriptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site scripts params
func (o *GetSiteScriptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site scripts params
func (o *GetSiteScriptsParams) WithHTTPClient(client *http.Client) *GetSiteScriptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site scripts params
func (o *GetSiteScriptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get site scripts params
func (o *GetSiteScriptsParams) WithPageRequestAfter(pageRequestAfter *string) *GetSiteScriptsParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get site scripts params
func (o *GetSiteScriptsParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get site scripts params
func (o *GetSiteScriptsParams) WithPageRequestFilter(pageRequestFilter *string) *GetSiteScriptsParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get site scripts params
func (o *GetSiteScriptsParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get site scripts params
func (o *GetSiteScriptsParams) WithPageRequestFirst(pageRequestFirst *string) *GetSiteScriptsParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get site scripts params
func (o *GetSiteScriptsParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get site scripts params
func (o *GetSiteScriptsParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetSiteScriptsParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get site scripts params
func (o *GetSiteScriptsParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithSiteID adds the siteID to the get site scripts params
func (o *GetSiteScriptsParams) WithSiteID(siteID string) *GetSiteScriptsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get site scripts params
func (o *GetSiteScriptsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the get site scripts params
func (o *GetSiteScriptsParams) WithStackID(stackID string) *GetSiteScriptsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get site scripts params
func (o *GetSiteScriptsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteScriptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageRequestAfter != nil {

		// query param page_request.after
		var qrPageRequestAfter string
		if o.PageRequestAfter != nil {
			qrPageRequestAfter = *o.PageRequestAfter
		}
		qPageRequestAfter := qrPageRequestAfter
		if qPageRequestAfter != "" {
			if err := r.SetQueryParam("page_request.after", qPageRequestAfter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFilter != nil {

		// query param page_request.filter
		var qrPageRequestFilter string
		if o.PageRequestFilter != nil {
			qrPageRequestFilter = *o.PageRequestFilter
		}
		qPageRequestFilter := qrPageRequestFilter
		if qPageRequestFilter != "" {
			if err := r.SetQueryParam("page_request.filter", qPageRequestFilter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFirst != nil {

		// query param page_request.first
		var qrPageRequestFirst string
		if o.PageRequestFirst != nil {
			qrPageRequestFirst = *o.PageRequestFirst
		}
		qPageRequestFirst := qrPageRequestFirst
		if qPageRequestFirst != "" {
			if err := r.SetQueryParam("page_request.first", qPageRequestFirst); err != nil {
				return err
			}
		}

	}

	if o.PageRequestSortBy != nil {

		// query param page_request.sort_by
		var qrPageRequestSortBy string
		if o.PageRequestSortBy != nil {
			qrPageRequestSortBy = *o.PageRequestSortBy
		}
		qPageRequestSortBy := qrPageRequestSortBy
		if qPageRequestSortBy != "" {
			if err := r.SetQueryParam("page_request.sort_by", qPageRequestSortBy); err != nil {
				return err
			}
		}

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
