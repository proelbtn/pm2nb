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
	"github.com/go-openapi/swag"
)

// NewIpamPrefixesDeleteParams creates a new IpamPrefixesDeleteParams object
// with the default values initialized.
func NewIpamPrefixesDeleteParams() *IpamPrefixesDeleteParams {
	var ()
	return &IpamPrefixesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesDeleteParamsWithTimeout creates a new IpamPrefixesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamPrefixesDeleteParamsWithTimeout(timeout time.Duration) *IpamPrefixesDeleteParams {
	var ()
	return &IpamPrefixesDeleteParams{

		timeout: timeout,
	}
}

// NewIpamPrefixesDeleteParamsWithContext creates a new IpamPrefixesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamPrefixesDeleteParamsWithContext(ctx context.Context) *IpamPrefixesDeleteParams {
	var ()
	return &IpamPrefixesDeleteParams{

		Context: ctx,
	}
}

// NewIpamPrefixesDeleteParamsWithHTTPClient creates a new IpamPrefixesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamPrefixesDeleteParamsWithHTTPClient(client *http.Client) *IpamPrefixesDeleteParams {
	var ()
	return &IpamPrefixesDeleteParams{
		HTTPClient: client,
	}
}

/*IpamPrefixesDeleteParams contains all the parameters to send to the API endpoint
for the ipam prefixes delete operation typically these are written to a http.Request
*/
type IpamPrefixesDeleteParams struct {

	/*ID
	  A unique integer value identifying this prefix.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithTimeout(timeout time.Duration) *IpamPrefixesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithContext(ctx context.Context) *IpamPrefixesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithHTTPClient(client *http.Client) *IpamPrefixesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithID(id int64) *IpamPrefixesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
