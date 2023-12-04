//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolUsagesServer is a fake server for instances of the armsynapse.SQLPoolUsagesClient type.
type SQLPoolUsagesServer struct {
	// NewListPager is the fake for method SQLPoolUsagesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, sqlPoolName string, options *armsynapse.SQLPoolUsagesClientListOptions) (resp azfake.PagerResponder[armsynapse.SQLPoolUsagesClientListResponse])
}

// NewSQLPoolUsagesServerTransport creates a new instance of SQLPoolUsagesServerTransport with the provided implementation.
// The returned SQLPoolUsagesServerTransport instance is connected to an instance of armsynapse.SQLPoolUsagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolUsagesServerTransport(srv *SQLPoolUsagesServer) *SQLPoolUsagesServerTransport {
	return &SQLPoolUsagesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsynapse.SQLPoolUsagesClientListResponse]](),
	}
}

// SQLPoolUsagesServerTransport connects instances of armsynapse.SQLPoolUsagesClient to instances of SQLPoolUsagesServer.
// Don't use this type directly, use NewSQLPoolUsagesServerTransport instead.
type SQLPoolUsagesServerTransport struct {
	srv          *SQLPoolUsagesServer
	newListPager *tracker[azfake.PagerResponder[armsynapse.SQLPoolUsagesClientListResponse]]
}

// Do implements the policy.Transporter interface for SQLPoolUsagesServerTransport.
func (s *SQLPoolUsagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLPoolUsagesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLPoolUsagesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsynapse.SQLPoolUsagesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}