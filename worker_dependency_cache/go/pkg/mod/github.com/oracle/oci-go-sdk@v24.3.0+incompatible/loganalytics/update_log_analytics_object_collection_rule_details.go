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

// UpdateLogAnalyticsObjectCollectionRuleDetails To update the attributes of an Object Storage based collection rule.
type UpdateLogAnalyticsObjectCollectionRuleDetails struct {

	// A string that describes the details of the rule.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Log Analytics Log group OCID to associate the processed logs with.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// Name of the Log Analytics Source to use for the processing.
	LogSourceName *string `mandatory:"false" json:"logSourceName"`

	// Log Analytics entity OCID. Associates the processed logs with the given entity (optional).
	EntityId *string `mandatory:"false" json:"entityId"`

	// An optional character encoding to aid in detecting the character encoding of the contents of the objects while processing.
	// It is recommended to set this value as ISO_8589_1 when configuring content of the objects having more numeric characters,
	// and very few alphabets.
	// For e.g. this applies when configuring VCN Flow Logs.
	CharEncoding *string `mandatory:"false" json:"charEncoding"`

	// Use this to override some property values which are defined at bucket level to the scope of object.
	// Supported propeties for override are, logSourceName, charEncoding.
	// Supported matchType for override are "contains".
	Overrides map[string][]PropertyOverride `mandatory:"false" json:"overrides"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateLogAnalyticsObjectCollectionRuleDetails) String() string {
	return common.PointerString(m)
}
