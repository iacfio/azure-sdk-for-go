package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// ExpressRouteCircuitPeeringsClient is the network Client
type ExpressRouteCircuitPeeringsClient struct {
	ManagementClient
}

// NewExpressRouteCircuitPeeringsClient creates an instance of the ExpressRouteCircuitPeeringsClient client.
func NewExpressRouteCircuitPeeringsClient(p pipeline.Pipeline) ExpressRouteCircuitPeeringsClient {
	return ExpressRouteCircuitPeeringsClient{NewManagementClient(p)}
}

// CreateOrUpdate the Put Pering operation creates/updates an peering in the specified ExpressRouteCircuits This method
// may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering. peeringParameters is parameters supplied to the create/update
// ExpressRouteCircuit Peering operation
func (client ExpressRouteCircuitPeeringsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering) (*ExpressRouteCircuitPeering, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, circuitName, peeringName, peeringParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitPeering), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitPeeringsClient) createOrUpdatePreparer(resourceGroupName string, circuitName string, peeringName string, peeringParameters ExpressRouteCircuitPeering) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(peeringParameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client ExpressRouteCircuitPeeringsClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitPeering{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete the delete peering operation deletes the specified peering from the ExpressRouteCircuit. This method may poll
// for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering.
func (client ExpressRouteCircuitPeeringsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client ExpressRouteCircuitPeeringsClient) deletePreparer(resourceGroupName string, circuitName string, peeringName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client ExpressRouteCircuitPeeringsClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get the GET peering operation retrieves the specified authorization from the ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the express route circuit.
// peeringName is the name of the peering.
func (client ExpressRouteCircuitPeeringsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (*ExpressRouteCircuitPeering, error) {
	req, err := client.getPreparer(resourceGroupName, circuitName, peeringName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitPeering), err
}

// getPreparer prepares the Get request.
func (client ExpressRouteCircuitPeeringsClient) getPreparer(resourceGroupName string, circuitName string, peeringName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client ExpressRouteCircuitPeeringsClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitPeering{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List the List peering operation retrieves all the peerings in an ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the name of the curcuit.
func (client ExpressRouteCircuitPeeringsClient) List(ctx context.Context, resourceGroupName string, circuitName string) (*ExpressRouteCircuitPeeringListResult, error) {
	req, err := client.listPreparer(resourceGroupName, circuitName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*ExpressRouteCircuitPeeringListResult), err
}

// listPreparer prepares the List request.
func (client ExpressRouteCircuitPeeringsClient) listPreparer(resourceGroupName string, circuitName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client ExpressRouteCircuitPeeringsClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &ExpressRouteCircuitPeeringListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
