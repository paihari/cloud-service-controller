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

// LogAnalyticsEntitySummary Summary of a log analytics entity.
type LogAnalyticsEntitySummary struct {

	// The log analytics entity OCID. This ID is a reference used by log analytics features and it represents
	// a resource that is provisioned and managed by the customer on their premises or on the cloud.
	Id *string `mandatory:"true" json:"id"`

	// Log analytics entity name. The name must be unique, within the tenancy, and cannot be changed.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Log analytics entity type name.
	EntityTypeName *string `mandatory:"true" json:"entityTypeName"`

	// Internal name for the log analytics entity type.
	EntityTypeInternalName *string `mandatory:"true" json:"entityTypeInternalName"`

	// The current state of the log analytics entity.
	LifecycleState EntityLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
	LifecycleDetails *string `mandatory:"true" json:"lifecycleDetails"`

	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID of the Management Agent.
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity
	// represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId *string `mandatory:"false" json:"cloudResourceId"`

	// The timezone region of the log analytics entity.
	TimezoneRegion *string `mandatory:"false" json:"timezoneRegion"`

	// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
	AreLogsCollected *bool `mandatory:"false" json:"areLogsCollected"`

	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId *string `mandatory:"false" json:"sourceId"`
}

func (m LogAnalyticsEntitySummary) String() string {
	return common.PointerString(m)
}
