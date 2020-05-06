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

// IpamVlanGroupsUpdateReader is a Reader for the IpamVlanGroupsUpdate structure.
type IpamVlanGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamVlanGroupsUpdateOK creates a IpamVlanGroupsUpdateOK with default headers values
func NewIpamVlanGroupsUpdateOK() *IpamVlanGroupsUpdateOK {
	return &IpamVlanGroupsUpdateOK{}
}

/*IpamVlanGroupsUpdateOK handles this case with default header values.

IpamVlanGroupsUpdateOK ipam vlan groups update o k
*/
type IpamVlanGroupsUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
