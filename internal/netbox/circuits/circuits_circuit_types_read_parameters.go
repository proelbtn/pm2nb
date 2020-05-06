// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// NewCircuitsCircuitTypesReadParams creates a new CircuitsCircuitTypesReadParams object
// with the default values initialized.
func NewCircuitsCircuitTypesReadParams() *CircuitsCircuitTypesReadParams {
	var ()
	return &CircuitsCircuitTypesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesReadParamsWithTimeout creates a new CircuitsCircuitTypesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesReadParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesReadParams {
	var ()
	return &CircuitsCircuitTypesReadParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesReadParamsWithContext creates a new CircuitsCircuitTypesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesReadParamsWithContext(ctx context.Context) *CircuitsCircuitTypesReadParams {
	var ()
	return &CircuitsCircuitTypesReadParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesReadParamsWithHTTPClient creates a new CircuitsCircuitTypesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesReadParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesReadParams {
	var ()
	return &CircuitsCircuitTypesReadParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesReadParams contains all the parameters to send to the API endpoint
for the circuits circuit types read operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesReadParams struct {

	/*ID
	  A unique integer value identifying this circuit type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithContext(ctx context.Context) *CircuitsCircuitTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) WithID(id int64) *CircuitsCircuitTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types read params
func (o *CircuitsCircuitTypesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
