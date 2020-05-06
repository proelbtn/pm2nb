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

// IpamVlanGroupsPartialUpdateReader is a Reader for the IpamVlanGroupsPartialUpdate structure.
type IpamVlanGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamVlanGroupsPartialUpdateOK creates a IpamVlanGroupsPartialUpdateOK with default headers values
func NewIpamVlanGroupsPartialUpdateOK() *IpamVlanGroupsPartialUpdateOK {
	return &IpamVlanGroupsPartialUpdateOK{}
}

/*IpamVlanGroupsPartialUpdateOK handles this case with default header values.

IpamVlanGroupsPartialUpdateOK ipam vlan groups partial update o k
*/
type IpamVlanGroupsPartialUpdateOK struct {
	Payload *models.VLANGroup
}

func (o *IpamVlanGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsPartialUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
