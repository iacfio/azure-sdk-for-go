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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// PublicIPAddressesGroupClient is the network Client
type PublicIPAddressesGroupClient struct {
	BaseClient
}

// NewPublicIPAddressesGroupClient creates an instance of the PublicIPAddressesGroupClient client.
func NewPublicIPAddressesGroupClient(subscriptionID string) PublicIPAddressesGroupClient {
	return NewPublicIPAddressesGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPublicIPAddressesGroupClientWithBaseURI creates an instance of the PublicIPAddressesGroupClient client.
func NewPublicIPAddressesGroupClientWithBaseURI(baseURI string, subscriptionID string) PublicIPAddressesGroupClient {
	return PublicIPAddressesGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a static or dynamic public IP address.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the public IP address.
// parameters is parameters supplied to the create or update public IP address operation.
func (client PublicIPAddressesGroupClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (result PublicIPAddressesGroupCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.PublicIPAddressesGroupClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, publicIPAddressName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PublicIPAddressesGroupClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) CreateOrUpdateSender(req *http.Request) (future PublicIPAddressesGroupCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) CreateOrUpdateResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified public IP address.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the subnet.
func (client PublicIPAddressesGroupClient) Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (result PublicIPAddressesGroupDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, publicIPAddressName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublicIPAddressesGroupClient) DeletePreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) DeleteSender(req *http.Request) (future PublicIPAddressesGroupDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified public IP address in a specified resource group.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the subnet. expand is
// expands referenced resources.
func (client PublicIPAddressesGroupClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result PublicIPAddress, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, publicIPAddressName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PublicIPAddressesGroupClient) GetPreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) GetResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetVirtualMachineScaleSetPublicIPAddress get the specified public IP address in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the name of the
// network interface. IPConfigurationName is the name of the IP configuration. publicIPAddressName is the name of
// the public IP Address. expand is expands referenced resources.
func (client PublicIPAddressesGroupClient) GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (result PublicIPAddress, err error) {
	req, err := client.GetVirtualMachineScaleSetPublicIPAddressPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName, publicIPAddressName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "GetVirtualMachineScaleSetPublicIPAddress", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVirtualMachineScaleSetPublicIPAddressSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "GetVirtualMachineScaleSetPublicIPAddress", resp, "Failure sending request")
		return
	}

	result, err = client.GetVirtualMachineScaleSetPublicIPAddressResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "GetVirtualMachineScaleSetPublicIPAddress", resp, "Failure responding to request")
	}

	return
}

// GetVirtualMachineScaleSetPublicIPAddressPreparer prepares the GetVirtualMachineScaleSetPublicIPAddress request.
func (client PublicIPAddressesGroupClient) GetVirtualMachineScaleSetPublicIPAddressPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":        autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"publicIpAddressName":        autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetVirtualMachineScaleSetPublicIPAddressSender sends the GetVirtualMachineScaleSetPublicIPAddress request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) GetVirtualMachineScaleSetPublicIPAddressSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetVirtualMachineScaleSetPublicIPAddressResponder handles the response to the GetVirtualMachineScaleSetPublicIPAddress request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) GetVirtualMachineScaleSetPublicIPAddressResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all public IP addresses in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client PublicIPAddressesGroupClient) List(ctx context.Context, resourceGroupName string) (result PublicIPAddressListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PublicIPAddressesGroupClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) ListResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesGroupClient) listNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesGroupClient) ListComplete(ctx context.Context, resourceGroupName string) (result PublicIPAddressListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all the public IP addresses in a subscription.
func (client PublicIPAddressesGroupClient) ListAll(ctx context.Context) (result PublicIPAddressListResultPage, err error) {
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client PublicIPAddressesGroupClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) ListAllResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesGroupClient) listAllNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesGroupClient) ListAllComplete(ctx context.Context) (result PublicIPAddressListResultIterator, err error) {
	result.page, err = client.ListAll(ctx)
	return
}

// ListVirtualMachineScaleSetPublicIPAddresses gets information about all public IP addresses on a virtual machine
// scale set level.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result PublicIPAddressListResultPage, err error) {
	result.fn = client.listVirtualMachineScaleSetPublicIPAddressesNextResults
	req, err := client.ListVirtualMachineScaleSetPublicIPAddressesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetPublicIPAddresses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetPublicIPAddressesSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetPublicIPAddresses", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListVirtualMachineScaleSetPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetPublicIPAddresses", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetPublicIPAddressesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetPublicIPAddressesSender sends the ListVirtualMachineScaleSetPublicIPAddresses request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetPublicIPAddressesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetPublicIPAddresses request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetPublicIPAddressesResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetPublicIPAddressesNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesGroupClient) listVirtualMachineScaleSetPublicIPAddressesNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetPublicIPAddressesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetPublicIPAddressesComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result PublicIPAddressListResultIterator, err error) {
	result.page, err = client.ListVirtualMachineScaleSetPublicIPAddresses(ctx, resourceGroupName, virtualMachineScaleSetName)
	return
}

// ListVirtualMachineScaleSetVMPublicIPAddresses gets information about all public IP addresses in a virtual machine IP
// configuration in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual
// machine scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the network
// interface name. IPConfigurationName is the IP configuration name.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetVMPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result PublicIPAddressListResultPage, err error) {
	result.fn = client.listVirtualMachineScaleSetVMPublicIPAddressesNextResults
	req, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetVMPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetVMPublicIPAddresses request.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetVMPublicIPAddressesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":        autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetVMPublicIPAddressesSender sends the ListVirtualMachineScaleSetVMPublicIPAddresses request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetVMPublicIPAddressesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetVMPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetVMPublicIPAddresses request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetVMPublicIPAddressesNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesGroupClient) listVirtualMachineScaleSetVMPublicIPAddressesNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesGroupClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetVMPublicIPAddressesComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesGroupClient) ListVirtualMachineScaleSetVMPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result PublicIPAddressListResultIterator, err error) {
	result.page, err = client.ListVirtualMachineScaleSetVMPublicIPAddresses(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName)
	return
}
