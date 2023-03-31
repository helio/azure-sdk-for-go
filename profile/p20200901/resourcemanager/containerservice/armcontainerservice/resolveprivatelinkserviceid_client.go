//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResolvePrivateLinkServiceIDClient contains the methods for the ResolvePrivateLinkServiceID group.
// Don't use this type directly, use NewResolvePrivateLinkServiceIDClient() instead.
type ResolvePrivateLinkServiceIDClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResolvePrivateLinkServiceIDClient creates a new instance of ResolvePrivateLinkServiceIDClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResolvePrivateLinkServiceIDClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResolvePrivateLinkServiceIDClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armcontainerservice.ResolvePrivateLinkServiceIDClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResolvePrivateLinkServiceIDClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// POST - Gets the private link service ID the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - Parameters (name, groupId) supplied in order to resolve a private link service ID.
//   - options - ResolvePrivateLinkServiceIDClientPOSTOptions contains the optional parameters for the ResolvePrivateLinkServiceIDClient.POST
//     method.
func (client *ResolvePrivateLinkServiceIDClient) POST(ctx context.Context, resourceGroupName string, resourceName string, parameters PrivateLinkResource, options *ResolvePrivateLinkServiceIDClientPOSTOptions) (ResolvePrivateLinkServiceIDClientPOSTResponse, error) {
	req, err := client.postCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return ResolvePrivateLinkServiceIDClientPOSTResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResolvePrivateLinkServiceIDClientPOSTResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResolvePrivateLinkServiceIDClientPOSTResponse{}, runtime.NewResponseError(resp)
	}
	return client.postHandleResponse(resp)
}

// postCreateRequest creates the POST request.
func (client *ResolvePrivateLinkServiceIDClient) postCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters PrivateLinkResource, options *ResolvePrivateLinkServiceIDClientPOSTOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/resolvePrivateLinkServiceId"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// postHandleResponse handles the POST response.
func (client *ResolvePrivateLinkServiceIDClient) postHandleResponse(resp *http.Response) (ResolvePrivateLinkServiceIDClientPOSTResponse, error) {
	result := ResolvePrivateLinkServiceIDClientPOSTResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return ResolvePrivateLinkServiceIDClientPOSTResponse{}, err
	}
	return result, nil
}