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
)

// NewSecretsSecretRolesListParams creates a new SecretsSecretRolesListParams object
// with the default values initialized.
func NewSecretsSecretRolesListParams() *SecretsSecretRolesListParams {
	var ()
	return &SecretsSecretRolesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesListParamsWithTimeout creates a new SecretsSecretRolesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretRolesListParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesListParams {
	var ()
	return &SecretsSecretRolesListParams{

		timeout: timeout,
	}
}

// NewSecretsSecretRolesListParamsWithContext creates a new SecretsSecretRolesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretRolesListParamsWithContext(ctx context.Context) *SecretsSecretRolesListParams {
	var ()
	return &SecretsSecretRolesListParams{

		Context: ctx,
	}
}

// NewSecretsSecretRolesListParamsWithHTTPClient creates a new SecretsSecretRolesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretRolesListParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesListParams {
	var ()
	return &SecretsSecretRolesListParams{
		HTTPClient: client,
	}
}

/*SecretsSecretRolesListParams contains all the parameters to send to the API endpoint
for the secrets secret roles list operation typically these are written to a http.Request
*/
type SecretsSecretRolesListParams struct {

	/*ID*/
	ID *string
	/*IDGt*/
	IDGt *string
	/*IDGte*/
	IDGte *string
	/*IDLt*/
	IDLt *string
	/*IDLte*/
	IDLte *string
	/*IDn*/
	IDn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string
	/*SlugIc*/
	SlugIc *string
	/*SlugIe*/
	SlugIe *string
	/*SlugIew*/
	SlugIew *string
	/*SlugIsw*/
	SlugIsw *string
	/*Slugn*/
	Slugn *string
	/*SlugNic*/
	SlugNic *string
	/*SlugNie*/
	SlugNie *string
	/*SlugNiew*/
	SlugNiew *string
	/*SlugNisw*/
	SlugNisw *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithContext(ctx context.Context) *SecretsSecretRolesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithID(id *string) *SecretsSecretRolesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithIDGt(iDGt *string) *SecretsSecretRolesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithIDGte(iDGte *string) *SecretsSecretRolesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithIDLt(iDLt *string) *SecretsSecretRolesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithIDLte(iDLte *string) *SecretsSecretRolesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithIDn(iDn *string) *SecretsSecretRolesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithLimit(limit *int64) *SecretsSecretRolesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithName(name *string) *SecretsSecretRolesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameIc(nameIc *string) *SecretsSecretRolesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameIe(nameIe *string) *SecretsSecretRolesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameIew(nameIew *string) *SecretsSecretRolesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameIsw(nameIsw *string) *SecretsSecretRolesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNamen(namen *string) *SecretsSecretRolesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameNic(nameNic *string) *SecretsSecretRolesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameNie(nameNie *string) *SecretsSecretRolesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameNiew(nameNiew *string) *SecretsSecretRolesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithNameNisw(nameNisw *string) *SecretsSecretRolesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithOffset(offset *int64) *SecretsSecretRolesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithQ(q *string) *SecretsSecretRolesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlug(slug *string) *SecretsSecretRolesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugIc(slugIc *string) *SecretsSecretRolesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugIe(slugIe *string) *SecretsSecretRolesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugIew(slugIew *string) *SecretsSecretRolesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugIsw(slugIsw *string) *SecretsSecretRolesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugn(slugn *string) *SecretsSecretRolesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugNic(slugNic *string) *SecretsSecretRolesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugNie(slugNie *string) *SecretsSecretRolesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugNiew(slugNiew *string) *SecretsSecretRolesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) WithSlugNisw(slugNisw *string) *SecretsSecretRolesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the secrets secret roles list params
func (o *SecretsSecretRolesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string
		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {
			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}

	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string
		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {
			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}

	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string
		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {
			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}

	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string
		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {
			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}

	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string
		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {
			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string
		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {
			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}

	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string
		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {
			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}

	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string
		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {
			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}

	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string
		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {
			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}

	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string
		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {
			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}

	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string
		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {
			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}

	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string
		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {
			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}

	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string
		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {
			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}

	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string
		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {
			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
