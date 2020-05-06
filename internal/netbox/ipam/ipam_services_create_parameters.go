// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

	"github.com/proelbtn/proxmox2netbox/models"
)

// NewIpamServicesCreateParams creates a new IpamServicesCreateParams object
// with the default values initialized.
func NewIpamServicesCreateParams() *IpamServicesCreateParams {
	var ()
	return &IpamServicesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesCreateParamsWithTimeout creates a new IpamServicesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamServicesCreateParamsWithTimeout(timeout time.Duration) *IpamServicesCreateParams {
	var ()
	return &IpamServicesCreateParams{

		timeout: timeout,
	}
}

// NewIpamServicesCreateParamsWithContext creates a new IpamServicesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamServicesCreateParamsWithContext(ctx context.Context) *IpamServicesCreateParams {
	var ()
	return &IpamServicesCreateParams{

		Context: ctx,
	}
}

// NewIpamServicesCreateParamsWithHTTPClient creates a new IpamServicesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamServicesCreateParamsWithHTTPClient(client *http.Client) *IpamServicesCreateParams {
	var ()
	return &IpamServicesCreateParams{
		HTTPClient: client,
	}
}

/*IpamServicesCreateParams contains all the parameters to send to the API endpoint
for the ipam services create operation typically these are written to a http.Request
*/
type IpamServicesCreateParams struct {

	/*Data*/
	Data *models.WritableService

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam services create params
func (o *IpamServicesCreateParams) WithTimeout(timeout time.Duration) *IpamServicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services create params
func (o *IpamServicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services create params
func (o *IpamServicesCreateParams) WithContext(ctx context.Context) *IpamServicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services create params
func (o *IpamServicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services create params
func (o *IpamServicesCreateParams) WithHTTPClient(client *http.Client) *IpamServicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services create params
func (o *IpamServicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services create params
func (o *IpamServicesCreateParams) WithData(data *models.WritableService) *IpamServicesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services create params
func (o *IpamServicesCreateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
