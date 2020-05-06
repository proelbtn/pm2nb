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

// DcimRackGroupsUpdateReader is a Reader for the DcimRackGroupsUpdate structure.
type DcimRackGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackGroupsUpdateOK creates a DcimRackGroupsUpdateOK with default headers values
func NewDcimRackGroupsUpdateOK() *DcimRackGroupsUpdateOK {
	return &DcimRackGroupsUpdateOK{}
}

/*DcimRackGroupsUpdateOK handles this case with default header values.

DcimRackGroupsUpdateOK dcim rack groups update o k
*/
type DcimRackGroupsUpdateOK struct {
	Payload *models.RackGroup
}

func (o *DcimRackGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-groups/{id}/][%d] dcimRackGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackGroupsUpdateOK) GetPayload() *models.RackGroup {
	return o.Payload
}

func (o *DcimRackGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
