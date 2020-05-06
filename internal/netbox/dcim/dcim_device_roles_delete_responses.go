// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceRolesDeleteReader is a Reader for the DcimDeviceRolesDelete structure.
type DcimDeviceRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceRolesDeleteNoContent creates a DcimDeviceRolesDeleteNoContent with default headers values
func NewDcimDeviceRolesDeleteNoContent() *DcimDeviceRolesDeleteNoContent {
	return &DcimDeviceRolesDeleteNoContent{}
}

/*DcimDeviceRolesDeleteNoContent handles this case with default header values.

DcimDeviceRolesDeleteNoContent dcim device roles delete no content
*/
type DcimDeviceRolesDeleteNoContent struct {
}

func (o *DcimDeviceRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcimDeviceRolesDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
