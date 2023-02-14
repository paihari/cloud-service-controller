// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListParserMetaPluginsRequest wrapper for the ListParserMetaPlugins operation
type ListParserMetaPluginsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// orderBy
	SortBy *string `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParserMetaPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParserMetaPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParserMetaPluginsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParserMetaPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListParserMetaPluginsResponse wrapper for the ListParserMetaPlugins operation
type ListParserMetaPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsParserMetaPluginCollection instances
	LogAnalyticsParserMetaPluginCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListParserMetaPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParserMetaPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParserMetaPluginsSortOrderEnum Enum with underlying type: string
type ListParserMetaPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListParserMetaPluginsSortOrderEnum
const (
	ListParserMetaPluginsSortOrderAsc  ListParserMetaPluginsSortOrderEnum = "ASC"
	ListParserMetaPluginsSortOrderDesc ListParserMetaPluginsSortOrderEnum = "DESC"
)

var mappingListParserMetaPluginsSortOrder = map[string]ListParserMetaPluginsSortOrderEnum{
	"ASC":  ListParserMetaPluginsSortOrderAsc,
	"DESC": ListParserMetaPluginsSortOrderDesc,
}

// GetListParserMetaPluginsSortOrderEnumValues Enumerates the set of values for ListParserMetaPluginsSortOrderEnum
func GetListParserMetaPluginsSortOrderEnumValues() []ListParserMetaPluginsSortOrderEnum {
	values := make([]ListParserMetaPluginsSortOrderEnum, 0)
	for _, v := range mappingListParserMetaPluginsSortOrder {
		values = append(values, v)
	}
	return values
}
