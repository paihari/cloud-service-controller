// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateProjectRequest wrapper for the UpdateProject operation
type UpdateProjectRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// DIS Project key
	ProjectKey *string `mandatory:"true" contributesTo:"path" name:"projectKey"`

	// The details needed to update a project.
	UpdateProjectDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Update and Delete operations should accept an optional If-Match header,
	// in which clients can send a previously-received ETag. When If-Match is
	// provided and its value does not exactly match the ETag of the resource
	// on the server, the request should fail with HTTP response status code 412
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateProjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateProjectRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateProjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateProjectResponse wrapper for the UpdateProject operation
type UpdateProjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Project instance
	Project `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateProjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateProjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
