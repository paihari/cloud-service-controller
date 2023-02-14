// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListUploadFilesRequest wrapper for the ListUploadFiles operation
type ListUploadFilesRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Unique internal identifier to refer to upload container
	UploadReference *string `mandatory:"true" contributesTo:"path" name:"uploadReference"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListUploadFilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	SortBy ListUploadFilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Search string
	SearchStr *string `mandatory:"false" contributesTo:"query" name:"searchStr"`

	// Status
	Status []ListUploadFilesStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUploadFilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUploadFilesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUploadFilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUploadFilesResponse wrapper for the ListUploadFiles operation
type ListUploadFilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UploadFileCollection instances
	UploadFileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUploadFilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUploadFilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUploadFilesSortOrderEnum Enum with underlying type: string
type ListUploadFilesSortOrderEnum string

// Set of constants representing the allowable values for ListUploadFilesSortOrderEnum
const (
	ListUploadFilesSortOrderAsc  ListUploadFilesSortOrderEnum = "ASC"
	ListUploadFilesSortOrderDesc ListUploadFilesSortOrderEnum = "DESC"
)

var mappingListUploadFilesSortOrder = map[string]ListUploadFilesSortOrderEnum{
	"ASC":  ListUploadFilesSortOrderAsc,
	"DESC": ListUploadFilesSortOrderDesc,
}

// GetListUploadFilesSortOrderEnumValues Enumerates the set of values for ListUploadFilesSortOrderEnum
func GetListUploadFilesSortOrderEnumValues() []ListUploadFilesSortOrderEnum {
	values := make([]ListUploadFilesSortOrderEnum, 0)
	for _, v := range mappingListUploadFilesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUploadFilesSortByEnum Enum with underlying type: string
type ListUploadFilesSortByEnum string

// Set of constants representing the allowable values for ListUploadFilesSortByEnum
const (
	ListUploadFilesSortByTimecreated ListUploadFilesSortByEnum = "timeCreated"
	ListUploadFilesSortByFilename    ListUploadFilesSortByEnum = "fileName"
	ListUploadFilesSortByLoggroup    ListUploadFilesSortByEnum = "logGroup"
	ListUploadFilesSortBySourcename  ListUploadFilesSortByEnum = "sourceName"
	ListUploadFilesSortByStatus      ListUploadFilesSortByEnum = "status"
)

var mappingListUploadFilesSortBy = map[string]ListUploadFilesSortByEnum{
	"timeCreated": ListUploadFilesSortByTimecreated,
	"fileName":    ListUploadFilesSortByFilename,
	"logGroup":    ListUploadFilesSortByLoggroup,
	"sourceName":  ListUploadFilesSortBySourcename,
	"status":      ListUploadFilesSortByStatus,
}

// GetListUploadFilesSortByEnumValues Enumerates the set of values for ListUploadFilesSortByEnum
func GetListUploadFilesSortByEnumValues() []ListUploadFilesSortByEnum {
	values := make([]ListUploadFilesSortByEnum, 0)
	for _, v := range mappingListUploadFilesSortBy {
		values = append(values, v)
	}
	return values
}

// ListUploadFilesStatusEnum Enum with underlying type: string
type ListUploadFilesStatusEnum string

// Set of constants representing the allowable values for ListUploadFilesStatusEnum
const (
	ListUploadFilesStatusInProgress ListUploadFilesStatusEnum = "IN_PROGRESS"
	ListUploadFilesStatusSuccessful ListUploadFilesStatusEnum = "SUCCESSFUL"
	ListUploadFilesStatusFailed     ListUploadFilesStatusEnum = "FAILED"
)

var mappingListUploadFilesStatus = map[string]ListUploadFilesStatusEnum{
	"IN_PROGRESS": ListUploadFilesStatusInProgress,
	"SUCCESSFUL":  ListUploadFilesStatusSuccessful,
	"FAILED":      ListUploadFilesStatusFailed,
}

// GetListUploadFilesStatusEnumValues Enumerates the set of values for ListUploadFilesStatusEnum
func GetListUploadFilesStatusEnumValues() []ListUploadFilesStatusEnum {
	values := make([]ListUploadFilesStatusEnum, 0)
	for _, v := range mappingListUploadFilesStatus {
		values = append(values, v)
	}
	return values
}
