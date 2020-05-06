// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// NewDcimDeviceBayTemplatesPartialUpdateParams creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the default values initialized.
func NewDcimDeviceBayTemplatesPartialUpdateParams() *DcimDeviceBayTemplatesPartialUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithTimeout creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesPartialUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithContext creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesPartialUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesPartialUpdateParams {
	var ()
	return &DcimDeviceBayTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceBayTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim device bay templates partial update operation typically these are written to a http.Request
*/
type DcimDeviceBayTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableDeviceBayTemplate
	/*ID
	  A unique integer value identifying this device bay template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithData(data *models.WritableDeviceBayTemplate) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetData(data *models.WritableDeviceBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithID(id int64) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
