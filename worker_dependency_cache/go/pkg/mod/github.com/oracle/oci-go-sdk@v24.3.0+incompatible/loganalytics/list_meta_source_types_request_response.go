// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListMetaSourceTypesRequest wrapper for the ListMetaSourceTypes operation
type ListMetaSourceTypesRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// orderBy
	SortBy *string `mandatory:"false" contributesTo:"query" name:"sortBy"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMetaSourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMetaSourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMetaSourceTypesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMetaSourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMetaSourceTypesResponse wrapper for the ListMetaSourceTypes operation
type ListMetaSourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsMetaSourceTypeCollection instances
	LogAnalyticsMetaSourceTypeCollection `presentIn:"body"`

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

func (response ListMetaSourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMetaSourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMetaSourceTypesSortOrderEnum Enum with underlying type: string
type ListMetaSourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListMetaSourceTypesSortOrderEnum
const (
	ListMetaSourceTypesSortOrderAsc  ListMetaSourceTypesSortOrderEnum = "ASC"
	ListMetaSourceTypesSortOrderDesc ListMetaSourceTypesSortOrderEnum = "DESC"
)

var mappingListMetaSourceTypesSortOrder = map[string]ListMetaSourceTypesSortOrderEnum{
	"ASC":  ListMetaSourceTypesSortOrderAsc,
	"DESC": ListMetaSourceTypesSortOrderDesc,
}

// GetListMetaSourceTypesSortOrderEnumValues Enumerates the set of values for ListMetaSourceTypesSortOrderEnum
func GetListMetaSourceTypesSortOrderEnumValues() []ListMetaSourceTypesSortOrderEnum {
	values := make([]ListMetaSourceTypesSortOrderEnum, 0)
	for _, v := range mappingListMetaSourceTypesSortOrder {
		values = append(values, v)
	}
	return values
}
