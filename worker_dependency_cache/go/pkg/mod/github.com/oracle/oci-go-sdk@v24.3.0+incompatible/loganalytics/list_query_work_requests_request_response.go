// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListQueryWorkRequestsRequest wrapper for the ListQueryWorkRequests operation
type ListQueryWorkRequestsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter based on job execution mode
	Mode ListQueryWorkRequestsModeEnum `mandatory:"false" contributesTo:"query" name:"mode" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListQueryWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeStarted is descending. If no value is specified timeStarted is default.
	SortBy ListQueryWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQueryWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQueryWorkRequestsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQueryWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListQueryWorkRequestsResponse wrapper for the ListQueryWorkRequests operation
type ListQueryWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueryWorkRequestCollection instances
	QueryWorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListQueryWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQueryWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQueryWorkRequestsModeEnum Enum with underlying type: string
type ListQueryWorkRequestsModeEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsModeEnum
const (
	ListQueryWorkRequestsModeAll        ListQueryWorkRequestsModeEnum = "ALL"
	ListQueryWorkRequestsModeForeground ListQueryWorkRequestsModeEnum = "FOREGROUND"
	ListQueryWorkRequestsModeBackground ListQueryWorkRequestsModeEnum = "BACKGROUND"
)

var mappingListQueryWorkRequestsMode = map[string]ListQueryWorkRequestsModeEnum{
	"ALL":        ListQueryWorkRequestsModeAll,
	"FOREGROUND": ListQueryWorkRequestsModeForeground,
	"BACKGROUND": ListQueryWorkRequestsModeBackground,
}

// GetListQueryWorkRequestsModeEnumValues Enumerates the set of values for ListQueryWorkRequestsModeEnum
func GetListQueryWorkRequestsModeEnumValues() []ListQueryWorkRequestsModeEnum {
	values := make([]ListQueryWorkRequestsModeEnum, 0)
	for _, v := range mappingListQueryWorkRequestsMode {
		values = append(values, v)
	}
	return values
}

// ListQueryWorkRequestsSortOrderEnum Enum with underlying type: string
type ListQueryWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsSortOrderEnum
const (
	ListQueryWorkRequestsSortOrderAsc  ListQueryWorkRequestsSortOrderEnum = "ASC"
	ListQueryWorkRequestsSortOrderDesc ListQueryWorkRequestsSortOrderEnum = "DESC"
)

var mappingListQueryWorkRequestsSortOrder = map[string]ListQueryWorkRequestsSortOrderEnum{
	"ASC":  ListQueryWorkRequestsSortOrderAsc,
	"DESC": ListQueryWorkRequestsSortOrderDesc,
}

// GetListQueryWorkRequestsSortOrderEnumValues Enumerates the set of values for ListQueryWorkRequestsSortOrderEnum
func GetListQueryWorkRequestsSortOrderEnumValues() []ListQueryWorkRequestsSortOrderEnum {
	values := make([]ListQueryWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListQueryWorkRequestsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListQueryWorkRequestsSortByEnum Enum with underlying type: string
type ListQueryWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsSortByEnum
const (
	ListQueryWorkRequestsSortByTimestarted ListQueryWorkRequestsSortByEnum = "timeStarted"
	ListQueryWorkRequestsSortByTimeexpires ListQueryWorkRequestsSortByEnum = "timeExpires"
)

var mappingListQueryWorkRequestsSortBy = map[string]ListQueryWorkRequestsSortByEnum{
	"timeStarted": ListQueryWorkRequestsSortByTimestarted,
	"timeExpires": ListQueryWorkRequestsSortByTimeexpires,
}

// GetListQueryWorkRequestsSortByEnumValues Enumerates the set of values for ListQueryWorkRequestsSortByEnum
func GetListQueryWorkRequestsSortByEnumValues() []ListQueryWorkRequestsSortByEnum {
	values := make([]ListQueryWorkRequestsSortByEnum, 0)
	for _, v := range mappingListQueryWorkRequestsSortBy {
		values = append(values, v)
	}
	return values
}
