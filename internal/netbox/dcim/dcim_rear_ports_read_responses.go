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

// DcimRearPortsReadReader is a Reader for the DcimRearPortsRead structure.
type DcimRearPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRearPortsReadOK creates a DcimRearPortsReadOK with default headers values
func NewDcimRearPortsReadOK() *DcimRearPortsReadOK {
	return &DcimRearPortsReadOK{}
}

/*DcimRearPortsReadOK handles this case with default header values.

DcimRearPortsReadOK dcim rear ports read o k
*/
type DcimRearPortsReadOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcimRearPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsReadOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
