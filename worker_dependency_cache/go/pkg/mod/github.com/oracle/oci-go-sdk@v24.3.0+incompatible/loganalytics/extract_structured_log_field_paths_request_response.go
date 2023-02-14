// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ExtractStructuredLogFieldPathsRequest wrapper for the ExtractStructuredLogFieldPaths operation
type ExtractStructuredLogFieldPathsRequest struct {

	// The Log Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// parser definition
	LoganParserDetails LogAnalyticsParser `contributesTo:"body"`

	// type - possible values are xml or json
	ParserType ExtractStructuredLogFieldPathsParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ExtractStructuredLogFieldPathsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ExtractStructuredLogFieldPathsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ExtractStructuredLogFieldPathsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ExtractStructuredLogFieldPathsResponse wrapper for the ExtractStructuredLogFieldPaths operation
type ExtractStructuredLogFieldPathsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExtractLogFieldResults instance
	ExtractLogFieldResults `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ExtractStructuredLogFieldPathsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ExtractStructuredLogFieldPathsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ExtractStructuredLogFieldPathsParserTypeEnum Enum with underlying type: string
type ExtractStructuredLogFieldPathsParserTypeEnum string

// Set of constants representing the allowable values for ExtractStructuredLogFieldPathsParserTypeEnum
const (
	ExtractStructuredLogFieldPathsParserTypeXml  ExtractStructuredLogFieldPathsParserTypeEnum = "XML"
	ExtractStructuredLogFieldPathsParserTypeJson ExtractStructuredLogFieldPathsParserTypeEnum = "JSON"
)

var mappingExtractStructuredLogFieldPathsParserType = map[string]ExtractStructuredLogFieldPathsParserTypeEnum{
	"XML":  ExtractStructuredLogFieldPathsParserTypeXml,
	"JSON": ExtractStructuredLogFieldPathsParserTypeJson,
}

// GetExtractStructuredLogFieldPathsParserTypeEnumValues Enumerates the set of values for ExtractStructuredLogFieldPathsParserTypeEnum
func GetExtractStructuredLogFieldPathsParserTypeEnumValues() []ExtractStructuredLogFieldPathsParserTypeEnum {
	values := make([]ExtractStructuredLogFieldPathsParserTypeEnum, 0)
	for _, v := range mappingExtractStructuredLogFieldPathsParserType {
		values = append(values, v)
	}
	return values
}
