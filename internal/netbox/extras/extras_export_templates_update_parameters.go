// Code generated by go-swagger; DO NOT EDIT.

package extras

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
	"github.com/go-openapi/swag"

	"github.com/proelbtn/proxmox2netbox/models"
)

// NewExtrasExportTemplatesUpdateParams creates a new ExtrasExportTemplatesUpdateParams object
// with the default values initialized.
func NewExtrasExportTemplatesUpdateParams() *ExtrasExportTemplatesUpdateParams {
	var ()
	return &ExtrasExportTemplatesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithTimeout creates a new ExtrasExportTemplatesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasExportTemplatesUpdateParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesUpdateParams {
	var ()
	return &ExtrasExportTemplatesUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithContext creates a new ExtrasExportTemplatesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasExportTemplatesUpdateParamsWithContext(ctx context.Context) *ExtrasExportTemplatesUpdateParams {
	var ()
	return &ExtrasExportTemplatesUpdateParams{

		Context: ctx,
	}
}

// NewExtrasExportTemplatesUpdateParamsWithHTTPClient creates a new ExtrasExportTemplatesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasExportTemplatesUpdateParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesUpdateParams {
	var ()
	return &ExtrasExportTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasExportTemplatesUpdateParams contains all the parameters to send to the API endpoint
for the extras export templates update operation typically these are written to a http.Request
*/
type ExtrasExportTemplatesUpdateParams struct {

	/*Data*/
	Data *models.WritableExportTemplate
	/*ID
	  A unique integer value identifying this export template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithContext(ctx context.Context) *ExtrasExportTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithData(data *models.WritableExportTemplate) *ExtrasExportTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetData(data *models.WritableExportTemplate) {
	o.Data = data
}

// WithID adds the id to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) WithID(id int64) *ExtrasExportTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates update params
func (o *ExtrasExportTemplatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
