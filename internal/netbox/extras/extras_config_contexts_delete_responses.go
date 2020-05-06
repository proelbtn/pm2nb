// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextsDeleteReader is a Reader for the ExtrasConfigContextsDelete structure.
type ExtrasConfigContextsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasConfigContextsDeleteNoContent creates a ExtrasConfigContextsDeleteNoContent with default headers values
func NewExtrasConfigContextsDeleteNoContent() *ExtrasConfigContextsDeleteNoContent {
	return &ExtrasConfigContextsDeleteNoContent{}
}

/*ExtrasConfigContextsDeleteNoContent handles this case with default header values.

ExtrasConfigContextsDeleteNoContent extras config contexts delete no content
*/
type ExtrasConfigContextsDeleteNoContent struct {
}

func (o *ExtrasConfigContextsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/{id}/][%d] extrasConfigContextsDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}