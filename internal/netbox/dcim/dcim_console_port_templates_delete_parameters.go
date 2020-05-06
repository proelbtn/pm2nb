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
)

// NewDcimConsolePortTemplatesDeleteParams creates a new DcimConsolePortTemplatesDeleteParams object
// with the default values initialized.
func NewDcimConsolePortTemplatesDeleteParams() *DcimConsolePortTemplatesDeleteParams {
	var ()
	return &DcimConsolePortTemplatesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesDeleteParamsWithTimeout creates a new DcimConsolePortTemplatesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesDeleteParams {
	var ()
	return &DcimConsolePortTemplatesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesDeleteParamsWithContext creates a new DcimConsolePortTemplatesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortTemplatesDeleteParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesDeleteParams {
	var ()
	return &DcimConsolePortTemplatesDeleteParams{

		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesDeleteParamsWithHTTPClient creates a new DcimConsolePortTemplatesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesDeleteParams {
	var ()
	return &DcimConsolePortTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortTemplatesDeleteParams contains all the parameters to send to the API endpoint
for the dcim console port templates delete operation typically these are written to a http.Request
*/
type DcimConsolePortTemplatesDeleteParams struct {

	/*ID
	  A unique integer value identifying this console port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) WithID(id int64) *DcimConsolePortTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates delete params
func (o *DcimConsolePortTemplatesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
