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

// DcimPowerPortsReadReader is a Reader for the DcimPowerPortsRead structure.
type DcimPowerPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortsReadOK creates a DcimPowerPortsReadOK with default headers values
func NewDcimPowerPortsReadOK() *DcimPowerPortsReadOK {
	return &DcimPowerPortsReadOK{}
}

/*DcimPowerPortsReadOK handles this case with default header values.

DcimPowerPortsReadOK dcim power ports read o k
*/
type DcimPowerPortsReadOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/][%d] dcimPowerPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsReadOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
