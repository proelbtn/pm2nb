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

// DcimPowerPortsTraceReader is a Reader for the DcimPowerPortsTrace structure.
type DcimPowerPortsTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortsTraceOK creates a DcimPowerPortsTraceOK with default headers values
func NewDcimPowerPortsTraceOK() *DcimPowerPortsTraceOK {
	return &DcimPowerPortsTraceOK{}
}

/*DcimPowerPortsTraceOK handles this case with default header values.

DcimPowerPortsTraceOK dcim power ports trace o k
*/
type DcimPowerPortsTraceOK struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-ports/{id}/trace/][%d] dcimPowerPortsTraceOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsTraceOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
