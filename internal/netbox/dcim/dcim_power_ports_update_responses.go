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

// DcimPowerPortsUpdateReader is a Reader for the DcimPowerPortsUpdate structure.
type DcimPowerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortsUpdateOK creates a DcimPowerPortsUpdateOK with default headers values
func NewDcimPowerPortsUpdateOK() *DcimPowerPortsUpdateOK {
	return &DcimPowerPortsUpdateOK{}
}

/*DcimPowerPortsUpdateOK handles this case with default header values.

DcimPowerPortsUpdateOK dcim power ports update o k
*/
type DcimPowerPortsUpdateOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/{id}/][%d] dcimPowerPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
