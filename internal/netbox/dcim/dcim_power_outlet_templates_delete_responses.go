// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerOutletTemplatesDeleteReader is a Reader for the DcimPowerOutletTemplatesDelete structure.
type DcimPowerOutletTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerOutletTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerOutletTemplatesDeleteNoContent creates a DcimPowerOutletTemplatesDeleteNoContent with default headers values
func NewDcimPowerOutletTemplatesDeleteNoContent() *DcimPowerOutletTemplatesDeleteNoContent {
	return &DcimPowerOutletTemplatesDeleteNoContent{}
}

/*DcimPowerOutletTemplatesDeleteNoContent handles this case with default header values.

DcimPowerOutletTemplatesDeleteNoContent dcim power outlet templates delete no content
*/
type DcimPowerOutletTemplatesDeleteNoContent struct {
}

func (o *DcimPowerOutletTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesDeleteNoContent ", 204)
}

func (o *DcimPowerOutletTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
