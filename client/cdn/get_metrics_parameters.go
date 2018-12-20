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

// NewGetMetricsParams creates a new GetMetricsParams object
// with the default values initialized.
func NewGetMetricsParams() *GetMetricsParams {
	var (
		granularityDefault = string("AUTO")
		groupByDefault     = string("NONE")
	)
	return &GetMetricsParams{
		Granularity: &granularityDefault,
		GroupBy:     &groupByDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetricsParamsWithTimeout creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMetricsParamsWithTimeout(timeout time.Duration) *GetMetricsParams {
	var (
		granularityDefault = string("AUTO")
		groupByDefault     = string("NONE")
	)
	return &GetMetricsParams{
		Granularity: &granularityDefault,
		GroupBy:     &groupByDefault,

		timeout: timeout,
	}
}

// NewGetMetricsParamsWithContext creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMetricsParamsWithContext(ctx context.Context) *GetMetricsParams {
	var (
		granularityDefault = string("AUTO")
		groupByDefault     = string("NONE")
	)
	return &GetMetricsParams{
		Granularity: &granularityDefault,
		GroupBy:     &groupByDefault,

		Context: ctx,
	}
}

// NewGetMetricsParamsWithHTTPClient creates a new GetMetricsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMetricsParamsWithHTTPClient(client *http.Client) *GetMetricsParams {
	var (
		granularityDefault = string("AUTO")
		groupByDefault     = string("NONE")
	)
	return &GetMetricsParams{
		Granularity: &granularityDefault,
		GroupBy:     &groupByDefault,
		HTTPClient:  client,
	}
}

/*GetMetricsParams contains all the parameters to send to the API endpoint
for the get metrics operation typically these are written to a http.Request
*/
type GetMetricsParams struct {

	/*BillingRegions*/
	BillingRegions *string
	/*EndDate*/
	EndDate *strfmt.DateTime
	/*Granularity*/
	Granularity *string
	/*GroupBy*/
	GroupBy *string
	/*Platforms*/
	Platforms *string
	/*Pops*/
	Pops *string
	/*Sites*/
	Sites *string
	/*StackID*/
	StackID string
	/*StartDate*/
	StartDate *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) WithTimeout(timeout time.Duration) *GetMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get metrics params
func (o *GetMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get metrics params
func (o *GetMetricsParams) WithContext(ctx context.Context) *GetMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get metrics params
func (o *GetMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) WithHTTPClient(client *http.Client) *GetMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get metrics params
func (o *GetMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBillingRegions adds the billingRegions to the get metrics params
func (o *GetMetricsParams) WithBillingRegions(billingRegions *string) *GetMetricsParams {
	o.SetBillingRegions(billingRegions)
	return o
}

// SetBillingRegions adds the billingRegions to the get metrics params
func (o *GetMetricsParams) SetBillingRegions(billingRegions *string) {
	o.BillingRegions = billingRegions
}

// WithEndDate adds the endDate to the get metrics params
func (o *GetMetricsParams) WithEndDate(endDate *strfmt.DateTime) *GetMetricsParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get metrics params
func (o *GetMetricsParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithGranularity adds the granularity to the get metrics params
func (o *GetMetricsParams) WithGranularity(granularity *string) *GetMetricsParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the get metrics params
func (o *GetMetricsParams) SetGranularity(granularity *string) {
	o.Granularity = granularity
}

// WithGroupBy adds the groupBy to the get metrics params
func (o *GetMetricsParams) WithGroupBy(groupBy *string) *GetMetricsParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get metrics params
func (o *GetMetricsParams) SetGroupBy(groupBy *string) {
	o.GroupBy = groupBy
}

// WithPlatforms adds the platforms to the get metrics params
func (o *GetMetricsParams) WithPlatforms(platforms *string) *GetMetricsParams {
	o.SetPlatforms(platforms)
	return o
}

// SetPlatforms adds the platforms to the get metrics params
func (o *GetMetricsParams) SetPlatforms(platforms *string) {
	o.Platforms = platforms
}

// WithPops adds the pops to the get metrics params
func (o *GetMetricsParams) WithPops(pops *string) *GetMetricsParams {
	o.SetPops(pops)
	return o
}

// SetPops adds the pops to the get metrics params
func (o *GetMetricsParams) SetPops(pops *string) {
	o.Pops = pops
}

// WithSites adds the sites to the get metrics params
func (o *GetMetricsParams) WithSites(sites *string) *GetMetricsParams {
	o.SetSites(sites)
	return o
}

// SetSites adds the sites to the get metrics params
func (o *GetMetricsParams) SetSites(sites *string) {
	o.Sites = sites
}

// WithStackID adds the stackID to the get metrics params
func (o *GetMetricsParams) WithStackID(stackID string) *GetMetricsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get metrics params
func (o *GetMetricsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WithStartDate adds the startDate to the get metrics params
func (o *GetMetricsParams) WithStartDate(startDate *strfmt.DateTime) *GetMetricsParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get metrics params
func (o *GetMetricsParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BillingRegions != nil {

		// query param billing_regions
		var qrBillingRegions string
		if o.BillingRegions != nil {
			qrBillingRegions = *o.BillingRegions
		}
		qBillingRegions := qrBillingRegions
		if qBillingRegions != "" {
			if err := r.SetQueryParam("billing_regions", qBillingRegions); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.DateTime
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.Granularity != nil {

		// query param granularity
		var qrGranularity string
		if o.Granularity != nil {
			qrGranularity = *o.Granularity
		}
		qGranularity := qrGranularity
		if qGranularity != "" {
			if err := r.SetQueryParam("granularity", qGranularity); err != nil {
				return err
			}
		}

	}

	if o.GroupBy != nil {

		// query param group_by
		var qrGroupBy string
		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := qrGroupBy
		if qGroupBy != "" {
			if err := r.SetQueryParam("group_by", qGroupBy); err != nil {
				return err
			}
		}

	}

	if o.Platforms != nil {

		// query param platforms
		var qrPlatforms string
		if o.Platforms != nil {
			qrPlatforms = *o.Platforms
		}
		qPlatforms := qrPlatforms
		if qPlatforms != "" {
			if err := r.SetQueryParam("platforms", qPlatforms); err != nil {
				return err
			}
		}

	}

	if o.Pops != nil {

		// query param pops
		var qrPops string
		if o.Pops != nil {
			qrPops = *o.Pops
		}
		qPops := qrPops
		if qPops != "" {
			if err := r.SetQueryParam("pops", qPops); err != nil {
				return err
			}
		}

	}

	if o.Sites != nil {

		// query param sites
		var qrSites string
		if o.Sites != nil {
			qrSites = *o.Sites
		}
		qSites := qrSites
		if qSites != "" {
			if err := r.SetQueryParam("sites", qSites); err != nil {
				return err
			}
		}

	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.DateTime
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
