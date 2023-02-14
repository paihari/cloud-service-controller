// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ProjectSummaryCollection A collection of project summaries. The collection can be lightweight details or full definitions.
type ProjectSummaryCollection struct {

	// The array of Project summaries
	Items []ProjectSummary `mandatory:"true" json:"items"`
}

func (m ProjectSummaryCollection) String() string {
	return common.PointerString(m)
}
