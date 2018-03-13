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

// InterfacesClient is the network Client
type InterfacesClient struct {
	ManagementClient
}

// NewInterfacesClient creates an instance of the InterfacesClient client.
func NewInterfacesClient(p pipeline.Pipeline) InterfacesClient {
	return InterfacesClient{NewManagementClient(p)}
}

// CreateOrUpdate the Put NetworkInterface operation creates/updates a networkInterface This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
// parameters is parameters supplied to the create/update NetworkInterface operation
func (client InterfacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters Interface) (*Interface, error) {
	req, err := client.createOrUpdatePreparer(resourceGroupName, networkInterfaceName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Interface), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InterfacesClient) createOrUpdatePreparer(resourceGroupName string, networkInterfaceName string, parameters Interface) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
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
func (client InterfacesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Interface{rawResponse: resp.Response()}
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

// Delete the delete netwokInterface operation deletes the specified netwokInterface. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfacesClient) Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, networkInterfaceName)
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
func (client InterfacesClient) deletePreparer(resourceGroupName string, networkInterfaceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"
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
func (client InterfacesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get the Get network interface operation retrieves information about the specified network interface.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
// expand is expand references resources.
func (client InterfacesClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand *string) (*Interface, error) {
	req, err := client.getPreparer(resourceGroupName, networkInterfaceName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Interface), err
}

// getPreparer prepares the Get request.
func (client InterfacesClient) getPreparer(resourceGroupName string, networkInterfaceName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client InterfacesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Interface{rawResponse: resp.Response()}
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

// GetEffectiveRouteTable the get effective routetable operation retrieves all the route tables applied on a
// networkInterface. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfacesClient) GetEffectiveRouteTable(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*EffectiveRouteListResult, error) {
	req, err := client.getEffectiveRouteTablePreparer(resourceGroupName, networkInterfaceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getEffectiveRouteTableResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*EffectiveRouteListResult), err
}

// getEffectiveRouteTablePreparer prepares the GetEffectiveRouteTable request.
func (client InterfacesClient) getEffectiveRouteTablePreparer(resourceGroupName string, networkInterfaceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/effectiveRouteTable"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getEffectiveRouteTableResponder handles the response to the GetEffectiveRouteTable request.
func (client InterfacesClient) getEffectiveRouteTableResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &EffectiveRouteListResult{rawResponse: resp.Response()}
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

// GetVirtualMachineScaleSetNetworkInterface the Get network interface operation retrieves information about the
// specified network interface in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the name of the network
// interface. expand is expand references resources.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand *string) (*Interface, error) {
	req, err := client.getVirtualMachineScaleSetNetworkInterfacePreparer(resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getVirtualMachineScaleSetNetworkInterfaceResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Interface), err
}

// getVirtualMachineScaleSetNetworkInterfacePreparer prepares the GetVirtualMachineScaleSetNetworkInterface request.
func (client InterfacesClient) getVirtualMachineScaleSetNetworkInterfacePreparer(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getVirtualMachineScaleSetNetworkInterfaceResponder handles the response to the GetVirtualMachineScaleSetNetworkInterface request.
func (client InterfacesClient) getVirtualMachineScaleSetNetworkInterfaceResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Interface{rawResponse: resp.Response()}
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

// List the List networkInterfaces operation retrieves all the networkInterfaces in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client InterfacesClient) List(ctx context.Context, resourceGroupName string) (*InterfaceListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*InterfaceListResult), err
}

// listPreparer prepares the List request.
func (client InterfacesClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces"
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
func (client InterfacesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &InterfaceListResult{rawResponse: resp.Response()}
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

// ListAll the List networkInterfaces operation retrieves all the networkInterfaces in a subscription.
func (client InterfacesClient) ListAll(ctx context.Context) (*InterfaceListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*InterfaceListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client InterfacesClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listAllResponder handles the response to the ListAll request.
func (client InterfacesClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &InterfaceListResult{rawResponse: resp.Response()}
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

// ListEffectiveNetworkSecurityGroups the list effective network security group operation retrieves all the network
// security groups applied on a networkInterface. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. networkInterfaceName is the name of the network interface.
func (client InterfacesClient) ListEffectiveNetworkSecurityGroups(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*EffectiveNetworkSecurityGroupListResult, error) {
	req, err := client.listEffectiveNetworkSecurityGroupsPreparer(resourceGroupName, networkInterfaceName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listEffectiveNetworkSecurityGroupsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*EffectiveNetworkSecurityGroupListResult), err
}

// listEffectiveNetworkSecurityGroupsPreparer prepares the ListEffectiveNetworkSecurityGroups request.
func (client InterfacesClient) listEffectiveNetworkSecurityGroupsPreparer(resourceGroupName string, networkInterfaceName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/effectiveNetworkSecurityGroups"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listEffectiveNetworkSecurityGroupsResponder handles the response to the ListEffectiveNetworkSecurityGroups request.
func (client InterfacesClient) listEffectiveNetworkSecurityGroupsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusAccepted)
	if resp == nil {
		return nil, err
	}
	result := &EffectiveNetworkSecurityGroupListResult{rawResponse: resp.Response()}
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

// ListVirtualMachineScaleSetNetworkInterfaces the list network interface operation retrieves information about all
// network interfaces in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*InterfaceListResult, error) {
	req, err := client.listVirtualMachineScaleSetNetworkInterfacesPreparer(resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listVirtualMachineScaleSetNetworkInterfacesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*InterfaceListResult), err
}

// listVirtualMachineScaleSetNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetNetworkInterfaces request.
func (client InterfacesClient) listVirtualMachineScaleSetNetworkInterfacesPreparer(resourceGroupName string, virtualMachineScaleSetName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listVirtualMachineScaleSetNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetNetworkInterfaces request.
func (client InterfacesClient) listVirtualMachineScaleSetNetworkInterfacesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &InterfaceListResult{rawResponse: resp.Response()}
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

// ListVirtualMachineScaleSetVMNetworkInterfaces the list network interface operation retrieves information about all
// network interfaces in a virtual machine from a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set. virtualmachineIndex is the virtual machine index.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (*InterfaceListResult, error) {
	req, err := client.listVirtualMachineScaleSetVMNetworkInterfacesPreparer(resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listVirtualMachineScaleSetVMNetworkInterfacesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*InterfaceListResult), err
}

// listVirtualMachineScaleSetVMNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetVMNetworkInterfaces request.
func (client InterfacesClient) listVirtualMachineScaleSetVMNetworkInterfacesPreparer(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", APIVersion)
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listVirtualMachineScaleSetVMNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetVMNetworkInterfaces request.
func (client InterfacesClient) listVirtualMachineScaleSetVMNetworkInterfacesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &InterfaceListResult{rawResponse: resp.Response()}
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
