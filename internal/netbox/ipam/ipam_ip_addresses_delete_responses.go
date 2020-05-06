// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamIPAddressesDeleteReader is a Reader for the IpamIPAddressesDelete structure.
type IpamIPAddressesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPAddressesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamIPAddressesDeleteNoContent creates a IpamIPAddressesDeleteNoContent with default headers values
func NewIpamIPAddressesDeleteNoContent() *IpamIPAddressesDeleteNoContent {
	return &IpamIPAddressesDeleteNoContent{}
}

/*IpamIPAddressesDeleteNoContent handles this case with default header values.

IpamIPAddressesDeleteNoContent ipam Ip addresses delete no content
*/
type IpamIPAddressesDeleteNoContent struct {
}

func (o *IpamIPAddressesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-addresses/{id}/][%d] ipamIpAddressesDeleteNoContent ", 204)
}

func (o *IpamIPAddressesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
