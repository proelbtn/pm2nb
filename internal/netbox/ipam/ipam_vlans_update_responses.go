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

// IpamVlansUpdateReader is a Reader for the IpamVlansUpdate structure.
type IpamVlansUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlansUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamVlansUpdateOK creates a IpamVlansUpdateOK with default headers values
func NewIpamVlansUpdateOK() *IpamVlansUpdateOK {
	return &IpamVlansUpdateOK{}
}

/*IpamVlansUpdateOK handles this case with default header values.

IpamVlansUpdateOK ipam vlans update o k
*/
type IpamVlansUpdateOK struct {
	Payload *models.VLAN
}

func (o *IpamVlansUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlans/{id}/][%d] ipamVlansUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlansUpdateOK) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}