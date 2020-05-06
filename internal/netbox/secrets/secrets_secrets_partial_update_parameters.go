// Code generated by go-swagger; DO NOT EDIT.

package secrets

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

// NewSecretsSecretsPartialUpdateParams creates a new SecretsSecretsPartialUpdateParams object
// with the default values initialized.
func NewSecretsSecretsPartialUpdateParams() *SecretsSecretsPartialUpdateParams {
	var ()
	return &SecretsSecretsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretsPartialUpdateParamsWithTimeout creates a new SecretsSecretsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretsPartialUpdateParamsWithTimeout(timeout time.Duration) *SecretsSecretsPartialUpdateParams {
	var ()
	return &SecretsSecretsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretsPartialUpdateParamsWithContext creates a new SecretsSecretsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretsPartialUpdateParamsWithContext(ctx context.Context) *SecretsSecretsPartialUpdateParams {
	var ()
	return &SecretsSecretsPartialUpdateParams{

		Context: ctx,
	}
}

// NewSecretsSecretsPartialUpdateParamsWithHTTPClient creates a new SecretsSecretsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretsPartialUpdateParamsWithHTTPClient(client *http.Client) *SecretsSecretsPartialUpdateParams {
	var ()
	return &SecretsSecretsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretsPartialUpdateParams contains all the parameters to send to the API endpoint
for the secrets secrets partial update operation typically these are written to a http.Request
*/
type SecretsSecretsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableSecret
	/*ID
	  A unique integer value identifying this secret.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) WithTimeout(timeout time.Duration) *SecretsSecretsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) WithContext(ctx context.Context) *SecretsSecretsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) WithHTTPClient(client *http.Client) *SecretsSecretsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) WithData(data *models.WritableSecret) *SecretsSecretsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) SetData(data *models.WritableSecret) {
	o.Data = data
}

// WithID adds the id to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) WithID(id int64) *SecretsSecretsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secrets partial update params
func (o *SecretsSecretsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
