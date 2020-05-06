// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceTypesDeleteReader is a Reader for the DcimDeviceTypesDelete structure.
type DcimDeviceTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceTypesDeleteNoContent creates a DcimDeviceTypesDeleteNoContent with default headers values
func NewDcimDeviceTypesDeleteNoContent() *DcimDeviceTypesDeleteNoContent {
	return &DcimDeviceTypesDeleteNoContent{}
}

/*DcimDeviceTypesDeleteNoContent handles this case with default header values.

DcimDeviceTypesDeleteNoContent dcim device types delete no content
*/
type DcimDeviceTypesDeleteNoContent struct {
}

func (o *DcimDeviceTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/{id}/][%d] dcimDeviceTypesDeleteNoContent ", 204)
}

func (o *DcimDeviceTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
