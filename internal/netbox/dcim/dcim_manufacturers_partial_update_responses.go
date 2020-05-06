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

// DcimManufacturersPartialUpdateReader is a Reader for the DcimManufacturersPartialUpdate structure.
type DcimManufacturersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimManufacturersPartialUpdateOK creates a DcimManufacturersPartialUpdateOK with default headers values
func NewDcimManufacturersPartialUpdateOK() *DcimManufacturersPartialUpdateOK {
	return &DcimManufacturersPartialUpdateOK{}
}

/*DcimManufacturersPartialUpdateOK handles this case with default header values.

DcimManufacturersPartialUpdateOK dcim manufacturers partial update o k
*/
type DcimManufacturersPartialUpdateOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/{id}/][%d] dcimManufacturersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}