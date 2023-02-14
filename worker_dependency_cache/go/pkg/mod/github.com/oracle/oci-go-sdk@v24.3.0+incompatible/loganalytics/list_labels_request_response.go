// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLabelsRequest wrapper for the ListLabels operation
type ListLabelsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// label name
	LabelName *string `mandatory:"false" contributesTo:"query" name:"labelName"`

	// search by label display name or description
	LabelDisplayText *string `mandatory:"false" contributesTo:"query" name:"labelDisplayText"`

	// Is system param of value (all, custom, sourceUsing)
	IsSystem ListLabelsIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// label priority
	LabelPriority ListLabelsLabelPriorityEnum `mandatory:"false" contributesTo:"query" name:"labelPriority" omitEmpty:"true"`

	// isCountPop
	IsCountPop *bool `mandatory:"false" contributesTo:"query" name:"isCountPop"`

	// isAliasPop
	IsAliasPop *bool `mandatory:"false" contributesTo:"query" name:"isAliasPop"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLabelsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// sort by label
	LabelSortBy ListLabelsLabelSortByEnum `mandatory:"false" contributesTo:"query" name:"labelSortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLabelsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLabelsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLabelsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLabelsResponse wrapper for the ListLabels operation
type ListLabelsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsLabelCollection instances
	LogAnalyticsLabelCollection `presentIn:"body"`

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

func (response ListLabelsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLabelsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLabelsIsSystemEnum Enum with underlying type: string
type ListLabelsIsSystemEnum string

// Set of constants representing the allowable values for ListLabelsIsSystemEnum
const (
	ListLabelsIsSystemAll     ListLabelsIsSystemEnum = "ALL"
	ListLabelsIsSystemCustom  ListLabelsIsSystemEnum = "CUSTOM"
	ListLabelsIsSystemBuiltIn ListLabelsIsSystemEnum = "BUILT_IN"
)

var mappingListLabelsIsSystem = map[string]ListLabelsIsSystemEnum{
	"ALL":      ListLabelsIsSystemAll,
	"CUSTOM":   ListLabelsIsSystemCustom,
	"BUILT_IN": ListLabelsIsSystemBuiltIn,
}

// GetListLabelsIsSystemEnumValues Enumerates the set of values for ListLabelsIsSystemEnum
func GetListLabelsIsSystemEnumValues() []ListLabelsIsSystemEnum {
	values := make([]ListLabelsIsSystemEnum, 0)
	for _, v := range mappingListLabelsIsSystem {
		values = append(values, v)
	}
	return values
}

// ListLabelsLabelPriorityEnum Enum with underlying type: string
type ListLabelsLabelPriorityEnum string

// Set of constants representing the allowable values for ListLabelsLabelPriorityEnum
const (
	ListLabelsLabelPriorityNone   ListLabelsLabelPriorityEnum = "NONE"
	ListLabelsLabelPriorityLow    ListLabelsLabelPriorityEnum = "LOW"
	ListLabelsLabelPriorityMedium ListLabelsLabelPriorityEnum = "MEDIUM"
	ListLabelsLabelPriorityHigh   ListLabelsLabelPriorityEnum = "HIGH"
)

var mappingListLabelsLabelPriority = map[string]ListLabelsLabelPriorityEnum{
	"NONE":   ListLabelsLabelPriorityNone,
	"LOW":    ListLabelsLabelPriorityLow,
	"MEDIUM": ListLabelsLabelPriorityMedium,
	"HIGH":   ListLabelsLabelPriorityHigh,
}

// GetListLabelsLabelPriorityEnumValues Enumerates the set of values for ListLabelsLabelPriorityEnum
func GetListLabelsLabelPriorityEnumValues() []ListLabelsLabelPriorityEnum {
	values := make([]ListLabelsLabelPriorityEnum, 0)
	for _, v := range mappingListLabelsLabelPriority {
		values = append(values, v)
	}
	return values
}

// ListLabelsSortOrderEnum Enum with underlying type: string
type ListLabelsSortOrderEnum string

// Set of constants representing the allowable values for ListLabelsSortOrderEnum
const (
	ListLabelsSortOrderAsc  ListLabelsSortOrderEnum = "ASC"
	ListLabelsSortOrderDesc ListLabelsSortOrderEnum = "DESC"
)

var mappingListLabelsSortOrder = map[string]ListLabelsSortOrderEnum{
	"ASC":  ListLabelsSortOrderAsc,
	"DESC": ListLabelsSortOrderDesc,
}

// GetListLabelsSortOrderEnumValues Enumerates the set of values for ListLabelsSortOrderEnum
func GetListLabelsSortOrderEnumValues() []ListLabelsSortOrderEnum {
	values := make([]ListLabelsSortOrderEnum, 0)
	for _, v := range mappingListLabelsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLabelsLabelSortByEnum Enum with underlying type: string
type ListLabelsLabelSortByEnum string

// Set of constants representing the allowable values for ListLabelsLabelSortByEnum
const (
	ListLabelsLabelSortByName        ListLabelsLabelSortByEnum = "name"
	ListLabelsLabelSortByPriority    ListLabelsLabelSortByEnum = "priority"
	ListLabelsLabelSortBySourceusing ListLabelsLabelSortByEnum = "sourceUsing"
)

var mappingListLabelsLabelSortBy = map[string]ListLabelsLabelSortByEnum{
	"name":        ListLabelsLabelSortByName,
	"priority":    ListLabelsLabelSortByPriority,
	"sourceUsing": ListLabelsLabelSortBySourceusing,
}

// GetListLabelsLabelSortByEnumValues Enumerates the set of values for ListLabelsLabelSortByEnum
func GetListLabelsLabelSortByEnumValues() []ListLabelsLabelSortByEnum {
	values := make([]ListLabelsLabelSortByEnum, 0)
	for _, v := range mappingListLabelsLabelSortBy {
		values = append(values, v)
	}
	return values
}
