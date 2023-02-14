// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// UploadLogFileRequest wrapper for the UploadLogFile operation
type UploadLogFileRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the upload. This can be considered as a container name where different kind of logs will be collected and searched together. This upload name/id can further be used for retrieving the details of the upload, including its status or deleting the upload.
	UploadName *string `mandatory:"true" contributesTo:"query" name:"uploadName"`

	// Name of the log source that will be used to process the files being uploaded.
	LogSourceName *string `mandatory:"true" contributesTo:"query" name:"logSourceName"`

	// The name of the file being uploaded. The extension of the filename part will be used to detect the type of file (like zip, tar).
	Filename *string `mandatory:"true" contributesTo:"query" name:"filename"`

	// The log group OCID to which the log data in this upload will be mapped to.
	//   Example: `ocid1.loganalyticsloggroup.oc1..aaaaaaaad3q4sosi5i7z7onw2kgbwyk1581620537198`
	OpcMetaLoggrpid *string `mandatory:"true" contributesTo:"header" name:"opc-meta-loggrpid"`

	// Log data
	UploadLogFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// Timezone to be used when processing log entries whose timestamps do not include an explicit timezone. When this property is not specified, the timezone of the entity specified is used. If the entity is also not specified or do not have a valid timezone then UTC is used
	Timezone *string `mandatory:"false" contributesTo:"query" name:"timezone"`

	// character Encoding
	CharEncoding *string `mandatory:"false" contributesTo:"query" name:"charEncoding"`

	// This property is used to specify the format of the date. This is to be used for ambiguous dates like 12/11/10. This property can take any of the following values -  MONTH_DAY_YEAR, DAY_MONTH_YEAR, YEAR_MONTH_DAY, MONTH_DAY, DAY_MONTH.
	DateFormat *string `mandatory:"false" contributesTo:"query" name:"dateFormat"`

	// Used to indicate the year where the log entries timestamp do not mention year (Ex: Nov  8 20:45:56).
	DateYear *string `mandatory:"false" contributesTo:"query" name:"dateYear"`

	// This property can be used to reset configuration cache in case of an issue with the upload.
	InvalidateCache *bool `mandatory:"false" contributesTo:"query" name:"invalidateCache"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The base-64 encoded MD5 hash of the body. If the Content-MD5 header is present, Log Analytics performs an integrity check
	// on the body of the HTTP request by computing the MD5 hash for the body and comparing it to the MD5 hash supplied in the header.
	// If the two hashes do not match, the log data is rejected and an HTTP-400 Unmatched Content MD5 error is returned with the message:
	// "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMd5 *string `mandatory:"false" contributesTo:"header" name:"content-md5"`

	// The content type of the log data. Defaults to 'application/octet-stream' if not overridden during the UploadLogFile call.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"content-type"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UploadLogFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UploadLogFileRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UploadLogFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UploadLogFileResponse wrapper for the UploadLogFile operation
type UploadLogFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Upload instance
	Upload `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	OpcContentMd5 *string `presentIn:"header" name:"opc-content-md5"`

	// Unique Oracle-assigned identifier for log data.
	OpcObjectId *string `presentIn:"header" name:"opc-object-id"`
}

func (response UploadLogFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UploadLogFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
