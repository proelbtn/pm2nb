// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new circuits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for circuits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CircuitsCircuitTerminationsCreate(params *CircuitsCircuitTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsCreateCreated, error)

	CircuitsCircuitTerminationsDelete(params *CircuitsCircuitTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsDeleteNoContent, error)

	CircuitsCircuitTerminationsList(params *CircuitsCircuitTerminationsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsListOK, error)

	CircuitsCircuitTerminationsPartialUpdate(params *CircuitsCircuitTerminationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsPartialUpdateOK, error)

	CircuitsCircuitTerminationsRead(params *CircuitsCircuitTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsReadOK, error)

	CircuitsCircuitTerminationsUpdate(params *CircuitsCircuitTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsUpdateOK, error)

	CircuitsCircuitTypesCreate(params *CircuitsCircuitTypesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesCreateCreated, error)

	CircuitsCircuitTypesDelete(params *CircuitsCircuitTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesDeleteNoContent, error)

	CircuitsCircuitTypesList(params *CircuitsCircuitTypesListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesListOK, error)

	CircuitsCircuitTypesPartialUpdate(params *CircuitsCircuitTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesPartialUpdateOK, error)

	CircuitsCircuitTypesRead(params *CircuitsCircuitTypesReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesReadOK, error)

	CircuitsCircuitTypesUpdate(params *CircuitsCircuitTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesUpdateOK, error)

	CircuitsCircuitsCreate(params *CircuitsCircuitsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsCreateCreated, error)

	CircuitsCircuitsDelete(params *CircuitsCircuitsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsDeleteNoContent, error)

	CircuitsCircuitsList(params *CircuitsCircuitsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsListOK, error)

	CircuitsCircuitsPartialUpdate(params *CircuitsCircuitsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsPartialUpdateOK, error)

	CircuitsCircuitsRead(params *CircuitsCircuitsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsReadOK, error)

	CircuitsCircuitsUpdate(params *CircuitsCircuitsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsUpdateOK, error)

	CircuitsProvidersCreate(params *CircuitsProvidersCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersCreateCreated, error)

	CircuitsProvidersDelete(params *CircuitsProvidersDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersDeleteNoContent, error)

	CircuitsProvidersGraphs(params *CircuitsProvidersGraphsParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersGraphsOK, error)

	CircuitsProvidersList(params *CircuitsProvidersListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersListOK, error)

	CircuitsProvidersPartialUpdate(params *CircuitsProvidersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersPartialUpdateOK, error)

	CircuitsProvidersRead(params *CircuitsProvidersReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersReadOK, error)

	CircuitsProvidersUpdate(params *CircuitsProvidersUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CircuitsCircuitTerminationsCreate circuits circuit terminations create API
*/
func (a *Client) CircuitsCircuitTerminationsCreate(params *CircuitsCircuitTerminationsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsDelete circuits circuit terminations delete API
*/
func (a *Client) CircuitsCircuitTerminationsDelete(params *CircuitsCircuitTerminationsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsList Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitTerminationsList(params *CircuitsCircuitTerminationsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsPartialUpdate circuits circuit terminations partial update API
*/
func (a *Client) CircuitsCircuitTerminationsPartialUpdate(params *CircuitsCircuitTerminationsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsRead Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitTerminationsRead(params *CircuitsCircuitTerminationsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTerminationsUpdate circuits circuit terminations update API
*/
func (a *Client) CircuitsCircuitTerminationsUpdate(params *CircuitsCircuitTerminationsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTerminationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTerminationsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-terminations_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-terminations/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTerminationsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTerminationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-terminations_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesCreate circuits circuit types create API
*/
func (a *Client) CircuitsCircuitTypesCreate(params *CircuitsCircuitTypesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesDelete circuits circuit types delete API
*/
func (a *Client) CircuitsCircuitTypesDelete(params *CircuitsCircuitTypesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesList Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitTypesList(params *CircuitsCircuitTypesListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesPartialUpdate circuits circuit types partial update API
*/
func (a *Client) CircuitsCircuitTypesPartialUpdate(params *CircuitsCircuitTypesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesRead Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitTypesRead(params *CircuitsCircuitTypesReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitTypesUpdate circuits circuit types update API
*/
func (a *Client) CircuitsCircuitTypesUpdate(params *CircuitsCircuitTypesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitTypesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitTypesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuit-types_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuit-types/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitTypesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitTypesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuit-types_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsCreate circuits circuits create API
*/
func (a *Client) CircuitsCircuitsCreate(params *CircuitsCircuitsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_create",
		Method:             "POST",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsDelete circuits circuits delete API
*/
func (a *Client) CircuitsCircuitsDelete(params *CircuitsCircuitsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsList Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitsList(params *CircuitsCircuitsListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_list",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsPartialUpdate circuits circuits partial update API
*/
func (a *Client) CircuitsCircuitsPartialUpdate(params *CircuitsCircuitsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsRead Call to super to allow for caching
*/
func (a *Client) CircuitsCircuitsRead(params *CircuitsCircuitsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_read",
		Method:             "GET",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsCircuitsUpdate circuits circuits update API
*/
func (a *Client) CircuitsCircuitsUpdate(params *CircuitsCircuitsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsCircuitsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsCircuitsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_circuits_update",
		Method:             "PUT",
		PathPattern:        "/circuits/circuits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsCircuitsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsCircuitsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_circuits_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersCreate circuits providers create API
*/
func (a *Client) CircuitsProvidersCreate(params *CircuitsProvidersCreateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_create",
		Method:             "POST",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersDelete circuits providers delete API
*/
func (a *Client) CircuitsProvidersDelete(params *CircuitsProvidersDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_delete",
		Method:             "DELETE",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersGraphs A convenience method for rendering graphs for a particular provider.
*/
func (a *Client) CircuitsProvidersGraphs(params *CircuitsProvidersGraphsParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersGraphsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersGraphsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_graphs",
		Method:             "GET",
		PathPattern:        "/circuits/providers/{id}/graphs/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersGraphsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersGraphsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_graphs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersList Call to super to allow for caching
*/
func (a *Client) CircuitsProvidersList(params *CircuitsProvidersListParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_list",
		Method:             "GET",
		PathPattern:        "/circuits/providers/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersPartialUpdate circuits providers partial update API
*/
func (a *Client) CircuitsProvidersPartialUpdate(params *CircuitsProvidersPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_partial_update",
		Method:             "PATCH",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersRead Call to super to allow for caching
*/
func (a *Client) CircuitsProvidersRead(params *CircuitsProvidersReadParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_read",
		Method:             "GET",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CircuitsProvidersUpdate circuits providers update API
*/
func (a *Client) CircuitsProvidersUpdate(params *CircuitsProvidersUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*CircuitsProvidersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCircuitsProvidersUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "circuits_providers_update",
		Method:             "PUT",
		PathPattern:        "/circuits/providers/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CircuitsProvidersUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CircuitsProvidersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for circuits_providers_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
