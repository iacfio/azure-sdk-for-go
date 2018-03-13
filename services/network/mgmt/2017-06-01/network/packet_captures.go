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

// PacketCapturesClient is the network Client
type PacketCapturesClient struct {
	ManagementClient
}

// NewPacketCapturesClient creates an instance of the PacketCapturesClient client.
func NewPacketCapturesClient(p pipeline.Pipeline) PacketCapturesClient {
	return PacketCapturesClient{NewManagementClient(p)}
}

// Create create and start a packet capture on the specified VM. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// packetCaptureName is the name of the packet capture session. parameters is parameters that define the create packet
// capture operation.
func (client PacketCapturesClient) Create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture) (*PacketCaptureResult, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.PacketCaptureParameters", name: null, rule: true,
				chain: []constraint{{target: "parameters.PacketCaptureParameters.Target", name: null, rule: true, chain: nil},
					{target: "parameters.PacketCaptureParameters.StorageLocation", name: null, rule: true, chain: nil},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createPreparer(resourceGroupName, networkWatcherName, packetCaptureName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PacketCaptureResult), err
}

// createPreparer prepares the Create request.
func (client PacketCapturesClient) createPreparer(resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
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

// createResponder handles the response to the Create request.
func (client PacketCapturesClient) createResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	result := &PacketCaptureResult{rawResponse: resp.Response()}
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

// Delete deletes the specified packet capture session. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// packetCaptureName is the name of the packet capture session.
func (client PacketCapturesClient) Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, networkWatcherName, packetCaptureName)
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
func (client PacketCapturesClient) deletePreparer(resourceGroupName string, networkWatcherName string, packetCaptureName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client PacketCapturesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusNoContent, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get gets a packet capture session by name.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// packetCaptureName is the name of the packet capture session.
func (client PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*PacketCaptureResult, error) {
	req, err := client.getPreparer(resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PacketCaptureResult), err
}

// getPreparer prepares the Get request.
func (client PacketCapturesClient) getPreparer(resourceGroupName string, networkWatcherName string, packetCaptureName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client PacketCapturesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PacketCaptureResult{rawResponse: resp.Response()}
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

// GetStatus query the status of a running packet capture session. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the Network Watcher resource.
// packetCaptureName is the name given to the packet capture session.
func (client PacketCapturesClient) GetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*PacketCaptureQueryStatusResult, error) {
	req, err := client.getStatusPreparer(resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getStatusResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PacketCaptureQueryStatusResult), err
}

// getStatusPreparer prepares the GetStatus request.
func (client PacketCapturesClient) getStatusPreparer(resourceGroupName string, networkWatcherName string, packetCaptureName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getStatusResponder handles the response to the GetStatus request.
func (client PacketCapturesClient) getStatusResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &PacketCaptureQueryStatusResult{rawResponse: resp.Response()}
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

// List lists all packet capture sessions within the specified resource group.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the Network Watcher resource.
func (client PacketCapturesClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string) (*PacketCaptureListResult, error) {
	req, err := client.listPreparer(resourceGroupName, networkWatcherName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PacketCaptureListResult), err
}

// listPreparer prepares the List request.
func (client PacketCapturesClient) listPreparer(resourceGroupName string, networkWatcherName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client PacketCapturesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PacketCaptureListResult{rawResponse: resp.Response()}
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

// Stop stops a specified packet capture session. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkWatcherName is the name of the network watcher.
// packetCaptureName is the name of the packet capture session.
func (client PacketCapturesClient) Stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Response, error) {
	req, err := client.stopPreparer(resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.stopResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// stopPreparer prepares the Stop request.
func (client PacketCapturesClient) stopPreparer(resourceGroupName string, networkWatcherName string, packetCaptureName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-06-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// stopResponder handles the response to the Stop request.
func (client PacketCapturesClient) stopResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}
