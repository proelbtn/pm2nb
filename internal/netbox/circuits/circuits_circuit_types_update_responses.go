// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/proelbtn/proxmox2netbox/models"
)

// CircuitsCircuitTypesUpdateReader is a Reader for the CircuitsCircuitTypesUpdate structure.
type CircuitsCircuitTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitTypesUpdateOK creates a CircuitsCircuitTypesUpdateOK with default headers values
func NewCircuitsCircuitTypesUpdateOK() *CircuitsCircuitTypesUpdateOK {
	return &CircuitsCircuitTypesUpdateOK{}
}

/*CircuitsCircuitTypesUpdateOK handles this case with default header values.

CircuitsCircuitTypesUpdateOK circuits circuit types update o k
*/
type CircuitsCircuitTypesUpdateOK struct {
	Payload *models.CircuitType
}

func (o *CircuitsCircuitTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
