// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSourcePatternsRequest wrapper for the ListSourcePatterns operation
type ListSourcePatternsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// source name
	SourceName *string `mandatory:"true" contributesTo:"path" name:"sourceName"`

	// is included source patterns
	IsInclude *bool `mandatory:"false" contributesTo:"query" name:"isInclude"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// sort by source pattern text
	SortBy ListSourcePatternsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourcePatternsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourcePatternsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourcePatternsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourcePatternsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSourcePatternsResponse wrapper for the ListSourcePatterns operation
type ListSourcePatternsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsSourcePatternCollection instances
	LogAnalyticsSourcePatternCollection `presentIn:"body"`

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

func (response ListSourcePatternsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourcePatternsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourcePatternsSortByEnum Enum with underlying type: string
type ListSourcePatternsSortByEnum string

// Set of constants representing the allowable values for ListSourcePatternsSortByEnum
const (
	ListSourcePatternsSortByPatterntext ListSourcePatternsSortByEnum = "patternText"
)

var mappingListSourcePatternsSortBy = map[string]ListSourcePatternsSortByEnum{
	"patternText": ListSourcePatternsSortByPatterntext,
}

// GetListSourcePatternsSortByEnumValues Enumerates the set of values for ListSourcePatternsSortByEnum
func GetListSourcePatternsSortByEnumValues() []ListSourcePatternsSortByEnum {
	values := make([]ListSourcePatternsSortByEnum, 0)
	for _, v := range mappingListSourcePatternsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSourcePatternsSortOrderEnum Enum with underlying type: string
type ListSourcePatternsSortOrderEnum string

// Set of constants representing the allowable values for ListSourcePatternsSortOrderEnum
const (
	ListSourcePatternsSortOrderAsc  ListSourcePatternsSortOrderEnum = "ASC"
	ListSourcePatternsSortOrderDesc ListSourcePatternsSortOrderEnum = "DESC"
)

var mappingListSourcePatternsSortOrder = map[string]ListSourcePatternsSortOrderEnum{
	"ASC":  ListSourcePatternsSortOrderAsc,
	"DESC": ListSourcePatternsSortOrderDesc,
}

// GetListSourcePatternsSortOrderEnumValues Enumerates the set of values for ListSourcePatternsSortOrderEnum
func GetListSourcePatternsSortOrderEnumValues() []ListSourcePatternsSortOrderEnum {
	values := make([]ListSourcePatternsSortOrderEnum, 0)
	for _, v := range mappingListSourcePatternsSortOrder {
		values = append(values, v)
	}
	return values
}
