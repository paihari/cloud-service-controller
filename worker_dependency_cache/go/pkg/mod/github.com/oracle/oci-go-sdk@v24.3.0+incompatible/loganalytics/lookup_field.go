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

// LookupField LookupField
type LookupField struct {

	// common field name
	CommonFieldName *string `mandatory:"false" json:"commonFieldName"`

	// default match value
	DefaultMatchValue *string `mandatory:"false" json:"defaultMatchValue"`

	// display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// is common field
	IsCommonField *bool `mandatory:"false" json:"isCommonField"`

	// match operator
	MatchOperator *string `mandatory:"false" json:"matchOperator"`

	// name
	Name *string `mandatory:"false" json:"name"`

	// position
	Position *int64 `mandatory:"false" json:"position"`
}

func (m LookupField) String() string {
	return common.PointerString(m)
}
