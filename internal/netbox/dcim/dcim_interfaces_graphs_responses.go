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

// DcimInterfacesGraphsReader is a Reader for the DcimInterfacesGraphs structure.
type DcimInterfacesGraphsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesGraphsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesGraphsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfacesGraphsOK creates a DcimInterfacesGraphsOK with default headers values
func NewDcimInterfacesGraphsOK() *DcimInterfacesGraphsOK {
	return &DcimInterfacesGraphsOK{}
}

/*DcimInterfacesGraphsOK handles this case with default header values.

DcimInterfacesGraphsOK dcim interfaces graphs o k
*/
type DcimInterfacesGraphsOK struct {
	Payload *models.DeviceInterface
}

func (o *DcimInterfacesGraphsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/graphs/][%d] dcimInterfacesGraphsOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesGraphsOK) GetPayload() *models.DeviceInterface {
	return o.Payload
}

func (o *DcimInterfacesGraphsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
