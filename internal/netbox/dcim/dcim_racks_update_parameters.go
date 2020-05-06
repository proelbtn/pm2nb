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

// NewDcimRacksUpdateParams creates a new DcimRacksUpdateParams object
// with the default values initialized.
func NewDcimRacksUpdateParams() *DcimRacksUpdateParams {
	var ()
	return &DcimRacksUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksUpdateParamsWithTimeout creates a new DcimRacksUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRacksUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksUpdateParams {
	var ()
	return &DcimRacksUpdateParams{

		timeout: timeout,
	}
}

// NewDcimRacksUpdateParamsWithContext creates a new DcimRacksUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRacksUpdateParamsWithContext(ctx context.Context) *DcimRacksUpdateParams {
	var ()
	return &DcimRacksUpdateParams{

		Context: ctx,
	}
}

// NewDcimRacksUpdateParamsWithHTTPClient creates a new DcimRacksUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRacksUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksUpdateParams {
	var ()
	return &DcimRacksUpdateParams{
		HTTPClient: client,
	}
}

/*DcimRacksUpdateParams contains all the parameters to send to the API endpoint
for the dcim racks update operation typically these are written to a http.Request
*/
type DcimRacksUpdateParams struct {

	/*Data*/
	Data *models.WritableRack
	/*ID
	  A unique integer value identifying this rack.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim racks update params
func (o *DcimRacksUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks update params
func (o *DcimRacksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks update params
func (o *DcimRacksUpdateParams) WithContext(ctx context.Context) *DcimRacksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks update params
func (o *DcimRacksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks update params
func (o *DcimRacksUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks update params
func (o *DcimRacksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks update params
func (o *DcimRacksUpdateParams) WithData(data *models.WritableRack) *DcimRacksUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks update params
func (o *DcimRacksUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WithID adds the id to the dcim racks update params
func (o *DcimRacksUpdateParams) WithID(id int64) *DcimRacksUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks update params
func (o *DcimRacksUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
