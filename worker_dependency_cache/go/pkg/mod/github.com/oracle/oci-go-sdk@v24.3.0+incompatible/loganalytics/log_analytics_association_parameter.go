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

// LogAnalyticsAssociationParameter LogAnalyticsAssociationParameter
type LogAnalyticsAssociationParameter struct {

	// agent guid
	AgentId *string `mandatory:"false" json:"agentId"`

	// entity type
	EntityType *string `mandatory:"false" json:"entityType"`

	// entity guid
	EntityId *string `mandatory:"false" json:"entityId"`

	// source name
	SourceId *string `mandatory:"false" json:"sourceId"`

	// source display name
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// source type
	SourceType *string `mandatory:"false" json:"sourceType"`

	// status
	Status LogAnalyticsAssociationParameterStatusEnum `mandatory:"false" json:"status,omitempty"`

	// missingProperties
	MissingProperties []string `mandatory:"false" json:"missingProperties"`

	// requiredProperties
	RequiredProperties []string `mandatory:"false" json:"requiredProperties"`
}

func (m LogAnalyticsAssociationParameter) String() string {
	return common.PointerString(m)
}

// LogAnalyticsAssociationParameterStatusEnum Enum with underlying type: string
type LogAnalyticsAssociationParameterStatusEnum string

// Set of constants representing the allowable values for LogAnalyticsAssociationParameterStatusEnum
const (
	LogAnalyticsAssociationParameterStatusSucceeded LogAnalyticsAssociationParameterStatusEnum = "SUCCEEDED"
	LogAnalyticsAssociationParameterStatusFailed    LogAnalyticsAssociationParameterStatusEnum = "FAILED"
)

var mappingLogAnalyticsAssociationParameterStatus = map[string]LogAnalyticsAssociationParameterStatusEnum{
	"SUCCEEDED": LogAnalyticsAssociationParameterStatusSucceeded,
	"FAILED":    LogAnalyticsAssociationParameterStatusFailed,
}

// GetLogAnalyticsAssociationParameterStatusEnumValues Enumerates the set of values for LogAnalyticsAssociationParameterStatusEnum
func GetLogAnalyticsAssociationParameterStatusEnumValues() []LogAnalyticsAssociationParameterStatusEnum {
	values := make([]LogAnalyticsAssociationParameterStatusEnum, 0)
	for _, v := range mappingLogAnalyticsAssociationParameterStatus {
		values = append(values, v)
	}
	return values
}
