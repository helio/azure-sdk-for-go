//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualEndpointsServer is a fake server for instances of the armpostgresqlflexibleservers.VirtualEndpointsClient type.
type VirtualEndpointsServer struct {
	// BeginCreate is the fake for method VirtualEndpointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters armpostgresqlflexibleservers.VirtualEndpointResource, options *armpostgresqlflexibleservers.VirtualEndpointsClientBeginCreateOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *armpostgresqlflexibleservers.VirtualEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *armpostgresqlflexibleservers.VirtualEndpointsClientGetOptions) (resp azfake.Responder[armpostgresqlflexibleservers.VirtualEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method VirtualEndpointsClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armpostgresqlflexibleservers.VirtualEndpointsClientListByServerOptions) (resp azfake.PagerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientListByServerResponse])

	// BeginUpdate is the fake for method VirtualEndpointsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters armpostgresqlflexibleservers.VirtualEndpointResourceForPatch, options *armpostgresqlflexibleservers.VirtualEndpointsClientBeginUpdateOptions) (resp azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualEndpointsServerTransport creates a new instance of VirtualEndpointsServerTransport with the provided implementation.
// The returned VirtualEndpointsServerTransport instance is connected to an instance of armpostgresqlflexibleservers.VirtualEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualEndpointsServerTransport(srv *VirtualEndpointsServer) *VirtualEndpointsServerTransport {
	return &VirtualEndpointsServerTransport{
		srv:                  srv,
		beginCreate:          newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientCreateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientDeleteResponse]](),
		newListByServerPager: newTracker[azfake.PagerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientListByServerResponse]](),
		beginUpdate:          newTracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientUpdateResponse]](),
	}
}

// VirtualEndpointsServerTransport connects instances of armpostgresqlflexibleservers.VirtualEndpointsClient to instances of VirtualEndpointsServer.
// Don't use this type directly, use NewVirtualEndpointsServerTransport instead.
type VirtualEndpointsServerTransport struct {
	srv                  *VirtualEndpointsServer
	beginCreate          *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientCreateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientDeleteResponse]]
	newListByServerPager *tracker[azfake.PagerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientListByServerResponse]]
	beginUpdate          *tracker[azfake.PollerResponder[armpostgresqlflexibleservers.VirtualEndpointsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualEndpointsServerTransport.
func (v *VirtualEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualEndpointsClient.BeginCreate":
		resp, err = v.dispatchBeginCreate(req)
	case "VirtualEndpointsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualEndpointsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualEndpointsClient.NewListByServerPager":
		resp, err = v.dispatchNewListByServerPager(req)
	case "VirtualEndpointsClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualEndpointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := v.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualendpoints/(?P<virtualEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.VirtualEndpointResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		virtualEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreate(req.Context(), resourceGroupNameParam, serverNameParam, virtualEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		v.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		v.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		v.beginCreate.remove(req)
	}

	return resp, nil
}

func (v *VirtualEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualendpoints/(?P<virtualEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		virtualEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, virtualEndpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualendpoints/(?P<virtualEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	virtualEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, virtualEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualEndpointsServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := v.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualendpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		v.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armpostgresqlflexibleservers.VirtualEndpointsClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		v.newListByServerPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualEndpointsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/flexibleServers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualendpoints/(?P<virtualEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.VirtualEndpointResourceForPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		virtualEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, serverNameParam, virtualEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}