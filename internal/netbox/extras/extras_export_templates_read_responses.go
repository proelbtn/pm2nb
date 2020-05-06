// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/proelbtn/proxmox2netbox/models"
)

// ExtrasExportTemplatesReadReader is a Reader for the ExtrasExportTemplatesRead structure.
type ExtrasExportTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesReadOK creates a ExtrasExportTemplatesReadOK with default headers values
func NewExtrasExportTemplatesReadOK() *ExtrasExportTemplatesReadOK {
	return &ExtrasExportTemplatesReadOK{}
}

/*ExtrasExportTemplatesReadOK handles this case with default header values.

ExtrasExportTemplatesReadOK extras export templates read o k
*/
type ExtrasExportTemplatesReadOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/export-templates/{id}/][%d] extrasExportTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesReadOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
