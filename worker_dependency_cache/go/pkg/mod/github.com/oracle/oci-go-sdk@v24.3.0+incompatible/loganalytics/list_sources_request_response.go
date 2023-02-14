// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSourcesRequest wrapper for the ListSources operation
type ListSourcesRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// entityType
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// search by source display name or description
	SourceDisplayText *string `mandatory:"false" contributesTo:"query" name:"sourceDisplayText"`

	// Is system param of value (all, custom, sourceUsing)
	IsSystem ListSourcesIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// auto association flag
	IsAutoAssociated *bool `mandatory:"false" contributesTo:"query" name:"isAutoAssociated"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// sort by source
	SortBy ListSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// is simplified
	IsSimplified *bool `mandatory:"false" contributesTo:"query" name:"isSimplified"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourcesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSourcesResponse wrapper for the ListSources operation
type ListSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsSourceCollection instances
	LogAnalyticsSourceCollection `presentIn:"body"`

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

func (response ListSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourcesIsSystemEnum Enum with underlying type: string
type ListSourcesIsSystemEnum string

// Set of constants representing the allowable values for ListSourcesIsSystemEnum
const (
	ListSourcesIsSystemAll     ListSourcesIsSystemEnum = "ALL"
	ListSourcesIsSystemCustom  ListSourcesIsSystemEnum = "CUSTOM"
	ListSourcesIsSystemBuiltIn ListSourcesIsSystemEnum = "BUILT_IN"
)

var mappingListSourcesIsSystem = map[string]ListSourcesIsSystemEnum{
	"ALL":      ListSourcesIsSystemAll,
	"CUSTOM":   ListSourcesIsSystemCustom,
	"BUILT_IN": ListSourcesIsSystemBuiltIn,
}

// GetListSourcesIsSystemEnumValues Enumerates the set of values for ListSourcesIsSystemEnum
func GetListSourcesIsSystemEnumValues() []ListSourcesIsSystemEnum {
	values := make([]ListSourcesIsSystemEnum, 0)
	for _, v := range mappingListSourcesIsSystem {
		values = append(values, v)
	}
	return values
}

// ListSourcesSortOrderEnum Enum with underlying type: string
type ListSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSourcesSortOrderEnum
const (
	ListSourcesSortOrderAsc  ListSourcesSortOrderEnum = "ASC"
	ListSourcesSortOrderDesc ListSourcesSortOrderEnum = "DESC"
)

var mappingListSourcesSortOrder = map[string]ListSourcesSortOrderEnum{
	"ASC":  ListSourcesSortOrderAsc,
	"DESC": ListSourcesSortOrderDesc,
}

// GetListSourcesSortOrderEnumValues Enumerates the set of values for ListSourcesSortOrderEnum
func GetListSourcesSortOrderEnumValues() []ListSourcesSortOrderEnum {
	values := make([]ListSourcesSortOrderEnum, 0)
	for _, v := range mappingListSourcesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSourcesSortByEnum Enum with underlying type: string
type ListSourcesSortByEnum string

// Set of constants representing the allowable values for ListSourcesSortByEnum
const (
	ListSourcesSortByName             ListSourcesSortByEnum = "name"
	ListSourcesSortByTimeupdated      ListSourcesSortByEnum = "timeUpdated"
	ListSourcesSortByAssociationcount ListSourcesSortByEnum = "associationCount"
	ListSourcesSortBySourcetype       ListSourcesSortByEnum = "sourceType"
)

var mappingListSourcesSortBy = map[string]ListSourcesSortByEnum{
	"name":             ListSourcesSortByName,
	"timeUpdated":      ListSourcesSortByTimeupdated,
	"associationCount": ListSourcesSortByAssociationcount,
	"sourceType":       ListSourcesSortBySourcetype,
}

// GetListSourcesSortByEnumValues Enumerates the set of values for ListSourcesSortByEnum
func GetListSourcesSortByEnumValues() []ListSourcesSortByEnum {
	values := make([]ListSourcesSortByEnum, 0)
	for _, v := range mappingListSourcesSortBy {
		values = append(values, v)
	}
	return values
}
