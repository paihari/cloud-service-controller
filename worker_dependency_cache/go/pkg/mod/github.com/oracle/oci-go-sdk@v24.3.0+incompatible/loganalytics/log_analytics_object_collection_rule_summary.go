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

// LogAnalyticsObjectCollectionRuleSummary The summary of an Object Storage based collection rule.
type LogAnalyticsObjectCollectionRuleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this rule.
	Id *string `mandatory:"true" json:"id"`

	// A unique name to the rule. The name must be unique, within the tenancy, and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to which this rule belongs.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Object Storage namespace.
	OsNamespace *string `mandatory:"true" json:"osNamespace"`

	// Name of the Object Storage bucket.
	OsBucketName *string `mandatory:"true" json:"osBucketName"`

	// The type of collection.
	// Accepted values are: LIVE.
	// Collection type LIVE indicates to enable log collection from the time of this rule creation,
	// and continue until the rule exists.
	CollectionType ObjectCollectionRuleCollectionTypesEnum `mandatory:"true" json:"collectionType"`

	// The current state of the rule.
	LifecycleState LogAnalyticsObjectCollectionRuleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when this rule was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this rule was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A unique name given to the rule. The name must be unique within the tenancy, and cannot be modified.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A detailed status of the life cycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m LogAnalyticsObjectCollectionRuleSummary) String() string {
	return common.PointerString(m)
}
