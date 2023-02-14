// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ScopeFilter Scope filter to reduce the scope of the query.
type ScopeFilter struct {

	// Field must be a valid enterprise logging out-of-the-box field.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Field values that will be used to filter the query scope. Please note all values should reflect the fields data type otherwise the query is subject to fail.
	Values []interface{} `mandatory:"true" json:"values"`
}

func (m ScopeFilter) String() string {
	return common.PointerString(m)
}
