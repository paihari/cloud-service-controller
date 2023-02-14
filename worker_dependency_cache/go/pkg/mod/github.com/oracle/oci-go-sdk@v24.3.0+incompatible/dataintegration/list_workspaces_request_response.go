// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkspacesRequest wrapper for the ListWorkspaces operation
type ListWorkspacesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Lifecycle state of the resource.
	LifecycleState WorkspaceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListWorkspacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkspacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkspacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkspacesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkspacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkspacesResponse wrapper for the ListWorkspaces operation
type ListWorkspacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkspaceSummary instances
	Items []WorkspaceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkspacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkspacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkspacesSortOrderEnum Enum with underlying type: string
type ListWorkspacesSortOrderEnum string

// Set of constants representing the allowable values for ListWorkspacesSortOrderEnum
const (
	ListWorkspacesSortOrderAsc  ListWorkspacesSortOrderEnum = "ASC"
	ListWorkspacesSortOrderDesc ListWorkspacesSortOrderEnum = "DESC"
)

var mappingListWorkspacesSortOrder = map[string]ListWorkspacesSortOrderEnum{
	"ASC":  ListWorkspacesSortOrderAsc,
	"DESC": ListWorkspacesSortOrderDesc,
}

// GetListWorkspacesSortOrderEnumValues Enumerates the set of values for ListWorkspacesSortOrderEnum
func GetListWorkspacesSortOrderEnumValues() []ListWorkspacesSortOrderEnum {
	values := make([]ListWorkspacesSortOrderEnum, 0)
	for _, v := range mappingListWorkspacesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkspacesSortByEnum Enum with underlying type: string
type ListWorkspacesSortByEnum string

// Set of constants representing the allowable values for ListWorkspacesSortByEnum
const (
	ListWorkspacesSortByTimeCreated ListWorkspacesSortByEnum = "TIME_CREATED"
	ListWorkspacesSortByDisplayName ListWorkspacesSortByEnum = "DISPLAY_NAME"
)

var mappingListWorkspacesSortBy = map[string]ListWorkspacesSortByEnum{
	"TIME_CREATED": ListWorkspacesSortByTimeCreated,
	"DISPLAY_NAME": ListWorkspacesSortByDisplayName,
}

// GetListWorkspacesSortByEnumValues Enumerates the set of values for ListWorkspacesSortByEnum
func GetListWorkspacesSortByEnumValues() []ListWorkspacesSortByEnum {
	values := make([]ListWorkspacesSortByEnum, 0)
	for _, v := range mappingListWorkspacesSortBy {
		values = append(values, v)
	}
	return values
}
