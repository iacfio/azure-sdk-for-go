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

// LocalNetworkGatewaysClient is the network Client
type LocalNetworkGatewaysClient struct {
	ManagementClient
}

// NewLocalNetworkGatewaysClient creates an instance of the LocalNetworkGatewaysClient client.
func NewLocalNetworkGatewaysClient(p pipeline.Pipeline) LocalNetworkGatewaysClient {
	return LocalNetworkGatewaysClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a local network gateway in the specified resource group. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway. parameters is parameters supplied to the create or update local network gateway operation.
func (client LocalNetworkGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (*LocalNetworkGateway, error) {
	if err := validate([]validation{
		{targetValue: localNetworkGatewayName,
			constraints: []constraint{{target: "localNetworkGatewayName", name: minLength, rule: 1, chain: nil}}},
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.LocalNetworkGatewayPropertiesFormat", name: null, rule: true, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LocalNetworkGateway), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) createOrUpdatePreparer(resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-09-01")
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

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client LocalNetworkGatewaysClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LocalNetworkGateway{rawResponse: resp.Response()}
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

// Delete deletes the specified local network gateway. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway.
func (client LocalNetworkGatewaysClient) Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*http.Response, error) {
	if err := validate([]validation{
		{targetValue: localNetworkGatewayName,
			constraints: []constraint{{target: "localNetworkGatewayName", name: minLength, rule: 1, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.deletePreparer(resourceGroupName, localNetworkGatewayName)
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
func (client LocalNetworkGatewaysClient) deletePreparer(resourceGroupName string, localNetworkGatewayName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-09-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client LocalNetworkGatewaysClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get gets the specified local network gateway in a resource group.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway.
func (client LocalNetworkGatewaysClient) Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (*LocalNetworkGateway, error) {
	if err := validate([]validation{
		{targetValue: localNetworkGatewayName,
			constraints: []constraint{{target: "localNetworkGatewayName", name: minLength, rule: 1, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.getPreparer(resourceGroupName, localNetworkGatewayName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LocalNetworkGateway), err
}

// getPreparer prepares the Get request.
func (client LocalNetworkGatewaysClient) getPreparer(resourceGroupName string, localNetworkGatewayName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-09-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client LocalNetworkGatewaysClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LocalNetworkGateway{rawResponse: resp.Response()}
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

// List gets all the local network gateways in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client LocalNetworkGatewaysClient) List(ctx context.Context, resourceGroupName string) (*LocalNetworkGatewayListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LocalNetworkGatewayListResult), err
}

// listPreparer prepares the List request.
func (client LocalNetworkGatewaysClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-09-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client LocalNetworkGatewaysClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LocalNetworkGatewayListResult{rawResponse: resp.Response()}
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

// UpdateTags updates a local network gateway tags. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. localNetworkGatewayName is the name of the local network
// gateway. parameters is parameters supplied to update local network gateway tags.
func (client LocalNetworkGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (*LocalNetworkGateway, error) {
	if err := validate([]validation{
		{targetValue: localNetworkGatewayName,
			constraints: []constraint{{target: "localNetworkGatewayName", name: minLength, rule: 1, chain: nil}}}}); err != nil {
		return nil, err
	}
	req, err := client.updateTagsPreparer(resourceGroupName, localNetworkGatewayName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateTagsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*LocalNetworkGateway), err
}

// updateTagsPreparer prepares the UpdateTags request.
func (client LocalNetworkGatewaysClient) updateTagsPreparer(resourceGroupName string, localNetworkGatewayName string, parameters TagsObject) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-09-01")
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

// updateTagsResponder handles the response to the UpdateTags request.
func (client LocalNetworkGatewaysClient) updateTagsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &LocalNetworkGateway{rawResponse: resp.Response()}
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
