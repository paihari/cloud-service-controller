// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteUploadFileRequest wrapper for the DeleteUploadFile operation
type DeleteUploadFileRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Unique internal identifier to refer to upload container
	UploadReference *string `mandatory:"true" contributesTo:"path" name:"uploadReference"`

	// Unique internal identifier to refer to upload file
	FileReference *string `mandatory:"true" contributesTo:"path" name:"fileReference"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteUploadFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteUploadFileRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteUploadFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteUploadFileResponse wrapper for the DeleteUploadFile operation
type DeleteUploadFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Deleted log records count.
	OpcDeletedLogCount *int64 `presentIn:"header" name:"opc-deleted-log-count"`

	// Deleted log files count.
	OpcDeletedLogfileCount *int64 `presentIn:"header" name:"opc-deleted-logfile-count"`
}

func (response DeleteUploadFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteUploadFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
