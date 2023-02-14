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

// UpdateLogAnalyticsEntityDetails Details of log analytics entity to be updated.
type UpdateLogAnalyticsEntityDetails struct {

	// Log analytics entity name. The name must be unique, within the tenancy, and cannot be changed.
	Name *string `mandatory:"false" json:"name"`

	// The OCID of the Management Agent.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The timezone region of the log analytics entity.
	TimezoneRegion *string `mandatory:"false" json:"timezoneRegion"`

	// The hostname where the entity represented here is actually present. This would be the output one would get if
	// they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from
	// management agents host since logs may be collected remotely.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties map[string]string `mandatory:"false" json:"properties"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateLogAnalyticsEntityDetails) String() string {
	return common.PointerString(m)
}
