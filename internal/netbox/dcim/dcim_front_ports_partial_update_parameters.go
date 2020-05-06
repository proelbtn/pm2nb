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

// NewDcimFrontPortsPartialUpdateParams creates a new DcimFrontPortsPartialUpdateParams object
// with the default values initialized.
func NewDcimFrontPortsPartialUpdateParams() *DcimFrontPortsPartialUpdateParams {
	var ()
	return &DcimFrontPortsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithTimeout creates a new DcimFrontPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimFrontPortsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsPartialUpdateParams {
	var ()
	return &DcimFrontPortsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithContext creates a new DcimFrontPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimFrontPortsPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsPartialUpdateParams {
	var ()
	return &DcimFrontPortsPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimFrontPortsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsPartialUpdateParams {
	var ()
	return &DcimFrontPortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimFrontPortsPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim front ports partial update operation typically these are written to a http.Request
*/
type DcimFrontPortsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableFrontPort
	/*ID
	  A unique integer value identifying this front port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WithID adds the id to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithID(id int64) *DcimFrontPortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
