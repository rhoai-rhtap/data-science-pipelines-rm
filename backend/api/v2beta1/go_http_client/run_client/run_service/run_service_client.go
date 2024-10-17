// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new run service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for run service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RunServiceArchiveRun archives a run in an experiment given by run ID and experiment ID
*/
func (a *Client) RunServiceArchiveRun(params *RunServiceArchiveRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceArchiveRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceArchiveRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_ArchiveRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/runs/{run_id}:archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceArchiveRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceArchiveRunOK), nil

}

/*
RunServiceCreateRun creates a new run in an experiment specified by experiment ID if experiment ID is not specified the run is created in the default experiment
*/
func (a *Client) RunServiceCreateRun(params *RunServiceCreateRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceCreateRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceCreateRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_CreateRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceCreateRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceCreateRunOK), nil

}

/*
RunServiceDeleteRun deletes a run in an experiment given by run ID and experiment ID
*/
func (a *Client) RunServiceDeleteRun(params *RunServiceDeleteRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceDeleteRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceDeleteRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_DeleteRun",
		Method:             "DELETE",
		PathPattern:        "/apis/v2beta1/runs/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceDeleteRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceDeleteRunOK), nil

}

/*
RunServiceGetRun finds a specific run by ID
*/
func (a *Client) RunServiceGetRun(params *RunServiceGetRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceGetRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceGetRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_GetRun",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/runs/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceGetRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceGetRunOK), nil

}

/*
RunServiceListRuns finds all runs in an experiment given by experiment ID if experiment id is not specified finds all runs across all experiments
*/
func (a *Client) RunServiceListRuns(params *RunServiceListRunsParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceListRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceListRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_ListRuns",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceListRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceListRunsOK), nil

}

/*
RunServiceReadArtifact finds artifact data in a run
*/
func (a *Client) RunServiceReadArtifact(params *RunServiceReadArtifactParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceReadArtifactOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceReadArtifactParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_ReadArtifact",
		Method:             "GET",
		PathPattern:        "/apis/v2beta1/runs/{run_id}/nodes/{node_id}/artifacts/{artifact_name}:read",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceReadArtifactReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceReadArtifactOK), nil

}

/*
RunServiceRetryRun res initiates a failed or terminated run
*/
func (a *Client) RunServiceRetryRun(params *RunServiceRetryRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceRetryRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceRetryRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_RetryRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/runs/{run_id}:retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceRetryRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceRetryRunOK), nil

}

/*
RunServiceTerminateRun terminates an active run
*/
func (a *Client) RunServiceTerminateRun(params *RunServiceTerminateRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceTerminateRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceTerminateRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_TerminateRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/runs/{run_id}:terminate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceTerminateRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceTerminateRunOK), nil

}

/*
RunServiceUnarchiveRun restores an archived run in an experiment given by run ID and experiment ID
*/
func (a *Client) RunServiceUnarchiveRun(params *RunServiceUnarchiveRunParams, authInfo runtime.ClientAuthInfoWriter) (*RunServiceUnarchiveRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunServiceUnarchiveRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RunService_UnarchiveRun",
		Method:             "POST",
		PathPattern:        "/apis/v2beta1/runs/{run_id}:unarchive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RunServiceUnarchiveRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RunServiceUnarchiveRunOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
