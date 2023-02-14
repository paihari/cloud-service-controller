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

// LogAnalyticsParameter LogAnalyticsParameter
type LogAnalyticsParameter struct {

	// default value
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// is active flag
	IsActive *bool `mandatory:"false" json:"isActive"`

	// parameter name
	Name *string `mandatory:"false" json:"name"`

	// source Id
	SourceId *int64 `mandatory:"false" json:"sourceId"`
}

func (m LogAnalyticsParameter) String() string {
	return common.PointerString(m)
}
