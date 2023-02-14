// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary A summary of a work request.
type WorkRequestSummary struct {

	// The OCID of the work request.
	Id *string `mandatory:"false" json:"id"`

	// The type of work the work request is doing.
	OperationType OperationTypesEnum `mandatory:"false" json:"operationType,omitempty"`

	// The current status of the work request.
	Status OperationStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The OCID of the work request’s compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The resources this work request affects.
	Resources []WorkRequestResource `mandatory:"false" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// The time the work request was accepted.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time the work request was started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the work request was finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
