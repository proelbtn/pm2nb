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

// DcimConsoleServerPortTemplatesUpdateReader is a Reader for the DcimConsoleServerPortTemplatesUpdate structure.
type DcimConsoleServerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsoleServerPortTemplatesUpdateOK creates a DcimConsoleServerPortTemplatesUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesUpdateOK() *DcimConsoleServerPortTemplatesUpdateOK {
	return &DcimConsoleServerPortTemplatesUpdateOK{}
}

/*DcimConsoleServerPortTemplatesUpdateOK handles this case with default header values.

DcimConsoleServerPortTemplatesUpdateOK dcim console server port templates update o k
*/
type DcimConsoleServerPortTemplatesUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
