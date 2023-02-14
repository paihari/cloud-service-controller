// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLogAnalyticsEntitiesRequest wrapper for the ListLogAnalyticsEntities operation
type ListLogAnalyticsEntitiesRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only log analytics entities whose name contains the name given. The match
	// is case-insensitive.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only log analytics entities whose entityTypeName matches the entire log analytics entity type name of
	// one of the entityTypeNames given in the list. The match is case-insensitive.
	EntityTypeName []string `contributesTo:"query" name:"entityTypeName" collectionFormat:"multi"`

	// A filter to return only log analytics entities whose cloudResourceId matches the cloudResourceId given.
	CloudResourceId *string `mandatory:"false" contributesTo:"query" name:"cloudResourceId"`

	// A filter to return only those log analytics entities with the specified lifecycle state. The state
	// value is case-insensitive.
	LifecycleState ListLogAnalyticsEntitiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only log analytics entities whose lifecycleDetails contains the specified string.
	LifecycleDetailsContains *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetailsContains"`

	// A filter to return only those log analytics entities whose managementAgentId is null or is not null.
	IsManagementAgentIdNull ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum `mandatory:"false" contributesTo:"query" name:"isManagementAgentIdNull" omitEmpty:"true"`

	// A filter to return only log analytics entities whose hostname matches the entire hostname given.
	Hostname *string `mandatory:"false" contributesTo:"query" name:"hostname"`

	// A filter to return only log analytics entities whose hostname contains the substring given.
	// The match is case-insensitive.
	HostnameContains *string `mandatory:"false" contributesTo:"query" name:"hostnameContains"`

	// A filter to return only log analytics entities whose sourceId matches the sourceId given.
	SourceId *string `mandatory:"false" contributesTo:"query" name:"sourceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort entities by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for entity name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEntitiesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsEntitiesResponse wrapper for the ListLogAnalyticsEntities operation
type ListLogAnalyticsEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEntityCollection instances
	LogAnalyticsEntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEntitiesLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesLifecycleStateEnum
const (
	ListLogAnalyticsEntitiesLifecycleStateActive  ListLogAnalyticsEntitiesLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsEntitiesLifecycleStateDeleted ListLogAnalyticsEntitiesLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsEntitiesLifecycleState = map[string]ListLogAnalyticsEntitiesLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsEntitiesLifecycleStateActive,
	"DELETED": ListLogAnalyticsEntitiesLifecycleStateDeleted,
}

// GetListLogAnalyticsEntitiesLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesLifecycleStateEnum
func GetListLogAnalyticsEntitiesLifecycleStateEnumValues() []ListLogAnalyticsEntitiesLifecycleStateEnum {
	values := make([]ListLogAnalyticsEntitiesLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum
const (
	ListLogAnalyticsEntitiesIsManagementAgentIdNullTrue  ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum = "true"
	ListLogAnalyticsEntitiesIsManagementAgentIdNullFalse ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum = "false"
)

var mappingListLogAnalyticsEntitiesIsManagementAgentIdNull = map[string]ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum{
	"true":  ListLogAnalyticsEntitiesIsManagementAgentIdNullTrue,
	"false": ListLogAnalyticsEntitiesIsManagementAgentIdNullFalse,
}

// GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum
func GetListLogAnalyticsEntitiesIsManagementAgentIdNullEnumValues() []ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum {
	values := make([]ListLogAnalyticsEntitiesIsManagementAgentIdNullEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesIsManagementAgentIdNull {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntitiesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesSortOrderEnum
const (
	ListLogAnalyticsEntitiesSortOrderAsc  ListLogAnalyticsEntitiesSortOrderEnum = "ASC"
	ListLogAnalyticsEntitiesSortOrderDesc ListLogAnalyticsEntitiesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEntitiesSortOrder = map[string]ListLogAnalyticsEntitiesSortOrderEnum{
	"ASC":  ListLogAnalyticsEntitiesSortOrderAsc,
	"DESC": ListLogAnalyticsEntitiesSortOrderDesc,
}

// GetListLogAnalyticsEntitiesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesSortOrderEnum
func GetListLogAnalyticsEntitiesSortOrderEnumValues() []ListLogAnalyticsEntitiesSortOrderEnum {
	values := make([]ListLogAnalyticsEntitiesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntitiesSortByEnum Enum with underlying type: string
type ListLogAnalyticsEntitiesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntitiesSortByEnum
const (
	ListLogAnalyticsEntitiesSortByTimecreated ListLogAnalyticsEntitiesSortByEnum = "timeCreated"
	ListLogAnalyticsEntitiesSortByTimeupdated ListLogAnalyticsEntitiesSortByEnum = "timeUpdated"
	ListLogAnalyticsEntitiesSortByName        ListLogAnalyticsEntitiesSortByEnum = "name"
)

var mappingListLogAnalyticsEntitiesSortBy = map[string]ListLogAnalyticsEntitiesSortByEnum{
	"timeCreated": ListLogAnalyticsEntitiesSortByTimecreated,
	"timeUpdated": ListLogAnalyticsEntitiesSortByTimeupdated,
	"name":        ListLogAnalyticsEntitiesSortByName,
}

// GetListLogAnalyticsEntitiesSortByEnumValues Enumerates the set of values for ListLogAnalyticsEntitiesSortByEnum
func GetListLogAnalyticsEntitiesSortByEnumValues() []ListLogAnalyticsEntitiesSortByEnum {
	values := make([]ListLogAnalyticsEntitiesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEntitiesSortBy {
		values = append(values, v)
	}
	return values
}
