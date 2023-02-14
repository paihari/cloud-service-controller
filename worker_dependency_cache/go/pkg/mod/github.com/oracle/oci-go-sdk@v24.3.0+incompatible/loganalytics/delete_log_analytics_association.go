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

// DeleteLogAnalyticsAssociation DeleteLogAnalyticsAssociation
type DeleteLogAnalyticsAssociation struct {

	// Lama Idf
	AgentId *string `mandatory:"false" json:"agentId"`

	// source name
	SourceName *string `mandatory:"false" json:"sourceName"`

	// source type internal name
	SourceTypeName *string `mandatory:"false" json:"sourceTypeName"`

	// entity GUID
	EntityId *string `mandatory:"false" json:"entityId"`

	// entity type internal name
	EntityTypeName *string `mandatory:"false" json:"entityTypeName"`

	// host name
	Host *string `mandatory:"false" json:"host"`

	// log group ocid
	LogGroupId *string `mandatory:"false" json:"logGroupId"`
}

func (m DeleteLogAnalyticsAssociation) String() string {
	return common.PointerString(m)
}
