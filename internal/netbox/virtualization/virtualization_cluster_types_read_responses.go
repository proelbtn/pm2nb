// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/proelbtn/proxmox2netbox/models"
)

// VirtualizationClusterTypesReadReader is a Reader for the VirtualizationClusterTypesRead structure.
type VirtualizationClusterTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterTypesReadOK creates a VirtualizationClusterTypesReadOK with default headers values
func NewVirtualizationClusterTypesReadOK() *VirtualizationClusterTypesReadOK {
	return &VirtualizationClusterTypesReadOK{}
}

/*VirtualizationClusterTypesReadOK handles this case with default header values.

VirtualizationClusterTypesReadOK virtualization cluster types read o k
*/
type VirtualizationClusterTypesReadOK struct {
	Payload *models.ClusterType
}

func (o *VirtualizationClusterTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-types/{id}/][%d] virtualizationClusterTypesReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesReadOK) GetPayload() *models.ClusterType {
	return o.Payload
}

func (o *VirtualizationClusterTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
