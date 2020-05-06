// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/proelbtn/proxmox2netbox/models"
)

// IpamVlansPartialUpdateReader is a Reader for the IpamVlansPartialUpdate structure.
type IpamVlansPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamVlansPartialUpdateOK creates a IpamVlansPartialUpdateOK with default headers values
func NewIpamVlansPartialUpdateOK() *IpamVlansPartialUpdateOK {
	return &IpamVlansPartialUpdateOK{}
}

/*IpamVlansPartialUpdateOK handles this case with default header values.

IpamVlansPartialUpdateOK ipam vlans partial update o k
*/
type IpamVlansPartialUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlans/{id}/][%d] ipamVlansPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlansPartialUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
