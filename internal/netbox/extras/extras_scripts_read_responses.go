// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasScriptsReadReader is a Reader for the ExtrasScriptsRead structure.
type ExtrasScriptsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasScriptsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasScriptsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasScriptsReadOK creates a ExtrasScriptsReadOK with default headers values
func NewExtrasScriptsReadOK() *ExtrasScriptsReadOK {
	return &ExtrasScriptsReadOK{}
}

/*ExtrasScriptsReadOK handles this case with default header values.

ExtrasScriptsReadOK extras scripts read o k
*/
type ExtrasScriptsReadOK struct {
}

func (o *ExtrasScriptsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/{id}/][%d] extrasScriptsReadOK ", 200)
}

func (o *ExtrasScriptsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
