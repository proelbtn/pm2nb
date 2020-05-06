// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/proelbtn/proxmox2netbox/models"
)

// DcimPowerOutletTemplatesCreateReader is a Reader for the DcimPowerOutletTemplatesCreate structure.
type DcimPowerOutletTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerOutletTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesCreateCreated creates a DcimPowerOutletTemplatesCreateCreated with default headers values
func NewDcimPowerOutletTemplatesCreateCreated() *DcimPowerOutletTemplatesCreateCreated {
	return &DcimPowerOutletTemplatesCreateCreated{}
}

/*DcimPowerOutletTemplatesCreateCreated handles this case with default header values.

DcimPowerOutletTemplatesCreateCreated dcim power outlet templates create created
*/
type DcimPowerOutletTemplatesCreateCreated struct {
	Payload *models.PowerOutletTemplate
}

func (o *DcimPowerOutletTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-outlet-templates/][%d] dcimPowerOutletTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerOutletTemplatesCreateCreated) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}